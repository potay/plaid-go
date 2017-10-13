package api

import "github.com/potay/plaid-go/plaid/client"

type API struct {
	Client *client.Client
}

func NewAPI(client *client.Client) *API {
	return &API{
		Client: client,
	}
}
