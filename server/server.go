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
	println("WORKED")
	s := "WORKED"
	w.Write([]byte(s))
}

type Number struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
}
