package cip

import (
	"fmt"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/*
DisableField
After disabling a field Sumo Logic will start dropping its incoming values at ingest. As a result, they won&#x27;t be searchable or usable. Historical values are not removed and remain searchable.
	id - Identifier of a field to disable.
*/
func (a *APIClient) DisableField(id string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/fields/{id}/disable"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode >= 400 {
			var v types.ErrorResponse
			err = a.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			if v.Errors[0].Meta.Reason != "" {
				newErr.error = v.Errors[0].Message + ": " + v.Errors[0].Meta.Reason
			} else {
				newErr.error = v.Errors[0].Message
			}
			return localVarHttpResponse, newErr
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
