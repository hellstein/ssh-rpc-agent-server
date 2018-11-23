package jobmgr
import (
    "golang.org/x/crypto/ssh"
    "github.com/gorilla/websocket"
    "log"
)


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
            //reader := os.Stdin
            buf := make([]byte, 1024)
            n, err := reader.Read(buf)
            if err != nil {
                log.Print(err)
                return
            }
            _, err = sessionWriter.Write(buf[:n])
            if err != nil {
                log.Print(err)
                conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
                return
            }
        }
    }(session, client, conn)


}
