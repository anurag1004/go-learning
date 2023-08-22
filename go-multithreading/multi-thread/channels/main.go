package main

import "fmt"

// simple example of using a channel (unbuffered)
func main() {
	ch := make(chan int) // unbuffered channel
	// NOTE: using an unbuffered channel is a blocking operation
	go foo(ch)
	for data := range ch { // will be receiving till the sender sending the data
		// listnening for data from the unbuffered channel is a blocking operation
		fmt.Println(data)
	}
}

func foo(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i // putting/passing data to an unbuffered channel
		// meaning- the receiving end should be ready to receive,
		// otherwise this statement will simple wait for receiving end to be ready and receive the data
	}
	close(ch) // signaling the channel that it is now closed i.e no data will be passed from now to this channel
}
