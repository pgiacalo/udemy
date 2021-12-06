package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error1:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)

	// bs := make([]byte, 99999)
	// n, err := resp.Body.Read(bs)
	// if err != nil && err != io.EOF {
	// 	fmt.Println("Error2:", n, err)
	// 	os.Exit(1)
	// }
	// fmt.Println(string(bs[:n]))
}
