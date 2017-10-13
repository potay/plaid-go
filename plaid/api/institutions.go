package api

import "github.com/potay/plaid-go/plaid/client"

// Institutions endpoints.
// (`HTTP docs <https://plaid.com/docs/api/#institutions>`__)
type Institutions struct {
	*API
}

type Institution struct {
	Credentials []struct {
		Label string `json:"label"`
		Name  string `json:"name"`
		Type  string `json:"type"`
	} `json:"credentials"`
	HasMfa        bool     `json:"has_mfa"`
	InstitutionID string   `json:"institution_id"`
	Mfa           []string `json:"mfa"`
	Name          string   `json:"name"`
	Products      []string `json:"products"`
}

func NewInstitutions(client *client.Client) *Institutions {
	return &Institutions{NewAPI(client)}
}

// Fetch all Plaid institutions, using /institutions/all.
// :param  int     count:      Number of institutions to fetch.
// :param  int     offset:     Number of institutions to skip.
func (a *Institutions) Get(count, offset int) ([]Institution, error) {
	var respData struct {
		Institutions []Institution `json:"institutions"`
		RequestID    string        `json:"request_id"`
		Total        int           `json:"total"`
	}
	err := a.Client.Post("/institutions/get", struct {
		Count  int `json:"count"`
		Offset int `json:"offset"`
	}{
		Count:  count,
		Offset: offset,
	}, &respData)
	if err != nil {
		return nil, err
	}

	return respData.Institutions, nil
}

// Fetch a single institution by id.
// :param  str     institution_id:
func (a *Institutions) GetByID(institutionID string) (Institution, error) {
	var respData struct {
		Institution `json:"institution"`
		RequestID   string `json:"request_id"`
	}
	err := a.Client.PostPublicKey("/institutions/get_by_id", struct {
		InstitutionID string `json:"institution_id"`
	}{
		InstitutionID: institutionID,
	}, &respData)
	if err != nil {
		return Institution{}, err
	}

	return respData.Institution, nil
}

// Search all institutions by name.
// :param  str      query:     Query against the full list of
//                             institutions.
// :param  [str]    products:  Filter FIs by available products. Optional.
func (a *Institutions) Search(query, products string) ([]Institution, error) {
	var respData struct {
		Institutions []Institution `json:"institutions"`
		RequestID    string        `json:"request_id"`
	}
	err := a.Client.PostPublicKey("/institutions/search", struct {
		Query    string `json:"query"`
		Products string `json:"products"`
	}{
		Query:    query,
		Products: products,
	}, &respData)
	if err != nil {
		return nil, err
	}

	return respData.Institutions, nil

}
