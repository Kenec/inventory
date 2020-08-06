package main

import (
	"log"
	"net/http"
)

type fooHandler struct {
	Message string
}

func(f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bar called"))
}

func main(){
	http.Handle("/", &fooHandler{Message: "Hello World"})
	http.HandleFunc("/foo", barHandler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
