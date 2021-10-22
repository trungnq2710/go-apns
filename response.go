// Created at 10/21/2021 5:26 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package go_apns

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Reason            string `json:"reason"`
	Timestamp         string `json:"timestamp"`
	ReasonDescription string
	StatusCode        int
	ApnsID            string
	NeedCleanUpToken  bool
}

func parseResponse(resp *http.Response) (*Response, error) {
	response := &Response{}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	response.StatusCode = resp.StatusCode
	response.ReasonDescription = ReasonDescription(response.Reason)
	response.ApnsID = resp.Header.Get("apns-id")

	switch response.StatusCode {
	case StatusCodeErrorDeviceToken:
		response.NeedCleanUpToken = true
	case StatusCodeBadRequest:
		switch response.Reason {
		case ReasonDeviceTokenNotForTopic,
			ReasonBadDeviceToken,
			ReasonExpiredProviderToken,
			ReasonInvalidProviderToken,
			ReasonUnregistered:
			response.NeedCleanUpToken = true
		}
	}

	return response, nil
}
