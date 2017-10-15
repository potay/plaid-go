package tokens

// Access token endpoints.
type AccessToken interface {
	// Rotate the access token for an item.
	// (`HTTP docs <https://plaid.com/docs/api/#rotate-access-token>`__)
	// :param  str     access_token:
	Invalidate(accessToken string) (InvalidateResponse, error)
	// Transition an access token to work with the current version of
	// the Plaid API
	// (`HTTP docs <https://plaid.com/docs/api/#update-access-token-
	// version>`__)
	// :param  str      access_token:
	UpdateVersion(accessToken string) (UpdateVersionResponse, error)
}

type InvalidateResponse struct {
	NewAccessToken string `json:"new_access_token"`
	RequestID      string `json:"request_id"`
}
type UpdateVersionResponse struct {
	AccessToken string `json:"access_token"`
	ItemID      string `json:"item_id"`
	RequestID   string `json:"request_id"`
}
