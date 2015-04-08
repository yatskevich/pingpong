package main

import (
	"fmt"
	"flag"
	"errors"
	"log"
	
	"net/http"
	"io/ioutil"
	
	"os"
	
	"github.com/benschw/dns-clb-go/dns"
	"github.com/benschw/dns-clb-go/clb"
)

type HandlerConfig struct {
	dnsServer dns.Lookup
	pongServiceDomain string
}

func handler(config HandlerConfig) func(http.ResponseWriter, *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		// NOTE: FOR DEMO PURPOSES ONLY. DO NOT USE THIS CODE IN PRODUCTION :)
		//
		// It seems that Client Load Balancer (CLB) doesn't handle errors correctly, so 
		// after the first error internal state will be broken. 
		// This is the reason to create CLB on every request.
		// 
		dnsResolver := clb.NewRandomClb(config.dnsServer)
		address, err := dnsResolver.GetAddress(config.pongServiceDomain)
		if err != nil {
			log.Printf("failed to resolve pong service domain: %v\n", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		
		response, err := http.Get(fmt.Sprintf("http://%s", address.String()))
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

func serve(port int, config HandlerConfig) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	
	http.HandleFunc("/ping", handler(config)) 
	
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
		
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func run() error {
	var portVar = flag.Int("port", 8888, "ping client port")
	var dnsVar = flag.String("dns", "172.17.42.1:53", "dns server used to resolve pong-service name")
	var pongServiceVar = flag.String("pong-service", "pong-server.service.consul", "pong service domain")

	flag.Parse()
	
	port := *portVar
	if port < 1 || port > 65535 {
		return errors.New("server port should be between 1 and 65535")
	}
	
	dnsServer := dns.NewLookupLib(*dnsVar)
	
	pongServiceDomain := *pongServiceVar
	return serve(port, HandlerConfig{ dnsServer, pongServiceDomain })
}

func main() {
    if err := run(); err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}