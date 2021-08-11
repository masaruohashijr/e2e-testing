package services

import (
	"net/http"
	"zarbat_test/internal/logging"
)

func MediaHandler(w http.ResponseWriter, r *http.Request) {
	logging.Debug.Println("MediaHandler")
	http.ServeFile(w, r, "media/Avaya.jpg")
}
