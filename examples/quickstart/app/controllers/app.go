package controllers

import (
	"encoding/json"
	"time"

	"github.com/potay/plaid-go/plaid"
	"github.com/potay/plaid-go/plaid/api"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func makePlaid() *plaid.Plaid {
	clientID := revel.Config.StringDefault("plaid.client_id", "")
	secret := revel.Config.StringDefault("plaid.secret", "")
	publicKey := revel.Config.StringDefault("plaid.public_key", "")
	environment := revel.Config.StringDefault("plaid.env", "sandbox")
	revel.INFO.Printf("clientID: %s, secret: %s, publicKey: %s, environment: %s", clientID, secret, publicKey, environment)
	return plaid.NewPlaid(clientID, secret, publicKey, environment,
		false, time.Duration(600)*time.Second)
}

func (c App) Index() revel.Result {
	c.ViewArgs["plaid_environment"] = revel.Config.StringDefault("plaid.env", "sandbox")
	c.ViewArgs["plaid_public_key"] = revel.Config.StringDefault("plaid.public_key", "")
	revel.INFO.Printf("publicKey: %s, environment: %s", c.ViewArgs["plaid_public_key"], c.ViewArgs["plaid_environment"])
	return c.Render()
}

func (c App) GetAccessToken() revel.Result {
	revel.TRACE.Printf("Performing an Accounts Request...")
	client := makePlaid()

	publicToken := c.Params.Form.Get("public_token")
	revel.TRACE.Printf("publicToken: %s", publicToken)
	exchangeResponse, err := client.PublicToken.Exchange(publicToken)
	if err != nil {
		revel.WARN.Printf("err: %s", err)
		return c.RenderError(err)
	}

	revel.TRACE.Printf("access token: %s", exchangeResponse.AccessToken)
	revel.TRACE.Printf("item ID: %s", exchangeResponse.ItemID)

	c.Session["access_token"] = exchangeResponse.AccessToken

	jsonStr, _ := json.MarshalIndent(exchangeResponse, "", "    ")
	revel.TRACE.Printf("response: %s", jsonStr)
	return c.RenderJSON(exchangeResponse)
}

func (c App) Accounts() revel.Result {
	revel.TRACE.Printf("Performing an Accounts Request...")
	client := makePlaid()

	accessToken := c.Session["access_token"]
	revel.TRACE.Printf("accessToken: %s", accessToken)
	response, err := client.Auth.Get(accessToken, nil)
	if err != nil {
		revel.WARN.Printf("err: %s", err)
		return c.RenderError(err)
	}

	jsonStr, _ := json.MarshalIndent(response, "", "    ")
	revel.TRACE.Printf("response: %s", jsonStr)
	return c.RenderJSON(response)
}

func (c App) Item() revel.Result {
	revel.TRACE.Printf("Performing an Item Request...")
	client := makePlaid()

	accessToken := c.Session["access_token"]
	revel.TRACE.Printf("accessToken: %s", accessToken)
	itemResponse, err := client.Item.Get(accessToken)
	if err != nil {
		revel.WARN.Printf("err: %s", err)
		return c.RenderError(err)
	}

	institutionResponse, err := client.Institutions.GetByID(itemResponse.Item.InstitutionID)
	if err != nil {
		revel.WARN.Printf("err: %s", err)
		return c.RenderError(err)
	}

	jsonStr, err := json.MarshalIndent(struct {
		Item        api.Item        `json:"item"`
		Institution api.Institution `json:"institution"`
	}{
		Item:        itemResponse.Item,
		Institution: institutionResponse.Institution,
	}, "", "    ")
	revel.TRACE.Printf("response: %s", jsonStr)
	return c.RenderJSON(struct {
		Item        api.Item        `json:"item"`
		Institution api.Institution `json:"institution"`
	}{
		Item:        itemResponse.Item,
		Institution: institutionResponse.Institution,
	})
}

func (c App) Transactions() revel.Result {
	revel.TRACE.Printf("Performing a Transaction Request...")
	client := makePlaid()

	startDate := time.Now().AddDate(0, 0, -30)
	endDate := time.Now()

	accessToken := c.Session["access_token"]
	revel.TRACE.Printf("startDate: %s, endDate: %s", startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))
	revel.TRACE.Printf("accessToken: %s", accessToken)
	response, err := client.Transactions.Get(accessToken, startDate, endDate, nil, 500, 0)
	if err != nil {
		revel.WARN.Printf("err: %s", err)
		return c.RenderError(err)
	}

	jsonStr, _ := json.MarshalIndent(response, "", "    ")
	revel.TRACE.Printf("response: %s", jsonStr)
	return c.RenderJSON(response)
}

func (c App) CreatePublicToken() revel.Result {
	revel.TRACE.Printf("Performing a Create Public Token Request...")
	client := makePlaid()

	accessToken := c.Session["access_token"]
	revel.TRACE.Printf("accessToken: %s", accessToken)
	response, err := client.PublicToken.Create(accessToken)
	if err != nil {
		revel.WARN.Printf("err: %s", err)
		return c.RenderError(err)
	}

	jsonStr, _ := json.MarshalIndent(response, "", "    ")
	revel.TRACE.Printf("response: %s", jsonStr)
	return c.RenderJSON(response)
}
