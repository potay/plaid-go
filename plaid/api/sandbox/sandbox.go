package sandbox

import (
	"errors"

	"github.com/potay/plaid-go/plaid/api"
	"github.com/potay/plaid-go/plaid/client"
)

// Sandbox item endpoints.
type Item struct {
	*api.API
}

func NewItem(client *client.Client) *Item {
	return &Item{api.NewAPI(client)}
}

// Put an item into an ITEM_LOGIN_REQUIRED error state.
// :param  str     access_token:
func (s *Item) ResetLogin(accessToken string) (interface{}, error) {
	// return self.client.post('/sandbox/item/reset_login', {
	//     'access_token': access_token,
	// })
	return nil, errors.New("Not implemented")
}

// Sandbox-only endpoints.
// (`HTTP docs <https://plaid.com/docs/api/#sandbox>`__)
// These endpoints may not be used in other environments.
// .. autoclass:: plaid.api.sandbox.Item
//     :members:
type Sandbox struct {
	*api.API
	Item *Item
}

func NewSandbox(client *client.Client) *Sandbox {
	return &Sandbox{api.NewAPI(client), NewItem(client)}
}
