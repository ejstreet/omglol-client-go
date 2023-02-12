package omglol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Create a PersistentURL object
func NewPersistentURL(Name string, URL string, Counter ...*int) *PersistentURL {
	var counter *int
	if len(Counter) > 0 {
		c := Counter[0]
		counter = c
	}

	return &PersistentURL{
		Name:    Name,
		URL:     URL,
		Counter: counter,
	}
}

// Returns a string representaion of a PersistentURL
func (p *PersistentURL) ToString() string {
	counter := "<nil>"
	if p.Counter != nil {
		counter = strconv.Itoa(*p.Counter)
	}
	return fmt.Sprintf("Name: %s, URL: %s, Counter: %s", p.Name, p.URL, counter)
}

// Create a new PersistentURL. See https://api.omg.lol/#token-post-purls-create-a-new-purl
func (c *Client) CreatePersistentURL(domain string, purl PersistentURL) error {
	jsonData, err := json.Marshal(purl)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/address/%s/purl", c.HostURL, domain), bytes.NewBuffer([]byte(jsonData)))

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

// Get a specific PURL. See https://api.omg.lol/#token-get-purls-retrieve-a-specific-purl
func (c *Client) GetPersistentURL(domain string, purlName string) (*PersistentURL, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/address/%s/purl/%s", c.HostURL, domain, purlName), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	type getPURLResponse struct {
		Request struct {
			StatusCode int  `json:"status_code"`
			Success    bool `json:"success"`
		} `json:"request"`
		Response struct {
			Message string `json:"message"`
			PURL    struct {
				Name    string `json:"name"`
				URL     string `json:"url"`
				Counter *int   `json:"counter"`
			} `json:"purl"`
		} `json:"response"`
	}

	var g getPURLResponse
	if err := json.Unmarshal(body, &g); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return NewPersistentURL(g.Response.PURL.Name, g.Response.PURL.URL, g.Response.PURL.Counter), nil
}

// Retrieve a list of PURLs associated with an address. See https://api.omg.lol/#token-get-purls-retrieve-a-list-of-purls-for-an-address
func (c *Client) ListPersistentURLs(address string) (*[]PersistentURL, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/address/%s/purls", c.HostURL, address), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	type listPURLResponse struct {
		Request struct {
			StatusCode int    `json:"status_code"`
			Success    bool   `json:"success"`
		} `json:"request"`
		Response struct {
			Message string  `json:"message"`
			PURLs   []struct {
				Name    string `json:"name"`
				URL     string `json:"url"`
				Counter *string `json:"counter"`
			} `json:"purls"`
		} `json:"response"`
	}

	var r listPURLResponse
	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	var p []PersistentURL

	for _, purl := range r.Response.PURLs {
		var x PersistentURL

		x.Name = purl.Name
		x.URL = purl.URL
		
		if purl.Counter == nil {
			x.Counter = nil
		} else {
			p, _ := strconv.Atoi(*purl.Counter)
			x.Counter = &p
		}

		p = append(p, x)
	}

	return &p, nil
}

// Permanently delete a PURL. See https://api.omg.lol/#token-delete-purls-delete-a-purl
func (c *Client) DeletePersistentURL(domain string, purlName string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/address/%s/purl/%s", c.HostURL, domain, purlName), nil)

	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	var response apiResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return err
	}

	return nil
}
