# \ViewsApi

All URIs are relative to *https://api.rs2.usw2.rockset.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateView**](ViewsApi.md#CreateView) | **Post** /v1/orgs/self/ws/{workspace}/views | Create View
[**DeleteView**](ViewsApi.md#DeleteView) | **Delete** /v1/orgs/self/ws/{workspace}/views/{view} | Delete View
[**GetView**](ViewsApi.md#GetView) | **Get** /v1/orgs/self/ws/{workspace}/views/{view} | Retrieve View
[**ListViews**](ViewsApi.md#ListViews) | **Get** /v1/orgs/self/views | List Views
[**UpdateView**](ViewsApi.md#UpdateView) | **Post** /v1/orgs/self/ws/{workspace}/views/{view} | Update View
[**WorkspaceViews**](ViewsApi.md#WorkspaceViews) | **Get** /v1/orgs/self/ws/{workspace}/views | List Views in Workspace



## CreateView

> CreateViewResponse CreateView(ctx, workspace).Body(body).Execute()

Create View



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
    body := *openapiclient.NewCreateViewRequest("myAwesomeView", "SELECT * FROM foo") // CreateViewRequest | JSON object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewsApi.CreateView(context.Background(), workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.CreateView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateView`: CreateViewResponse
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.CreateView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiCreateViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateViewRequest**](CreateViewRequest.md) | JSON object | 

### Return type

[**CreateViewResponse**](CreateViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteView

> DeleteViewResponse DeleteView(ctx, workspace, view).Execute()

Delete View



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
    view := "view_example" // string | name of the view

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewsApi.DeleteView(context.Background(), workspace, view).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.DeleteView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteView`: DeleteViewResponse
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.DeleteView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**view** | **string** | name of the view | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteViewResponse**](DeleteViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetView

> GetViewResponse GetView(ctx, workspace, view).Execute()

Retrieve View



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
    view := "view_example" // string | name of the view

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewsApi.GetView(context.Background(), workspace, view).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.GetView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetView`: GetViewResponse
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.GetView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**view** | **string** | name of the view | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetViewResponse**](GetViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListViews

> ListViewsResponse ListViews(ctx).Execute()

List Views



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
    resp, r, err := apiClient.ViewsApi.ListViews(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.ListViews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListViews`: ListViewsResponse
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.ListViews`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListViewsRequest struct via the builder pattern


### Return type

[**ListViewsResponse**](ListViewsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateView

> UpdateViewResponse UpdateView(ctx, workspace, view).Body(body).Execute()

Update View



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
    view := "view_example" // string | name of the view
    body := *openapiclient.NewUpdateViewRequest("SELECT * FROM foo") // UpdateViewRequest | JSON object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewsApi.UpdateView(context.Background(), workspace, view).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.UpdateView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateView`: UpdateViewResponse
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.UpdateView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**view** | **string** | name of the view | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateViewRequest**](UpdateViewRequest.md) | JSON object | 

### Return type

[**UpdateViewResponse**](UpdateViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspaceViews

> ListViewsResponse WorkspaceViews(ctx, workspace).Execute()

List Views in Workspace



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewsApi.WorkspaceViews(context.Background(), workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.WorkspaceViews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspaceViews`: ListViewsResponse
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.WorkspaceViews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListViewsResponse**](ListViewsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

