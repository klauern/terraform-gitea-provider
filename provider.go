package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITEA_TOKEN", nil),
				Description: "Token key for Gitea API access",
			},
			"base_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITEA_URL", nil),
				Description: "URL to the Gitea Server (https://gitea.user.com:8888)",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"resource_user": resourceUser(),
		},
		// ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	// TODO
	return nil, nil
}
