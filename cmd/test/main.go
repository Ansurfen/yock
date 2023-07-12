package main

import (
	"flag"
	"fmt"
	"net/http"
)

var port = flag.Int("p", 0, "")

func main() {
	flag.Parse()
	if *port == 0 {
		panic("invalid port")
	}
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
