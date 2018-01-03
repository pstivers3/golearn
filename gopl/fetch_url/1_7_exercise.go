// prints the content found at a URL
// use io.Copy instead of ioutil.ReadAll
// usage:
// $ go run filename.go <url>, for example http://gopl.io
// or
// $ go build filename.go
// $ ./filename <url> 

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err :=http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetc: %v\n", err)
			os.Exit(1)
		}
		_, err2 := io.Copy(os.Stdout, resp.Body) // gives "no new variable" error if use err as var name
		if err2 != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

