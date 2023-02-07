package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default omg.lol URL
const HostURL string = "https://api.omg.lol"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Auth       AuthStruct
}

// AuthStruct -
type AuthStruct struct {
	Email  string `json:"email"`
	ApiKey string `json:"api_key"`
}

// NewClient -
func NewClient(email string, api_key string, host ...string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default omg.lol URL
		HostURL: HostURL,
	}

	if len(host) > 0 {
		c.HostURL = host[0]
	}

	// If email or api_key not provided, return empty client
	if email == "" || api_key == "" {
		return &c, nil
	}

	c.Auth = AuthStruct{
		Email:  email,
		ApiKey: api_key,
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	token := c.Auth.ApiKey
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, nil
}


