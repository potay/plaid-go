package income

import "github.com/potay/plaid-go/plaid/api"

// Income endpoints.
// (`HTTP docs <https://plaid.com/docs/api/#income>`__)
type Income interface {
	// Retrieve income data associated with an item.
	// Adds the income product to the item if it does not already have it.
	// :param str access_token:
	Get(accessToken string) (GetResponse, error)
}

type GetResponse struct {
	Item      api.Item   `json:"item"`
	Income    api.Income `json:"income"`
	RequestID string     `json:"request_id"`
}
