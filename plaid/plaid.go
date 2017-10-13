package plaid

import (
	"time"

	"github.com/potay/plaid-go/plaid/api"
	"github.com/potay/plaid-go/plaid/client"
)

type Plaid struct {
	client       *client.Client
	Institutions *api.Institutions
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
		api.NewInstitutions(client),
	}

	return plaid
}
