package pkg

import (
	"bytes"
	"encoding/json"
	telegraph_errors "go-telegraph/api/errors"
	telegraph "go-telegraph/api/types"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"gopkg.in/validator.v2"
)

// Method that validate params, make request and return valid structure as a result
func PerformRequest(domainUrl string, apiReq *telegraph.ApiReq) (result interface{}, err error) {
	// Validate input
	err = validator.Validate(apiReq.Params)
	if err != nil {
		return result, &telegraph_errors.ValidationError{Message: err.Error()}
	}

	// Make request
	res, err := Post(domainUrl, apiReq.Endpoint, &apiReq.Params)
	if err != nil {
		return
	}

	// Parse response
	response := telegraph.ApiRes{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return
	}

	// Validate if response is OK
	if !response.OK {
		return result, &telegraph_errors.APIError{Message: response.Error}
	}

	// Get the type of the Model field
	modelType := reflect.TypeOf(apiReq.Model)

	// Create a new instance of the Model type
	model := reflect.New(modelType).Interface()

	// Convert response to our Model struct
	temporaryVariable, err := json.Marshal(response.Result)
	if err != nil {
		return
	}

	err = json.Unmarshal(temporaryVariable, &model)
	if err != nil {
		return
	}

	return model, nil
}

// Base method to make API calls
func Post(domainUrl string, endpoint string, params interface{}) (result []byte, err error) {
	// Request url
	url, err := url.JoinPath(domainUrl, endpoint)
	if err != nil {
		return
	}

	// Marshal it into JSON prior to requesting
	reqParamsJSON, err := json.Marshal(params)
	if err != nil {
		return
	}

	// Prepare request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqParamsJSON))
	if err != nil {
		return
	}

	// Make request
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return body, nil
}
