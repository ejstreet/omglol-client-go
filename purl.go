package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Create a new PURL. See https://api.omg.lol/#token-post-purls-create-a-new-purl
func (c *Client) CreatePersistentURL(domain string, purlName string, url string) (*PersistantURL, error) {
	jsonData, err := json.Marshal(map[string]string{"name": purlName, "url": url})

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/address/%s/purl", c.HostURL, domain), bytes.NewBuffer([]byte(jsonData)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var p PersistantURL
	if err := json.Unmarshal(body, &p); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &p, nil
}

// Get a specific PURL. See https://api.omg.lol/#token-get-purls-retrieve-a-specific-purl
func (c *Client) GetPersistentURL(domain string, purlName string) (*PersistantURL, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/address/%s/purl/%s", c.HostURL, domain, purlName), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var p PersistantURL
	if err := json.Unmarshal(body, &p); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &p, nil
}

// Retrieve a list of PURLs associated with an address. See https://api.omg.lol/#token-get-purls-retrieve-a-list-of-purls-for-an-address
func (c *Client) ListPersistentURLs(address string) (*PersistantURLs, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/address/%s/purls", c.HostURL, address), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var p PersistantURLs
	if err := json.Unmarshal(body, &p); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &p, nil
}

// Permanently delete a PURL. See https://api.omg.lol/#token-delete-purls-delete-a-purl
func (c *Client) DeletePersistentURL(domain string, purlName string) (*MessageResponse, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/address/%s/purl/%s", c.HostURL, domain, purlName), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response MessageResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &response, nil
}
