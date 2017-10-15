package identity

import "github.com/potay/plaid-go/plaid/api"

// Identity endpoints.
// (`HTTP docs <https://plaid.com/docs/api/#identity>`__)
type Identity interface {
	// Retrieve information about an item.
	// :param  str     access_token:
	Get(accessToken string) (GetResponse, error)
}

type GetResponse struct {
	Accounts  []api.Account `json:"accounts"`
	Identity  api.Identity  `json:"identity"`
	Item      api.Item      `json:"item"`
	RequestID string        `json:"request_id"`
}
