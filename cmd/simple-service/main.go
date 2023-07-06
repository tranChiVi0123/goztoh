package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	greetingHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Greet, buddie!\n")
	}

	http.HandleFunc("/greeting", greetingHandler)

	log.Println("Listening for request at http://localhost:2327/greeting")
	log.Fatal(http.ListenAndServe(":2327", nil))
}
