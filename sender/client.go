package sender

import (
	"github.com/afosto/sendcloud-go"
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

func (c *Client) GetAddresses() ([]*sendcloud.Sender, error) {
	address := sendcloud.SenderResponseContainer{}
	_, err := sendcloud.Request("GET", "/api/v2/user/addresses/sender", nil, c.apiKey, c.apiSecret, &address)
	if err != nil {
		return nil, err
	}

	return address.GetResponse().([]*sendcloud.Sender), nil
}