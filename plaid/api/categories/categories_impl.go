package categories

import (
	"github.com/potay/plaid-go/plaid/api"
	"github.com/potay/plaid-go/plaid/client"
)

type categories struct {
	*api.API
}

func NewCategories(client *client.Client) Categories {
	return &categories{api.NewAPI(client)}
}

func (c *categories) Get() (GetResponse, error) {
	var respData GetResponse
	err := c.Client.PostPublic("/categories/get",
		nil,
		&respData)
	if err != nil {
		return GetResponse{}, err
	}

	return respData, nil
}
