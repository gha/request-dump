package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	httpListenAddr = flag.String("httpListenAddr", ":80", "HTTP Listen Address")
)

func main() {
	flag.Parse()

	log.Printf("request-dump: starting server on %s", *httpListenAddr)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(*httpListenAddr, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("------------------------------------------")
	log.Printf("New request at %s", time.Now().Format(time.UnixDate))
	log.Print("------------------------------------------")
	log.Printf("Protocol: %s", r.Proto)
	log.Printf("Method: %s", r.Method)
	log.Printf("Host: %s", r.Host)
	log.Printf("Request URI: %s", r.RequestURI)
	log.Printf("Content-Length: %d", r.ContentLength)
	log.Printf("Remote Address: %s", r.RemoteAddr)
	log.Print("")

	// Headers
	log.Printf("Headers (%d):", len(r.Header))
	for k, v := range r.Header {
		log.Printf(" > %s: %s", k, v[0])
	}
	log.Print("")

	// Post Data
	if err := r.ParseForm(); err != nil {
		log.Printf("Post data: Error parsing form data: %v", err)
	} else {
		log.Printf("Post data (%d):", len(r.PostForm))
		for k, v := range r.PostForm {
			log.Printf(" > %s: %s", k, v[0])
		}
	}
	log.Print("")

	// Request Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print("Body: Error reading body: %v", err)
	} else {
		log.Print("Body:")
		log.Print(string(body))
	}
	log.Print("------------------------------------------")
	log.Print("")
	log.Print("")
	log.Print("")
}
