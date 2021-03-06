package gitea

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider is the Terraform provider for Gitea.
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
			"gitea_user": resourceUser(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &Config{
		Token:   d.Get("token").(string),
		BaseURL: d.Get("base_url").(string),
	}

	return config.Client(), nil
}
