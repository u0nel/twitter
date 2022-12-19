package twitter

type client struct {
	token string
}

func NewClient() (c client, err error) {
	c.token, err = fetchToken()
	return
}
