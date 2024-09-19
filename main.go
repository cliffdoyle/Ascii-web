package main

import (
	"fmt"
	"log"
	"net/http"

	backend "ascii-art/Backend"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", backend.Webhandler)
	mux.HandleFunc("/download",backend.DownloadHandler)
	mux.HandleFunc("/submit", backend.HandleRequest)
	fmt.Println("Server running on http://localhost:5000/")
	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}
