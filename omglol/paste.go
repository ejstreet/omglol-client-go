package omglol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Create a Paste object
func NewPaste(Title, Content string, ModifiedOn int64) *Paste {
	return &Paste{
		Title:      Title,
		Content:    Content,
		ModifiedOn: ModifiedOn,
	}
}

// Returns a string representaion of a Paste
func (p *Paste) String() string {
	return fmt.Sprintf("Title: %s, Content: %s, ModifiedOn: %d", p.Title, p.Content, p.ModifiedOn)
}

// Create a new Paste. See https://api.omg.lol/#token-post-pastes-create-a-new-paste
func (c *Client) CreatePaste(domain string, paste Paste) error {
	type pasteRequest struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	p := pasteRequest{
		Title:   paste.Title,
		Content: paste.Content,
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/address/%s/pastebin", c.HostURL, domain), bytes.NewBuffer([]byte(jsonData)))
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

// Get a specific paste. See https://api.omg.lol/#noauth-get-pastebin-retrieve-a-specific-paste
func (c *Client) GetPaste(domain string, pasteTitle string) (*Paste, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/address/%s/pastebin/%s", c.HostURL, domain, pasteTitle), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	type getpasteResponse struct {
		Request struct {
			StatusCode int64 `json:"status_code"`
			Success    bool  `json:"success"`
		} `json:"request"`
		Response struct {
			Message string `json:"message"`
			Paste   struct {
				Title      string `json:"title"`
				Content    string `json:"content"`
				ModifiedOn int64  `json:"modified_on"`
			} `json:"paste"`
		} `json:"response"`
	}

	var g getpasteResponse
	if err := json.Unmarshal(body, &g); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return NewPaste(g.Response.Paste.Title, g.Response.Paste.Content, g.Response.Paste.ModifiedOn), nil
}

// Retrieve a list of pastes associated with an address. See https://api.omg.lol/#token-get-pastebin-retrieve-an-entire-pastebin
func (c *Client) ListPastes(address string) (*[]Paste, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/address/%s/pastebin", c.HostURL, address), nil)
	if err != nil {
		return nil, err
	}

	var p []Paste

	body, err := c.doRequest(req)
	if err != nil {
		if strings.Contains(err.Error(), "status: 404") {
			return &p, nil
		}
		return nil, err
	}

	type listpasteResponse struct {
		Request struct {
			StatusCode int64 `json:"status_code"`
			Success    bool  `json:"success"`
		} `json:"request"`
		Response struct {
			Message  string `json:"message"`
			Pastebin []struct {
				Title      string `json:"title"`
				Content    string `json:"content"`
				ModifiedOn int64  `json:"modified_on"`
			} `json:"pastebin"`
		} `json:"response"`
	}

	var r listpasteResponse
	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	for _, paste := range r.Response.Pastebin {
		var x Paste

		x.Title = paste.Title
		x.Content = paste.Content
		x.ModifiedOn = paste.ModifiedOn
		p = append(p, x)
	}

	return &p, nil
}

// Permanently delete a paste. See https://api.omg.lol/#token-delete-pastebin-delete-a-paste-from-a-pastebin
func (c *Client) DeletePaste(domain string, pasteTitle string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/address/%s/pastebin/%s", c.HostURL, domain, pasteTitle), nil)
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
