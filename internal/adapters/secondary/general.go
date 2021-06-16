package secondary

import (
	"encoding/base64"
	"net/http"
	"net/url"
)

type ZangClient struct {
	HttpClient      *http.Client
	BaseURL         *url.URL
	AccountSid      string
	AuthToken       string
	IsAuthenticated bool
}

func EncodeToBasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
