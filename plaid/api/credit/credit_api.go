package credit

// Credit details endpoints.
// (`HTTP docs <https://plaid.com/docs/api/#credit-details>`__)
type Credit interface {
	// Retrieve credit details associated with an item.
	// :param  str     access_token:
	Get(accessToken string) (GetResponse, error)
}

type GetResponse struct{}
