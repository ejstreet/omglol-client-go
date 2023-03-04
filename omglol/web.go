package omglol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Get webpage content for a domain. See https://api.omg.lol/#token-get-web-retrieve-web-page-content
func (c *Client) GetWeb(domain string) (*Web, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/address/%s/web", c.HostURL, domain), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	type webResponse struct {
		Request  request `json:"request"`
		Response Web     `json:"response"`
	}

	var w webResponse
	if err = json.Unmarshal(body, &w); err != nil {
		fmt.Printf("error unmarshalling response: %v\n", err)
		return nil, err
	}
	w.Response.ContentBytes = []byte(w.Response.Content)
	return &w.Response, nil
}

// Update webpage content for a domain. See https://api.omg.lol/#token-post-web-update-web-page-content-and-publish
func (c *Client) SetWeb(domain string, content []byte, publish bool) (bool, error) {
	type setWebInput struct {
		Content string `json:"content"`
		Publish bool   `json:"publish,omitempty"`
	}

	input := setWebInput{string(content), publish}
	jsonData, err := json.Marshal(input)
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/address/%s/web", c.HostURL, domain), bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		return false, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return false, fmt.Errorf("sent: %s, error: %w", jsonData, err)
	}

	var r apiResponse
	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Printf("error unmarshalling response: %v\n", err)
		return false, err
	}

	type decodedResponse struct {
		Message string `json:"message"`
	}
	var m decodedResponse
	if err := json.Unmarshal(r.Response, &m); err != nil {
		fmt.Printf("error unmarshalling response: %v\n", err)
		return false, err
	}

	if m.Message == "Your web content has been saved and published." {
		return true, nil
	} else if m.Message == "Your web content has been saved." {
		return false, nil
	} else {
		return false, fmt.Errorf("unexpected response: %s", m.Message)
	}
}
