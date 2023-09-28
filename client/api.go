package client

import (
	"github.com/lpieri/sendcloud/integration"
	"github.com/lpieri/sendcloud/method"
	"github.com/lpieri/sendcloud/parcel"
	"github.com/lpieri/sendcloud/sender"
	"github.com/lpieri/sendcloud/servicepoint"
)

type API struct {
	Parcel       *parcel.Client
	Method       *method.Client
	Sender       *sender.Client
	ServicePoint *servicepoint.Client
	Integration  *integration.Client
}

// Initialize the client
func (a *API) Init(apiKey string, apiSecret string) {
	a.Parcel = parcel.New(apiKey, apiSecret)
	a.Method = method.New(apiKey, apiSecret)
	a.Sender = sender.New(apiKey, apiSecret)
	a.ServicePoint = servicepoint.New(apiKey, apiSecret)
	a.Integration = integration.New(apiKey, apiSecret)
}
