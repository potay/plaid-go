package balance

import (
	"github.com/potay/plaid-go/plaid/api"
	"github.com/potay/plaid-go/plaid/client"
)

type balance struct {
	*api.API
}

func NewBalance(client *client.Client) Balance {
	return &balance{api.NewAPI(client)}
}

func (b *balance) Get(accessToken string, accountIDs []string) (GetResponse, error) {
	options := make(map[string]interface{})
	if len(accountIDs) > 0 {
		options["account_ids"] = accountIDs
	}

	var respData GetResponse
	err := b.Client.Post("/accounts/balance/get", struct {
		AccessToken string                 `json:"access_token"`
		Options     map[string]interface{} `json:"options"`
	}{
		AccessToken: accessToken,
		Options:     options,
	}, &respData)
	if err != nil {
		return GetResponse{}, err
	}

	return respData, nil
}
