package provider

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"path": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_ADDRESS", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"sample_item": resourceFile(),
		},
		ConfigureFunc: providerConfigure,
	}
}

type ProviderData struct {
	Path string
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	log.Println("[TF-SAMPLE] providerConfigure", d)
	path := d.Get("path").(string)

	log.Println("[TF-SAMPLE] path:", path)
	return &ProviderData{Path: path}, nil
}
