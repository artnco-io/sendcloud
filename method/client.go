package method

import (
	"fmt"
	"strconv"

	"github.com/lpieri/sendcloud"
)

type Client struct {
	apiKey    string
	apiSecret string
}

func New(apiKey string, apiSecret string) *Client {
	return &Client{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

func (c *Client) GetMethodsWithArgs(fromPostalCode string, toPostalCode string, isReturn bool, senderAddress string, servicePointID int64) ([]*sendcloud.Method, error) {
	smr := sendcloud.MethodListResponseContainer{}

	URL := "/api/v2/shipping_methods?to_country=FR"
	if fromPostalCode != "" {
		URL = fmt.Sprint(URL, "&from_postal_code=", fromPostalCode)
	}
	if toPostalCode != "" {
		URL = fmt.Sprint(URL, "&to_postal_code=", fromPostalCode)
	}
	if isReturn {
		URL = fmt.Sprint(URL, "&is_return=", isReturn)
	}
	if senderAddress != "" {
		URL = fmt.Sprint(URL, "&sender_address=", senderAddress)
	}
	if servicePointID > 0 {
		URL = fmt.Sprint(URL, "&service_point_id=", servicePointID)
	}

	err := sendcloud.Request(
		"GET",
		URL,
		nil,
		c.apiKey,
		c.apiSecret,
		&smr,
	)
	if err != nil {
		return nil, err
	}
	return smr.GetResponse().([]*sendcloud.Method), nil
}

// Get all shipment methods
func (c *Client) GetMethods() ([]*sendcloud.Method, error) {
	smr := sendcloud.MethodListResponseContainer{}
	err := sendcloud.Request("GET", "/api/v2/shipping_methods", nil, c.apiKey, c.apiSecret, &smr)
	if err != nil {
		return nil, err
	}
	return smr.GetResponse().([]*sendcloud.Method), nil
}

// Get a single method
func (c *Client) GetMethod(id int64) (*sendcloud.Method, error) {
	mr := sendcloud.MethodResponseContainer{}
	err := sendcloud.Request("GET", "/api/v2/shipping_methods/"+strconv.Itoa(int(id)), nil, c.apiKey, c.apiSecret, &mr)
	if err != nil {
		return nil, err
	}
	return mr.GetResponse().(*sendcloud.Method), nil
}
