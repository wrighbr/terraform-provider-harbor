package provider

import (
	"encoding/json"
	"log"
	"strconv"

	"bitbucket.org/bestsellerit/terraform-provider-harbor/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var pathProjects = "/api/projects"

type project struct {
	ProjectName string   `json:"project_name"`
	Metadata    metadata `json:"metadata"`
}

type metadata struct {
	AutoScan string `json:"auto_scan"`
}

type projects struct {
	Name      string `json:"name"`
	ProjectID int    `json:"project_id"`
}

func resourceProject() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"project_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
		Create: resourceProjectCreate,
		Read:   resourceProjectRead,
		// Update: resourceProjectUpdate,
		Delete: resourceProjectDelete,
	}
}

func resourceProjectCreate(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	body := project{
		ProjectName: d.Get("name").(string),
		Metadata:    metadata{AutoScan: "true"},
	}

	apiClient.SendRequest("POST", pathProjects, body)

	id := RandomString(15)
	d.SetId(id)
	return resourceProjectRead(d, m)
}

func resourceProjectRead(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	body := project{
		ProjectName: d.Get("name").(string),
	}

	resp, err := apiClient.SendRequest("GET", pathProjects+"?name="+body.ProjectName, nil)
	if err != nil {
		log.Println(err)
	}

	var jsonData []projects

	json.Unmarshal([]byte(resp), &jsonData)

	projectid := jsonData[0].ProjectID

	d.Set("project_id", projectid)
	return nil
}

func resourceProjectUpdate(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	body := project{
		ProjectName: d.Get("name").(string),
	}

	_, err := apiClient.SendRequest("PUT", pathProjects, body)
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}

	return resourceProjectRead(d, m)
}

func resourceProjectDelete(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	id := d.Get("project_id").(int)

	apiClient.SendRequest("DELETE", pathProjects+"/"+strconv.Itoa(id), nil)
	return nil
}
