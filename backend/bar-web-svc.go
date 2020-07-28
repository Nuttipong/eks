package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Say hello from Bar service"))
	})

	err := http.ListenAndServe(":5002", nil)
	if err != nil {
		log.Fatal(err)
	}
}
