package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	var port = flag.String("port", ":8080", "Server's port.")
	http.HandleFunc("/", handler)

	log.Println("Starting web server on", *port)
	if err := http.ListenAndServe(*port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
