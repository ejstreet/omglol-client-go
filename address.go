package omglol

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Get information about the availability of an address. See https://api.omg.lol/#noauth-get-address-retrieve-address-availability
func (c *Client) GetAddressAvailability(address string) (*Address, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/address/%s/availability", c.HostURL, address), nil)
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

	var a Address
	if err := json.Unmarshal(response.Response, &a); err != nil {
		fmt.Printf("Error unmarshaling address: %v\n", err)
		return nil, err
	}

	return &a, nil
}

// Get the expiration date for an address. See https://api.omg.lol/#noauth-get-address-retrieve-address-expiration
func (c *Client) GetAddressExpiration(address string) (*AddressExpiration, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/address/%s/availability", c.HostURL, address), nil)
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

	var e AddressExpiration
	if err := json.Unmarshal(response.Response, &e); err != nil {
		fmt.Printf("Error unmarshaling address expiration: %v\n", err)
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

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("Error unmarshaling response: %v\n", err)
		return nil, err
	}

	var i AddressInfo
	if err := json.Unmarshal(response.Response, &i); err != nil {
		fmt.Printf("Error unmarshaling address detail: %v\n", err)
		return nil, err
	}

	return &i, nil
}

