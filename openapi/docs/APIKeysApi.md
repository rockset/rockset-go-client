# \APIKeysApi

All URIs are relative to *https://api.rs2.usw2.rockset.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiKey**](APIKeysApi.md#CreateApiKey) | **Post** /v1/orgs/self/users/self/apikeys | Create API Key
[**DeleteApiKey**](APIKeysApi.md#DeleteApiKey) | **Delete** /v1/orgs/self/users/self/apikeys/{name} | Delete API Key
[**DeleteApiKeyAdmin**](APIKeysApi.md#DeleteApiKeyAdmin) | **Delete** /v1/orgs/self/users/{user}/apikeys/{name} | Delete API Key (any user)
[**GetApiKey**](APIKeysApi.md#GetApiKey) | **Get** /v1/orgs/self/users/self/apikeys/{name} | Get API Key
[**GetApiKeyAdmin**](APIKeysApi.md#GetApiKeyAdmin) | **Get** /v1/orgs/self/users/{user}/apikeys/{name} | Get any API Key
[**ListApiKeys**](APIKeysApi.md#ListApiKeys) | **Get** /v1/orgs/self/users/self/apikeys | List API Keys
[**ListApiKeysAdmin**](APIKeysApi.md#ListApiKeysAdmin) | **Get** /v1/orgs/self/users/{user}/apikeys | List API Keys (any user)



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
    body := *openapiclient.NewCreateApiKeyRequest("event-logger") // CreateApiKeyRequest | JSON object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.APIKeysApi.CreateApiKey(context.Background()).Body(body).Execute()
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

> DeleteApiKeyResponse DeleteApiKey(ctx, name).Execute()

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
    name := "name_example" // string | name of the API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.APIKeysApi.DeleteApiKey(context.Background(), name).Execute()
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
**name** | **string** | name of the API key | 

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


## DeleteApiKeyAdmin

> DeleteApiKeyResponse DeleteApiKeyAdmin(ctx, name, user).Execute()

Delete API Key (any user)



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
    name := "name_example" // string | name of the API key
    user := "user_example" // string | user email

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.APIKeysApi.DeleteApiKeyAdmin(context.Background(), name, user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysApi.DeleteApiKeyAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteApiKeyAdmin`: DeleteApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `APIKeysApi.DeleteApiKeyAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | name of the API key | 
**user** | **string** | user email | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiKeyAdminRequest struct via the builder pattern


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

> GetApiKeyResponse GetApiKey(ctx, name).Execute()

Get API Key



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
    name := "name_example" // string | name of the API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.APIKeysApi.GetApiKey(context.Background(), name).Execute()
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
**name** | **string** | name of the API key | 

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


## GetApiKeyAdmin

> GetApiKeyResponse GetApiKeyAdmin(ctx, user, name).Execute()

Get any API Key



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
    user := "user_example" // string | email of the api key owner
    name := "name_example" // string | name of the API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.APIKeysApi.GetApiKeyAdmin(context.Background(), user, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysApi.GetApiKeyAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiKeyAdmin`: GetApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `APIKeysApi.GetApiKeyAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** | email of the api key owner | 
**name** | **string** | name of the API key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeyAdminRequest struct via the builder pattern


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

> ListApiKeysResponse ListApiKeys(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.APIKeysApi.ListApiKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysApi.ListApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiKeys`: ListApiKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `APIKeysApi.ListApiKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiKeysRequest struct via the builder pattern


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


## ListApiKeysAdmin

> ListApiKeysResponse ListApiKeysAdmin(ctx, user).Execute()

List API Keys (any user)



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
    user := "user_example" // string | user email

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.APIKeysApi.ListApiKeysAdmin(context.Background(), user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysApi.ListApiKeysAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiKeysAdmin`: ListApiKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `APIKeysApi.ListApiKeysAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** | user email | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApiKeysAdminRequest struct via the builder pattern


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

