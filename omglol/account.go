package omglol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Get information about your account. See https://api.omg.lol/#token-get-account-retrieve-account-information
func (c *Client) GetAccountInfo() (*Account, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/account/%s/info", c.HostURL, c.Auth.Email), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	type accountResponse struct {
		Request request `json:"request"`
		Account Account `json:"response"`
	}

	var r accountResponse
	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &r.Account, nil
}

// Get all addresses associated with your account. See https://api.omg.lol/#token-get-account-retrieve-addresses-for-an-account
func (c *Client) GetAccountAddresses() (*[]Address, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/account/%s/addresses", c.HostURL, c.Auth.Email), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	type addressesResponse struct {
		Request   request   `json:"request"`
		Addresses []Address `json:"response"`
	}

	var r addressesResponse
	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &r.Addresses, nil
}

// Get the name associated with the account. See https://api.omg.lol/#token-get-account-retrieve-the-account-name
func (c *Client) GetAccountName() (*string, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/account/%s/name", c.HostURL, c.Auth.Email), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	type nameResponse struct {
		Request  request `json:"request"`
		Response struct {
			Message string `json:"message"`
			Name    string `json:"name"`
		} `json:"response"`
	}

	var r nameResponse
	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &r.Response.Name, nil
}

// Set the name associated with the account. See https://api.omg.lol/#token-post-account-set-the-account-name
func (c *Client) SetAccountName(name string) error {
	jsonData := fmt.Sprintf(`{"name": "%s"}`, name)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/account/%s/name", c.HostURL, c.Auth.Email), bytes.NewBuffer([]byte(jsonData)))

	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return fmt.Errorf("Sent: %s, Error: %w", jsonData, err)
	}

	var r apiResponse
	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return err
	}

	return nil
}

// Get all sessions associated with the account. See https://api.omg.lol/#token-get-account-retrieve-active-sessions
func (c *Client) GetActiveSessions() (*[]ActiveSession, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/account/%s/sessions", c.HostURL, c.Auth.Email), nil)
	if err != nil {
		return nil, err
	}

	var Sessions []ActiveSession

	body, err := c.doRequest(req)
	if err != nil {
		if strings.Contains(err.Error(), "status: 404") {
			return &Sessions, nil
		}
		return nil, err
	}

	var r apiResponse
	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}	

	if err := json.Unmarshal(r.Response, &Sessions); err != nil {
		fmt.Printf("Error unmarshalling sessions: %v\n", err)
		return nil, err
	}

	return &Sessions, nil
}

// Delete a session. See https://api.omg.lol/#token-delete-account-remove-a-session
func (c *Client) DeleteActiveSession(sessionID string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/account/%s/sessions/%s", c.HostURL, c.Auth.Email, sessionID), nil)
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

// Get settings associated with the account. See https://api.omg.lol/#token-get-account-retrieve-account-settings
func (c *Client) GetAccountSettings() (*AccountSettings, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/account/%s/settings", c.HostURL, c.Auth.Email), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	type settingsResponse struct {
		Request  request `json:"request"`
		Response struct {
			Message  string          `json:"message"`
			Settings AccountSettings `json:"settings"`
		} `json:"response"`
	}

	var s settingsResponse
	if err = json.Unmarshal(body, &s); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &s.Response.Settings, nil
}

// Update settings associated with the account. See https://api.omg.lol/#token-post-account-set-account-settings
func (c *Client) SetAccountSettings(settings map[string]string) error {
	jsonData, err := json.Marshal(settings)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/account/%s/settings", c.HostURL, c.Auth.Email), bytes.NewBuffer([]byte(jsonData)))

	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return fmt.Errorf("Sent: %s, Error: %w", jsonData, err)
	}

	var r apiResponse
	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return err
	}

	return nil
}
