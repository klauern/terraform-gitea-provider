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

type User struct {
	ID        int    `json:"id,omitempty"`
	Login     string `json:"login,omitempty"`
	FullName  string `json:"full_name,omitempty"`
	Email     string `json:"email,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	Username  string `json:"username,omitempty"`
}

func (u *User) Create(gitea *Client) (bool, error) {
	req := http.Request{
		RemoteAddr: "",
		Method:
	}
	gitea.client.D
}
