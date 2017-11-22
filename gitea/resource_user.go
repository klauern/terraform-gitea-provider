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

	return setUserResourceData(d, user)
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*gitea.Client)
	user, err := client.GetUserInfo(d.Get("username").(string))
	if err != nil {
		return errors.WithMessage(err, "unable to retrieve user "+d.Get("username").(string))
	}

	return setUserResourceData(d, user)
}

func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*gitea.Client)
	isAdmin := d.Get("is_admin_user").(bool)
	edit := gitea.EditUserOption{
		Admin:     &isAdmin,
		Email:     d.Get("email").(string),
		FullName:  d.Get("full_name").(string),
		LoginName: d.Get("login").(string),
		Password:  d.Get("password").(string),
	}

	err := client.AdminEditUser(d.Get("username").(string), edit)
	if err != nil {
		return errors.WithMessage(err, "unable to edit user "+d.Get("username").(string))
	}
	return nil
}

func resourceUserDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*gitea.Client)
	return client.AdminDeleteUser(d.Get("username").(string))
}

func setUserResourceData(d *schema.ResourceData, u *gitea.User) error {
	if err := d.Set("avatar_url", u.AvatarURL); err != nil {
		return errors.WithMessage(err, "cannot set avatar URL")
	}
	if err := d.Set("email", u.Email); err != nil {
		return errors.WithMessage(err, "cannot set email")
	}
	if err := d.Set("full_name", u.FullName); err != nil {
		return errors.WithMessage(err, "cannot set full name")
	}
	if err := d.Set("id", u.ID); err != nil {
		return errors.WithMessage(err, "cannot set id")
	}
	if err := d.Set("username", u.UserName); err != nil {
		return errors.WithMessage(err, "cannot set username")
	}
	return nil
}
