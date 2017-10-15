package income

import (
	"github.com/potay/plaid-go/plaid/api"
	"github.com/potay/plaid-go/plaid/client"
)

type income struct {
	*api.API
}

func NewIncome(client *client.Client) Income {
	return &income{api.NewAPI(client)}
}

func (i *income) Get(accessToken string) (GetResponse, error) {
	var respData GetResponse
	err := i.Client.Post("/income/get", struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: accessToken,
	}, &respData)
	if err != nil {
		return GetResponse{}, err
	}

	return respData, nil
}
