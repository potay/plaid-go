package processor

// Endpoints for creating processor tokens.
type Processor interface {
	// Create a Stripe bank_account token for a given account ID
	// (`HTTP docs <https://plaid.com/docs/link/stripe>`__)
	// :param  str     public_token:
	// :param  str     account_id:
	StripeBankAccountTokenCreate(accessToken, accountID string) (StripeBankAccountTokenCreateResponse, error)
}

type StripeBankAccountTokenCreateResponse struct{}
