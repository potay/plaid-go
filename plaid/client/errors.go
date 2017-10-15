package client

import (
	"fmt"
)

type plaidClientError struct {
	// List of all errors: https://plaid.com/docs/api/#errors-overview
	ErrorType      string `json:"error_type"`
	ErrorCode      string `json:"error_code"`
	ErrorMessage   string `json:"error_message"`
	DisplayMessage string `json:"display_message"`

	// StatusCode needs to manually set from the http response
	StatusCode int
}

func (e plaidClientError) Error() string {
	return fmt.Sprintf("Plaid Error - http status: %d, type: %s, code: %s, message: %s, resolve: %s",
		e.StatusCode, e.ErrorType, e.ErrorCode, e.ErrorMessage, e.DisplayMessage)
}
