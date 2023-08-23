# \SourcesApi

All URIs are relative to *https://api.use1a1.rockset.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSource**](SourcesApi.md#CreateSource) | **Post** /v1/orgs/self/ws/{workspace}/collections/{collection}/sources | Create a source
[**DeleteSource**](SourcesApi.md#DeleteSource) | **Delete** /v1/orgs/self/ws/{workspace}/collections/{collection}/sources/{source} | Delete Collection source
[**GetSource**](SourcesApi.md#GetSource) | **Get** /v1/orgs/self/ws/{workspace}/collections/{collection}/sources/{source} | Retrieve source
[**ListCollectionSources**](SourcesApi.md#ListCollectionSources) | **Get** /v1/orgs/self/ws/{workspace}/collections/{collection}/sources | List sources in collection
[**ResumeSource**](SourcesApi.md#ResumeSource) | **Post** /v1/orgs/self/ws/{workspace}/collections/{collection}/sources/{source}/resume | Resume source ingest
[**SuspendSource**](SourcesApi.md#SuspendSource) | **Post** /v1/orgs/self/ws/{workspace}/collections/{collection}/sources/{source}/suspend | Suspend source ingest



## CreateSource

> GetSourceResponse CreateSource(ctx, workspace, collection).Body(body).Execute()

Create a source



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/rockset/rockset-go-client"
)

func main() {
    workspace := "workspace_example" // string | name of the workspace (default to "commons")
    collection := "collection_example" // string | name of the collection
    body := *openapiclient.NewSource() // Source | JSON object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourcesApi.CreateSource(context.Background(), workspace, collection).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.CreateSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSource`: GetSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `SourcesApi.CreateSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**collection** | **string** | name of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**Source**](Source.md) | JSON object | 

### Return type

[**GetSourceResponse**](GetSourceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSource

> DeleteSourceResponse DeleteSource(ctx, workspace, collection, source).Execute()

Delete Collection source



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/rockset/rockset-go-client"
)

func main() {
    workspace := "workspace_example" // string | name of the workspace (default to "commons")
    collection := "collection_example" // string | name of the collection
    source := "source_example" // string | id of source

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourcesApi.DeleteSource(context.Background(), workspace, collection, source).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.DeleteSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSource`: DeleteSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `SourcesApi.DeleteSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**collection** | **string** | name of the collection | 
**source** | **string** | id of source | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DeleteSourceResponse**](DeleteSourceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSource

> GetSourceResponse GetSource(ctx, workspace, collection, source).Execute()

Retrieve source



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/rockset/rockset-go-client"
)

func main() {
    workspace := "workspace_example" // string | name of the workspace (default to "commons")
    collection := "collection_example" // string | name of the collection
    source := "source_example" // string | id of source

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourcesApi.GetSource(context.Background(), workspace, collection, source).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.GetSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSource`: GetSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `SourcesApi.GetSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**collection** | **string** | name of the collection | 
**source** | **string** | id of source | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetSourceResponse**](GetSourceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollectionSources

> ListSourcesResponse ListCollectionSources(ctx, workspace, collection).Execute()

List sources in collection



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/rockset/rockset-go-client"
)

func main() {
    workspace := "workspace_example" // string | name of the workspace (default to "commons")
    collection := "collection_example" // string | name of the collection

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourcesApi.ListCollectionSources(context.Background(), workspace, collection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.ListCollectionSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCollectionSources`: ListSourcesResponse
    fmt.Fprintf(os.Stdout, "Response from `SourcesApi.ListCollectionSources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**collection** | **string** | name of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListSourcesResponse**](ListSourcesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeSource

> GetSourceResponse ResumeSource(ctx, workspace, collection, source).Execute()

Resume source ingest



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/rockset/rockset-go-client"
)

func main() {
    workspace := "workspace_example" // string | name of the workspace (default to "commons")
    collection := "collection_example" // string | name of the collection
    source := "source_example" // string | id of source

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourcesApi.ResumeSource(context.Background(), workspace, collection, source).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.ResumeSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeSource`: GetSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `SourcesApi.ResumeSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**collection** | **string** | name of the collection | 
**source** | **string** | id of source | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetSourceResponse**](GetSourceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuspendSource

> GetSourceResponse SuspendSource(ctx, workspace, collection, source).Execute()

Suspend source ingest



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/rockset/rockset-go-client"
)

func main() {
    workspace := "workspace_example" // string | name of the workspace (default to "commons")
    collection := "collection_example" // string | name of the collection
    source := "source_example" // string | id of source

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourcesApi.SuspendSource(context.Background(), workspace, collection, source).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.SuspendSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuspendSource`: GetSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `SourcesApi.SuspendSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**collection** | **string** | name of the collection | 
**source** | **string** | id of source | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetSourceResponse**](GetSourceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

