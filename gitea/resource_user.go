package gitea

import "github.com/hashicorp/terraform/helper/schema"

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
