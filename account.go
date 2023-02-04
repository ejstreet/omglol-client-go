package omglol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Get information about your account. See https://api.omg.lol/#token-get-account-retrieve-account-information
func (c *Client) GetAccountInfo() (*Account, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/account/%s/info", c.HostURL, c.Auth.Email), nil)
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

	var account Account
	if err := json.Unmarshal(response.Response, &account); err != nil {
		fmt.Printf("Error unmarshaling account: %v\n", err)
		return nil, err
	}

	return &account, nil
}

// Get all addresses associated with your account. See https://api.omg.lol/#token-get-account-retrieve-addresses-for-an-account
func (c *Client) GetAccountAddresses() (*Addresses, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/account/%s/addresses", c.HostURL, c.Auth.Email), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var addresses Addresses
	err = json.Unmarshal(body, &addresses)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling body to Addresses: %v", err)
	}

	return &addresses, nil
}

// Get the name associated with the account. See https://api.omg.lol/#token-get-account-retrieve-the-account-name
func (c *Client) GetAccountName() (*AccountName, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/account/%s/name", c.HostURL, c.Auth.Email), nil)
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

	var name AccountName
	if err := json.Unmarshal(response.Response, &name); err != nil {
		fmt.Printf("Error unmarshaling account: %v\n", err)
		return nil, err
	}

	return &name, nil
}

// Set the name associated with the account. See https://api.omg.lol/#token-post-account-set-the-account-name
func (c *Client) SetAccountName(name string) (*AccountName, error) {
	jsonData := fmt.Sprintf(`{"name": "%s"}`, name)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/account/%s/name", c.HostURL, c.Auth.Email), bytes.NewBuffer([]byte(jsonData)))
	
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

	var account_name AccountName
	if err := json.Unmarshal(response.Response, &account_name); err != nil {
		fmt.Printf("Error unmarshaling account: %v\n", err)
		return nil, err
	}

	return &account_name, nil
}

// Get all sessions associated with the account. See https://api.omg.lol/#token-get-account-retrieve-active-sessions
func (c *Client) GetActiveSessions() (*ActiveSessions, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/account/%s/sessions", c.HostURL, c.Auth.Email), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var sessions ActiveSessions
	err = json.Unmarshal(body, &sessions)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling body to ActiveSessions: %v", err)
	}

	return &sessions, nil
}

// Delete a session. See https://api.omg.lol/#token-delete-account-remove-a-session
func (c *Client) DeleteActiveSession(sessionID string) (*MessageResponse, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/account/%s/sessions/%s", c.HostURL, c.Auth.Email, sessionID), nil)
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

// Get settings associated with the account. See https://api.omg.lol/#token-get-account-retrieve-account-settings
func (c *Client) GetAccountSettings() (*AccountSettings, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/account/%s/settings", c.HostURL, c.Auth.Email), nil)
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

	var settings AccountSettings
	if err := json.Unmarshal(response.Response, &settings); err != nil {
		fmt.Printf("Error unmarshaling account: %v\n", err)
		return nil, err
	}

	return &settings, nil
}

// Update settings associated with the account. See https://api.omg.lol/#token-post-account-set-account-settings
func (c *Client) SetAccountSettings(settings map[string]string) (*MessageResponse, error) {
	jsonData, err := json.Marshal(settings)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/account/%s/settings", c.HostURL, c.Auth.Email), bytes.NewBuffer([]byte(jsonData)))
	
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