package route

import (
    . "../http"
)

func image(r Request) []byte {
    p := r.Query["path"]
    m := ResponseImage(p)
    return m
}

func index(r Request) []byte {
    name := "index.html"
    m := ResponseTemplate(name)
    //fmt.Printf(string(m))
    return m
}

// todo, 该函数应该移入 http 包
func ErrorResponse(code int) []byte {
    // todo, 根据 code 值返回不同错误响应
    s := "HTTP/1.1 404 NOT FOUND\r\nContent-Type: text/html\r\n\r\n404 NOT FOUND!"
    return []byte(s)
}

func doge(r Request) []byte {
    name := "doge.html"
    m := ResponseTemplate(name)
    return m
}

var RouteIndex = map[string]func(Request) []byte{
    "/doge": doge,
    "/":     index,
    "/img":  image,
}
