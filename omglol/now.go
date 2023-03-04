package omglol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Get Now page content for a domain. See https://api.omg.lol/#noauth-get-now-page-retrieve-/now-page
func (c *Client) GetNow(domain string) (*Now, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/address/%s/now", c.HostURL, domain), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	type nowResponse struct {
		Request  request `json:"request"`
		Response struct {
			Message string `json:"message"`
			Now     Now    `json:"now"`
		}
	}

	var r nowResponse
	if err = json.Unmarshal(body, &r); err != nil {
		fmt.Printf("error unmarshalling response: %v\n", err)
		return nil, err
	}
	r.Response.Now.ContentBytes = []byte(r.Response.Now.Content)
	return &r.Response.Now, nil
}

// Update Now page content for a domain. See https://api.omg.lol/#token-post-now-page-update-/now-page
func (c *Client) SetNow(domain string, content []byte, listed bool) error {
	type setNowInput struct {
		Content string `json:"content"`
		Listed  int    `json:"listed"`
	}

	var listedInt int
	if listed {
		listedInt = 1
	} else {
		listedInt = 0
	}
	input := setNowInput{string(content), listedInt}
	jsonData, err := json.Marshal(input)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/address/%s/now", c.HostURL, domain), bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return fmt.Errorf("sent: %s, error: %w", jsonData, err)
	}

	var r apiResponse
	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Printf("error unmarshalling response: %v\n", err)
		return err
	}

	return nil
}

// Retrieve the Now garden. See https://api.omg.lol/#noauth-get-now-page-retrieve-the-now.garden-listing
func (c *Client) GetNowGarden() (*[]NowGardenEntry, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/now/garden", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	type gardenResponse struct {
		Request  request `json:"request"`
		Response struct {
			Message string           `json:"message"`
			Garden  []NowGardenEntry `json:"garden"`
		} `json:"response"`
	}

	var g gardenResponse
	if err := json.Unmarshal(body, &g); err != nil {
		fmt.Printf("error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &g.Response.Garden, nil
}
