package transactions

import (
	"time"

	"github.com/potay/plaid-go/plaid/api"
	"github.com/potay/plaid-go/plaid/client"
)

type transactions struct {
	*api.API
}

func NewTransactions(client *client.Client) Transactions {
	return &transactions{api.NewAPI(client)}
}

func (t *transactions) Get(accessToken string,
	startDate, endDate time.Time,
	accountIDs []string,
	count, offset int) (GetResponse, error) {
	options := make(map[string]interface{})
	if len(accountIDs) > 0 {
		options["account_ids"] = accountIDs
	}
	if count > 0 {
		options["count"] = count
	}
	if offset > 0 {
		options["offset"] = offset
	}

	var respData GetResponse
	err := t.Client.Post("/transactions/get", struct {
		AccessToken string                 `json:"access_token"`
		StartDate   string                 `json:"start_date"`
		EndDate     string                 `json:"end_date"`
		Options     map[string]interface{} `json:"options"`
	}{
		AccessToken: accessToken,
		StartDate:   startDate.Format("2006-01-02"),
		EndDate:     startDate.Format("2006-01-02"),
		Options:     options,
	}, &respData)
	if err != nil {
		return GetResponse{}, err
	}

	return respData, nil
}
