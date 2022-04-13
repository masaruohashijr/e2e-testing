package helper

import "net/http"

func EnsureCors(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Request-Headers", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, PUT, POST, GET, OPTIONS, PATCH")
	w.Header().Add("Content-Type", "application/json,text/plain")
	return w
}
