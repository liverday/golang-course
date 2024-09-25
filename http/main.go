package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type LogWriter struct{}

func main() {
	r, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("An error ocurred: ", err)
		os.Exit(1)
	}

	io.Copy(&LogWriter{}, r.Body)
}

func (LogWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
