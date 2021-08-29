package cip

import (
	"fmt"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

/*
UploadFile
Create a request to populate a lookup table with a CSV file.
	file - A CSV file
	id - Identifier of the lookup table to populate.
	optional - nil or *types.LookupTableUploadFileOpts - Optional Parameters:
		Merge (optional.Bool) - This indicates whether the file contents will be merged with existing data in the lookup table or not. If this is true then data with the same primary keys will be updated while the rest of the rows will be appended. By default, merge is false. The response includes a request identifier that you need to use in the RequestJobStatus function to track the status of the upload request.
		FileEncoding (optional.String) - File encoding of file being uploaded.
*/
func (a *APIClient) UploadFile(file *os.File, id string, localVarOptionals *types.LookupTableUploadFileOpts) (types.LookupRequestToken, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue types.LookupRequestToken
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/lookupTables/{id}/upload"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Merge.IsSet() {
		localVarQueryParams.Add("merge", parameterToString(localVarOptionals.Merge.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileEncoding.IsSet() {
		localVarQueryParams.Add("fileEncoding", parameterToString(localVarOptionals.FileEncoding.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"multipart/form-data"}

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
	localVarFile := file
	if localVarFile != nil {
		fbs, err := io.ReadAll(localVarFile)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
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
	} else if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v types.LookupRequestToken
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
