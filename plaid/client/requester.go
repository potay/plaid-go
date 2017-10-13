package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var ALLOWED_METHODS = []string{"POST"}
var DEFAULT_TIMEOUT = 600 // 10 minutes

func requestsHttpRequest(url, method string, data interface{}, timeout time.Duration) (*http.Response, error) {
	normalizedMethod := strings.ToUpper(method)
	for _, value := range ALLOWED_METHODS {
		if value == normalizedMethod {
			jsonData, err := json.Marshal(data)
			request, err := http.NewRequest(normalizedMethod, url, bytes.NewBuffer(jsonData))
			if err != nil {
				return nil, err
			}
			request.Header.Add("Content-Type", "application/json")
			request.Header.Add("User-Agent", "Plaid Go v"+version)

			return (&http.Client{Timeout: timeout}).Do(request)
		}
	}
	return nil, fmt.Errorf("Invalid request method %s", method)
}

func httpRequest(url, method string, data interface{}, timeout time.Duration) (interface{}, error) {
	response, err := requestsHttpRequest(url, method, data, timeout)
	if err != nil {
		return nil, err
	}

	raw, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	switch {
	case response.StatusCode == 200:
		// Successful response.
		var respData interface{}
		if err = json.Unmarshal(raw, &respData); err != nil {
			return nil, err
		}
		return respData, nil
	case response.StatusCode >= 400:
		// Attempt to unmarshal into Plaid error format
		var plaidErr plaidClientError
		if err = json.Unmarshal(raw, &plaidErr); err != nil {
			return nil, err
		}
		plaidErr.StatusCode = response.StatusCode
		return nil, plaidErr
	}

	return nil, errors.New("Unknown Plaid Error - Status:" + string(response.StatusCode))
}

func postRequest(url string, data interface{}, timeout time.Duration) (interface{}, error) {
	return httpRequest(url, "POST", data, timeout)
}
