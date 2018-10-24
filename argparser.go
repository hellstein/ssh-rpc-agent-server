package main
import (
    "net/http"
    "bytes"
    "io"
    "log"
)

func GetConf(r *http.Request) (map[string][]byte, error) {
    mconf, merr := getContentFromFile(r, "machinefile")
    if merr != nil {
        log.Println(merr)
        return nil, merr
    }
    tconf, terr := getContentFromFile(r, "taskfile")
    if terr != nil {
        log.Println(terr)
        return nil, terr
    }

    conf := map[string][]byte{"machine": mconf, "tasks": tconf}
    return conf, nil
}


func getContentFromFile(r *http.Request, filekey string) ([]byte, error) {
    var Buf bytes.Buffer
    file, _, err := r.FormFile(filekey)
    if err != nil {
       return nil, err
    }
    defer file.Close()
    // Copy the file data to my buffer
    io.Copy(&Buf, file)
    // do something with the contents...
    // I normally have a struct defined and unmarshal into a struct, but this will
    // work as an example
    // contents := Buf.String()
    // fmt.Fprintf(w, contents)
    contents := Buf.Bytes()
    // I reset the buffer in case I want to use it again
    // reduces memory allocations in more intense projects
    Buf.Reset()
    return contents, nil
}

