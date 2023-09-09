package main

import (
	"fmt"
	"time"
)

func main() {
	timeChan := timer(200)
	for {
		select {
		case <-timeChan:
			fmt.Println("timeChan returned!")
			return
		default:
			fmt.Println(".")
		}
		time.Sleep(20 * time.Millisecond)
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
