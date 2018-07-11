package http

import (
    "strings"
    "bytes"
)

// todo, response 中应为 setcookie 值
type Response struct {
    Protocol  string
    Status    string
    Headers   map[string]string
    SetCookie map[string]string
    Body      []byte
    Mime      string
}

func (r *Response) AddHeader(k string, v string) {
    r.Headers[k] = v
}

func (r *Response) Init() {
    r.SetCookie = make(map[string]string)
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
    for k, v := range r.SetCookie {
        s := k + "=" + v
        cookies = append(cookies, s)
    }
    s := strings.Join(cookies, ": ")
    if s != "" {
        cookie := "Set-Cookie: " + s
        headers = append(headers, cookie)
    }
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
        r = responseByCode("200")
    } else {
        r = responseByCode("404")
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

// todo, error page 不一定存在, 应当使用默认页面
func ResponseError(code string) []byte {
    r := responseByCode(code)
    r.Body, _ = ErrorPageByCode(code)
    s := r.Bytes()
    return []byte(s)
}

func responseByCode(code string) Response {
    r := Response{}
    r.Init()
    r.Status = code
    return r
}
