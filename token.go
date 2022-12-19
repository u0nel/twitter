package twitter

import (
	"encoding/json"
	"errors"
	"net/http"
)

func fetchToken() (string, error) {
	h := http.Header{}
	//	h.Set("accept", "application/json,text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	//	h.Set("accept-encoding", "gzip")
	//	h.Set("accept-language", "en-US,en;q=0.5")
	//	h.Set("connection", "keep-alive")
	h.Set("authorization", auth)

	req, _ := http.NewRequest("POST", activate, nil)
	req.Header = h

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", errors.New("Error doing request: " + err.Error())
	}

	var v struct {
		GuestToken string `json:"guest_token"`
	}
	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		return "", errors.New("Error decoding json: " + err.Error())
	}
	return v.GuestToken, nil
}
