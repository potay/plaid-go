package processor

import (
	"errors"

	"github.com/potay/plaid-go/plaid/api"
	"github.com/potay/plaid-go/plaid/client"
)

type processor struct {
	*api.API
}

func NewProcessor(client *client.Client) Processor {
	return &processor{api.NewAPI(client)}
}

func (p *processor) StripeBankAccountTokenCreate(accessToken, accountID string) (StripeBankAccountTokenCreateResponse, error) {
	// return self.client.post('/processor/stripe/bank_account_token/create',
	//                               {
	//                                   'access_token': access_token,
	//                                   'account_id': account_id,
	return StripeBankAccountTokenCreateResponse{}, errors.New("Not implemented")
}
