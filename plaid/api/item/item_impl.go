package item

import (
	"errors"
	"time"

	"github.com/potay/plaid-go/plaid/api"
	"github.com/potay/plaid-go/plaid/client"
)

type item struct {
	*api.API
}

func NewItem(client *client.Client) Item {
	return &item{
		api.NewAPI(client),
	}
}

func (i *item) Create(credentials api.Credentials,
	institutionID string, initialProducts []string,
	transactionsStartDate, transactionsEndDate time.Time,
	transactionsAwaitResults bool,
	webhook string) (CreateResponse, error) {

	transactionOptions := make(map[string]interface{})
	if !transactionsStartDate.IsZero() {
		transactionOptions["start_date"] = transactionsStartDate
	}
	if !transactionsEndDate.IsZero() {
		transactionOptions["end_date"] = transactionsEndDate
	}
	transactionOptions["await_results"] = transactionsAwaitResults

	options := make(map[string]interface{})
	if len(transactionOptions) > 0 {
		options["transactions"] = transactionOptions
	}
	if webhook != "" {
		options["webhook"] = webhook
	}

	var respData CreateResponse
	err := i.Client.Post("/item/create", struct {
		Credentials     api.Credentials        `json:"credentials"`
		InstitutionID   string                 `json:"institution_id"`
		InitialProducts []string               `json:"initial_products"`
		Options         map[string]interface{} `json:"options"`
	}{
		Credentials:     credentials,
		InstitutionID:   institutionID,
		InitialProducts: initialProducts,
		Options:         options,
	}, nil)
	if err != nil {
		return CreateResponse{}, err
	}

	return respData, errors.New("Not implemented")
}

func (i *item) MFA(accessToken, mfaType string, responses []string) (MFAResponse, error) {
	var respData MFAResponse
	err := i.Client.Post("/item/mfa", struct {
		AccessToken string   `json:"access_token"`
		MfaType     string   `json:"mfa_type"`
		Responses   []string `json:"responses"`
	}{
		AccessToken: accessToken,
		MfaType:     mfaType,
		Responses:   responses,
	}, &respData)
	if err != nil {
		return MFAResponse{}, err
	}

	return MFAResponse{}, errors.New("Not implemented")
}

func (i *item) Get(accessToken string) (GetResponse, error) {
	var respData GetResponse
	err := i.Client.Post("/item/get", struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: accessToken,
	}, &respData)
	if err != nil {
		return GetResponse{}, err
	}

	return respData, nil
}

func (i *item) Delete(accessToken string) (DeleteResponse, error) {
	var respData DeleteResponse
	err := i.Client.Post("/item/delete", struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: accessToken,
	}, &respData)
	if err != nil {
		return DeleteResponse{}, err
	}

	return respData, nil
}
