package main

import (
	"net/http"
	"zarbat_test/internal/logging"

	"github.com/gorilla/mux"
)

func main() {
	logging.Debug.Println("Zarbat Router running")
	router := mux.NewRouter()
	router.HandleFunc("/Zarbat", ZarbatTest).Methods("POST", "GET")
	http.Handle("/", router)
	routerPort := ":5010"
	http.ListenAndServe(routerPort, nil)
}

func ZarbatTest(w http.ResponseWriter, r *http.Request) {
	s := "Hello World !!! - It Worked!!!"
	logging.Debug.Println(s)
	w.Write([]byte(s))
}
