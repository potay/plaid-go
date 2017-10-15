package institutions

import "github.com/potay/plaid-go/plaid/api"

// Institutions endpoints.
// (`HTTP docs <https://plaid.com/docs/api/#institutions>`__)
type Institutions interface {
	// Fetch all Plaid institutions, using /institutions/all.
	// :param  int     count:      Number of institutions to fetch.
	// :param  int     offset:     Number of institutions to skip.
	Get(count, offset int) (GetResponse, error)
	// Fetch a single institution by id.
	// :param  str     institution_id:
	GetByID(institutionID string) (GetByIDResponse, error)
	// Search all institutions by name.
	// :param  str      query:     Query against the full list of
	//                             institutions.
	// :param  [str]    products:  Filter FIs by available products. Optional.
	Search(query, products string) (SearchResponse, error)
}

type GetResponse struct {
	Institutions []api.Institution `json:"institutions"`
	RequestID    string            `json:"request_id"`
	Total        int               `json:"total"`
}

type GetByIDResponse struct {
	Institution api.Institution `json:"institution"`
	RequestID   string          `json:"request_id"`
}
type SearchResponse struct {
	Institutions []api.Institution `json:"institutions"`
	RequestID    string            `json:"request_id"`
}
