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
