// Fetch prints the content found at a URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const httpPrefix = "http://"

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, httpPrefix) {
			url = strings.Join([]string{httpPrefix, url}, "")
		}
		resp, err := http.Get(url)
		defer resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(resp.Status)
		for {
			written, err := io.Copy(os.Stdout, resp.Body)
			if written == 0 {
				break
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: io.Copy %s: %v\n", url, err)
				os.Exit(1)
			}
		}
	}
}
