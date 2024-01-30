package main

import (
	"net/http"

	"github.com/altbankdesafio/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Welcome API MATH - Alt Bank Challenge - </h1>"))
	})

	http.HandleFunc("/prime", handlers.PrimeHandler)

	http.ListenAndServe(":8080", nil)
}
