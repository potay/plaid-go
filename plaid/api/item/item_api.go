package item

import (
	"time"

	"github.com/potay/plaid-go/plaid/api"
)

// Item endpoints.
// (`HTTP docs <https://plaid.com/docs/api/#item-management>`__)
// .. autoclass:: plaid.api.item.AccessToken
//     :members:
// .. autoclass:: plaid.api.item.PublicToken
//     :members:
// .. autoclass:: plaid.api.item.Credentials
//     :members:
// .. autoclass:: plaid.api.item.Webhook
//     :members:
type Item interface {
	// Add a bank account user/login to Plaid.
	// Returns either a successful response or a response indicating that MFA
	// (Multi Factor Authentication) is required.
	// :param  dict    credentials:                    A dictionary containing
	//                                                 the fields
	//                                                 ``username``,
	//                                                 ``password``, and
	//                                                 (optionally) ``pin``.
	// :param  str     institution_id:
	// :param  list    initial_products:               A list containing
	//                                                 product names.
	// :param  str     transactions__start_date:       The date to begin the
	//                                                 item's initial
	//                                                 transaction pull.
	// :param  str     transactions__end_date:         The date to end the
	//                                                 item's initial
	//                                                 transaction pull.
	// :param  str     transactions__await_results:    If ``True``, wait for
	//                                                 the initial
	//                                                 transaction pull to
	//                                                 complete before
	//                                                 returning. Will
	//                                                 increase the user wait
	//                                                 time.
	// :param  str     webhook:                        The URL for a webhook
	//                                                 associated with the
	//                                                 item.
	// All dates should be formatted as ``YYYY-MM-DD``.
	Create(credentials api.Credentials,
		institutionID string, initialProducts []string,
		transactionsStartDate, transactionsEndDate time.Time,
		transactionsAwaitResults bool,
		webhook string) (CreateResponse, error)
	// Provide an MFA (Multi-Factor Authentication) response for an item.
	// :param  str     access_token:
	// :param  str     mfa_type:       The type of mfa answered (e.g. device)
	// :param  [str]   responses:      The MFA response(s)
	MFA(accessToken, mfaType string, responses []string) (MFAResponse, error)
	// Get information about the status of an item.
	// (`HTTP docs <https://plaid.com/docs/api/#get-item>`__)
	// :param  str     access_token:
	Get(accessToken string) (GetResponse, error)
	// Delete an item.
	// (`HTTP docs <https://plaid.com/docs/api/#delete-an-item>`__)
	// This also deactivates the access_token.
	// :param  str     access_token:
	Delete(accessToken string) (DeleteResponse, error)
}

type CreateResponse struct{}

type MFAResponse struct{}

type GetResponse struct {
	Item      api.Item `json:"item"`
	RequestID string   `json:"request_id"`
}

type DeleteResponse struct {
	Deleted   bool   `json:"deleted"`
	RequestID string `json:"request_id"`
}
