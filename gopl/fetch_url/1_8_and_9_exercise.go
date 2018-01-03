// prints the content found at a URL
// add the prefix http:// to each argument URL if it is missing
// also print the HTTP status code, found in resp.Status
// usage:
// $ go run filename.go <url>, for example http://gopl.io
// or 
// $ go build filename.go
// $ ./filename <url>

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/ryanuber/go-glob"
)

func main() {
	for _, url := range os.Args[1:] {
		if !glob.Glob("http://*", url) {
			url = "http://" + url
		}
		resp, err :=http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetc: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s%s\n\n", "HTTP status code is: ", resp.Status)
		fmt.Printf("%s", b)
	}
}

