package api

import (
	"encoding/json"
	"net/http"
	"zarbat_test/internal/logging"
)

func RunSingleTestHandler(w http.ResponseWriter, r *http.Request) {

	logging.Debug.Println("RUN SINGLE")
	var run TestRun
	_ = json.NewDecoder(r.Body).Decode(&run)
	bytes, _ := json.Marshal(run)
	logging.Debug.Println(string(bytes))
	executed := RunSingleTest(run)
	response, _ := json.Marshal(executed)
	println(string(response))
	w.Write(response)
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.Header().Add("Access-Control-Allow-Methods", "POST")
	w.Header().Add("Access-Control-Allow-Methods", "OPTIONS")
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Request-Headers", "*")
}
