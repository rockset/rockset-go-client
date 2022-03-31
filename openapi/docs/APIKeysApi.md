# \APIKeysApi

All URIs are relative to *https://api.rs2.usw2.rockset.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiKey**](APIKeysApi.md#CreateApiKey) | **Post** /v1/orgs/self/users/self/apikeys | Create API Key
[**DeleteApiKey**](APIKeysApi.md#DeleteApiKey) | **Delete** /v1/orgs/self/users/{user}/apikeys/{name} | Delete API Key
[**GetApiKey**](APIKeysApi.md#GetApiKey) | **Get** /v1/orgs/self/users/{user}/apikeys/{name} | Retrieve API Key
[**ListApiKeys**](APIKeysApi.md#ListApiKeys) | **Get** /v1/orgs/self/users/{user}/apikeys | List API Keys
[**UpdateApiKey**](APIKeysApi.md#UpdateApiKey) | **Post** /v1/orgs/self/users/{user}/apikeys/{name} | Update API Key State



## CreateApiKey

> CreateApiKeyResponse CreateApiKey(ctx).Body(body).Execute()

Create API Key



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
    body := *openapiclient.NewCreateApiKeyRequest("my-app") // CreateApiKeyRequest | JSON object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeysApi.CreateApiKey(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysApi.CreateApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiKey`: CreateApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `APIKeysApi.CreateApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateApiKeyRequest**](CreateApiKeyRequest.md) | JSON object | 

### Return type

[**CreateApiKeyResponse**](CreateApiKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiKey

> DeleteApiKeyResponse DeleteApiKey(ctx, name, user).Execute()

Delete API Key



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
    name := "my-key" // string | Name of the API key.
    user := "admin@me.com" // string | Email of the API key owner. Use `self` to specify the currently authenticated user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeysApi.DeleteApiKey(context.Background(), name, user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysApi.DeleteApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteApiKey`: DeleteApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `APIKeysApi.DeleteApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the API key. | 
**user** | **string** | Email of the API key owner. Use &#x60;self&#x60; to specify the currently authenticated user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteApiKeyResponse**](DeleteApiKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiKey

> GetApiKeyResponse GetApiKey(ctx, user, name).Execute()

Retrieve API Key



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
    user := "admin@me.com" // string | Email of the API key owner. Use `self` to specify the currently authenticated user.
    name := "my-key" // string | Name of the API key.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeysApi.GetApiKey(context.Background(), user, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysApi.GetApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiKey`: GetApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `APIKeysApi.GetApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** | Email of the API key owner. Use &#x60;self&#x60; to specify the currently authenticated user. | 
**name** | **string** | Name of the API key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetApiKeyResponse**](GetApiKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiKeys

> ListApiKeysResponse ListApiKeys(ctx, user).Execute()

List API Keys



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
    user := "admin@me.com" // string | Email of the API key owner. Use `self` to specify the currently authenticated user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeysApi.ListApiKeys(context.Background(), user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysApi.ListApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiKeys`: ListApiKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `APIKeysApi.ListApiKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** | Email of the API key owner. Use &#x60;self&#x60; to specify the currently authenticated user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListApiKeysResponse**](ListApiKeysResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiKey

> UpdateApiKeyResponse UpdateApiKey(ctx, name, user).Body(body).Execute()

Update API Key State



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
    name := "my-key" // string | Name of the API key.
    user := "admin@me.com" // string | Email of the API key owner. Use `self` to specify the currently authenticated user.
    body := *openapiclient.NewUpdateApiKeyRequest() // UpdateApiKeyRequest | JSON object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeysApi.UpdateApiKey(context.Background(), name, user).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysApi.UpdateApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiKey`: UpdateApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `APIKeysApi.UpdateApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the API key. | 
**user** | **string** | Email of the API key owner. Use &#x60;self&#x60; to specify the currently authenticated user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateApiKeyRequest**](UpdateApiKeyRequest.md) | JSON object | 

### Return type

[**UpdateApiKeyResponse**](UpdateApiKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

