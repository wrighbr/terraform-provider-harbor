package provider

import (
	"log"

	"bitbucket.org/bestsellerit/terraform-provider-harbor/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var path = "/api/configurations"

type system struct {
	ProjectCreationRestriction string `json:"project_creation_restriction"`
	ReadOnly                   string `json:"read_only,omitempty"`

	// EmailVerifyCert string `json:"email_verify_cert,omitempty"`
}

func resourceConfigSystem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"project_creation_restriction": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"read_only": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
		Create: resourceConfigSystemCreate,
		Read:   resourceConfigSystemRead,
		Update: resourceConfigSystemUpdate,
		Delete: resourceConfigSystemDelete,
	}
}

func resourceConfigSystemCreate(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	body := system{
		ProjectCreationRestriction: d.Get("project_creation_restriction").(string),
		ReadOnly:                   d.Get("read_only").(string),
	}

	id := RandomString(15)

	err := apiClient.PutRequest(path, body)

	if err != nil {
		log.Printf("[ERROR] %s", err)
	}

	d.SetId(id)
	// return resourceConfigSystemRead(d, m)
	return nil
}

func resourceConfigSystemRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceConfigSystemUpdate(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	body := system{
		ProjectCreationRestriction: d.Get("project_creation_restriction").(string),
		ReadOnly:                   d.Get("Read_only").(string),
	}

	err := apiClient.PutRequest(path, body)

	if err != nil {
		log.Printf("[ERROR] %s", err)
	}

	return resourceConfigSystemRead(d, m)
}

func resourceConfigSystemDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
