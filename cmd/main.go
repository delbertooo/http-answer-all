package main

import (
	"log"
	"net/http"
	"os"
)

var response = os.Getenv("APP_RESPONSE")
var contentType = os.Getenv("APP_CONTENT_TYPE")
var listenAddr = os.Getenv("APP_LISTEN_ADDR")

func answerAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", contentType)

	w.Write([]byte(response))
}

func main() {
	if contentType == "" {
		contentType = "text/html"
	}
	if listenAddr == "" {
		listenAddr = ":8090"
	}

	log.Printf("Answering all requests with 200 OK, Content-Type: %s and Response \"%s\"", contentType, response)
	http.HandleFunc("/", answerAll)

	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatal(err)
	}
}
