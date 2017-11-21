package gitea

import "net/http"
import "time"

// Client represents an HTTP client to communicate with the Gitea server's API with.
type Client struct {
	token string
	client *http.Client
	baseUrl string
}

func NewClient(token string, baseUrl string) (*Client) {
	return &Client {
		token: token,
		client: &http.Client {
			Timeout: time.Second * 10,
		},
		baseUrl: baseUrl,
	}
}
