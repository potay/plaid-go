package auth

import "github.com/potay/plaid-go/plaid/api"

// Auth endpoints.
type Auth interface {
	// Retrieve account and routing numbers for checking and savings accounts.
	// (`HTTP docs <https://plaid.com/docs/api/#auth>`__)
	// :param  str     access_token:
	// :param  [str]   account_ids:    A list of account_ids to retrieve for
	//                                 the item. Optional.
	Get(accessToken string, accountIDs []string) (GetResponse, error)
}

type GetResponse struct {
	Accounts  []api.Account `json:"accounts"`
	Numbers   []api.Number  `json:"numbers"`
	Item      api.Item      `json:"item"`
	RequestID string        `json:"request_id"`
}
