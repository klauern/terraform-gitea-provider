package main

import "github.com/hashicorp/terraform/helper/schema"

func resourceAdminUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceAdminUserCreate,
		Read:   resourceAdminUserRead,
		Update: resourceAdminUserUpdate,
		Delete: resourceAdminUserDelete,

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
		},
	}
}

func resourceAdminUserCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}
func resourceAdminUserRead(d *schema.ResourceData, m interface{}) error {
	return nil
}
func resourceAdminUserUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}
func resourceAdminUserDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
