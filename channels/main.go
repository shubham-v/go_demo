package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for {
	// 	fmt.Println(<-c)
	// }

	// for {
	// 	go checkLink(<-c, c)
	// }

	// for l := range c {
	// 	// time.Sleep(5 * time.Second)
	// 	go checkLink(l, c)
	// }

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c) // l is global reference value set by chanel c
		}(l) // l is an argument passed to funtion literal (annonymous function)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// fmt.Println(<-c) // receiving a message from channel is blocking
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// fmt.Println(<-c) // it hangs here as no message is received by channel

}

func checkLink(link string, c chan string) {
	time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if nil != err {
		fmt.Println(link, " might be down!")
		// c <- "Might be down"
		c <- link
		return
	}
	fmt.Println(link, " is up!")
	// c <- "It'up"
	c <- link
}

// func checkLink(link string, c chan string) {
// 	_, err := http.Get(link)
// 	if nil != err {
// 		c <- link + " is down"
// 		return
// 	}
// 	c <- link + " is up"
// 	time.Sleep(5 * time.Second)
// 	checkLink(link, c)
// }
