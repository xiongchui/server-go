package route

import (
	. "../http"
)

// todo, 新增 Response 类，避免路由函数硬编码
func image(r Request) []byte {
	p := r.Query["path"]
	b := Image(p)
	s := []byte("HTTP/1.1 200 OK\r\nContent-Type: Image/gif\r\n\r\n")
	if b == nil {
		s = ErrorResponse(404)
		return s
	}
	m := append([]byte(s), b...)
	return m
}

func index(r Request) []byte {
	s := "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\n\r\n"
	name := "index.html"
	t := Template(name)
	m := append([]byte(s), t...)
	return m
}

// todo, 该函数应该移入 http 包
func ErrorResponse(code int) []byte {
	// todo, 根据 code 值返回不同错误响应
	s := "HTTP/1.1 404 NOT FOUND\r\nContent-Type: text/html\r\n\r\n404 NOT FOUND!"
	return []byte(s)
}

func doge(r Request) []byte {
	s := "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\n\r\n"
	name := "doge.html"
	t := Template(name)
	m := append([]byte(s), t...)
	return m
}

var RouteIndex = map[string]func(Request) []byte{
	"/doge": doge,
	"/":     index,
	"/img":  image,
}
