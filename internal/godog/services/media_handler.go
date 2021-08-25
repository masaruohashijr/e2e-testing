package services

import (
	"net/http"
	l "zarbat_test/internal/logging"
)

func MediaHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("MediaHandler")
	http.ServeFile(w, r, "media/Avaya.jpg")
}
