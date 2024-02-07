# \ScheduledLambdasApi

All URIs are relative to *https://api.use1a1.rockset.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateScheduledLambda**](ScheduledLambdasApi.md#CreateScheduledLambda) | **Post** /v1/orgs/self/ws/{workspace}/scheduled_lambdas | Create a Scheduled Lambda mapping
[**DeleteScheduledLambda**](ScheduledLambdasApi.md#DeleteScheduledLambda) | **Delete** /v1/orgs/self/ws/{workspace}/scheduled_lambdas/{scheduledLambdaId} | Delete a Scheduled Lambda mapping
[**UpdateScheduledLambda**](ScheduledLambdasApi.md#UpdateScheduledLambda) | **Post** /v1/orgs/self/ws/{workspace}/scheduled_lambdas/{scheduledLambdaId} | Update a Scheduled Lambda mapping



## CreateScheduledLambda

> ScheduledLambdaResponse CreateScheduledLambda(ctx, workspace).Body(body).Execute()

Create a Scheduled Lambda mapping



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
    body := *openapiclient.NewCreateScheduledLambdaRequest("* * * * *", "ql_name") // CreateScheduledLambdaRequest | JSON Object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledLambdasApi.CreateScheduledLambda(context.Background(), workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledLambdasApi.CreateScheduledLambda``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateScheduledLambda`: ScheduledLambdaResponse
    fmt.Fprintf(os.Stdout, "Response from `ScheduledLambdasApi.CreateScheduledLambda`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiCreateScheduledLambdaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateScheduledLambdaRequest**](CreateScheduledLambdaRequest.md) | JSON Object | 

### Return type

[**ScheduledLambdaResponse**](ScheduledLambdaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScheduledLambda

> ScheduledLambdaResponse DeleteScheduledLambda(ctx, workspace, scheduledLambdaId).Execute()

Delete a Scheduled Lambda mapping



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
    scheduledLambdaId := "scheduledLambdaId_example" // string | Scheduled Lambda RRN

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledLambdasApi.DeleteScheduledLambda(context.Background(), workspace, scheduledLambdaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledLambdasApi.DeleteScheduledLambda``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteScheduledLambda`: ScheduledLambdaResponse
    fmt.Fprintf(os.Stdout, "Response from `ScheduledLambdasApi.DeleteScheduledLambda`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**scheduledLambdaId** | **string** | Scheduled Lambda RRN | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScheduledLambdaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ScheduledLambdaResponse**](ScheduledLambdaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScheduledLambda

> ScheduledLambdaResponse UpdateScheduledLambda(ctx, workspace, scheduledLambdaId).Body(body).Execute()

Update a Scheduled Lambda mapping



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
    scheduledLambdaId := "scheduledLambdaId_example" // string | Scheduled Lambda RRN
    body := *openapiclient.NewUpdateScheduledLambdaRequest() // UpdateScheduledLambdaRequest | JSON Object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledLambdasApi.UpdateScheduledLambda(context.Background(), workspace, scheduledLambdaId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledLambdasApi.UpdateScheduledLambda``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateScheduledLambda`: ScheduledLambdaResponse
    fmt.Fprintf(os.Stdout, "Response from `ScheduledLambdasApi.UpdateScheduledLambda`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**scheduledLambdaId** | **string** | Scheduled Lambda RRN | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScheduledLambdaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateScheduledLambdaRequest**](UpdateScheduledLambdaRequest.md) | JSON Object | 

### Return type

[**ScheduledLambdaResponse**](ScheduledLambdaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

