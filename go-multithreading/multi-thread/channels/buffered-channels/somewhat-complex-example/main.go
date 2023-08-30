package main

import (
	"fmt"
	"sync"
	"time"
)

type Request struct {
	num int
	msg string
}

var reqChan chan *Request
var resChan chan *Request
var wg sync.WaitGroup
var sema chan int

const maxOutStanding = 3

func init() {
	sema = make(chan int, maxOutStanding)
	reqChan = make(chan *Request)
	resChan = make(chan *Request)
}
func requestEmitter(numOfReqs int) {
	for i := 0; i < numOfReqs; i++ {
		reqChan <- &Request{num: i, msg: "client request"}
	}
	close(reqChan)
}
func responseGrabber() {
	for res := range resChan {
		fmt.Printf("RES: %+v\n", *res)
		wg.Done() //mark the end of each req's response
	}
	close(resChan)
}
func sendResponse(res *Request, wg *sync.WaitGroup) {
	fmt.Printf("sendResponse: res:%+v [%s]\n", *res, time.Now().Format("2006-01-02 15:04:05"))
	res.msg = "Server response"
	resChan <- res
}
func serve(wg *sync.WaitGroup) {
	for req := range reqChan {
		// all the handle routines are sharing same req object, its better to create a copy of it
		req := req
		go handle(req, wg) // handle the req in seperate go routine
	}

}
func handle(newReq *Request, wg *sync.WaitGroup) { // at max 5 reqs will be handled parallely
	fmt.Printf("handle: req:%+v Waiting...[%s]\n", *newReq, time.Now().Format("2006-01-02 15:04:05"))

	sema <- 1
	fmt.Printf("handle: req:%+v Inside...[%s]\n", *newReq, time.Now().Format("2006-01-02 15:04:05"))
	val := process(newReq.num)
	newReq.num = val
	go sendResponse(newReq, wg) // send response using seperate go routine
	<-sema

	fmt.Printf("handle: req:%+v COMPLETED...[%s]\n", *newReq, time.Now().Format("2006-01-02 15:04:05"))
}
func process(num int) int {
	// sleep, for simulating some time taking task
	time.Sleep(1 * time.Second)
	return num * num
}
func main() {
	numOfReqs := 5
	wg.Add(numOfReqs)
	go requestEmitter(numOfReqs)
	go responseGrabber()
	go serve(&wg)
	wg.Wait()
}
