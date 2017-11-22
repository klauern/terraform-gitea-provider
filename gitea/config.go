package gitea

import (
	"code.gitea.io/sdk/gitea"
)

type Config struct {
	Token   string
	BaseURL string
}

func (c *Config) Client() *gitea.Client {
	return gitea.NewClient(c.BaseURL, c.Token)
}
