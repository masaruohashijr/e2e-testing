package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zarbat_test/internal/logging"
	"zarbat_test/server/api/helper"
)

func RunSingleTestHandler(w http.ResponseWriter, r *http.Request) {

	logging.Debug.Println("RUN SINGLE")
	var run TestRun
	_ = json.NewDecoder(r.Body).Decode(&run)
	bytes, _ := json.Marshal(run)
	logging.Debug.Println(string(bytes))
	executed := RunSingleTest(run)
	response, _ := json.Marshal(executed)
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}
