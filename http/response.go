package http

import (
    "strings"
    "bytes"
)

type Response struct {
    Protocol string
    Status   string
    Headers  map[string]string
    Cookies  map[string]string
    Body     []byte
    Mime     string
}

func (r *Response) AddHeader(k string, v string) {
    r.Headers[k] = v
}

func (r *Response) Init() {
    r.Cookies = make(map[string]string)
    r.Headers = make(map[string]string)
    r.Protocol = "HTTP/1.1"
}

func (r *Response) Bytes() []byte {
    arr := []string{r.Protocol, r.Status, mapStatus[r.Status]}
    header := strings.Join(arr, " ")
    headers := []string{header}
    for k, v := range r.Headers {
        s := k + ": " + v
        headers = append(headers, s)
    }
    var cookies []string
    for k, v := range r.Cookies {
        s := k + "=" + v
        cookies = append(cookies, s)
    }
    cookie := "Cookie" + ": " + strings.Join(cookies, ": ")
    headers = append(headers, cookie)
    m := []byte(strings.Join(headers, "\r\n"))
    bb := [][]byte{m, r.Body}
    data := bytes.Join(bb, []byte("\r\n\r\n"))
    return data
}

var mapMime = map[string]string{
    "html": "text/html",
    "jpg":  "image/jpg",
    "gif":  "image/gif",
}

var mapFunc = map[string]func(string) ([]byte, error){
    "html": Template,
    "jpg":  Image,
    "gif":  Image,
}

var mapStatus = map[string]string{
    "200": "OK",
    "404": "Not Found",
    "302": "Moved Temporarily",
}

func ResponseFile(name string) []byte {
    var fn func(string) ([]byte, error)
    for k, v := range mapFunc {
        if strings.HasSuffix(name, k) {
            fn = v
            break
        }
    }
    b, err := fn(name)
    var r Response
    if err == nil {
        r = responseSuccess()
    } else {
        r = responseFailure()
    }
    r.Body = b
    suf := strings.Split(name, ".")[1]
    mapHeaders := map[string]string{
        "Content-Type": mapMime[suf],
    }
    for k, v := range mapHeaders {
        r.Headers[k] = v
    }
    d := r.Bytes()
    return d
}

func responseSuccess() Response {
    r := Response{}
    r.Init()
    r.Status = "200"
    return r
}

func responseFailure() Response {
    r := Response{}
    r.Init()
    r.Status = "404"
    return r
}
