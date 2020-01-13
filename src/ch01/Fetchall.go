package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	time2 "time"
)

func main() {
	time := time2.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Printf(<-ch)
	}

	fmt.Printf("get all time:%.2f\n", time2.Since(time).Seconds())
}

func fetch(url string, ch chan<- string) {
	time := time2.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("get %s error: %v\n", url, err)
		return
	}
	len, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	ch <- fmt.Sprintf("get %s length:%d,time:%.2f second\n", url, len, time2.Since(time).Seconds())
}
