package gitea

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