package services

import (
	"bytes"
	"ego-api/loan/httperror"
	"ego-api/loan/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//HTTPClient handles http requests to remote endpoints
type HTTPClient struct {
	logger logger.Logger
}

//NewHTTPClient instantiates an HTTPClient
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		logger: logger.NewLogger(),
	}
}

func (hc *HTTPClient) transferFunds(senderID string, receiverID string, amount int64) (map[string]interface{}, *httperror.HTTPError) {
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}
	url := os.Getenv("INTER_WALLET_TRANSFER_URL")
	if url == "" {
		hc.logger.Error("INTER_WALLET_TRANSFER_URL env variable not defined")

		return nil, &httperror.HTTPError{
			Message: http.StatusText(http.StatusInternalServerError),
			Code:    http.StatusInternalServerError,
		}
	}
	url = fmt.Sprintf("%s/%s/%s", url, senderID, receiverID)
	log.Println(url)
	// Intialize new request
	requestBody, err := json.Marshal(map[string]int64{
		"amount": amount,
	})

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		hc.logger.Error(err.Error())

		return nil, &httperror.HTTPError{
			Message: http.StatusText(http.StatusInternalServerError),
			Code:    http.StatusInternalServerError,
		}
	}

	//Set headers and make remote request
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		hc.logger.Error(err.Error())
		return nil, &httperror.HTTPError{
			Message: http.StatusText(http.StatusInternalServerError),
			Code:    http.StatusInternalServerError,
		}
	}

	defer response.Body.Close()

	// Parse response
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		hc.logger.Error(err.Error())
		return nil, &httperror.HTTPError{
			Message: http.StatusText(http.StatusInternalServerError),
			Code:    http.StatusInternalServerError,
		}
	}
	log.Println(string(data))
	return hc.parseEGOAPIResponse(data, response.StatusCode)
}

func (hc *HTTPClient) parseEGOAPIResponse(data []byte, statusCode int) (map[string]interface{}, *httperror.HTTPError) {
	res := make(map[string]interface{})
	json.Unmarshal(data, &res)

	//TODO: Account for 400 errors.
	if statusCode != http.StatusOK {
		return nil, &httperror.HTTPError{
			Message: res["error"].(string),
			Code:    statusCode,
		}
	}

	return res, nil
}
