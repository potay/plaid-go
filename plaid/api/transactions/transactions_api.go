package transactions

import (
	"time"

	"github.com/potay/plaid-go/plaid/api"
)

// Transactions endpoints.
type Transactions interface {
	// Return accounts and transactions for an item.
	// (`HTTP docs <https://plaid.com/docs/api/#transactions>`__)
	// The transactions in the response are paginated -- compare the number of
	// transactions received so far against response['total_transactions'] to
	// determine whether to fetch another page.
	// :param  str     access_token:
	// :param  str     start_date:     The earliest date for transactions.
	// :param  str     end_date:       The latest date for transactions.
	// :param  [str]   account_ids:    A list of account_ids to retrieve for
	//                                 the item. Optional.
	// :param  int     count:          The number of transactions to fetch.
	//                                 Optional.
	// :param  int     offset:         The number of transactions to skip from
	//                                 the beginning of the fetch. Optional.
	// All date should be formatted as ``YYYY-MM-DD``.
	Get(accessToken string,
		startDate, endDate time.Time,
		accountIDs []string,
		count, offset int) (GetResponse, error)
}

type GetResponse struct {
	Accounts          []api.Account     `json:"accounts"`
	Transactions      []api.Transaction `json:"transactions"`
	Item              api.Item          `json:"item"`
	TotalTransactions int               `json:"total_transactions"`
	RequestID         string            `json:"request_id"`
}
