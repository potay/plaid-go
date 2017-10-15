package credentials

import "github.com/potay/plaid-go/plaid/api"

// Item Credentials endpoints.
type Credentials interface {
	// Update the credentials associated with an item.
	// Returns either a successful response or a response indicating that MFA
	// (Multi Factor Authentication) is required.
	// :param  dict     credentials:   See ``Item.create`` for details.
	Update(accessToken string, credentials api.Credentials) (UpdateResponse, error)
}

type UpdateResponse struct{}
