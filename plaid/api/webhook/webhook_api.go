package webhook

import "github.com/potay/plaid-go/plaid/api"

// Webhook endpoints.
type Webhook interface {
	// Update the webhook for an Item.
	// (`HTTP docs <https://plaid.com/docs/api/#update-webhook>`__)
	// :param  str     access_token:
	// :param  str     webhook: The URL of the webhook to associate.
	Update(accessToken, webhook string) (UpdateResponse, error)
}

type UpdateResponse struct {
	Item      api.Item `json:"item"`
	RequestID string   `json:"request_id"`
}
