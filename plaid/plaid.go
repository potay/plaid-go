// Package plaid implements a Go client for the Plaid API (https://plaid.com/docs)
package plaid

import (
	"time"

	"github.com/potay/plaid-go/plaid/api/accounts"
	"github.com/potay/plaid-go/plaid/api/auth"
	"github.com/potay/plaid-go/plaid/api/balance"
	"github.com/potay/plaid-go/plaid/api/categories"
	"github.com/potay/plaid-go/plaid/api/credentials"
	"github.com/potay/plaid-go/plaid/api/credit"
	"github.com/potay/plaid-go/plaid/api/identity"
	"github.com/potay/plaid-go/plaid/api/income"
	"github.com/potay/plaid-go/plaid/api/institutions"
	"github.com/potay/plaid-go/plaid/api/item"
	"github.com/potay/plaid-go/plaid/api/processor"
	"github.com/potay/plaid-go/plaid/api/sandbox"
	"github.com/potay/plaid-go/plaid/api/tokens"
	"github.com/potay/plaid-go/plaid/api/transactions"
	"github.com/potay/plaid-go/plaid/api/webhook"
	"github.com/potay/plaid-go/plaid/client"
)

// Go Plaid API client.
// See official documentation at: https://plaid.com/docs.
// All of the endpoints documented under the ``plaid.api``
// module may be called from a ``plaid.Plaid`` instance.

type Plaid struct {
	client *client.Client

	// Mirror the HTTP API hierarchy
	AccessToken  tokens.AccessToken
	Accounts     accounts.Accounts
	Auth         auth.Auth
	Balance      balance.Balance
	Categories   categories.Categories
	Credentials  credentials.Credentials
	Credit       credit.Credit
	Identity     identity.Identity
	Income       income.Income
	Institutions institutions.Institutions
	Item         item.Item
	Processor    processor.Processor
	PublicToken  tokens.PublicToken
	Sandbox      *sandbox.Sandbox
	Transactions transactions.Transactions
	Webhook      webhook.Webhook
}

func NewPlaid(clientID, secret, publicKey string, environment string, suppressWarnings bool, timeout time.Duration) *Plaid {
	// Initialize a client with credentials.
	// :param  str     client_id:          Your Plaid client ID
	// :arg    str     secret:             Your Plaid secret
	// :arg    str     public_key:         Your Plaid public key
	// :arg    str     environment:        One of ``sandbox``,
	//                                     ``development``, or ``production``.
	// :arg    bool    suppress_warnings:  Suppress Plaid warnings.
	// :arg    int     timeout:            Timeout for API requests.
	client := client.NewClient(clientID, secret, publicKey, environment, suppressWarnings, timeout)
	plaid := &Plaid{
		client,
		tokens.NewAccessToken(client),
		accounts.NewAccounts(client),
		auth.NewAuth(client),
		balance.NewBalance(client),
		categories.NewCategories(client),
		credentials.NewCredentials(client),
		credit.NewCredit(client),
		identity.NewIdentity(client),
		income.NewIncome(client),
		institutions.NewInstitutions(client),
		item.NewItem(client),
		processor.NewProcessor(client),
		tokens.NewPublicToken(client),
		sandbox.NewSandbox(client),
		transactions.NewTransactions(client),
		webhook.NewWebhook(client),
	}

	return plaid
}
