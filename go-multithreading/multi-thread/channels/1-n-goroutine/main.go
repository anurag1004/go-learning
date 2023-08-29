package main

import (
	"fmt"
	"sync"
)

/*
1-N pattern
One Producer, N consumers
Here N = 10
Notice we've a third go routine, at line 25 (

	its job is to wait for N number of done/s and calls wg.Done(),
	that'll signal main thread to proceed after wg.Wait()

)
If num of consumers spawn < N (line 22, Example: go consume(N-5, ch, done), here num of consumers are 5 )

	➡️ we'll have a deadlock, since the third go routine is waiting for 5 extra done/s to complete

if num of consumers spawn > N (line 22, Example: go consume(N+5, ch, done), here num of consumers are 15)

	➡️ It means we are spawning 5 extra consumers, though if you update the code and run it,
		it'll run without any issues.. hmmm😶‍🌫️ why?
		➡️ See what we doing in third go routine, we are not marking done channel as close.. So,
		those 5 extra consumer are able to write even after our main thread is completed
		Also we are not essentially waiting for those extra consumers to finish and even if we have waited,
		we could easily wait for those to complete because consumers are able to put data to done channel

		However if you close it, then after completion of 10 consumers, extra 5 consumers will not be able
		to write to done channel and we'll get an error: data passed to a closed channel ❌
*/
const N = 10

func main() {
	ch := make(chan int)
	done := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)
	go producer(ch)
	go consume(N, ch, done)
	go func() {
		for i := 0; i < N; i++ {
			<-done
		}
		close(done)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("DONE")
}

func producer(ch chan int) {
	for i := 0; i < 20; i++ {
		ch <- i
	}
	close(ch)
}
func consume(n int, ch chan int, done chan bool) {
	for i := 0; i < n; i++ {
		go consumer(i, ch, done)
	}
}
func consumer(id int, ch chan int, done chan bool) {
	for data := range ch {
		fmt.Printf("CONSUMER id:%d, data:%d\n", id, data)
	}
	done <- true
}
