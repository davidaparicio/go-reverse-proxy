package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/http2"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
	CERT_CRT  = "cert.pem"
	CERT_KEY  = "key.pem"
	PROXY_URL = "https://www.google.com/"
)

func init() {
	// https://youtu.be/tWSmUsYLiE4
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func main() {
	//demoURL, err := url.Parse("http://neverssl.com")
	demoURL, err := url.Parse(PROXY_URL)
	if err != nil {
		log.Fatal(err)
	}
	//proxy := httputil.NewSingleHostReverseProxy(demoURL)
	proxy := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		req.Host = demoURL.Host
		req.URL.Host = demoURL.Host
		req.URL.Scheme = demoURL.Scheme
		req.RequestURI = ""
		s, _, _ := net.SplitHostPort(req.RemoteAddr) //8m
		req.Header.Add("X-Forward-For", s)

		http2.ConfigureTransport(http.DefaultTransport.(*http.Transport))

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(rw, err)
			return
		}
		// To maintain the correct Content-Type/Length etc..
		for key, values := range resp.Header {
			for _, value := range values {
				rw.Header().Set(key, value)
			}
		}

		// To handle data streams
		ticker := time.NewTicker(10 * time.Millisecond)
		done := make(chan bool)
		go func() {
			for {
				select {
				// Use Ticker instead of Tick, to avoid goroutine leak
				// More info on the issue : https://github.com/golang/go/issues/37144
				case <-ticker.C: // select case must be send or receive
					rw.(http.Flusher).Flush()
				case <-done:
					return
				}
			}
		}()

		// To handle trailer header
		trailerKeys := []string{}
		for key := range resp.Trailer {
			trailerKeys = append(trailerKeys, key)
		}

		rw.Header().Set("Trailer", strings.Join(trailerKeys, ","))

		rw.WriteHeader(resp.StatusCode)
		io.Copy(rw, resp.Body)

		for key, values := range resp.Trailer {
			for _, value := range values {
				rw.Header().Set(key, value)
			}
		}
		ticker.Stop()
		close(done) //end the go routine
	})
	// HTTP 1
	//http.ListenAndServe(":8080", proxy)
	// HTTP 2
	log.Println("Proxy starting to listen on " + CONN_HOST + ":" + CONN_PORT + ", ready to forward the demoURL " + PROXY_URL)
	errServ := http.ListenAndServeTLS(CONN_HOST+":"+CONN_PORT, CERT_CRT, CERT_KEY, proxy)
	if errServ != nil {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", errServ.Error())
	}
	// Example with context https://pkg.go.dev/net/http#Server.Shutdown
}
