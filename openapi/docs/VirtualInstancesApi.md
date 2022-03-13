# \VirtualInstancesApi

All URIs are relative to *https://api.rs2.usw2.rockset.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVirtualInstance**](VirtualInstancesApi.md#GetVirtualInstance) | **Get** /v1/orgs/self/virtualinstances/{virtualInstanceId} | Retrieve Virtual Instance
[**ListVirtualInstances**](VirtualInstancesApi.md#ListVirtualInstances) | **Get** /v1/orgs/self/virtualinstances | List Virtual Instances
[**SetVirtualInstance**](VirtualInstancesApi.md#SetVirtualInstance) | **Post** /v1/orgs/self/virtualinstances/{virtualInstanceId} | Update Virtual Instance



## GetVirtualInstance

> GetVirtualInstanceResponse GetVirtualInstance(ctx, virtualInstanceId).Execute()

Retrieve Virtual Instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    virtualInstanceId := "virtualInstanceId_example" // string | uuid of the virtual instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualInstancesApi.GetVirtualInstance(context.Background(), virtualInstanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualInstancesApi.GetVirtualInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualInstance`: GetVirtualInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualInstancesApi.GetVirtualInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualInstanceId** | **string** | uuid of the virtual instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetVirtualInstanceResponse**](GetVirtualInstanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualInstances

> ListVirtualInstancesResponse ListVirtualInstances(ctx).Execute()

List Virtual Instances



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualInstancesApi.ListVirtualInstances(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualInstancesApi.ListVirtualInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVirtualInstances`: ListVirtualInstancesResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualInstancesApi.ListVirtualInstances`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualInstancesRequest struct via the builder pattern


### Return type

[**ListVirtualInstancesResponse**](ListVirtualInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVirtualInstance

> UpdateVirtualInstanceResponse SetVirtualInstance(ctx, virtualInstanceId).Body(body).Execute()

Update Virtual Instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    virtualInstanceId := "virtualInstanceId_example" // string | uuid of the virtual instance
    body := *openapiclient.NewUpdateVirtualInstanceRequest() // UpdateVirtualInstanceRequest | JSON object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualInstancesApi.SetVirtualInstance(context.Background(), virtualInstanceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualInstancesApi.SetVirtualInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetVirtualInstance`: UpdateVirtualInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualInstancesApi.SetVirtualInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualInstanceId** | **string** | uuid of the virtual instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVirtualInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateVirtualInstanceRequest**](UpdateVirtualInstanceRequest.md) | JSON object | 

### Return type

[**UpdateVirtualInstanceResponse**](UpdateVirtualInstanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

