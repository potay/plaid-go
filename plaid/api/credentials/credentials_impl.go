package credentials

import (
	"github.com/potay/plaid-go/plaid/api"
	"github.com/potay/plaid-go/plaid/client"
)

type credentials struct {
	*api.API
}

func NewCredentials(client *client.Client) Credentials {
	return &credentials{api.NewAPI(client)}
}

func (c *credentials) Update(accessToken string, credentials api.Credentials) (UpdateResponse, error) {
	var respData UpdateResponse
	err := c.Client.Post("/item/credentials/update", struct {
		AccessToken string          `json:"access_token"`
		Credentials api.Credentials `json:"credentials"`
	}{
		AccessToken: accessToken,
		Credentials: credentials,
	}, &respData)
	if err != nil {
		return UpdateResponse{}, err
	}

	return respData, nil
}
