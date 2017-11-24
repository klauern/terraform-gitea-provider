package gitea

import (
	"fmt"
	"strings"
	"testing"

	"code.gitea.io/sdk/gitea"
)

func TestGiteaUserDelete(t *testing.T) {
	client := gitea.NewClient("http://localhost:3000", "978fd3ee4f71d6cf8a87556201d37293d719230d")
	if err := client.AdminDeleteUser("random"); err != nil {
		if strings.Contains(err.Error(), "404") {
			t.Log("User already deleted")
		} else {
			t.Fatal(err)
		}
	}
}

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

func TestGiteaUserEdit(t *testing.T) {
	client := gitea.NewClient("http://localhost:3000", "978fd3ee4f71d6cf8a87556201d37293d719230d")
	disabled := false
	err := client.AdminEditUser("random", gitea.EditUserOption{
		Active: &disabled,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGiteaUserRead(t *testing.T) {
	client := gitea.NewClient("http://localhost:3000", "978fd3ee4f71d6cf8a87556201d37293d719230d")
	user, err := client.GetUserInfo("random")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", user)
}
