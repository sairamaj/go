package provider

import (
	"log"
	"os"

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
	providerData := m.(*ProviderData)
	name := d.Get("name").(string)
	filename := d.Get("filename").(string)
	content := d.Get("content").(string)

	filePath := providerData.Path + filename
	log.Println("[TF-SAMPLE] resourceCreateItem ", filename, content, "fullpath:", filePath)

	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(content)
	d.SetId(name)
	log.Println("[TF-SAMPLE] resourceCreateItem returning...")
	return nil
}

func resourceReadItem(d *schema.ResourceData, m interface{}) error {
	log.Println("[TF-SAMPLE] resourceReadItem...")
	d.SetId("test")
	// d.Set("name", "test")
	// d.Set("filename", "filename")
	// d.Set("cotent", "content")
	log.Println("[TF-SAMPLE] resourceReadItem returning")
	return nil
}

func resourceUpdateItem(d *schema.ResourceData, m interface{}) error {
	log.Println("[TF-SAMPLE] resourceUpdateItem...")
	return nil
}

func resourceDeleteItem(d *schema.ResourceData, m interface{}) error {
	log.Println("[TF-SAMPLE] resourceDeleteItem...")
	providerData := m.(*ProviderData)
	filename := d.Get("filename").(string)

	filePath := providerData.Path + filename
	log.Println("[TF-SAMPLE] resourceDeleteItem ", filename, "fullpath:", filePath)
	if _, err := os.Stat(filePath); err == nil {
		log.Println("[TF-SAMPLE] resourceDeleteItem ", filePath, "exists and hence removing.")
		os.Remove(filePath)
	}else{
		log.Println("[TF-SAMPLE] resourceDeleteItem err:", err)
	}
	d.SetId("")
	log.Println("[TF-SAMPLE] resourceDeleteItem returning...")
	return nil
}

func resourceExistsItem(d *schema.ResourceData, m interface{}) (bool, error) {
	log.Println("[TF-SAMPLE] resourceExistsItem...")
	providerData := m.(*ProviderData)
	filename := d.Get("filename").(string)
	filePath := providerData.Path + filename
	log.Println("[TF-SAMPLE] resourceDeleteItem ", filename, "fullpath:", filePath)
	if _, err := os.Stat(filePath); err == nil {
		log.Println("[TF-SAMPLE] resourceExistsItem returning true")
		return true, nil
	}else{
		log.Println("[TF-SAMPLE] resourceExistsItem err:", err)
	}

	log.Println("[TF-SAMPLE] resourceExistsItem returning false.")
	return false, nil
}
