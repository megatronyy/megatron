package main

import (
	"net/http"
	"fmt"
)

func chain(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("first invoke!")

		handler.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/echo", chain(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("second invoke!")
	})))

	http.ListenAndServe("0.0.0.0:9001", mux)
}
