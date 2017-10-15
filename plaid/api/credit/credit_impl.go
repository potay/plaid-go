package credit

import (
	"github.com/potay/plaid-go/plaid/api"
	"github.com/potay/plaid-go/plaid/client"
)

type credit struct {
	*api.API
}

func NewCredit(client *client.Client) Credit {
	return &credit{api.NewAPI(client)}
}

func (cd *credit) Get(accessToken string) (GetResponse, error) {
	var respData GetResponse
	err := cd.Client.Post("/credit_details/get", struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: accessToken,
	}, &respData)
	if err != nil {
		return GetResponse{}, err
	}

	return respData, nil
}
