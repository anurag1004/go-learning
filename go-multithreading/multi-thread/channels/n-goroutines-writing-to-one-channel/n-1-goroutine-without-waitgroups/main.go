package main

import "fmt"

func main() {
	N := 4
	ch := make(chan int)
	done := make(chan bool)
	go foo(1, ch, done)
	go foo(2, ch, done)
	go foo(3, ch, done)
	go foo(4, ch, done)

	go func() {
		// instead of writing n number of <-done
		// we can have a variable N
		// <-done
		// <-done
		// <-done
		for i := 0; i < N; i++ { // its a semaphore pattern
			// notice we actually dont know which goroutine has finished
			// this for loop  means, we'll wait for any N goroutines to finish their operation
			// after that its done, we'll not wait for any routines
			// and close the channel (where data was being passed by the goroutines)
			// look how we are controlling the access for the func foo()
			// if there are more than N goroutines, then after closing the channel if they'll try to put
			// data back to channel, they will be essentially putting data to a closed channel,
			// meaning- they'll not be able to put or pass data into the channel
			/*
				 if N=No of goroutines spawned ✅
				 N<No of goroutines spawned ❌ (error: data passed to a closed channel)
				 N>No of goroutines spawned ❌ (error: all goroutines are asleep - deadlock)
				 Explanation for the last one-
				 	after spawning routines less then N, we'll wait for N routines to finish (this is what we wrote in the for loop)
					say no of goroutines spawned is M
					Here, M < N
					meaning we are essentially waiting for (N-M) extra goroutines to finish, but we've only spawn M routines,
					(which is less then N).
					So at last our for loop will be blocked forever (since there are no goroutines to say done<-true)
					Meaning the close operation on channel will never happen
					Comming to our main goroutine, the range loop will also get blocked because ch was never closed
					So essentially one of our go routine is blocked as well as our main goroutine is also blocked
					= DEADLOCK
			*/
			<-done
		}
		close(ch)

	}()
	for data := range ch {
		fmt.Println(data)
	}
	fmt.Println("DONE")
}
func foo(id int, ch chan int, done chan bool) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Thread %d, ch<-%d\n", id, i)
		ch <- i
	}
	done <- true
}
