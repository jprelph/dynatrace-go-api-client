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

// TopologySmartscapeServiceApiService TopologySmartscapeServiceApi service
type TopologySmartscapeServiceApiService service

type ApiGetBaselineValuesForSingleServiceRequest struct {
	ctx _context.Context
	ApiService *TopologySmartscapeServiceApiService
	meIdentifier string
}


func (r ApiGetBaselineValuesForSingleServiceRequest) Execute() (ServiceBaselineValues, *_nethttp.Response, error) {
	return r.ApiService.GetBaselineValuesForSingleServiceExecute(r)
}

/*
 * GetBaselineValuesForSingleService Gets baseline data for the specified service | maturity=EARLY_ADOPTER
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param meIdentifier The Dynatrace entity ID of the required service.
 * @return ApiGetBaselineValuesForSingleServiceRequest
 */
func (a *TopologySmartscapeServiceApiService) GetBaselineValuesForSingleService(ctx _context.Context, meIdentifier string) ApiGetBaselineValuesForSingleServiceRequest {
	return ApiGetBaselineValuesForSingleServiceRequest{
		ApiService: a,
		ctx: ctx,
		meIdentifier: meIdentifier,
	}
}

/*
 * Execute executes the request
 * @return ServiceBaselineValues
 */
func (a *TopologySmartscapeServiceApiService) GetBaselineValuesForSingleServiceExecute(r ApiGetBaselineValuesForSingleServiceRequest) (ServiceBaselineValues, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ServiceBaselineValues
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologySmartscapeServiceApiService.GetBaselineValuesForSingleService")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/services/{meIdentifier}/baseline"
	localVarPath = strings.Replace(localVarPath, "{"+"meIdentifier"+"}", _neturl.PathEscape(parameterToString(r.meIdentifier, "")), -1)

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

type ApiGetServicesRequest struct {
	ctx _context.Context
	ApiService *TopologySmartscapeServiceApiService
	startTimestamp *int64
	endTimestamp *int64
	relativeTime *string
	tag *[]string
	entity *[]string
	managementZone *int64
	includeDetails *bool
	pageSize *int32
	nextPageKey *string
}

func (r ApiGetServicesRequest) StartTimestamp(startTimestamp int64) ApiGetServicesRequest {
	r.startTimestamp = &startTimestamp
	return r
}
func (r ApiGetServicesRequest) EndTimestamp(endTimestamp int64) ApiGetServicesRequest {
	r.endTimestamp = &endTimestamp
	return r
}
func (r ApiGetServicesRequest) RelativeTime(relativeTime string) ApiGetServicesRequest {
	r.relativeTime = &relativeTime
	return r
}
func (r ApiGetServicesRequest) Tag(tag []string) ApiGetServicesRequest {
	r.tag = &tag
	return r
}
func (r ApiGetServicesRequest) Entity(entity []string) ApiGetServicesRequest {
	r.entity = &entity
	return r
}
func (r ApiGetServicesRequest) ManagementZone(managementZone int64) ApiGetServicesRequest {
	r.managementZone = &managementZone
	return r
}
func (r ApiGetServicesRequest) IncludeDetails(includeDetails bool) ApiGetServicesRequest {
	r.includeDetails = &includeDetails
	return r
}
func (r ApiGetServicesRequest) PageSize(pageSize int32) ApiGetServicesRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiGetServicesRequest) NextPageKey(nextPageKey string) ApiGetServicesRequest {
	r.nextPageKey = &nextPageKey
	return r
}

func (r ApiGetServicesRequest) Execute() ([]Service, *_nethttp.Response, error) {
	return r.ApiService.GetServicesExecute(r)
}

/*
 * GetServices Lists all available services in your environment
 * You can narrow down the output by specifying filtering parameters for the request. 

You can additionally limit the output by using pagination: 
1. Specify the number of results per page in the **pageSize** query parameter. 
2. Then use the URL-encoded cursor from the **Next-Page-Key** response header in the **nextPageKey** query parameter to obtain subsequent pages.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetServicesRequest
 */
func (a *TopologySmartscapeServiceApiService) GetServices(ctx _context.Context) ApiGetServicesRequest {
	return ApiGetServicesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []Service
 */
func (a *TopologySmartscapeServiceApiService) GetServicesExecute(r ApiGetServicesRequest) ([]Service, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Service
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologySmartscapeServiceApiService.GetServices")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/services"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.startTimestamp != nil {
		localVarQueryParams.Add("startTimestamp", parameterToString(*r.startTimestamp, ""))
	}
	if r.endTimestamp != nil {
		localVarQueryParams.Add("endTimestamp", parameterToString(*r.endTimestamp, ""))
	}
	if r.relativeTime != nil {
		localVarQueryParams.Add("relativeTime", parameterToString(*r.relativeTime, ""))
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
	if r.entity != nil {
		t := *r.entity
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("entity", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("entity", parameterToString(t, "multi"))
		}
	}
	if r.managementZone != nil {
		localVarQueryParams.Add("managementZone", parameterToString(*r.managementZone, ""))
	}
	if r.includeDetails != nil {
		localVarQueryParams.Add("includeDetails", parameterToString(*r.includeDetails, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.nextPageKey != nil {
		localVarQueryParams.Add("nextPageKey", parameterToString(*r.nextPageKey, ""))
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type ApiGetSingleServiceRequest struct {
	ctx _context.Context
	ApiService *TopologySmartscapeServiceApiService
	meIdentifier string
}


func (r ApiGetSingleServiceRequest) Execute() (Service, *_nethttp.Response, error) {
	return r.ApiService.GetSingleServiceExecute(r)
}

/*
 * GetSingleService Gets parameters of the specified service
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param meIdentifier The Dynatrace entity ID of the required service.
 * @return ApiGetSingleServiceRequest
 */
func (a *TopologySmartscapeServiceApiService) GetSingleService(ctx _context.Context, meIdentifier string) ApiGetSingleServiceRequest {
	return ApiGetSingleServiceRequest{
		ApiService: a,
		ctx: ctx,
		meIdentifier: meIdentifier,
	}
}

/*
 * Execute executes the request
 * @return Service
 */
func (a *TopologySmartscapeServiceApiService) GetSingleServiceExecute(r ApiGetSingleServiceRequest) (Service, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Service
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologySmartscapeServiceApiService.GetSingleService")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/services/{meIdentifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"meIdentifier"+"}", _neturl.PathEscape(parameterToString(r.meIdentifier, "")), -1)

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

type ApiUpdateServiceRequest struct {
	ctx _context.Context
	ApiService *TopologySmartscapeServiceApiService
	meIdentifier string
	updateEntity *UpdateEntity
}

func (r ApiUpdateServiceRequest) UpdateEntity(updateEntity UpdateEntity) ApiUpdateServiceRequest {
	r.updateEntity = &updateEntity
	return r
}

func (r ApiUpdateServiceRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UpdateServiceExecute(r)
}

/*
 * UpdateService Updates parameters of the specified service
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param meIdentifier The Dynatrace entity ID of the service you're inquiring.
 * @return ApiUpdateServiceRequest
 */
func (a *TopologySmartscapeServiceApiService) UpdateService(ctx _context.Context, meIdentifier string) ApiUpdateServiceRequest {
	return ApiUpdateServiceRequest{
		ApiService: a,
		ctx: ctx,
		meIdentifier: meIdentifier,
	}
}

/*
 * Execute executes the request
 */
func (a *TopologySmartscapeServiceApiService) UpdateServiceExecute(r ApiUpdateServiceRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologySmartscapeServiceApiService.UpdateService")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/services/{meIdentifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"meIdentifier"+"}", _neturl.PathEscape(parameterToString(r.meIdentifier, "")), -1)

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
	localVarPostBody = r.updateEntity
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
