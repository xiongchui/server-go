package http

import (
    "io/ioutil"
    "fmt"
    "os"
    "path"
)

func Image(name string) ([]byte, error) {
    name = path.Join("static", "img", name)
    buf, err := ioutil.ReadFile(name)
    if err != nil {
        fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
    }
    return buf, err
}

func Template(name string) ([]byte, error) {
    name = path.Join("templates", name)
    buf, err := ioutil.ReadFile(name)
    if err != nil {
        fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
    }
    return buf, err
}

// todo, 错误码和错误页面对应，消除硬编码
func ErrorPageByCode(code string) ([]byte, error) {
    name := path.Join("templates", code) + ".html"
    buf, err := ioutil.ReadFile(name)
    if err != nil {
        fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
    }
    return buf, err
}
