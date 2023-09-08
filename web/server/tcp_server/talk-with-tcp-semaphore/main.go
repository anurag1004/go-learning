package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
)

const min = 100001
const max = 999999

type Server struct {
	MaxConn int
	Id      int
	sema    chan int
}

func (srv *Server) Start(host string) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	ln, err := net.Listen("tcp4", host)
	if err != nil {
		panic(err)
	}
	log.Printf("TCP server - %d started at : %s\n", srv.Id, host)
	srv.handle(ln)
}
func (srv *Server) handle(listner net.Listener) {
	if srv.sema == nil {
		fmt.Println("sema is nil")
		srv.sema = make(chan int, srv.MaxConn)
	}
	for {
		conn, err := listner.Accept()
		if err != nil {
			panic(err)
		}
		srv.sema <- 1 // aquire semaphore
		reqId := rand.Intn(max-min+1) + min
		go func(conn net.Conn, reqId int) {
			defer func() {
				err := conn.Close()
				if err != nil {
					panic(err)
				}
				log.Printf("connection closed with client: %v", reqId)
			}()
			srv.handleHelper(conn, reqId)
			<-srv.sema // release
		}(conn, reqId)
	}
}
func (srv *Server) handleHelper(conn net.Conn, reqId int) {
	log.Printf("connection established with client: %v", reqId)
	fmt.Fprintf(conn, "connection accepted!")
	str := ""
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		inputBhai := scanner.Text()
		if strings.Compare(inputBhai, "exit") == 0 {
			fmt.Fprintf(conn, "Bye bhai...")
			break
		}
		fmt.Fprintf(conn, "Mil gya bhai...%s\n", str)
		str += " dobara"
	}
}
func CreateNewServer(maxConn, id int) *Server {
	return &Server{
		MaxConn: maxConn,
		Id:      id,
		sema:    make(chan int, maxConn),
	}
}
func main() {
	srv := CreateNewServer(5, 1)
	srv.Start("localhost:5000")
}
