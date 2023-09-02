package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Request struct {
	num int
	msg string
	id  int
}

var reqChan chan *Request
var resChan chan *Request
var numOfReqs int
var serverInstances int
var wg sync.WaitGroup

var maxOutStanding int
var fakeProcessTime time.Duration

func init() {
	reqChan = make(chan *Request)
	resChan = make(chan *Request)
	numOfReqs = 10
	serverInstances = 1
	maxOutStanding = 5
	fakeProcessTime = 3 * time.Second
}
func requestEmitter(numOfReqs int) {
	for i := 1; i <= numOfReqs; i++ {
		reqChan <- &Request{num: i, msg: "client request", id: i}
	}
	close(reqChan)
}
func responseGrabber() {
	for res := range resChan {
		log.Printf("RES: %+v\n", *res)
		wg.Done() //mark the end of each req's response
	}
	close(resChan)
}
func responseEmitter(res *Request) {
	log.Printf("responseEmitter: resId:%v [%s]\n", (*res).id, time.Now().Format("2006-01-02 15:04:05"))
	res.msg = "Server response"
	resChan <- res
}
func serve(wg *sync.WaitGroup, serveId int) {
	sema := make(chan int, maxOutStanding)
	for req := range reqChan {
		// all the handle routines are sharing same req object, its better to create a copy of it
		req := req
		log.Printf("serve %d: reqId:%v Waiting...[%s]\n", serveId, (*req).id, time.Now().Format("2006-01-02 15:04:05"))
		/*
		 Observation:
		 earlier, we were spawning a new go routine (handle routine)for each new req
		 Even though we knew "handle" can handle max of maxOutStanding(say 5) reqs at a time
		 By doing that, after 5 requests new routines will be simply blocked (waiting for semaphore to release)
		 So if we have 1000 req at a time..then we are spawning 1000 new goroutines!
		 Which is not a good idea...
		 Instead of spawning unlimited goroutines for every reqs, we can limit the number of goroutines spawning to maxOutStanding
		*/
		sema <- 1 // acquire semaphore
		go func() {
			handle(req)
			<-sema // release semaphore after handle is done
		}()
	}

}
func handle(newReq *Request) {
	log.Printf("handle: reqId:%v Inside...[%s]\n", (*newReq).id, time.Now().Format("2006-01-02 15:04:05"))
	// blocking process
	val := process(newReq.num)
	newReq.num = val
	go responseEmitter(newReq) // send response using seperate go routine
	log.Printf("handle: reqId:%v COMPLETED...[%s]\n", (*newReq).id, time.Now().Format("2006-01-02 15:04:05"))
}
func process(num int) int {
	// sleep, for simulating some time taking task
	time.Sleep(fakeProcessTime)
	return num * num
}
func main() {
	fmt.Printf("No of reqs:%d\nNo of servers: %d\nRate-Limit for each server: %d\n", numOfReqs, serverInstances, maxOutStanding)
	start := time.Now()
	wg.Add(numOfReqs)
	go requestEmitter(numOfReqs)
	go responseGrabber()
	// spawn N (defined by serverInstances variable) go routines to handle incoming reqs
	// each go routine will handle maxOutStanding (defined by maxOutStanding variable) reqs at a time
	// so total number of reqs handled at a time will be N * maxOutStanding
	for i := 0; i < serverInstances; i++ {
		go serve(&wg, i+1) // this single go routine will handle max N (defined by maxOutStanding variable) reqs concurrently at a time
	}
	wg.Wait()
	fmt.Printf("Elapsed time: %0.5f\n", time.Since(start).Seconds())
}
