package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"io/ioutil"
)

type Request struct {
	raw     string
	method  string
	path    string
	query   map[string]string
	body    string
	headers map[string]string
	cookies map[string]string
}

func (r *Request) Init(raw string) {
	r.raw = raw
	f := strings.Split(raw, "\r\n")[0]
	es := strings.Split(f, " ")
	method, path := es[0], es[1]
	r.method = method
	r.body = strings.Split(raw, "\r\n\r\n")[1]
	r.headers = make(map[string]string)
	r.query = make(map[string]string)
	r.AddCookies()
	r.AddHeaders()
	r.ParseQuery(path)
}

func (r *Request) AddCookies() {
	e := r.headers["Cookie"]
	kvs := strings.Split(e, ": ")
	for _, s := range kvs {
		if strings.Contains(s, "=") {
			kv := strings.Split(s, "=")
			k, v := kv[0], kv[1]
			r.cookies[k] = v
		}
	}
}

func (r *Request) AddHeaders() {
	raw := r.raw
	s := strings.Split(raw, "\r\n\r\n")[0]
	hs := strings.Split(s, "\r\n")[1:]
	for _, s := range hs {
		kv := strings.Split(s, ": ")
		k, v := kv[0], kv[1]
		r.headers[k] = v
	}
}

func (r *Request) ParseQuery(query string) {
	if strings.Contains(query, "?") {
		es := strings.Split(query, "?")
		r.path = es[0]
		q := es[1]
		ms := strings.Split(q, "&")
		for _, m := range ms {
			kv := strings.Split(m, "=")
			k, v := kv[0], kv[1]
			r.query[k] = v
		}
	} else {
		r.path = query
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

func handleClient(conn net.Conn) {
	request := make([]byte, 1024)
	defer conn.Close()
	num, err := conn.Read(request)
	checkError(err)
	raw := string(request[:num])
	r := Request{}
	r.Init(raw)
	path := r.path
	fmt.Println("path", path)
	s := responseForPath(path, r)
	conn.Write(s)
	request = make([]byte, 1024)
}

func responseForPath(path string, r Request) []byte {
	m := map[string]func(Request) []byte{
		"/":     index,
		"/doge": doge,
		"/img": responseImage,
	}
	fn, ok := m[path]
	var s []byte
	if !ok {
		s = errorResponse(404)
	} else {
		s = fn(r)
	}
	return s
}

func image(path string) []byte {
	buf, err := ioutil.ReadFile(path)
	if err == nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	return buf
}

func responseImage(r Request) []byte {
	p := r.query["path"]
	b := image(p)
	s := []byte("HTTP/1.1 200 OK\r\nContent-Type: image/gif\r\n\r\n")
	if b == nil {
		s = errorResponse(404)
		return s
	}
	m := append([]byte(s), b...)
	return m
}

func index(r Request) []byte {
	s := "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\n\r\nHello World!"
	return []byte(s)
}

func errorResponse(code int) []byte {
	s := "HTTP/1.1 404 NOT FOUND\r\nContent-Type: text/html\r\n\r\n404 NOT FOUND!"
	return []byte(s)
}

func doge(r Request) []byte {
	s := "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\n\r\nHello Doge!<img src=\"img?path=doge0.gif\"><img src=\"img?path=doge1.gif\">"
	return []byte(s)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
