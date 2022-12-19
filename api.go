package twitter

import (
	"net/http"
)

func headers(token string) (h http.Header) {
	h = make(http.Header)
	h.Add("x-guest-token", token)
	h.Add("authorization", auth)
	return
}
