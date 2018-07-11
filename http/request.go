package http

import "strings"

type Request struct {
	Raw     string
	Method  string
	Path    string
	Query   map[string]string
	Body    string
	Headers map[string]string
	Cookies map[string]string
}

func (r *Request) Init(raw string) {
	r.Raw = raw
	f := strings.Split(raw, "\r\n")[0]
	es := strings.Split(f, " ")
	method, path := es[0], es[1]
	r.Method = method
	r.Body = strings.Split(raw, "\r\n\r\n")[1]
	r.Headers = make(map[string]string)
	r.Query = make(map[string]string)
	r.AddCookies()
	r.AddHeaders()
	r.ParseQuery(path)
}

func (r *Request) AddCookies() {
	e := r.Headers["Cookie"]
	kvs := strings.Split(e, ": ")
	for _, s := range kvs {
		if strings.Contains(s, "=") {
			kv := strings.Split(s, "=")
			k, v := kv[0], kv[1]
			r.Cookies[k] = v
		}
	}
}

func (r *Request) AddHeaders() {
	raw := r.Raw
	s := strings.Split(raw, "\r\n\r\n")[0]
	hs := strings.Split(s, "\r\n")[1:]
	for _, s := range hs {
		kv := strings.Split(s, ": ")
		k, v := kv[0], kv[1]
		r.Headers[k] = v
	}
}

func (r *Request) ParseQuery(query string) {
	if strings.Contains(query, "?") {
		es := strings.Split(query, "?")
		r.Path = es[0]
		q := es[1]
		ms := strings.Split(q, "&")
		for _, m := range ms {
			kv := strings.Split(m, "=")
			k, v := kv[0], kv[1]
			r.Query[k] = v
		}
	} else {
		r.Path = query
	}
}
