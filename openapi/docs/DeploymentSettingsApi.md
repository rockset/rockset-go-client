# \DeploymentSettingsApi

All URIs are relative to *https://api.use1a1.rockset.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeploymentSettings**](DeploymentSettingsApi.md#GetDeploymentSettings) | **Get** /v1/orgs/self/deploymentsettings | Retrieve Deployment Settings
[**UpdateDeploymentSettings**](DeploymentSettingsApi.md#UpdateDeploymentSettings) | **Put** /v1/orgs/self/deploymentsettings | Update Deployment Settings



## GetDeploymentSettings

> DeploymentSettingsResponse GetDeploymentSettings(ctx).Execute()

Retrieve Deployment Settings



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
    resp, r, err := apiClient.DeploymentSettingsApi.GetDeploymentSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentSettingsApi.GetDeploymentSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentSettings`: DeploymentSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `DeploymentSettingsApi.GetDeploymentSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentSettingsRequest struct via the builder pattern


### Return type

[**DeploymentSettingsResponse**](DeploymentSettingsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeploymentSettings

> DeploymentSettingsResponse UpdateDeploymentSettings(ctx).Body(body).Execute()

Update Deployment Settings



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
    body := *openapiclient.NewUpdateDeploymentSettingsRequest() // UpdateDeploymentSettingsRequest | JSON object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentSettingsApi.UpdateDeploymentSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentSettingsApi.UpdateDeploymentSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeploymentSettings`: DeploymentSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `DeploymentSettingsApi.UpdateDeploymentSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateDeploymentSettingsRequest**](UpdateDeploymentSettingsRequest.md) | JSON object | 

### Return type

[**DeploymentSettingsResponse**](DeploymentSettingsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

