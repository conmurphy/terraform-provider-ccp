package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CCP_USERNAME", nil),
				Description: "Username used to access Cisco Container Platform",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CCP_PASSWORD", nil),
				Description: "Password used to access Cisco Container Platform",
			},
			"base_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CCP_URL", nil),
				Description: "URL to the Cisco Container Platform",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ccp_user":    resourceUser(),
			"ccp_cluster": resourceCluster(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &Config{
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		Base_url: d.Get("base_url").(string),
	}

	return config.Client(), nil
}
