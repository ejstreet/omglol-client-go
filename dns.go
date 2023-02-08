package client

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

	var d DNSRecords
	if err := json.Unmarshal(body, &d); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &d, nil
}

// Find a single DNS record from its attributes.
func (c *Client) FilterDNSRecord(address string, filterCriteria map[string]interface{}) (*DNSRecord, error) {
	records, err := c.GetDNSRecords(address)
	if err != nil {
		return nil, err
	}

	var matchedRecord DNSRecord
	var matchCount int
	for _, record := range records.Response.DNS {
		match := true
		for key, value := range filterCriteria {
			switch key {
			case "ID":
				if *record.ID != value.(string) {
					match = false
					break
				}
			case "Type":
				if *record.Type != value.(string) {
					match = false
					break
				}
			case "Name":
				if *record.Name != value.(string) {
					match = false
					break
				}
			case "Data":
				if *record.Data != value.(string) {
					match = false
					break
				}
			case "Priority":
				priority, ok := value.(*int)
				if !ok {
					match = false
					break
				}
				if *record.Priority == nil || *record.Priority != priority {
					match = false
					break
				}
			case "TTL":
				if *record.TTL != value.(string) {
					match = false
					break
				}
			case "CreatedAt":
				if *record.CreatedAt != value.(string) {
					match = false
					break
				}
			case "UpdatedAt":
				if *record.UpdatedAt != value.(string) {
					match = false
					break
				}
			default:
				return nil, fmt.Errorf("Invalid filter criteria key: %s", key)
			}
			if !match {
				break
			}
		}
		if match {
			matchedRecord = record
			matchCount++
			if matchCount > 1 {
				return nil, fmt.Errorf("More than one record matches the filter criteria")
			}
		}
	}
	if matchCount == 0 {
		return nil, fmt.Errorf("No records match the filter criteria")
	}
	return &matchedRecord, nil
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

	var d DNSChangeResponse
	if err := json.Unmarshal(body, &d); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &d, nil
}

// Update an existing DNS record. See https://api.omg.lol/#token-patch-dns-edit-an-existing-dns-record
// Note this method does not work at time of writing due to an API bug, see https://github.com/neatnik/omg.lol/issues/584
// Suggested workaround is to use the Replace function instead, which uses the same interface.
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

	var d DNSChangeResponse
	if err := json.Unmarshal(body, &d); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
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
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &response, nil
}

// Delete a record with a given ID, then Create a new one using the provided values.
func (c *Client) ReplaceDNSRecord(domain string, record map[string]string, record_id int) (*DNSChangeResponse, error) {
	_, err := c.DeleteDNSRecord(domain, record_id)
	if err != nil {
		return nil, err
	}

	r, err := c.CreateDNSRecord(domain, record)
	if err != nil {
		return nil, err
	}

	return r, nil
}
