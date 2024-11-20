package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{ id int }

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{id: 1}
	io.Copy(&lw, resp.Body)
	fmt.Printf("%+v \n", lw)
}

func (*logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
