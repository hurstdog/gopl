// Fetchall fetches URLs in parallel and reports their times and sizes
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	// Receive as many times as we started, no need to use the results from range
	// here.
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	defer resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	fn := urlToFilename(url)
	err = ioutil.WriteFile(fn, b, 0644)
	if err != nil {
		ch <- fmt.Sprintf("while writing %s: %v", fn, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%7d\t%s", secs, len(b), url)
}

// Returns a url with all slashes replaces with _
func urlToFilename(url string) string {
	return strings.Replace(url, "/", "_", -1)
}
