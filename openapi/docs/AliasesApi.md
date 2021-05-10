# \AliasesApi

All URIs are relative to *https://api.rs2.usw2.rockset.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlias**](AliasesApi.md#CreateAlias) | **Post** /v1/orgs/self/ws/{workspace}/aliases | Create Alias
[**DeleteAlias**](AliasesApi.md#DeleteAlias) | **Delete** /v1/orgs/self/ws/{workspace}/aliases/{alias} | Delete Alias
[**GetAlias**](AliasesApi.md#GetAlias) | **Get** /v1/orgs/self/ws/{workspace}/aliases/{alias} | Get Alias
[**ListAliases**](AliasesApi.md#ListAliases) | **Get** /v1/orgs/self/aliases | List Aliases
[**ListQueryLambdasWithAlias**](AliasesApi.md#ListQueryLambdasWithAlias) | **Get** /v1/orgs/self/ws/{workspace}/aliases/{alias}/lambdas | Get Query Lambdas with Alias
[**UpdateAlias**](AliasesApi.md#UpdateAlias) | **Post** /v1/orgs/self/ws/{workspace}/aliases/{alias} | Update Alias
[**WorkspaceAliases**](AliasesApi.md#WorkspaceAliases) | **Get** /v1/orgs/self/ws/{workspace}/aliases | List Aliases for Workspace



## CreateAlias

> CreateAliasResponse CreateAlias(ctx, workspace).Body(body).Execute()

Create Alias



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
    workspace := "workspace_example" // string | name of the workspace (default to "commons")
    body := *openapiclient.NewCreateAliasRequest("aliasName", []string{"Collections_example"}) // CreateAliasRequest | JSON object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AliasesApi.CreateAlias(context.Background(), workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasesApi.CreateAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlias`: CreateAliasResponse
    fmt.Fprintf(os.Stdout, "Response from `AliasesApi.CreateAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateAliasRequest**](CreateAliasRequest.md) | JSON object | 

### Return type

[**CreateAliasResponse**](CreateAliasResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlias

> DeleteAliasResponse DeleteAlias(ctx, workspace, alias).Execute()

Delete Alias



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
    workspace := "workspace_example" // string | name of the workspace (default to "commons")
    alias := "alias_example" // string | name of the alias

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AliasesApi.DeleteAlias(context.Background(), workspace, alias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasesApi.DeleteAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAlias`: DeleteAliasResponse
    fmt.Fprintf(os.Stdout, "Response from `AliasesApi.DeleteAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**alias** | **string** | name of the alias | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteAliasResponse**](DeleteAliasResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlias

> GetAliasResponse GetAlias(ctx, workspace, alias).Execute()

Get Alias



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
    workspace := "workspace_example" // string | name of the workspace (default to "commons")
    alias := "alias_example" // string | name of the alias

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AliasesApi.GetAlias(context.Background(), workspace, alias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasesApi.GetAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlias`: GetAliasResponse
    fmt.Fprintf(os.Stdout, "Response from `AliasesApi.GetAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**alias** | **string** | name of the alias | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAliasResponse**](GetAliasResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAliases

> ListAliasesResponse ListAliases(ctx).Execute()

List Aliases



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
    resp, r, err := api_client.AliasesApi.ListAliases(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasesApi.ListAliases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAliases`: ListAliasesResponse
    fmt.Fprintf(os.Stdout, "Response from `AliasesApi.ListAliases`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAliasesRequest struct via the builder pattern


### Return type

[**ListAliasesResponse**](ListAliasesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQueryLambdasWithAlias

> ListQueryLambdasResponse ListQueryLambdasWithAlias(ctx, workspace, alias).Execute()

Get Query Lambdas with Alias



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
    workspace := "workspace_example" // string | 
    alias := "alias_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AliasesApi.ListQueryLambdasWithAlias(context.Background(), workspace, alias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasesApi.ListQueryLambdasWithAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListQueryLambdasWithAlias`: ListQueryLambdasResponse
    fmt.Fprintf(os.Stdout, "Response from `AliasesApi.ListQueryLambdasWithAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** |  | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListQueryLambdasWithAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListQueryLambdasResponse**](ListQueryLambdasResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlias

> GetAliasResponse UpdateAlias(ctx, workspace, alias).Body(body).Execute()

Update Alias



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
    workspace := "workspace_example" // string | name of the workspace (default to "commons")
    alias := "alias_example" // string | name of the alias
    body := *openapiclient.NewUpdateAliasRequest([]string{"Collections_example"}) // UpdateAliasRequest | JSON object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AliasesApi.UpdateAlias(context.Background(), workspace, alias).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasesApi.UpdateAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlias`: GetAliasResponse
    fmt.Fprintf(os.Stdout, "Response from `AliasesApi.UpdateAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**alias** | **string** | name of the alias | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateAliasRequest**](UpdateAliasRequest.md) | JSON object | 

### Return type

[**GetAliasResponse**](GetAliasResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspaceAliases

> ListAliasesResponse WorkspaceAliases(ctx, workspace).Execute()

List Aliases for Workspace



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
    workspace := "workspace_example" // string | name of the workspace (default to "commons")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AliasesApi.WorkspaceAliases(context.Background(), workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasesApi.WorkspaceAliases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspaceAliases`: ListAliasesResponse
    fmt.Fprintf(os.Stdout, "Response from `AliasesApi.WorkspaceAliases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceAliasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListAliasesResponse**](ListAliasesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

