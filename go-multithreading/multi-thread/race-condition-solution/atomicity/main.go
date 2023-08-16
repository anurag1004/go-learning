package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go foo(1)
	go foo(2)
	wg.Wait()
	fmt.Printf("Final Counter :%d\n", counter)
}
func foo(id int) {
	fmt.Printf("TID: %d, waiting to increment atomic counter, current:%d, [%s]\n", id, atomic.LoadInt64(&counter), time.Now().Format("2006-01-02 15:04:05"))
	// cs start
	time.Sleep(10 * time.Second)
	atomic.AddInt64(&counter, 1)
	// cs end
	fmt.Printf("TID: %d, incremented the atomic counter, new:%d, [%s]\n", id, atomic.LoadInt64(&counter), time.Now().Format("2006-01-02 15:04:05"))
	wg.Done()
}

/*
	Observation:
	If I swap line 23 & 24 and run the code, the output will look like this..

		TID: 1, waiting to increment atomic counter, current:0, [2023-08-16 20:05:50]
		TID: 2, waiting to increment atomic counter, current:0, [2023-08-16 20:05:50]
	👉  TID: 2, incremented the atomic counter, new:2, [2023-08-16 20:06:00]
	👉  TID: 1, incremented the atomic counter, new:2, [2023-08-16 20:06:00]
		Final Counter :2
	Notice 👉:
		we are seeing new counter value as 2, for both the TIDs. Why?
		First u need to see atomic variable to be a single entity or isolated entity.
		Whatever we write or load, we do atomically meaning if one operation we want to perform
		say "read" then other threads need to wait for that operation to complete. Similar
		for write operation.
		Now comin back to our strange output.
			.
			.
			// cs start
				atomic.AddInt64(&counter, 1)
				time.Sleep(10 * time.Second)
			// cs end
			.
			.
		so suppose TID1 enters, if it sees no thread is currently doing any operation into this var
		it'll go ahead with increment operation. Meanwhile TID2 comes and see some operation is being
		performed by other thread, so it waits for the increment operation to completes.
		As soon it completes, TID1 goes to sleep for 10secs and TID2 aquire the counter and starts
		incrementing operation then TID2 goes to sleep for 10secs. Now this all process happened very fast.
		Both the threads are asleep. after 10sec timer finishes, they both are instructed to print counter value
		At this point of time, the counter value is 2 for both the threads (although TID1 prev saw counter to be 1)
		Thats why we saw output to be same for TID1 and 2

*/
