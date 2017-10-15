package tokens

// Endpoints for translating between public tokens and access tokens.
type PublicToken interface {
	// Exchange a Link public_token for an API access_token.
	// (`HTTP docs <https://plaid.com/docs/api/#exchange-token-flow>`__)
	// :param  str     public_token:
	Exchange(publicToken string) (ExchangeResponse, error)
	// Create a Link public_token for an API access_token.
	// (`HTTP docs <https://plaid.com/docs/api/#create-public-token>`__)
	// :param  str     access_token:
	Create(accessToken string) (CreateResponse, error)
}

type ExchangeResponse struct {
	AccessToken string `json:"access_token"`
	ItemID      string `json:"item_id"`
	RequestID   string `json:"request_id"`
}

type CreateResponse struct {
	PublicToken string `json:"public_token"`
	RequestID   string `json:"request_id"`
}
