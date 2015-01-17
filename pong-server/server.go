package main

import (
	"fmt"
	"strings"
	"flag"
	"errors"
	
	"net"
	"net/http"
	
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	addresses, _ := net.LookupHost(hostname)
	fmt.Fprintf(w, "pong / %s / %s", hostname, strings.Join(addresses, ","))
}

func serve(port int) error {
	http.HandleFunc("/", handler)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func run() error {
	var portVar = flag.Int("port", 8000, "pong server port")
	flag.Parse()
	
	port := *portVar
	if port < 1 || port > 65535 {
		return errors.New("server port should be between 1 and 65535")
	}
	
	return serve(port)
}

func main() {
    if err := run(); err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}