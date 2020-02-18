package provider

import (
	"log"
	"github.com/hashicorp/terraform/helper/schema"
	_ "github.com/hashicorp/terraform/terraform"
)

func resourceFile() *schema.Resource {
	log.Println("[TF-SAMPLE] resourceFile...")
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the resource, also acts as it's unique ID",
				ForceNew:    true,
			},
			"filename": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the resource, also acts as it's unique ID",
			},
			"content": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A description of an item",
			},
		},
		Create: resourceCreateItem,
		Read:   resourceReadItem,
		Update: resourceUpdateItem,
		Delete: resourceDeleteItem,
		Exists: resourceExistsItem,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateItem(d *schema.ResourceData, m interface{}) error {
	log.Println("[TF-SAMPLE] resourceCreateItem...")

	name := d.Get("name").(string)
	filename := d.Get("filename").(string)
	content := d.Get("content").(string)

	log.Println("[TF-SAMPLE] resourceCreateItem ", filename, content)
	d.SetId(name)
	return nil
}

func resourceReadItem(d *schema.ResourceData, m interface{}) error {
	log.Println("[TF-SAMPLE] resourceReadItem...")
	d.SetId("test")
	d.Set("name", "test")
	d.Set("filename", "filename")
	d.Set("cotent", "content")
	return nil
}

func resourceUpdateItem(d *schema.ResourceData, m interface{}) error {
	log.Println("[TF-SAMPLE] resourceUpdateItem...")
	return nil
}

func resourceDeleteItem(d *schema.ResourceData, m interface{}) error {
	log.Println("[TF-SAMPLE] resourceDeleteItem...")
	d.SetId("")
	return nil
}

func resourceExistsItem(d *schema.ResourceData, m interface{}) (bool, error) {
	log.Println("[TF-SAMPLE] resourceExistsItem...")
	return false, nil
}
