package http

import (
	"io/ioutil"
	"fmt"
	"os"
	"path"
)

func Image(name string) []byte {
	name = path.Join("static", "img", name)
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	return buf
}

func Template(name string) []byte {
	name = path.Join("templates", name)
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	return buf
}
