package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	channel := make(chan string)
	links := []string{"https://www.google.com", "https://www.facebook.com", "https://www.youtube.com", "https://www.wikipedia.com"}

	checkStatus(links, channel)
	for ch := range channel {
		// never reference same var to different child routines
		// if u dont provide link and refer to ch
		// since ch is changing frequently, by that time inside getStatus child routine, it is also changing
		// since both are pointing to same variable
		go func(link string) {
			time.Sleep(3 * time.Second)
			getStatus(link, channel)
		}(ch)
	}
}

func checkStatus(links []string, channel chan string) {
	for i := 0; i < len(links); i++ {
		go getStatus(links[i], channel)
	}
}

func getStatus(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("Link: %v, is down!\n", link)
		channel <- link
		return
	}
	fmt.Printf("Link: %v, is Up!\n", link)
	channel <- link
}
