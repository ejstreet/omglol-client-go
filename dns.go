package omglol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Get a list of all of your DNS records for an address. See https://api.omg.lol/#token-get-dns-retrieve-dns-records-for-an-address
func (c *Client) GetDNSRecords(address string) (*DNSRecords, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/address/%s/dns", c.HostURL, address), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("Error unmarshaling response: %v\n", err)
		return nil, err
	}

	var d DNSRecords
	if err := json.Unmarshal(response.Response, &d); err != nil {
		fmt.Printf("Error unmarshaling DNS records: %v\n", err)
		return nil, err
	}

	return &d, nil
}


// Add a new DNS record. See https://api.omg.lol/#token-post-dns-create-a-new-dns-record
func (c *Client) CreateDNSRecord(domain string, record map[string]string) (*DNSChangeResponse, error) {
	jsonData, err := json.Marshal(record)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/address/%s/dns", c.HostURL, domain), bytes.NewBuffer([]byte(jsonData)))
	
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("Error unmarshaling response: %v\n", err)
		return nil, err
	}

	var d DNSChangeResponse
	if err := json.Unmarshal(response.Response, &d); err != nil {
		fmt.Printf("Error unmarshaling account: %v\n", err)
		return nil, err
	}

	return &d, nil
}

// Update an existing DNS record. See https://api.omg.lol/#token-patch-dns-edit-an-existing-dns-record
// Note this method does not work at time of writing due to an API bug, see https://github.com/neatnik/omg.lol/issues/584
// Suggested workaround is to invoke Delete followed by Create.
func (c *Client) UpdateDNSRecord(domain string, record map[string]string, record_id int) (*DNSChangeResponse, error) {
	jsonData, err := json.Marshal(record)

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/address/%s/dns/%d", c.HostURL, domain, record_id), bytes.NewBuffer([]byte(jsonData)))
	
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("Error unmarshaling response: %v\n", err)
		return nil, err
	}

	var d DNSChangeResponse
	if err := json.Unmarshal(response.Response, &d); err != nil {
		fmt.Printf("Error unmarshaling account: %v\n", err)
		return nil, err
	}

	return &d, nil
}

// Delete a DNS record. See https://api.omg.lol/#token-delete-dns-delete-a-dns-record
func (c *Client) DeleteDNSRecord(domain string, record_id int) (*MessageResponse, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/address/%s/dns/%d", c.HostURL, domain, record_id), nil)
	
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response MessageResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("Error unmarshaling response: %v\n", err)
		return nil, err
	}

	return &response, nil
}