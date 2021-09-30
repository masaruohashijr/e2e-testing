package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	println("Zarbat Router running")
	router := mux.NewRouter()
	router.HandleFunc("/Zarbat", ZarbatTest).Methods("POST", "GET")
	http.Handle("/", router)
	routerPort := ":5000"
	http.ListenAndServe(routerPort, nil)
}

func ZarbatTest(w http.ResponseWriter, r *http.Request) {
	s := "Hello World !!! - It Worked!!!"
	println(s)
	w.Write([]byte(s))
}
