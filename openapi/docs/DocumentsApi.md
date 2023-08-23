# \DocumentsApi

All URIs are relative to *https://api.use1a1.rockset.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDocuments**](DocumentsApi.md#AddDocuments) | **Post** /v1/orgs/self/ws/{workspace}/collections/{collection}/docs | Add Documents
[**DeleteDocuments**](DocumentsApi.md#DeleteDocuments) | **Delete** /v1/orgs/self/ws/{workspace}/collections/{collection}/docs | Delete Documents
[**PatchDocuments**](DocumentsApi.md#PatchDocuments) | **Patch** /v1/orgs/self/ws/{workspace}/collections/{collection}/docs | Patch Documents



## AddDocuments

> AddDocumentsResponse AddDocuments(ctx, workspace, collection).Body(body).Execute()

Add Documents



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
    workspace := "workspace_example" // string | Name of the workspace. (default to "commons")
    collection := "collection_example" // string | Name of the collection.
    body := *openapiclient.NewAddDocumentsRequest([]map[string]interface{}{map[string]interface{}(123)}) // AddDocumentsRequest | JSON object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.AddDocuments(context.Background(), workspace, collection).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.AddDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDocuments`: AddDocumentsResponse
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.AddDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | Name of the workspace. | [default to &quot;commons&quot;]
**collection** | **string** | Name of the collection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**AddDocumentsRequest**](AddDocumentsRequest.md) | JSON object | 

### Return type

[**AddDocumentsResponse**](AddDocumentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDocuments

> DeleteDocumentsResponse DeleteDocuments(ctx, workspace, collection).Body(body).Execute()

Delete Documents



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
    workspace := "workspace_example" // string | Name of the workspace. (default to "commons")
    collection := "collection_example" // string | Name of the collection.
    body := *openapiclient.NewDeleteDocumentsRequest([]openapiclient.DeleteDocumentsRequestData{*openapiclient.NewDeleteDocumentsRequestData("2cd61e3b")}) // DeleteDocumentsRequest | JSON object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.DeleteDocuments(context.Background(), workspace, collection).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.DeleteDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDocuments`: DeleteDocumentsResponse
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.DeleteDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | Name of the workspace. | [default to &quot;commons&quot;]
**collection** | **string** | Name of the collection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DeleteDocumentsRequest**](DeleteDocumentsRequest.md) | JSON object | 

### Return type

[**DeleteDocumentsResponse**](DeleteDocumentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDocuments

> PatchDocumentsResponse PatchDocuments(ctx, workspace, collection).Body(body).Execute()

Patch Documents



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
    workspace := "workspace_example" // string | Name of the workspace. (default to "commons")
    collection := "collection_example" // string | Name of the collection.
    body := *openapiclient.NewPatchDocumentsRequest([]openapiclient.PatchDocument{*openapiclient.NewPatchDocument("ca2d6832-1bfd-f88f-0620-d2aa27a5d86c", []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("ADD", "/foo/bar")})}) // PatchDocumentsRequest | JSON Patch objects

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.PatchDocuments(context.Background(), workspace, collection).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.PatchDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchDocuments`: PatchDocumentsResponse
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.PatchDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | Name of the workspace. | [default to &quot;commons&quot;]
**collection** | **string** | Name of the collection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**PatchDocumentsRequest**](PatchDocumentsRequest.md) | JSON Patch objects | 

### Return type

[**PatchDocumentsResponse**](PatchDocumentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

