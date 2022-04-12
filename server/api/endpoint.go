package api

import (
	"net/http"
	"zarbat_test/internal/logging"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/intern/run", RunSingleTestHandler).Methods("POST", "OPTIONS")
	// ZarbatWebTester 4200
	// ZarbatData 5002
	// ZarbatTester 5003
	// Api CPaaS 5004
	addr := ":5003"
	http.Handle("/intern/", router)
	println("Zarbat Tester API Server")
	logging.Debug.Println("Zarbat Tester API Server")
	http.ListenAndServe(addr, router)
}
