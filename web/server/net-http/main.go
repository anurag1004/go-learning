package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type rootHandler struct{}

func (rh *rootHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Server is UP (❁´◡`❁)")
}
func sayHello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "helloooo")
}
func sayGoodbye(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "good bye")
}
func sayWhatIsSent(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		err := req.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "unable to parse the form! - %v\n", err)
		} else {
			formValueMap := req.Form
			var buff bytes.Buffer
			i := 0
			for key, val := range formValueMap {
				buff.WriteString(fmt.Sprintf("%d.) %v: %v\n", i, key, val))
				i++
			}
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Header().Set("Content-Length", strconv.Itoa(buff.Len()))
			w.WriteHeader(200)
			w.Write(buff.Bytes())
		}
	} else {
		w.WriteHeader(405)
		fmt.Fprintf(w, "INVALID METHOD")
	}
}
func getThisManAHtml(w http.ResponseWriter, req *http.Request) {
	htmlBytesDoc, err := os.ReadFile("sample.html")
	if err != nil {
		fmt.Fprintln(w, "unable to parse sample.html")
	} else {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Content-Length", strconv.Itoa(len(htmlBytesDoc)))
		w.WriteHeader(200)
		w.Write(htmlBytesDoc)
	}
}
func main() {

	http.Handle("/", &rootHandler{})
	// curl -v localhost:3000/bye
	http.HandleFunc("/bye", sayGoodbye) // get
	// curl -v localhost:3000/hello
	http.HandleFunc("/hello", sayHello) // get
	// curl -d "fname=Anurag&lname=Verma" localhost:3000/what-is-sent
	http.HandleFunc("/what-is-sent", sayWhatIsSent) // POST
	// curl -v localhost:3000/get-me-html
	http.HandleFunc("/get-me-html", getThisManAHtml)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("error starting the server at port 3000 - %v\n", err)
	}
}
