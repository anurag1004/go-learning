package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	if ln, err := net.Listen("tcp4", ":5000"); err != nil {
		log.Println(err)
	} else {
		for {
			conn, err := ln.Accept() // blocking operation, can handle one conn at a time
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("🫡")
			time.Sleep(10 * time.Second)
			conn.Close()
		}
	}

}
