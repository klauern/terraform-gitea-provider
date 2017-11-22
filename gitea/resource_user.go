package gitea

import "github.com/hashicorp/terraform/helper/schema"
import "code.gitea.io/sdk/gitea"
import "github.com/pkg/errors"

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "",
			},
			"login": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
			"full_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
			"email": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
			"avatar_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    false,
				Description: "",
			},
			"is_admin_user": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "",
				Default:     false,
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*gitea.Client)
	create := gitea.CreateUserOption{
		Email:      d.Get("email").(string),
		FullName:   d.Get("full_name").(string),
		LoginName:  d.Get("login").(string),
		Password:   d.Get("password").(string),
		SendNotify: false,
		Username:   d.Get("username").(string),
	}

	user, err := client.AdminCreateUser(create)
	if err != nil {
		return errors.WithMessage(err, "unable to create user")
	}

	d.Set("id", user.ID)
	d.Set("avatar_url", user.AvatarURL)
	d.Set("username", user.UserName)
	d.Set("email", user.Email)
	d.Set("full_name", user.FullName)
	return nil
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {
	return nil
}
func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}
func resourceUserDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
