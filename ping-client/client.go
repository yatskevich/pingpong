package main

import (
	"fmt"
	"flag"
	"errors"
	"log"
	
	"net/url"
	"net/http"
	"io/ioutil"
	
	"os"
)

func handler(pongServer string) func(http.ResponseWriter, *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		response, err := http.Get(pongServer)
		if err != nil {
			log.Printf("request failed with error: %v\n", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()
		
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("response parsing failed with error: %v\n", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		
		fmt.Fprint(w, string(body))
	}	
} 

func serve(port int, pongServer *url.URL) error {
	http.HandleFunc("/", handler(pongServer.String()))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func run() error {
	var portVar = flag.Int("port", 8888, "ping client port")
	var pongServerVar = flag.String("pong-server", "http://pong:8000", "pong server URL")
	flag.Parse()
	
	port := *portVar
	if port < 1 || port > 65535 {
		return errors.New("server port should be between 1 and 65535")
	}
	
	pongServerUrl, err := url.Parse(*pongServerVar)
	if err != nil {
		return err
	}
	
	return serve(port, pongServerUrl)
}

func main() {
    if err := run(); err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}