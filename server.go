package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	println("Run Server")
	r := mux.NewRouter()
	r.HandleFunc("/InboundXml", NewFunction).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":5000", nil)
}

func NewFunction(w http.ResponseWriter, r *http.Request) {
	response := "<?xml version=\"1.0\"?><Response><Pause length=\"10\"></Pause></Response>"
	println(response)
	w.Write([]byte(response))
}
