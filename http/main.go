package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if nil != err {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// fmt.Println(resp)

	// byteStream := make([]byte, 99999)
	// resp.Body.Read(byteStream)
	// fmt.Println(string(byteStream))

	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
