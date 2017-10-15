package categories

import "github.com/potay/plaid-go/plaid/api"

// Categories endpoints.
//     (`HTTP docs <https://plaid.com/docs/api/#categories>`__)
type Categories interface {
	// Fetch all plaid categories.
	Get() (GetResponse, error)
}

type GetResponse struct {
	Categories []api.Category `json:"categories"`
	RequestID  string         `json:"request_id"`
}
