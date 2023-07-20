package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Printf("Encountered with an error while making a request: %v\n", err)
		os.Exit(-1)
	}
	// defer resp.Body.Close()
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	// u can only use the stream only once.. onnce its read u cannot read it again
	// io.Copy(os.Stdout, resp.Body)

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		fmt.Printf("Encountered with an error while copying response body: %v\n", err)
		return
	}
	fmt.Println(buf.String())
}
