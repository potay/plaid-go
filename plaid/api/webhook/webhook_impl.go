package webhook

import (
	"github.com/potay/plaid-go/plaid/api"
	"github.com/potay/plaid-go/plaid/client"
)

type webhook struct {
	*api.API
}

func NewWebhook(client *client.Client) Webhook {
	return &webhook{api.NewAPI(client)}
}

func (w *webhook) Update(accessToken, webhook string) (UpdateResponse, error) {
	var respData UpdateResponse
	err := w.Client.Post("/item/webhook/update", struct {
		AccessToken string `json:"access_token"`
		Webhook     string `json:"webhook"`
	}{
		AccessToken: accessToken,
		Webhook:     webhook,
	}, &respData)
	if err != nil {
		return UpdateResponse{}, err
	}

	return respData, nil
}
