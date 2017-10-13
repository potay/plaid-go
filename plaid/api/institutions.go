package api

import "github.com/potay/plaid-go/plaid/client"

// Institutions endpoints.
// (`HTTP docs <https://plaid.com/docs/api/#institutions>`__)
type Institutions struct {
	*API
}

func NewInstitutions(client *client.Client) *Institutions {
	return &Institutions{NewAPI(client)}
}

// Fetch all Plaid institutions, using /institutions/all.
// :param  int     count:      Number of institutions to fetch.
// :param  int     offset:     Number of institutions to skip.
func (a *Institutions) Get(count, offset int) (map[string]interface{}, error) {
	return a.Client.Post("/institutions/get", struct {
		Count  int `json:"count"`
		Offset int `json:"offset"`
	}{
		Count:  count,
		Offset: offset,
	})
}

// Fetch a single institution by id.
// :param  str     institution_id:
func (a *Institutions) GetByID(institutionID string) (map[string]interface{}, error) {
	return a.Client.PostPublicKey("/institutions/get_by_id", struct {
		InstitutionID string `json:"institution_id"`
	}{
		InstitutionID: institutionID,
	})
}

// Search all institutions by name.
// :param  str      query:     Query against the full list of
//                             institutions.
// :param  [str]    products:  Filter FIs by available products. Optional.
func (a *Institutions) Search(query, products string) (map[string]interface{}, error) {

	return a.Client.PostPublicKey("/institutions/search", struct {
		Query    string `json:"query"`
		Products string `json:"products"`
	}{
		Query:    query,
		Products: products,
	})
}
