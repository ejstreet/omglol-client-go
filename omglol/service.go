package omglol

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Get statistics about omg.lol. See https://api.omg.lol/#noauth-get-service-retrieve-service-information-and-statistics
func (c *Client) GetServiceInfo() (*ServiceInfo, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/service/info", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	type serviceResponse struct {
		Request  request     `json:"request"`
		Response ServiceInfo `json:"response"`
	}

	var r serviceResponse
	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	return &r.Response, nil
}
