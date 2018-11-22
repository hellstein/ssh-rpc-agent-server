package jobmgr

import (
    "log"
    "golang.org/x/crypto/ssh"
    "github.com/gorilla/websocket"
//    "bytes"
)
func syncIO (session *ssh.Session, conn *websocket.Conn, client *ssh.Client) {
    go func(*ssh.Session, *websocket.Conn, *ssh.Client) {
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

//        var b bytes.Buffer
//        session.Stdout = &b
        for {
            // set io.Writer of websocket
            outbuf := make([]byte, 8192)
            outn, err := sessionReader.Read(outbuf)
            if err != nil {
                log.Println("sshReader: ", err)
                return
            }
            /*
            log.Println(b.Len())
            outn, err := (&b).Read(outbuf)
            if err != nil {
                log.Println("sshReader: ", err)
                return
            }
            */
            log.Println(outn)
            log.Println(string(outbuf[:outn]))
            err = conn.WriteMessage(websocket.TextMessage, outbuf[:outn])
            if err != nil {
                log.Println("connWriter: ", err)
                return
            }
        }
    }(session, conn, client)

    go func(*ssh.Session, *websocket.Conn, *ssh.Client) {
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
            buf := make([]byte, 4096)
            n, err := reader.Read(buf)
            if err != nil {
                log.Print("connReader: ", err)
                return
            }
//            log.Println(string(buf[:n]))
            _, err = sessionWriter.Write(buf[:n])
            if err != nil {
                log.Print("sshWriter: ", err)
                conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
                return
            }
        }
    }(session, conn, client)
}
