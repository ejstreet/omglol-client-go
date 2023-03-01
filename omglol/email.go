package omglol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Get email forwarding address(es) for a domain. See https://api.omg.lol/#token-get-email-retrieve-forwarding-addresses
func (c *Client) GetEmail(domain string) ([]string, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/address/%s/email", c.HostURL, domain), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var r emailResponse
	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Printf("error unmarshalling response: %v\n", err)
		return nil, err
	}

	return r.Response.DestinationArray, nil
}

// Set email forwarding address(es) for a domain. See https://api.omg.lol/#token-post-email-set-forwarding-addresses
func (c *Client) SetEmail(domain string, destination []string) error {
	jsonData := fmt.Sprintf(`{"destination": "%s"}`, strings.Join(destination, ", "))
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/address/%s/email", c.HostURL, domain), bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return fmt.Errorf("sent: %s, error: %w", jsonData, err)
	}

	var r emailResponse
	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Printf("error unmarshalling response: %v\n", err)
		return err
	}

	return nil
}
