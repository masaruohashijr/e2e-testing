package general

import (
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	println("******************************** Ping")
	Ch <- "Pinged"
}
