package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Hello world"))
	}).Methods(http.MethodGet)

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		panic(err)
	}
}
