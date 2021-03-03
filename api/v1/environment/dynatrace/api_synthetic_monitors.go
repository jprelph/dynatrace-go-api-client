/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace Environment API v1. To read about use cases and examples, refer to the [help page](https://dt-url.net/xc03k3c).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// SyntheticMonitorsApiService SyntheticMonitorsApi service
type SyntheticMonitorsApiService service

type ApiAddMonitorRequest struct {
	ctx _context.Context
	ApiService *SyntheticMonitorsApiService
	syntheticMonitorUpdate *SyntheticMonitorUpdate
}

func (r ApiAddMonitorRequest) SyntheticMonitorUpdate(syntheticMonitorUpdate SyntheticMonitorUpdate) ApiAddMonitorRequest {
	r.syntheticMonitorUpdate = &syntheticMonitorUpdate
	return r
}

func (r ApiAddMonitorRequest) Execute() (EntityIdDto, *_nethttp.Response, error) {
	return r.ApiService.AddMonitorExecute(r)
}

/*
 * AddMonitor Creates a new synthetic monitor
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiAddMonitorRequest
 */
func (a *SyntheticMonitorsApiService) AddMonitor(ctx _context.Context) ApiAddMonitorRequest {
	return ApiAddMonitorRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return EntityIdDto
 */
func (a *SyntheticMonitorsApiService) AddMonitorExecute(r ApiAddMonitorRequest) (EntityIdDto, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EntityIdDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyntheticMonitorsApiService.AddMonitor")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/synthetic/monitors"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json; charset=utf-8"}

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
	// body params
	localVarPostBody = r.syntheticMonitorUpdate
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

type ApiDeleteMonitorRequest struct {
	ctx _context.Context
	ApiService *SyntheticMonitorsApiService
	monitorId string
}


func (r ApiDeleteMonitorRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteMonitorExecute(r)
}

/*
 * DeleteMonitor Deletes the specified synthetic monitor
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param monitorId The ID of the synthetic monitor to be deleted.
 * @return ApiDeleteMonitorRequest
 */
func (a *SyntheticMonitorsApiService) DeleteMonitor(ctx _context.Context, monitorId string) ApiDeleteMonitorRequest {
	return ApiDeleteMonitorRequest{
		ApiService: a,
		ctx: ctx,
		monitorId: monitorId,
	}
}

/*
 * Execute executes the request
 */
func (a *SyntheticMonitorsApiService) DeleteMonitorExecute(r ApiDeleteMonitorRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyntheticMonitorsApiService.DeleteMonitor")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/synthetic/monitors/{monitorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"monitorId"+"}", _neturl.PathEscape(parameterToString(r.monitorId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetMonitorRequest struct {
	ctx _context.Context
	ApiService *SyntheticMonitorsApiService
	monitorId string
}


func (r ApiGetMonitorRequest) Execute() (SyntheticMonitor, *_nethttp.Response, error) {
	return r.ApiService.GetMonitorExecute(r)
}

/*
 * GetMonitor Gets parameters of the specified synthetic monitor
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param monitorId The ID of the required synthetic monitor
 * @return ApiGetMonitorRequest
 */
func (a *SyntheticMonitorsApiService) GetMonitor(ctx _context.Context, monitorId string) ApiGetMonitorRequest {
	return ApiGetMonitorRequest{
		ApiService: a,
		ctx: ctx,
		monitorId: monitorId,
	}
}

/*
 * Execute executes the request
 * @return SyntheticMonitor
 */
func (a *SyntheticMonitorsApiService) GetMonitorExecute(r ApiGetMonitorRequest) (SyntheticMonitor, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SyntheticMonitor
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyntheticMonitorsApiService.GetMonitor")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/synthetic/monitors/{monitorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"monitorId"+"}", _neturl.PathEscape(parameterToString(r.monitorId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type ApiGetMonitorsCollectionRequest struct {
	ctx _context.Context
	ApiService *SyntheticMonitorsApiService
	managementZone *int64
	tag *[]string
	location *string
	assignedApps *[]string
	type_ *string
	enabled *bool
	credentialId *string
	credentialOwner *string
}

func (r ApiGetMonitorsCollectionRequest) ManagementZone(managementZone int64) ApiGetMonitorsCollectionRequest {
	r.managementZone = &managementZone
	return r
}
func (r ApiGetMonitorsCollectionRequest) Tag(tag []string) ApiGetMonitorsCollectionRequest {
	r.tag = &tag
	return r
}
func (r ApiGetMonitorsCollectionRequest) Location(location string) ApiGetMonitorsCollectionRequest {
	r.location = &location
	return r
}
func (r ApiGetMonitorsCollectionRequest) AssignedApps(assignedApps []string) ApiGetMonitorsCollectionRequest {
	r.assignedApps = &assignedApps
	return r
}
func (r ApiGetMonitorsCollectionRequest) Type_(type_ string) ApiGetMonitorsCollectionRequest {
	r.type_ = &type_
	return r
}
func (r ApiGetMonitorsCollectionRequest) Enabled(enabled bool) ApiGetMonitorsCollectionRequest {
	r.enabled = &enabled
	return r
}
func (r ApiGetMonitorsCollectionRequest) CredentialId(credentialId string) ApiGetMonitorsCollectionRequest {
	r.credentialId = &credentialId
	return r
}
func (r ApiGetMonitorsCollectionRequest) CredentialOwner(credentialOwner string) ApiGetMonitorsCollectionRequest {
	r.credentialOwner = &credentialOwner
	return r
}

func (r ApiGetMonitorsCollectionRequest) Execute() (Monitors, *_nethttp.Response, error) {
	return r.ApiService.GetMonitorsCollectionExecute(r)
}

/*
 * GetMonitorsCollection Lists all synthetic monitors in your Dynatrace environment
 * The full list can be lengthy, but you can narrow it down by specifying filter query parameters.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetMonitorsCollectionRequest
 */
func (a *SyntheticMonitorsApiService) GetMonitorsCollection(ctx _context.Context) ApiGetMonitorsCollectionRequest {
	return ApiGetMonitorsCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return Monitors
 */
func (a *SyntheticMonitorsApiService) GetMonitorsCollectionExecute(r ApiGetMonitorsCollectionRequest) (Monitors, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Monitors
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyntheticMonitorsApiService.GetMonitorsCollection")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/synthetic/monitors"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.managementZone != nil {
		localVarQueryParams.Add("managementZone", parameterToString(*r.managementZone, ""))
	}
	if r.tag != nil {
		t := *r.tag
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("tag", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("tag", parameterToString(t, "multi"))
		}
	}
	if r.location != nil {
		localVarQueryParams.Add("location", parameterToString(*r.location, ""))
	}
	if r.assignedApps != nil {
		t := *r.assignedApps
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("assignedApps", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("assignedApps", parameterToString(t, "multi"))
		}
	}
	if r.type_ != nil {
		localVarQueryParams.Add("type", parameterToString(*r.type_, ""))
	}
	if r.enabled != nil {
		localVarQueryParams.Add("enabled", parameterToString(*r.enabled, ""))
	}
	if r.credentialId != nil {
		localVarQueryParams.Add("credentialId", parameterToString(*r.credentialId, ""))
	}
	if r.credentialOwner != nil {
		localVarQueryParams.Add("credentialOwner", parameterToString(*r.credentialOwner, ""))
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

type ApiReplaceMonitorRequest struct {
	ctx _context.Context
	ApiService *SyntheticMonitorsApiService
	monitorId string
	syntheticMonitorUpdate *SyntheticMonitorUpdate
}

func (r ApiReplaceMonitorRequest) SyntheticMonitorUpdate(syntheticMonitorUpdate SyntheticMonitorUpdate) ApiReplaceMonitorRequest {
	r.syntheticMonitorUpdate = &syntheticMonitorUpdate
	return r
}

func (r ApiReplaceMonitorRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ReplaceMonitorExecute(r)
}

/*
 * ReplaceMonitor Updates parameters of the specified synthetic monitor
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param monitorId The ID of the synthetic monitor to be updated.
 * @return ApiReplaceMonitorRequest
 */
func (a *SyntheticMonitorsApiService) ReplaceMonitor(ctx _context.Context, monitorId string) ApiReplaceMonitorRequest {
	return ApiReplaceMonitorRequest{
		ApiService: a,
		ctx: ctx,
		monitorId: monitorId,
	}
}

/*
 * Execute executes the request
 */
func (a *SyntheticMonitorsApiService) ReplaceMonitorExecute(r ApiReplaceMonitorRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyntheticMonitorsApiService.ReplaceMonitor")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/synthetic/monitors/{monitorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"monitorId"+"}", _neturl.PathEscape(parameterToString(r.monitorId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json; charset=utf-8"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.syntheticMonitorUpdate
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
