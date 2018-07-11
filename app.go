package main

import (
    "fmt"
    "net"
    "os"
    . "./http"
    . "./route"
)

func handleClient(conn net.Conn) {
    request := make([]byte, 1024)
    defer conn.Close()
    num, err := conn.Read(request)
    checkError(err)
    raw := string(request[:num])
    r := Request{}
    r.Init(raw)
    path := r.Path
    fmt.Println("path", path)
    s := responseForPath(path, r)
    conn.Write(s)
    request = make([]byte, 1024)
}

func responseForPath(path string, r Request) []byte {
    m := map[string]func(Request) []byte{}
    // todo, add route 函数
    for k, v := range RouteIndex {
        m[k] = v
    }
    fn, ok := m[path]
    var s []byte
    if !ok {
        s = ErrorResponse(404)
    } else {
        s = fn(r)
    }
    return s
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}

func main() {
    service := ":3000"
    s, e := net.ResolveTCPAddr("tcp4", service)
    checkError(e)
    l, e := net.ListenTCP("tcp", s)
    checkError(e)
    for {
        conn, err := l.Accept()
        if err != nil {
            continue
        }
        go handleClient(conn)
    }
}
