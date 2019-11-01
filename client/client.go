package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	url        string
	username   string
	password   string
	httpClient *http.Client
}

// NewClient creates common settings
func NewClient(url string, username string, password string) *Client {

	return &Client{
		url:        url,
		username:   username,
		password:   password,
		httpClient: &http.Client{},
	}
}

func (c *Client) PutRequest(path string, payload interface{}) error {
	url := c.url + path

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(payload)
	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, b)
	if err != nil {
		log.Panicln(err)
	}

	req.SetBasicAuth(c.username, c.password)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	resp.Body.Close()

	fmt.Println(string(body))

	return nil
}

func (c *Client) SendRequest(method string, path string, payload interface{}) (value string, err error) {
	url := c.url + path

	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(payload)
	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, b)
	if err != nil {
		log.Println(err)
	}

	req.SetBasicAuth(c.username, c.password)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)

	}
	resp.Body.Close()

	strbody := string(body)

	return strbody, nil
}
