# \IPAllowlistApi

All URIs are relative to *https://api.rs2.usw2.rockset.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAllowlistIp**](IPAllowlistApi.md#CreateAllowlistIp) | **Post** /v1/orgs/self/ip/allowlist | Create IP Allowlist Network Policy
[**DeleteAllowlistIp**](IPAllowlistApi.md#DeleteAllowlistIp) | **Delete** /v1/orgs/self/ip/allowlist/{name} | Delete IP Allowlist Network Policy
[**GetAllowlistIp**](IPAllowlistApi.md#GetAllowlistIp) | **Get** /v1/orgs/self/ip/allowlist/{name} | Get IP Allowlist Network Policy
[**ListAllowlistIps**](IPAllowlistApi.md#ListAllowlistIps) | **Get** /v1/orgs/self/ip/allowlist | List IP Allowlist Entries



## CreateAllowlistIp

> CreateIpAllowlistResponse CreateAllowlistIp(ctx).Body(body).Execute()

Create IP Allowlist Network Policy



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
    body := *openapiclient.NewCreateIpAllowlistRequest("Office", "82.217.192.0/18") // CreateIpAllowlistRequest | JSON object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPAllowlistApi.CreateAllowlistIp(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAllowlistApi.CreateAllowlistIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAllowlistIp`: CreateIpAllowlistResponse
    fmt.Fprintf(os.Stdout, "Response from `IPAllowlistApi.CreateAllowlistIp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAllowlistIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateIpAllowlistRequest**](CreateIpAllowlistRequest.md) | JSON object | 

### Return type

[**CreateIpAllowlistResponse**](CreateIpAllowlistResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllowlistIp

> DeleteIpAllowlistResponse DeleteAllowlistIp(ctx, name).Execute()

Delete IP Allowlist Network Policy



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
    name := "name_example" // string | name of the IP allowlist network policy

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPAllowlistApi.DeleteAllowlistIp(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAllowlistApi.DeleteAllowlistIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAllowlistIp`: DeleteIpAllowlistResponse
    fmt.Fprintf(os.Stdout, "Response from `IPAllowlistApi.DeleteAllowlistIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | name of the IP allowlist network policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllowlistIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteIpAllowlistResponse**](DeleteIpAllowlistResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllowlistIp

> GetIpAllowlistResponse GetAllowlistIp(ctx, name).Execute()

Get IP Allowlist Network Policy



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
    name := "name_example" // string | name of the IP Allowlist network policy

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPAllowlistApi.GetAllowlistIp(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAllowlistApi.GetAllowlistIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllowlistIp`: GetIpAllowlistResponse
    fmt.Fprintf(os.Stdout, "Response from `IPAllowlistApi.GetAllowlistIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | name of the IP Allowlist network policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllowlistIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetIpAllowlistResponse**](GetIpAllowlistResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllowlistIps

> ListIpAllowlistsResponse ListAllowlistIps(ctx).Execute()

List IP Allowlist Entries

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPAllowlistApi.ListAllowlistIps(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAllowlistApi.ListAllowlistIps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllowlistIps`: ListIpAllowlistsResponse
    fmt.Fprintf(os.Stdout, "Response from `IPAllowlistApi.ListAllowlistIps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAllowlistIpsRequest struct via the builder pattern


### Return type

[**ListIpAllowlistsResponse**](ListIpAllowlistsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

