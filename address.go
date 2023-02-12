package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Get information about the availability of an address. See https://api.omg.lol/#noauth-get-address-retrieve-address-availability
func (c *Client) GetAddressAvailability(address string) (*AddressAvailability, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/address/%s/availability", c.HostURL, address), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	type availabilityResponse struct {
		Request      request             `json:"request"`
		Availability AddressAvailability `json:"response"`
	}

	var a availabilityResponse
	if err := json.Unmarshal(body, &a); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &a.Availability, nil
}

// Get the expiration date for an address. See https://api.omg.lol/#noauth-get-address-retrieve-address-expiration
func (c *Client) GetAddressExpiration(address string) (*AddressExpiration, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/address/%s/expiration", c.HostURL, address), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	type expirationResponse struct {
		Request    request           `json:"request"`
		Expiration AddressExpiration `json:"response"`
	}

	var e expirationResponse
	if err := json.Unmarshal(body, &e); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &e.Expiration, nil
}

// Get comprehensive information about an address. See https://api.omg.lol/#token-get-address-retrieve-private-information-about-an-address
func (c *Client) GetAddressInfo(address string) (*AddressInfo, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/address/%s/availability", c.HostURL, address), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	type infoResponse struct {
		Request request     `json:"request"`
		Info    AddressInfo `json:"response"`
	}

	var i infoResponse
	if err := json.Unmarshal(body, &i); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &i.Info, nil
}

// Retrieve the address directory. See https://api.omg.lol/#noauth-get-directory-retreive-the-address-directory
func (c *Client) GetAddressDirectory() (*AddressDirectory, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/directory", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	type directoryResponse struct {
		Request   request          `json:"request"`
		Directory AddressDirectory `json:"response"`
	}

	var d directoryResponse
	if err := json.Unmarshal(body, &d); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &d.Directory, nil
}
