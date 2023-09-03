package main

import (
	"fmt"
	"time"
)

/*
Imp things about channels
A send to a nil channel blocks forever
A receive from a nil channel blocks forever
A send to a closed channel panics
A receive from a closed channel returns the zero value immediately

what is a select?
just like a switch statement BUT for channels

	-> The select statement lets a goroutine wait on multiple communication operations.
	-> A select blocks until one of its cases can run, then it executes that case.
	   ✅ It chooses one at random if multiple are ready.
*/
func main() {
	msgChan := make(chan string)
	go func() {
		/*
			A select statement without a default case is blocking until a read or write
			in at least one of the case statements can be executed
			That's why ive written it in a seperate go routine

		*/
		select {
		case msg := <-msgChan:
			fmt.Println("recvd data from msgChan:" + msg)
		default:
			/*
				Basic sends and receives on channels(unbuffered) are blocking.
				However, we can use select with a default clause to implement non-blocking sends,
				receives, and even non-blocking multi-way selects.
			*/
			fmt.Println("no msg recvd!")
		}
	}()
	msgChan <- "hellow"
	// go func() {
	// 	for { // run this select in continous for loop because select statement as is itself excuted only once
	// 		select {
	// 		case msgChan <- "a msg from 🌕": // for this statement to run, since msgChan is unbuffered,
	// 			// a receiver of this channel should exist,
	// 			// otherwise this send operation will get blocked and the default statement will excute
	// 			fmt.Println("msg sent ✅")
	// 		default:
	// 			fmt.Println("send ops blocked ❌")
	// 		}
	// 		time.Sleep(50 * time.Millisecond) // to prevent excessive default satements
	// 	}
	// }()
	// time.Sleep(100 * time.Millisecond) // wait for some time.. to see default in action
	// fmt.Println("recvd: " + <-msgChan) // rcv the msg

	// last example.. multiple send/recv, comment the above code from line 44-58
	go func() {
		for {
			select {
			case msgChan <- "a msg from 🌕":
				fmt.Println("msg sent ✅")
			case newMsg := <-msgChan:
				fmt.Println("recvd data from msgChan:" + newMsg)
			default:
				fmt.Println(".") // no comms or both the channels are blocked
			}
			time.Sleep(50 * time.Millisecond)
		}
	}()
	msgChan <- "1st msg 🫡" // send msg
	time.Sleep(100 * time.Millisecond)
	fmt.Println("recvd: " + <-msgChan) // rcv the msg

	msgChan <- "2st msg 🫡"
	time.Sleep(200 * time.Millisecond)
	fmt.Println("recvd: " + <-msgChan) // rcv the msg

}
