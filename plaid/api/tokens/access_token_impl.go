package tokens

import (
	"github.com/potay/plaid-go/plaid/api"
	"github.com/potay/plaid-go/plaid/client"
)

type accessToken struct {
	*api.API
}

func NewAccessToken(client *client.Client) AccessToken {
	return &accessToken{api.NewAPI(client)}
}

func (at *accessToken) Invalidate(accessToken string) (InvalidateResponse, error) {
	var respData InvalidateResponse
	err := at.Client.Post("/item/access_token/invalidate", struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: accessToken,
	}, &respData)
	if err != nil {
		return InvalidateResponse{}, err
	}

	return respData, nil
}

func (at *accessToken) UpdateVersion(accessToken string) (UpdateVersionResponse, error) {
	var respData UpdateVersionResponse
	err := at.Client.Post("/item/access_token/update_version", struct {
		AccessTokenV1 string `json:"access_token_v1"`
	}{
		AccessTokenV1: accessToken,
	}, &respData)
	if err != nil {
		return UpdateVersionResponse{}, err
	}

	return respData, nil
}
