package auth

import (
	"github.com/potay/plaid-go/plaid/api"
	"github.com/potay/plaid-go/plaid/client"
)

type auth struct {
	*api.API
}

func NewAuth(client *client.Client) Auth {
	return &auth{api.NewAPI(client)}
}

func (a *auth) Get(accessToken string, accountIDs []string) (GetResponse, error) {
	options := make(map[string]interface{})
	if len(accountIDs) > 0 {
		options["account_ids"] = accountIDs
	}

	var respData GetResponse
	err := a.Client.Post("/auth/get", struct {
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
