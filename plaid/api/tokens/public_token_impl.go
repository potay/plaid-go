package tokens

import (
	"github.com/potay/plaid-go/plaid/api"
	"github.com/potay/plaid-go/plaid/client"
)

type publicToken struct {
	*api.API
}

func NewPublicToken(client *client.Client) PublicToken {
	return &publicToken{api.NewAPI(client)}
}

func (pt *publicToken) Exchange(publicToken string) (ExchangeResponse, error) {
	var respData ExchangeResponse
	err := pt.Client.Post("/item/public_token/exchange", struct {
		PublicToken string `json:"public_token"`
	}{
		PublicToken: publicToken,
	}, &respData)
	if err != nil {
		return ExchangeResponse{}, err
	}

	return respData, nil
}

func (pt *publicToken) Create(accessToken string) (CreateResponse, error) {
	var respData CreateResponse
	err := pt.Client.Post("/item/public_token/create", struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: accessToken,
	}, &respData)
	if err != nil {
		return CreateResponse{}, err
	}

	return respData, nil
}
