package api

import (
	"net/http"
	"zarbat_test/internal/logging"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/intern/run", RunSingleTestHandler).Methods("POST", "OPTIONS")
	addr := ":5003"
	http.Handle("/intern/", router)
	println("Zarbat Tester API Server")
	logging.Debug.Println("Zarbat Tester API Server")
	http.ListenAndServe(addr, router)
}
