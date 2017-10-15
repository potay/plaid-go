package institutions

import (
	"github.com/potay/plaid-go/plaid/api"
	"github.com/potay/plaid-go/plaid/client"
)

type institutions struct {
	*api.API
}

func NewInstitutions(client *client.Client) Institutions {
	return &institutions{api.NewAPI(client)}
}

func (a *institutions) Get(count, offset int) (GetResponse, error) {
	var respData GetResponse
	err := a.Client.Post("/institutions/get", struct {
		Count  int `json:"count"`
		Offset int `json:"offset"`
	}{
		Count:  count,
		Offset: offset,
	}, &respData)
	if err != nil {
		return GetResponse{}, err
	}

	return respData, nil
}

func (a *institutions) GetByID(institutionID string) (GetByIDResponse, error) {
	var respData GetByIDResponse
	err := a.Client.PostPublicKey("/institutions/get_by_id", struct {
		InstitutionID string `json:"institution_id"`
	}{
		InstitutionID: institutionID,
	}, &respData)
	if err != nil {
		return GetByIDResponse{}, err
	}

	return respData, nil
}

func (a *institutions) Search(query, products string) (SearchResponse, error) {
	var respData SearchResponse
	err := a.Client.PostPublicKey("/institutions/search", struct {
		Query    string `json:"query"`
		Products string `json:"products"`
	}{
		Query:    query,
		Products: products,
	}, &respData)
	if err != nil {
		return SearchResponse{}, err
	}

	return respData, nil

}
