package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello. Web application version 1.0"))
	})

	err := http.ListenAndServe(":4001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
