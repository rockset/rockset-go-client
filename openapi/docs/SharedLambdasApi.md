# \SharedLambdasApi

All URIs are relative to *https://api.use1a1.rockset.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecutePublicQueryLambda**](SharedLambdasApi.md#ExecutePublicQueryLambda) | **Get** /v1/public/shared_lambdas/{public_access_id} | Execute a Public Query Lambda



## ExecutePublicQueryLambda

> QueryResponse ExecutePublicQueryLambda(ctx, publicAccessId).Execute()

Execute a Public Query Lambda



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
    publicAccessId := "publicAccessId_example" // string | public access ID of the query lambda

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SharedLambdasApi.ExecutePublicQueryLambda(context.Background(), publicAccessId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedLambdasApi.ExecutePublicQueryLambda``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecutePublicQueryLambda`: QueryResponse
    fmt.Fprintf(os.Stdout, "Response from `SharedLambdasApi.ExecutePublicQueryLambda`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicAccessId** | **string** | public access ID of the query lambda | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecutePublicQueryLambdaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueryResponse**](QueryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

