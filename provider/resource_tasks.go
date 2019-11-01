package provider

import (
	"encoding/json"
	"fmt"

	"bitbucket.org/bestsellerit/terraform-provider-harbor/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var pathVuln = "/api/system/scanAll/schedule"
var TypeStr string
var CronStr string

type schedule struct {
	Schedule cron `json:"schedule`
}

type cron struct {
	Type string `json:"type"`
	Cron string `json:"cron`
}

type Schedule2 struct {
	Type string `json:"type"`
	Cron string `json:"cron"`
}
type Info struct {
	Schedule Schedule2 `json:schedule`
}

func resourceTasks() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"vulnerability_scan_policy": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		Create: resourceTasksCreate,
		Read:   resourceTasksRead,
		Update: resourceTasksUpdate,
		Delete: resourceTasksDelete,
	}
}

func resourceTasksCreate(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	id := RandomString(15)

	vulnShedule := d.Get("vulnerability_scan_policy").(string)
	switch vulnShedule {
	case "hourly":
		TypeStr = "Hourly"
		CronStr = "0 0 * * * *"
	case "daily":
		TypeStr = "Daily"
		CronStr = "0 0 0 * * *"
	case "weekly":
		TypeStr = "Weekly"
		CronStr = "0 0 0 * * 0"
	}

	body := schedule{
		Schedule: cron{
			Type: TypeStr,
			Cron: CronStr,
		},
	}

	resp, _ := apiClient.SendRequest("GET", pathVuln, nil)
	var jsonData Info

	json.Unmarshal([]byte(resp), &jsonData)

	time := jsonData.Schedule.Type
	if time != "" {
		fmt.Println("Shedule found performing PUT request")
		apiClient.SendRequest("PUT", pathVuln, body)
	} else {
		fmt.Println("No shedule found performing POST request")
		apiClient.SendRequest("POST", pathVuln, body)
	}

	d.SetId(id)
	return resourceTasksRead(d, m)
}

func resourceTasksRead(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceTasksUpdate(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	vulnShedule := d.Get("vulnerability_scan_policy").(string)
	switch vulnShedule {
	case "hourly":
		TypeStr = "Hourly"
		CronStr = "0 0 * * * *"
	case "daily":
		TypeStr = "Daily"
		CronStr = "0 0 0 * * *"
	case "weekly":
		TypeStr = "Weekly"
		CronStr = "0 0 0 * * 0"
	}

	body := schedule{
		Schedule: cron{
			Type: TypeStr,
			Cron: CronStr,
		},
	}

	apiClient.SendRequest("PUT", pathVuln, body)

	return resourceTasksRead(d, m)
}

func resourceTasksDelete(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	body := schedule{
		Schedule: cron{Cron: ""},
	}

	apiClient.SendRequest("PUT", pathVuln, body)
	return nil
}
