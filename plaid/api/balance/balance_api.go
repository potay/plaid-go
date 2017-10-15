package balance

import "github.com/potay/plaid-go/plaid/api"

// Accounts balance endpoint.
type Balance interface {
	// Retrieve real-time balance information for accounts.
	// :param  str     accessToken:
	// :param  [str]   accountIDs:    A list of account_ids to retrieve for
	//                                 the item. Optional.
	Get(accessToken string, accountIDs []string) (GetResponse, error)
}

type GetResponse struct {
	Accounts  []api.Account `json:"accounts"`
	Item      api.Item      `json:"item"`
	RequestID string        `json:"request_id"`
}
