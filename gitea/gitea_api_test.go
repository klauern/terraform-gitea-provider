package gitea

import (
	"fmt"
	"testing"

	"code.gitea.io/sdk/gitea"
)

func TestGiteaUserCreate(t *testing.T) {
	client := gitea.NewClient("http://localhost:3000", "978fd3ee4f71d6cf8a87556201d37293d719230d")
	user, err := client.AdminCreateUser(gitea.CreateUserOption{
		Email:     "random@example.com",
		LoginName: "random",
		Username:  "random",
	})

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", user)

}
