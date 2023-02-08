package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Get information about the availability of an address. See https://api.omg.lol/#noauth-get-address-retrieve-address-availability
func (c *Client) GetAddressAvailability(address string) (*AddressAvailability, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/address/%s/availability", c.HostURL, address), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var a AddressAvailability
	if err := json.Unmarshal(body, &a); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &a, nil
}

// Get the expiration date for an address. See https://api.omg.lol/#noauth-get-address-retrieve-address-expiration
func (c *Client) GetAddressExpiration(address string) (*AddressExpiration, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/address/%s/expiration", c.HostURL, address), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var e AddressExpiration
	if err := json.Unmarshal(body, &e); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &e, nil
}

// Get comprehensive information about an address. See https://api.omg.lol/#token-get-address-retrieve-private-information-about-an-address
func (c *Client) GetAddressInfo(address string) (*AddressInfo, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/address/%s/availability", c.HostURL, address), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var i AddressInfo
	if err := json.Unmarshal(body, &i); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &i, nil
}

// Retrieve the address directory. See https://api.omg.lol/#noauth-get-directory-retreive-the-address-directory
func (c *Client) GetAddressDirectory() (*AddressDirectory, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/directory", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var directory AddressDirectory
	if err := json.Unmarshal(body, &directory); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &directory, nil
}
