/*
 * Dynatrace Environment API
 *
 *  Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// SyntheticHTTPMonitorExecutionApiService SyntheticHTTPMonitorExecutionApi service
type SyntheticHTTPMonitorExecutionApiService service

type ApiGetExecutionResultRequest struct {
	ctx _context.Context
	ApiService *SyntheticHTTPMonitorExecutionApiService
	monitorId string
	resultType string
	locationId *string
}

func (r ApiGetExecutionResultRequest) LocationId(locationId string) ApiGetExecutionResultRequest {
	r.locationId = &locationId
	return r
}

func (r ApiGetExecutionResultRequest) Execute() (MonitorExecutionResults, *_nethttp.Response, error) {
	return r.ApiService.GetExecutionResultExecute(r)
}

/*
 * GetExecutionResult Returns detailed information about last HTTP monitor execution. | maturity=EARLY_ADOPTER
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param monitorId Identifier of the HTTP monitor for which last execution result is returned.
 * @param resultType Defines the result type of the last HTTP monitor's execution.
 * @return ApiGetExecutionResultRequest
 */
func (a *SyntheticHTTPMonitorExecutionApiService) GetExecutionResult(ctx _context.Context, monitorId string, resultType string) ApiGetExecutionResultRequest {
	return ApiGetExecutionResultRequest{
		ApiService: a,
		ctx: ctx,
		monitorId: monitorId,
		resultType: resultType,
	}
}

/*
 * Execute executes the request
 * @return MonitorExecutionResults
 */
func (a *SyntheticHTTPMonitorExecutionApiService) GetExecutionResultExecute(r ApiGetExecutionResultRequest) (MonitorExecutionResults, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MonitorExecutionResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyntheticHTTPMonitorExecutionApiService.GetExecutionResult")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/synthetic/execution/{monitorId}/{resultType}"
	localVarPath = strings.Replace(localVarPath, "{"+"monitorId"+"}", _neturl.PathEscape(parameterToString(r.monitorId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resultType"+"}", _neturl.PathEscape(parameterToString(r.resultType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.locationId != nil {
		localVarQueryParams.Add("locationId", parameterToString(*r.locationId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; charset=utf-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Api-Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
