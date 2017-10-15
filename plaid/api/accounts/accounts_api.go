package accounts

import "github.com/potay/plaid-go/plaid/api"

// Accounts endpoints.
// (`HTTP docs <https://plaid.com/docs/api/#accounts>`__)
// .. autoclass:: plaid.api.accounts.Balance
//     :members:
type Accounts interface {
	// Retrieve high-level account information for an Item.
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
