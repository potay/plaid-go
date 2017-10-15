package identity

import (
	"github.com/potay/plaid-go/plaid/api"
	"github.com/potay/plaid-go/plaid/client"
)

type identity struct {
	*api.API
}

func NewIdentity(client *client.Client) Identity {
	return &identity{api.NewAPI(client)}
}

func (i *identity) Get(accessToken string) (GetResponse, error) {
	var respData GetResponse
	err := i.Client.Post("/identity/get", struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: accessToken,
	}, &respData)
	if err != nil {
		return GetResponse{}, err
	}

	return respData, nil
}
