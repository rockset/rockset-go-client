# \QueriesApi

All URIs are relative to *https://api.use1a1.rockset.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelQuery**](QueriesApi.md#CancelQuery) | **Delete** /v1/orgs/self/queries/{queryId} | Cancel Query
[**GetQuery**](QueriesApi.md#GetQuery) | **Get** /v1/orgs/self/queries/{queryId} | Retrieve Query
[**GetQueryResults**](QueriesApi.md#GetQueryResults) | **Get** /v1/orgs/self/queries/{queryId}/pages | Retrieve Query Results Page
[**ListActiveQueries**](QueriesApi.md#ListActiveQueries) | **Get** /v1/orgs/self/queries | List Queries
[**Query**](QueriesApi.md#Query) | **Post** /v1/orgs/self/queries | Execute SQL Query
[**Validate**](QueriesApi.md#Validate) | **Post** /v1/orgs/self/queries/validations | Validate Query



## CancelQuery

> CancelQueryResponse CancelQuery(ctx, queryId).Execute()

Cancel Query



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
    queryId := "queryId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueriesApi.CancelQuery(context.Background(), queryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueriesApi.CancelQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelQuery`: CancelQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `QueriesApi.CancelQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CancelQueryResponse**](CancelQueryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuery

> GetQueryResponse GetQuery(ctx, queryId).Execute()

Retrieve Query



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
    queryId := "queryId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueriesApi.GetQuery(context.Background(), queryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueriesApi.GetQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQuery`: GetQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `QueriesApi.GetQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetQueryResponse**](GetQueryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueryResults

> QueryPaginationResponse GetQueryResults(ctx, queryId).Cursor(cursor).Docs(docs).Offset(offset).Execute()

Retrieve Query Results Page



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
    queryId := "queryId_example" // string | 
    cursor := "cursor_example" // string | Cursor to current page. If unset, will default to the first page. (optional)
    docs := int32(56) // int32 | Number of documents to fetch. (optional)
    offset := int32(56) // int32 | Offset from the cursor of the first document to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueriesApi.GetQueryResults(context.Background(), queryId).Cursor(cursor).Docs(docs).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueriesApi.GetQueryResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQueryResults`: QueryPaginationResponse
    fmt.Fprintf(os.Stdout, "Response from `QueriesApi.GetQueryResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueryResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Cursor to current page. If unset, will default to the first page. | 
 **docs** | **int32** | Number of documents to fetch. | 
 **offset** | **int32** | Offset from the cursor of the first document to be returned | 

### Return type

[**QueryPaginationResponse**](QueryPaginationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActiveQueries

> ListQueriesResponse ListActiveQueries(ctx).Execute()

List Queries



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueriesApi.ListActiveQueries(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueriesApi.ListActiveQueries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListActiveQueries`: ListQueriesResponse
    fmt.Fprintf(os.Stdout, "Response from `QueriesApi.ListActiveQueries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListActiveQueriesRequest struct via the builder pattern


### Return type

[**ListQueriesResponse**](ListQueriesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Query

> QueryResponse Query(ctx).Body(body).Execute()

Execute SQL Query



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
    body := *openapiclient.NewQueryRequest(*openapiclient.NewQueryRequestSql("SELECT * FROM foo where _id = :_id")) // QueryRequest | JSON object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueriesApi.Query(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueriesApi.Query``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Query`: QueryResponse
    fmt.Fprintf(os.Stdout, "Response from `QueriesApi.Query`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**QueryRequest**](QueryRequest.md) | JSON object | 

### Return type

[**QueryResponse**](QueryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Validate

> ValidateQueryResponse Validate(ctx).Body(body).Execute()

Validate Query



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
    body := *openapiclient.NewQueryRequest(*openapiclient.NewQueryRequestSql("SELECT * FROM foo where _id = :_id")) // QueryRequest | JSON object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueriesApi.Validate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueriesApi.Validate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Validate`: ValidateQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `QueriesApi.Validate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**QueryRequest**](QueryRequest.md) | JSON object | 

### Return type

[**ValidateQueryResponse**](ValidateQueryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

