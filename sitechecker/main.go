package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	websiteName := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.reddit.com",
		"https://www.wikipedia.org",
		"https://www.linkedin.com",
		"https://www.instagram.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.quora.com",
		"https://www.bing.com",
		"https://www.pinterest.com",
		"https://www.tumblr.com",
		"https://www.spotify.com",
		"https://www.flickr.com",
		"https://www.imdb.com",
		"https://www.ask.com",
	}

	c := make(chan string)

	for _, website := range websiteName {
		go checkWebsite(website, c)
	}

	//fmt.Println(<-c)
	//for {
	//	//fmt.Println(<-c)
	//	go checkWebsite(<-c, c)
	//}

	for l := range c {
		//fmt.Println(<-c)
		//l := l
		go func(l string) {
			time.Sleep(5 * time.Second)
			go checkWebsite(l, c)
		}(l)

	}

	//for i := 0; i < len(websiteName); i++ {
	//	//fmt.Println(<-c)
	//	go checkWebsite(<-c, c)
	//}

}

func checkWebsite(link string, c chan string) {
	//time.Sleep(time.Minute)
	res, err := http.Get(string(link))
	if err != nil {
		c <- err.Error()
		fmt.Println("Error: ", err)
	} else {
		if res.StatusCode == 200 {
			//c <- "Website is up and running"

			c <- link
			fmt.Println("Website: ", link, " is up and running")
		} else {
			c <- link
			//c <- "Website is down"
			fmt.Println("Website: ", link, " is down")
		}
	}
}
