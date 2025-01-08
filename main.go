package main

import (
	"fmt"
	"log"
	"net/http"
)

// function to echo hello world
func echoHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello PNT Class")
}

func main() {
	// create new server
	srv := &http.Server{Addr: ":8080"}

	// assuming we want to echo hello world in the root path
	http.HandleFunc("/", echoHelloWorld)

	fmt.Println("Server starting. Head to localhost:8080")

	// run server with error handling
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("ListenAndServe erorr: %s", err)
	}
}
