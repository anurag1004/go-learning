package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"text/template"
	"time"
)

func handlHttpReq(lsnr net.Listener) {
	for {
		conn, err := lsnr.Accept()
		if err != nil {
			fmt.Printf("❌ error occured while accepting a request: %v\nClosing http handler...", err.Error())
			break
		}

		go func(conn net.Conn) {
			defer func() {
				fmt.Printf("❌ closing connection for %s\n", conn.RemoteAddr().String())
				if err := conn.Close(); err != nil {
					fmt.Printf("🛑 unable to close the connection for %s : %s\n", conn.RemoteAddr().String(), err.Error())
				} else {
					fmt.Printf("🛑 closed the connection for %s\n", conn.RemoteAddr().String())
				}
			}()
			handleReqHelper(conn)
		}(conn)
	}
}
func handleReqHelper(conn net.Conn) {
	fmt.Printf("✅ connection accepted from %s\n", conn.RemoteAddr().String())
	reqHeaders := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanLines)
	isHttpReqLineProcessed := false
	var method, path, httpVer string
	var reqBody bytes.Buffer
	for scanner.Scan() {
		txt := scanner.Text()
		if !isHttpReqLineProcessed {
			// HTTP_METHOD PATH HTTP_VERSION
			split := strings.Split(txt, " ")
			method = split[0]
			path = split[1]
			httpVer = split[2]
			isHttpReqLineProcessed = true
		} else {
			/*
				The request headers and the request body are typically separated by an empty line
				(CRLF, which stands for carriage return and line feed, represented as "\r\n")
				Since we are already reading by line by line.. here after analysis body is seperated from headers with empty string
			*/
			if strings.Compare(txt, "") == 0 {
				// CRLF
				// process the body if exist
				if strings.Compare(method, "POST") == 0 || strings.Compare(method, "PUT") == 0 {
					timeOutChan := timer(100) // 100ms
					go func() {
						for scanner.Scan() {
							bodyStr := strings.TrimSpace(scanner.Text())
							reqBody.WriteString(bodyStr)
						}
					}()
					if <-timeOutChan {
						// add last curly braces
						reqBody.WriteString("}")
					}
				}
				mux(conn, reqHeaders, method, path, httpVer, reqBody.Bytes())
				break // breakout from scanner
			} else {
				// headers
				split := strings.Split(txt, ":")
				key := strings.TrimSpace(split[0])
				value := ""
				for i := 1; i < len(split); i++ {
					value += split[i]
				}
				reqHeaders[key] = strings.TrimSpace(value)
			}

		}
	}
}
func timer(ms int64) <-chan bool {
	out := make(chan bool)
	go func() {
		start := time.Now()
		for {
			elapsed := time.Now().Sub(start)
			if elapsed.Milliseconds() >= ms {
				out <- true
				return
			}
		}
	}()
	return out
}
func mux(conn net.Conn, reqHeaders map[string]string, method string, path string, httpVer string, body []byte) {
	var response bytes.Buffer
	switch method {
	case "GET":
		// handle GET request
		htmlDocBytes, _ := os.ReadFile("./htmls/get.html")
		contentLength := len(htmlDocBytes)
		contentLengthLine := fmt.Sprintf("Content-Length: %d\n", contentLength)
		constructBasicResponseHeader(&response, contentLengthLine, "200", httpVer, "OK")
		response.Write(htmlDocBytes)

		conn.Write(response.Bytes())
	case "POST":
		// handle POST request
		var jsonObject interface{}
		err := json.Unmarshal(body, &jsonObject)
		if err != nil {
			fmt.Println("unable to unmarshal json string")
			constructBasicResponseHeader(&response, "Content-Length: 0", "500", httpVer, "Internal Server Error") // internal server error
			conn.Write(response.Bytes())
			return
		}
		if jsonMap, ok := jsonObject.(map[string]interface{}); !ok {
			fmt.Println("unable to create object map from json string..")
			constructBasicResponseHeader(&response, "Content-Length: 0", "500", httpVer, "Internal Server Error") // internal server error
			conn.Write(response.Bytes())
			return
		} else {
			//json map
			tmpl, _ := template.ParseFiles("./htmls/post.gohtml")
			file, _ := os.Create("./out/parsed_post.html")
			defer func() {
				file.Close()
			}()
			_ = tmpl.ExecuteTemplate(file, "post.gohtml", jsonMap)
			htmlDocBytes, _ := os.ReadFile("./out/parsed_post.html")
			contentLength := len(htmlDocBytes)
			contentLengthLine := fmt.Sprintf("Content-Length: %d\n", contentLength)
			constructBasicResponseHeader(&response, contentLengthLine, "200", httpVer, "OK") // OK
			response.Write(htmlDocBytes)
			conn.Write(response.Bytes())
		}
	default:
		// handle unknown request method
		return
	}

}
func constructBasicResponseHeader(response *bytes.Buffer, contentLengthLine, status, httpVer, statusDes string) {
	resLine := fmt.Sprintf("%s %s %s\n", httpVer, status, statusDes)
	dateLine := fmt.Sprintf("Date: %s\n", time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT"))
	serverLine := "Server: GoServer\n"
	acceptRangeLine := "Accept-Ranges: bytes\n"
	contentTypeLine := "Content-Type: text/html\n"

	response.WriteString(resLine)
	response.WriteString(dateLine)
	response.WriteString(serverLine)
	response.WriteString(acceptRangeLine)
	response.WriteString(contentTypeLine)
	response.WriteString(contentLengthLine)
	response.WriteString("\r\n")
}
func main() {
	if lsnr, err := net.Listen("tcp4", "localhost:3000"); err != nil {
		fmt.Printf("❌ error while creating tcp server : %s\n", err.Error())
	} else {
		fmt.Println("🤖 TCP server started at : " + lsnr.Addr().String())
		handlHttpReq(lsnr)
	}
}
