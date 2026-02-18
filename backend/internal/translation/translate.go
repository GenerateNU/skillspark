package translations

import (
	"net/http"
)

type TranslateClient struct {
	Client *http.Client
}

func NewClient(client *http.Client) *TranslateClient {

	return &TranslateClient{
		Client: client,
	}
}
