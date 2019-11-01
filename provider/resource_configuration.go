package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceConfigurationCreate,
		Read:   resourceConfigurationRead,
		Update: resourceConfigurationUpdate,
		Delete: resourceConfigurationDelete,

		Schema: map[string]*schema.Schema{
			"auth_mode": &schema.Schema{
				Type: schema.TypeString,
				// Required: true,
			},
		},
	}
}

func resourceConfigurationCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)

	d.SetId(name)
	return resourceConfigurationRead(d, m)
}

func resourceConfigurationRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceConfigurationUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceConfigurationRead(d, m)
}

func resourceConfigurationDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
