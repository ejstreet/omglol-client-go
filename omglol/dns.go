package omglol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
			Message string `json:"message"`
			DNS     []struct {
				ID        string  `json:"id"`
				Type      string  `json:"type"`
				Name      string  `json:"name"`
				Data      string  `json:"data"`
				Priority  *string `json:"priority"`
				TTL       string  `json:"ttl"`
				CreatedAt string  `json:"created_at"`
				UpdatedAt string  `json:"updated_at"`
			} `json:"dns"`
		} `json:"response"`
	}

	var r recordsResponse
	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	// Some type conversion is required, as the API is a little inconsistent with its return types

	var d []DNSRecord

	for _, record := range r.Response.DNS {
		id_int, _ := strconv.Atoi(record.ID)
		id := int64(id_int)
		ttl_int, _ := strconv.Atoi(record.TTL)
		ttl := int64(ttl_int)

		var p *int64
		if record.Priority != nil {
			p_int, _ := strconv.Atoi(*record.Priority)
			p_int64 := int64(p_int)
			p = &p_int64
		}

		x := newDNSRecord(id, ttl, record.Type, record.Name, record.Data, record.CreatedAt, record.UpdatedAt, p)

		d = append(d, *x)
	}

	return &d, nil
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
				var id int64
				switch v := value.(type) {
				case int:
					id = int64(v)
				case int64:
					id = v
				default:
					match = false
					break
				}
				if id != *record.ID {
					match = false
					break
				}
			case "Type":
				if *record.Type != value.(string) {
					match = false
					break
				}
			case "Name":
				if strings.Contains(*record.Name, ".") && *record.Name != value.(string) {
					match = false
					break
				} else if !strings.Contains(*record.Name, ".") && "@" != value.(string) {
					match = false
					break
				}
			case "Data":
				if *record.Data != value.(string) {
					match = false
					break
				}
			case "Priority":
				switch v := value.(type) {
				case int:
					if record.Priority != nil && *record.Priority != int64(v) {
						match = false
						break
					}
				case int64:
					if record.Priority != nil && *record.Priority != v {
						match = false
						break
					}
				case nil:
					if record.Priority != nil {
						match = false
						break
					}
				default:
					match = false
					break
				}
			case "TTL":
				var ttl int64
				switch v := value.(type) {
				case int:
					ttl = int64(v)
				case int64:
					ttl = v
				default:
					match = false
					break
				}
				if ttl != *record.TTL {
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

func NewDNSEntry(Type, Name, Data string, TTL int64, Priority ...int64) *DNSEntry {
	var priority *int64
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

func newDNSRecord(id, ttl int64, typ, name, data, createdAt, updatedAt string, priority *int64) *DNSRecord {
	return &DNSRecord{
		ID:        &id,
		Type:      &typ,
		Name:      &name,
		Data:      &data,
		Priority:  priority,
		TTL:       &ttl,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
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
		priority = strconv.Itoa(int(*d.Priority))
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
		return nil, fmt.Errorf("Sent: %s, Error: %w", jsonData, err)
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
func (c *Client) UpdateDNSRecord(domain string, record DNSEntry, record_id int64) (*DNSRecord, error) {
	jsonData, err := json.Marshal(record)

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/address/%s/dns/%d", c.HostURL, domain, record_id), bytes.NewBuffer([]byte(jsonData)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("Sent: %s, Error: %w", jsonData, err)
	}

	var r dnsChangeResponse
	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return convertRecordResponse(r.Response.ResponseReceived.Data), nil
}

// Delete a DNS record. See https://api.omg.lol/#token-delete-dns-delete-a-dns-record
func (c *Client) DeleteDNSRecord(domain string, record_id int64) error {
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
func (c *Client) ReplaceDNSRecord(domain string, record DNSEntry, record_id int64) (*DNSRecord, error) {
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
