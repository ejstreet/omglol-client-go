package omglol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Get a list of all of your DNS records for an address. See https://api.omg.lol/#token-get-dns-retrieve-dns-records-for-an-address
func (c *Client) ListDNSRecords(address string) (*[]DNSRecord, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/address/%s/dns", c.HostURL, address), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	type recordsResponse struct {
		Request  request `json:"request"`
		Response struct {
			Message string      `json:"message"`
			DNS     []DNSRecord `json:"dns"`
		} `json:"response"`
	}

	var r recordsResponse
	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &r.Response.DNS, nil
}

// Find a single DNS record from its attributes.
func (c *Client) FilterDNSRecord(address string, filterCriteria map[string]interface{}) (*DNSRecord, error) {
	records, err := c.ListDNSRecords(address)
	if err != nil {
		return nil, err
	}

	var matchedRecord DNSRecord
	var matchCount int
	for _, record := range *records {
		match := true
		for key, value := range filterCriteria {
			switch key {
			case "ID":
				if *record.ID != value.(int) {
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
				if record.Priority != priority {
					match = false
					break
				}
			case "TTL":
				if *record.TTL != value.(int) {
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

func NewDNSEntry(Type string, Name string, Data string, TTL int, Priority ...int) *DNSEntry {
	var priority *int
	if len(Priority) > 0 {
		p := Priority[0]
		priority = &p
	}
	return &DNSEntry{
		Type:     &Type,
		Name:     &Name,
		Data:     &Data,
		TTL:      &TTL,
		Priority: priority,
	}
}

// This function resolves an inconsistency in the current version of the API response. Hopefully it won't be required for too long.
func convertRecordResponse(r dnsRecordContent) *DNSRecord {
	var d DNSRecord

	d.ID = r.ID
	d.Type = r.Type
	d.Name = r.Name
	d.Data = r.Content
	d.Priority = r.Priority
	d.TTL = r.TTL
	d.CreatedAt = r.CreatedAt
	d.UpdatedAt = r.UpdatedAt

	return &d
}

// Returns a string representaion of a DNS record
func (d *DNSRecord) ToString() string {

	priority := "<nil>"
	if d.Priority != nil {
		priority = strconv.Itoa(*d.Priority)
	}
	return fmt.Sprintf("ID: %d, Type: %s, Name: %s, Data: %s, Priority: %s, TTL: %d, CreatedAt: %s, UpdatedAt: %s", *d.ID, *d.Type, *d.Name, *d.Data, priority, *d.TTL, *d.CreatedAt, *d.UpdatedAt)
}

// Add a new DNS record. See https://api.omg.lol/#token-post-dns-create-a-new-dns-record
func (c *Client) CreateDNSRecord(domain string, record DNSEntry) (*DNSRecord, error) {
	jsonData, err := json.Marshal(record)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/address/%s/dns", c.HostURL, domain), bytes.NewBuffer([]byte(jsonData)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var r dnsChangeResponse
	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return convertRecordResponse(r.Response.ResponseReceived.Data), nil
}

// Update an existing DNS record. See https://api.omg.lol/#token-patch-dns-edit-an-existing-dns-record
// Note this method does not work at time of writing due to an API bug, see https://github.com/neatnik/omg.lol/issues/584
// Suggested workaround is to use the Replace function instead, which uses the same interface.
func (c *Client) UpdateDNSRecord(domain string, record DNSEntry, record_id int) (*DNSRecord, error) {
	jsonData, err := json.Marshal(record)

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/address/%s/dns/%d", c.HostURL, domain, record_id), bytes.NewBuffer([]byte(jsonData)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var r dnsChangeResponse
	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return convertRecordResponse(r.Response.ResponseReceived.Data), nil
}

// Delete a DNS record. See https://api.omg.lol/#token-delete-dns-delete-a-dns-record
func (c *Client) DeleteDNSRecord(domain string, record_id int) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/address/%s/dns/%d", c.HostURL, domain, record_id), nil)

	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	var r apiResponse
	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return err
	}

	return nil
}

// Delete a record with a given ID, then Create a new one using the provided values.
func (c *Client) ReplaceDNSRecord(domain string, record DNSEntry, record_id int) (*DNSRecord, error) {
	err := c.DeleteDNSRecord(domain, record_id)
	if err != nil {
		return nil, err
	}

	r, err := c.CreateDNSRecord(domain, record)
	if err != nil {
		return nil, err
	}

	return r, nil
}
