package cip

import (
	"fmt"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/*
GetFolderAsyncExportResult
Get folder results from content export job for the given job identifier. The results from this export are incompatible with the Library import feature in the Sumo user interface.
	contentId - The identifier of the exported content item.
	jobId - The identifier of the asynchronous job.
	optional - nil or *types.ContentManagementApiGetAsyncExportResultOpts - Optional Parameters:
		IsAdminMode (optional.String) - Set this to true if you want to perform the request as a Content Administrator.
*/
func (a *APIClient) GetFolderAsyncExportResult(contentId string, jobId string, localVarOptionals *types.ContentOpts) (types.FolderSyncDefinition, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue types.FolderSyncDefinition
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v2/content/{contentId}/export/{jobId}/result"
	localVarPath = strings.Replace(localVarPath, "{"+"contentId"+"}", fmt.Sprintf("%v", contentId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", fmt.Sprintf("%v", jobId), -1)

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
	if localVarOptionals != nil && localVarOptionals.IsAdminMode.IsSet() {
		localVarHeaderParams["isAdminMode"] = parameterToString(localVarOptionals.IsAdminMode.Value(), "")
	}
	r, err := a.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
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
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v types.FolderSyncDefinition
			err = a.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v types.ErrorResponse
			err = a.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
