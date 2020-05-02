package services

import (
	"ego-api/wallet/httperror"
	"ego-api/wallet/wlogger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//HTTPClient handles http requests to remote endpoints
type HTTPClient struct {
	logger wlogger.Logger
}

//NewHTTPClient instantiates an HTTPClient
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		logger: wlogger.NewLogger(),
	}
}

func (hc *HTTPClient) verifyWithPaystack(reference string) (map[string]interface{}, *httperror.HTTPError) {
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}
	url := fmt.Sprintf("https://api.paystack.co/transaction/verify/%s", reference)

	// Intialize new request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		hc.logger.Error(err.Error())

		return nil, &httperror.HTTPError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	// Check if PAYSTACK_SECRET env variable is set
	var key string
	if key = os.Getenv("PAYSTACK_SECRET"); key == "" {
		hc.logger.Error("PAYSTACK_SECRET env variable not set")

		return nil, &httperror.HTTPError{
			Message: httperror.ServerError,
			Code:    http.StatusInternalServerError,
		}
	}
	auth := fmt.Sprintf("Bearer %s", key)

	//Set headers and make remote request
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("Authorization", auth)

	response, err := client.Do(request)
	if err != nil {
		hc.logger.Error(err.Error())
		return nil, &httperror.HTTPError{
			Message: httperror.ServerError,
			Code:    http.StatusInternalServerError,
		}
	}

	defer response.Body.Close()

	// Parse response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		hc.logger.Error(err.Error())
		return nil, &httperror.HTTPError{
			Message: httperror.ServerError,
			Code:    http.StatusInternalServerError,
		}
	}

	res := make(map[string]interface{})
	json.Unmarshal(body, &res)

	return res, nil
}
