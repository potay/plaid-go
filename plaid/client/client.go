// Package plaid implements a Go client for the Plaid API (https://plaid.com/docs)
package client

import (
	"reflect"
	"time"

	log "github.com/sirupsen/logrus"
)

// Go Plaid API client.
// See official documentation at: https://plaid.com/docs.
// All of the endpoints documented under the ``plaid.api``
// module may be called from a ``plaid.Client`` instance.

type Client struct {
	clientID         string
	secret           string
	publicKey        string
	environment      string
	suppressWarnings bool
	timeout          time.Duration
	// Mirror the HTTP API hierarchy
}

func NewClient(clientID, secret, publicKey string, environment string, suppressWarnings bool, timeout time.Duration) *Client {
	// Initialize a client with credentials.
	// :param  str     client_id:          Your Plaid client ID
	// :arg    str     secret:             Your Plaid secret
	// :arg    str     public_key:         Your Plaid public key
	// :arg    str     environment:        One of ``sandbox``,
	//                                     ``development``, or ``production``.
	// :arg    bool    suppress_warnings:  Suppress Plaid warnings.
	// :arg    int     timeout:            Timeout for API requests.
	if environment == "development" && !suppressWarnings {
		log.Warn(`
                Development is not intended for production usage.
                Swap out url for https://api.plaid.com
                via Client.config before switching to production
            `)
	}
	client := &Client{
		clientID,
		secret,
		publicKey,
		environment,
		suppressWarnings,
		timeout,
	}

	return client
}

func (c *Client) Post(path string, data interface{}, respData interface{}) error {
	// Make a post request with clientID and secret key.
	postDataType := reflect.StructOf([]reflect.StructField{
		{
			Name:      "Data",
			Type:      reflect.TypeOf(data),
			Anonymous: true,
		},
		{
			Name: "ClientID",
			Type: reflect.TypeOf(c.clientID),
			Tag:  `json:"client_id"`,
		},
		{
			Name: "Secret",
			Type: reflect.TypeOf(c.secret),
			Tag:  `json:"secret"`,
		},
	})

	postData := reflect.New(postDataType).Elem()
	postData.Field(0).Set(reflect.ValueOf(data))
	postData.FieldByName("ClientID").SetString(c.clientID)
	postData.FieldByName("Secret").SetString(c.secret)

	return c.post(path, postData.Interface(), respData)
}

func (c *Client) PostPublic(path string, data interface{}, respData interface{}) error {
	// Make a post request requiring no auth.
	return c.post(path, data, respData)
}

func (c *Client) PostPublicKey(path string, data interface{}, respData interface{}) error {
	// Make a post request using a public key.
	postDataType := reflect.StructOf([]reflect.StructField{
		{
			Name:      "Data",
			Type:      reflect.TypeOf(data),
			Anonymous: true,
		},
		{
			Name: "PublicKey",
			Type: reflect.TypeOf(c.publicKey),
			Tag:  `json:"public_key"`,
		},
	})

	postData := reflect.New(postDataType).Elem()
	postData.Field(0).Set(reflect.ValueOf(data))
	postData.FieldByName("PublicKey").SetString(c.publicKey)

	return c.post(path, postData.Interface(), respData)
}

func (c *Client) post(path string, data interface{}, respData interface{}) error {
	return postRequest(
		"https://"+string(c.environment)+".plaid.com"+string(path),
		data,
		c.timeout,
		respData)
}
