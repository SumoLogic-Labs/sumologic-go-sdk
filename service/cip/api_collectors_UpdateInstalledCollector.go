package cip

import (
	"fmt"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

/*
UpdateInstalledCollector
Update an installed collector in the organization.
	body - Information to update about the collector.
	id - Identifier of the installed collector to update.
*/
func (a *APIClient) UpdateInstalledCollector(body types.UpdateInstalledCollectorDefinition, id string) (types.CollectorModel, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Put")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue types.CollectorModel
	)

	localVarPath := a.Cfg.BasePath + "/v1/collectors/{collectorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"collectorId"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

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
	etag, err := a.getCollectorEtag(localVarPath)
	if err != nil {
		return localVarReturnValue, nil, err
	}
	localVarHeaderParams["If-Match"] = etag[0]

	// body params
	collectorInfo, response, err := a.GetCollectorById(id)
	if err != nil {
		return localVarReturnValue, response, err
	}
	if collectorInfo.Collector.Category != body.Collector.Category {
		collectorInfo.Collector.Category = body.Collector.Category
	}
	if collectorInfo.Collector.Description != body.Collector.Description {
		collectorInfo.Collector.Description = body.Collector.Description
	}
	if collectorInfo.Collector.Ephemeral != body.Collector.Ephemeral {
		collectorInfo.Collector.Ephemeral = body.Collector.Ephemeral
	}
	if !reflect.DeepEqual(collectorInfo.Collector.Fields, body.Collector.Fields) {
		collectorInfo.Collector.Fields = body.Collector.Fields
	}
	if collectorInfo.Collector.HostName != body.Collector.HostName {
		collectorInfo.Collector.HostName = body.Collector.HostName
	}
	if collectorInfo.Collector.Name != body.Collector.Name && body.Collector.Name != "" {
		collectorInfo.Collector.Name = body.Collector.Name
	}
	if collectorInfo.Collector.SourceSyncMode != body.Collector.SourceSyncMode {
		collectorInfo.Collector.SourceSyncMode = body.Collector.SourceSyncMode
	}
	if collectorInfo.Collector.TargetCpu != int(body.Collector.TargetCPU) {
		collectorInfo.Collector.TargetCpu = int(body.Collector.TargetCPU)
	}
	if collectorInfo.Collector.TimeZone != body.Collector.TimeZone {
		collectorInfo.Collector.TimeZone = body.Collector.TimeZone
	}
	localVarPostBody = &collectorInfo

	r, err := a.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	} else if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v types.CollectorModel
			err = a.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		} else if localVarHttpResponse.StatusCode >= 400 {
			var v types.ErrorResponse
			err = a.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			if v.Errors[0].Meta.Reason != "" {
				newErr.error = v.Errors[0].Message + ": " + v.Errors[0].Meta.Reason
			} else {
				newErr.error = v.Errors[0].Message
			}
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
