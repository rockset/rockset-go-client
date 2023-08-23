# \SharedLambdasApi

All URIs are relative to *https://api.use1a1.rockset.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecutePublicQueryLambdaWithParams**](SharedLambdasApi.md#ExecutePublicQueryLambdaWithParams) | **Post** /v1/public/shared_lambdas/{public_access_id} | Execute a Public Query Lambda



## ExecutePublicQueryLambdaWithParams

> QueryResponse ExecutePublicQueryLambdaWithParams(ctx, publicAccessId).Body(body).Execute()

Execute a Public Query Lambda



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
    publicAccessId := "publicAccessId_example" // string | public access ID of the query lambda
    body := *openapiclient.NewExecutePublicQueryLambdaRequest() // ExecutePublicQueryLambdaRequest | JSON object (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SharedLambdasApi.ExecutePublicQueryLambdaWithParams(context.Background(), publicAccessId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedLambdasApi.ExecutePublicQueryLambdaWithParams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecutePublicQueryLambdaWithParams`: QueryResponse
    fmt.Fprintf(os.Stdout, "Response from `SharedLambdasApi.ExecutePublicQueryLambdaWithParams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicAccessId** | **string** | public access ID of the query lambda | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecutePublicQueryLambdaWithParamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ExecutePublicQueryLambdaRequest**](ExecutePublicQueryLambdaRequest.md) | JSON object | 

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

