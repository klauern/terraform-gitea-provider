package gitea

import "github.com/hashicorp/terraform/helper/schema"

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"login": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
			"full_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
			"email": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
			"avatar_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"is_admin_user": &schema.Schema{
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
