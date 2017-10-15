package accounts

import (
	"github.com/potay/plaid-go/plaid/api"
	"github.com/potay/plaid-go/plaid/api/balance"
	"github.com/potay/plaid-go/plaid/client"
)

type accounts struct {
	*api.API
	Balance balance.Balance
}

func NewAccounts(client *client.Client) Accounts {
	return &accounts{api.NewAPI(client), balance.NewBalance(client)}
}

func (a *accounts) Get(accessToken string, accountIDs []string) (GetResponse, error) {
	options := make(map[string]interface{})
	if len(accountIDs) > 0 {
		options["account_ids"] = accountIDs
	}

	var respData GetResponse
	err := a.Client.Post("/accounts/get", struct {
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
