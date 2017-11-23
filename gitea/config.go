package gitea

import (
	"code.gitea.io/sdk/gitea"
)

// Config represents the configuration necessary to connect to a Gitea API Server.
type Config struct {
	Token   string
	BaseURL string
}

// Client will create a new Gitea HTTP client based on the Config.
func (c *Config) Client() *gitea.Client {
	return gitea.NewClient(c.BaseURL, c.Token)
}
