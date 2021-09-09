package cip

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

/*
SendMessage
Allows you to send data to a Sumo Logic HTTP source endpoint. If sending the data to the endpoint fails it will retry every 10 seconds
based on the value of maxRetries.
	maxRetries - determines how many times the message will attempt to be sent if there is an error sending.
	message - data to send to the Sumo Logic source.
*/
func (a *APIClient) SendMessage(maxRetries int, message []byte) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path for source URL
	localVarpath := a.Cfg.SourceUrl

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarPostBody = &message
	r, err := a.prepareRequest(localVarpath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		count := 1
		for count <= maxRetries {
			fmt.Println("failed to send data to Sumo Logic retrying...attempt " + strconv.Itoa(count))
			localVarHttpResponse, err := a.callAPI(r)
			if err != nil || localVarHttpResponse == nil {
				return localVarHttpResponse, err
			}
			if localVarHttpResponse.StatusCode >= 300 {
				count++
			} else {
				fmt.Println("data successfully sent to Sumo Logic")
				break
			}
			time.Sleep(10 * time.Second)
		}
		if count > maxRetries {
			return localVarHttpResponse, err
		}
	} else {
		fmt.Println("data successfully sent to Sumo Logic")
	}
	return localVarHttpResponse, err
}
