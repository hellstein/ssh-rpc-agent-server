package jobmgr
import (
    "golang.org/x/crypto/ssh"
    "github.com/gorilla/websocket"
    "log"
    "encoding/json"
)

/*
    Bind output of ssh session with websocket input
    Bind websocket output with input of ssh session
*/
func syncIO(session *ssh.Session, client *ssh.Client, conn *websocket.Conn) {
    go func(*ssh.Session, *ssh.Client, *websocket.Conn) {
        sessionReader, err := session.StdoutPipe()
        if err != nil {
          log.Fatal(err)
        }
        log.Println("======================== Sync session output ======================")
        defer func() {
            log.Println("======================== output: end ======================")
            conn.Close()
            client.Close()
            session.Close()
        }()

        for {
            // set io.Writer of websocket
            outbuf := make([]byte, 8192)
            outn, err := sessionReader.Read(outbuf)
            if err != nil {
                log.Println("sshReader: ", err)
                return
            }
//           fmt.Fprint(os.Stdout, string(outbuf[:outn]))
           err = conn.WriteMessage(websocket.TextMessage, outbuf[:outn])
           if err != nil {
               log.Println("connWriter: ", err)
               return
           }
       }
    }(session, client, conn)

    go func(*ssh.Session, *ssh.Client, *websocket.Conn) {
        sessionWriter, err := session.StdinPipe()
        if err != nil {
            log.Fatal(err)
        }

        log.Println("======================== Sync session input ======================")
        defer func() {
            log.Println("======================== input: end ======================")
            conn.Close()
            client.Close()
            session.Close()
        }()

        for {
            // set up io.Reader of websocket
            _, reader, err := conn.NextReader()
            if err != nil {
                log.Println("connReaderCreator: ", err)
                return
            }

            dataTypeBuf := make([]byte, 1)
            _, err = reader.Read(dataTypeBuf)
            if err != nil {
                log.Print(err)
                return
            }

            buf := make([]byte, 1024)
            n, err := reader.Read(buf)
            if err != nil {
                log.Print(err)
                return
            }

            switch dataTypeBuf[0] {
            // when pass data
            case '1':
                _, err = sessionWriter.Write(buf[:n])
                if err != nil {
                    log.Print(err)
                    conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
                    return
                }
             // when resize terminal
             case '0':
                 resizeMessage := WindowSize{}
                 err = json.Unmarshal(buf[:n], &resizeMessage)
                 if err != nil {
                     log.Print(err.Error())
                     continue
                 }
                 err = session.WindowChange(resizeMessage.Height, resizeMessage.Width)
                 if err != nil {
                     log.Print(err.Error())
                     conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
                     return
                 }
             // unexpected data
             default:
                 log.Print("Unexpected data type")
             }
        }
    }(session, client, conn)
}
