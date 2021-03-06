# \QueryLambdasApi

All URIs are relative to *https://api.rs2.usw2.rockset.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQueryLambda**](QueryLambdasApi.md#CreateQueryLambda) | **Post** /v1/orgs/self/ws/{workspace}/lambdas | Create Query Lambda
[**CreateQueryLambdaTag**](QueryLambdasApi.md#CreateQueryLambdaTag) | **Post** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda}/tags | Create Query Lambda Tag
[**DeleteQueryLambda**](QueryLambdasApi.md#DeleteQueryLambda) | **Delete** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda} | Delete Query Lambda
[**DeleteQueryLambdaTag**](QueryLambdasApi.md#DeleteQueryLambdaTag) | **Delete** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda}/tags/{tag} | Delete Query Lambda Tag Version
[**DeleteQueryLambdaVersion**](QueryLambdasApi.md#DeleteQueryLambdaVersion) | **Delete** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda}/version/{version} | Delete Query Lambda Version
[**ExecuteQueryLambda**](QueryLambdasApi.md#ExecuteQueryLambda) | **Post** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda}/versions/{version} | Execute Query Lambda
[**ExecuteQueryLambdaByTag**](QueryLambdasApi.md#ExecuteQueryLambdaByTag) | **Post** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda}/tags/{tag} | Execute Query Lambda By Tag
[**GetQueryLambdaTagVersion**](QueryLambdasApi.md#GetQueryLambdaTagVersion) | **Get** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda}/tags/{tag} | Get Query Lambda Tag
[**GetQueryLambdaVersion**](QueryLambdasApi.md#GetQueryLambdaVersion) | **Get** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda}/versions/{version} | Get Query Lambda Version
[**ListAllQueryLambdas**](QueryLambdasApi.md#ListAllQueryLambdas) | **Get** /v1/orgs/self/lambdas | List Query Lambdas
[**ListOrganizationTags**](QueryLambdasApi.md#ListOrganizationTags) | **Get** /v1/orgs/self/lambdas/tags | List All Query Lambda Tags
[**ListQueryLambdaTagVersions**](QueryLambdasApi.md#ListQueryLambdaTagVersions) | **Get** /v1/orgs/self/lambdas/tags/{tag} | List Query Lambda Tag Versions
[**ListQueryLambdaTags**](QueryLambdasApi.md#ListQueryLambdaTags) | **Get** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda}/tags | List Query Lambda Tags
[**ListQueryLambdaVersions**](QueryLambdasApi.md#ListQueryLambdaVersions) | **Get** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda}/versions | List Query Lambda Versions
[**ListQueryLambdasInWorkspace**](QueryLambdasApi.md#ListQueryLambdasInWorkspace) | **Get** /v1/orgs/self/ws/{workspace}/lambdas | List Query Lambdas in Workspace
[**UpdateQueryLambda**](QueryLambdasApi.md#UpdateQueryLambda) | **Post** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda}/versions | Update Query Lambda



## CreateQueryLambda

> QueryLambdaVersionResponse CreateQueryLambda(ctx, workspace).Body(body).Execute()

Create Query Lambda



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
    body := *openapiclient.NewCreateQueryLambdaRequest("myQueryLambda", *openapiclient.NewQueryLambdaSql("SELECT 'Foo'")) // CreateQueryLambdaRequest | JSON object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QueryLambdasApi.CreateQueryLambda(context.Background(), workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryLambdasApi.CreateQueryLambda``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateQueryLambda`: QueryLambdaVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `QueryLambdasApi.CreateQueryLambda`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiCreateQueryLambdaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateQueryLambdaRequest**](CreateQueryLambdaRequest.md) | JSON object | 

### Return type

[**QueryLambdaVersionResponse**](QueryLambdaVersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQueryLambdaTag

> QueryLambdaTagResponse CreateQueryLambdaTag(ctx, workspace, queryLambda).Body(body).Execute()

Create Query Lambda Tag



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
    queryLambda := "queryLambda_example" // string | name of the Query Lambda
    body := *openapiclient.NewCreateQueryLambdaTagRequest("production", "123ABC") // CreateQueryLambdaTagRequest | JSON object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QueryLambdasApi.CreateQueryLambdaTag(context.Background(), workspace, queryLambda).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryLambdasApi.CreateQueryLambdaTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateQueryLambdaTag`: QueryLambdaTagResponse
    fmt.Fprintf(os.Stdout, "Response from `QueryLambdasApi.CreateQueryLambdaTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**queryLambda** | **string** | name of the Query Lambda | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateQueryLambdaTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**CreateQueryLambdaTagRequest**](CreateQueryLambdaTagRequest.md) | JSON object | 

### Return type

[**QueryLambdaTagResponse**](QueryLambdaTagResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQueryLambda

> DeleteQueryLambdaResponse DeleteQueryLambda(ctx, workspace, queryLambda).Execute()

Delete Query Lambda



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
    queryLambda := "queryLambda_example" // string | name of the Query Lambda

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QueryLambdasApi.DeleteQueryLambda(context.Background(), workspace, queryLambda).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryLambdasApi.DeleteQueryLambda``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteQueryLambda`: DeleteQueryLambdaResponse
    fmt.Fprintf(os.Stdout, "Response from `QueryLambdasApi.DeleteQueryLambda`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**queryLambda** | **string** | name of the Query Lambda | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQueryLambdaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteQueryLambdaResponse**](DeleteQueryLambdaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQueryLambdaTag

> QueryLambdaTagResponse DeleteQueryLambdaTag(ctx, workspace, queryLambda, tag).Execute()

Delete Query Lambda Tag Version



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
    queryLambda := "queryLambda_example" // string | name of the Query Lambda
    tag := "tag_example" // string | name of the tag

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QueryLambdasApi.DeleteQueryLambdaTag(context.Background(), workspace, queryLambda, tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryLambdasApi.DeleteQueryLambdaTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteQueryLambdaTag`: QueryLambdaTagResponse
    fmt.Fprintf(os.Stdout, "Response from `QueryLambdasApi.DeleteQueryLambdaTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**queryLambda** | **string** | name of the Query Lambda | 
**tag** | **string** | name of the tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQueryLambdaTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**QueryLambdaTagResponse**](QueryLambdaTagResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQueryLambdaVersion

> QueryLambdaVersionResponse DeleteQueryLambdaVersion(ctx, workspace, queryLambda, version).Execute()

Delete Query Lambda Version



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
    queryLambda := "queryLambda_example" // string | name of the Query Lambda
    version := "version_example" // string | version

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QueryLambdasApi.DeleteQueryLambdaVersion(context.Background(), workspace, queryLambda, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryLambdasApi.DeleteQueryLambdaVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteQueryLambdaVersion`: QueryLambdaVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `QueryLambdasApi.DeleteQueryLambdaVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**queryLambda** | **string** | name of the Query Lambda | 
**version** | **string** | version | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQueryLambdaVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**QueryLambdaVersionResponse**](QueryLambdaVersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteQueryLambda

> QueryResponse ExecuteQueryLambda(ctx, workspace, queryLambda, version).Body(body).Execute()

Execute Query Lambda



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
    queryLambda := "queryLambda_example" // string | name of the Query Lambda
    version := "version_example" // string | version
    body := *openapiclient.NewExecuteQueryLambdaRequest() // ExecuteQueryLambdaRequest | JSON object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QueryLambdasApi.ExecuteQueryLambda(context.Background(), workspace, queryLambda, version).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryLambdasApi.ExecuteQueryLambda``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteQueryLambda`: QueryResponse
    fmt.Fprintf(os.Stdout, "Response from `QueryLambdasApi.ExecuteQueryLambda`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**queryLambda** | **string** | name of the Query Lambda | 
**version** | **string** | version | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteQueryLambdaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**ExecuteQueryLambdaRequest**](ExecuteQueryLambdaRequest.md) | JSON object | 

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


## ExecuteQueryLambdaByTag

> QueryResponse ExecuteQueryLambdaByTag(ctx, workspace, queryLambda, tag).Body(body).Execute()

Execute Query Lambda By Tag



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
    queryLambda := "queryLambda_example" // string | name of the Query Lambda
    tag := "tag_example" // string | tag
    body := *openapiclient.NewExecuteQueryLambdaRequest() // ExecuteQueryLambdaRequest | JSON object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QueryLambdasApi.ExecuteQueryLambdaByTag(context.Background(), workspace, queryLambda, tag).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryLambdasApi.ExecuteQueryLambdaByTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteQueryLambdaByTag`: QueryResponse
    fmt.Fprintf(os.Stdout, "Response from `QueryLambdasApi.ExecuteQueryLambdaByTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**queryLambda** | **string** | name of the Query Lambda | 
**tag** | **string** | tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteQueryLambdaByTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**ExecuteQueryLambdaRequest**](ExecuteQueryLambdaRequest.md) | JSON object | 

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


## GetQueryLambdaTagVersion

> QueryLambdaTagResponse GetQueryLambdaTagVersion(ctx, workspace, queryLambda, tag).Execute()

Get Query Lambda Tag



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
    queryLambda := "queryLambda_example" // string | name of the Query Lambda
    tag := "tag_example" // string | name of the tag

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QueryLambdasApi.GetQueryLambdaTagVersion(context.Background(), workspace, queryLambda, tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryLambdasApi.GetQueryLambdaTagVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQueryLambdaTagVersion`: QueryLambdaTagResponse
    fmt.Fprintf(os.Stdout, "Response from `QueryLambdasApi.GetQueryLambdaTagVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**queryLambda** | **string** | name of the Query Lambda | 
**tag** | **string** | name of the tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueryLambdaTagVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**QueryLambdaTagResponse**](QueryLambdaTagResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueryLambdaVersion

> QueryLambdaVersionResponse GetQueryLambdaVersion(ctx, workspace, queryLambda, version).Execute()

Get Query Lambda Version



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
    queryLambda := "queryLambda_example" // string | name of the Query Lambda
    version := "version_example" // string | version

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QueryLambdasApi.GetQueryLambdaVersion(context.Background(), workspace, queryLambda, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryLambdasApi.GetQueryLambdaVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQueryLambdaVersion`: QueryLambdaVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `QueryLambdasApi.GetQueryLambdaVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**queryLambda** | **string** | name of the Query Lambda | 
**version** | **string** | version | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueryLambdaVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**QueryLambdaVersionResponse**](QueryLambdaVersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllQueryLambdas

> ListQueryLambdasResponse ListAllQueryLambdas(ctx).Execute()

List Query Lambdas



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
    resp, r, err := api_client.QueryLambdasApi.ListAllQueryLambdas(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryLambdasApi.ListAllQueryLambdas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllQueryLambdas`: ListQueryLambdasResponse
    fmt.Fprintf(os.Stdout, "Response from `QueryLambdasApi.ListAllQueryLambdas`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAllQueryLambdasRequest struct via the builder pattern


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


## ListOrganizationTags

> ListQueryLambdaTagsResponse ListOrganizationTags(ctx).Execute()

List All Query Lambda Tags



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
    resp, r, err := api_client.QueryLambdasApi.ListOrganizationTags(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryLambdasApi.ListOrganizationTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrganizationTags`: ListQueryLambdaTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `QueryLambdasApi.ListOrganizationTags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationTagsRequest struct via the builder pattern


### Return type

[**ListQueryLambdaTagsResponse**](ListQueryLambdaTagsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQueryLambdaTagVersions

> ListQueryLambdaVersionsResponse ListQueryLambdaTagVersions(ctx, tag).Execute()

List Query Lambda Tag Versions



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
    tag := "tag_example" // string | name of the tag

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QueryLambdasApi.ListQueryLambdaTagVersions(context.Background(), tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryLambdasApi.ListQueryLambdaTagVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListQueryLambdaTagVersions`: ListQueryLambdaVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `QueryLambdasApi.ListQueryLambdaTagVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tag** | **string** | name of the tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiListQueryLambdaTagVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListQueryLambdaVersionsResponse**](ListQueryLambdaVersionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQueryLambdaTags

> ListQueryLambdaTagsResponse ListQueryLambdaTags(ctx, workspace, queryLambda).Execute()

List Query Lambda Tags



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
    queryLambda := "queryLambda_example" // string | name of the Query Lambda

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QueryLambdasApi.ListQueryLambdaTags(context.Background(), workspace, queryLambda).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryLambdasApi.ListQueryLambdaTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListQueryLambdaTags`: ListQueryLambdaTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `QueryLambdasApi.ListQueryLambdaTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**queryLambda** | **string** | name of the Query Lambda | 

### Other Parameters

Other parameters are passed through a pointer to a apiListQueryLambdaTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListQueryLambdaTagsResponse**](ListQueryLambdaTagsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQueryLambdaVersions

> ListQueryLambdaVersionsResponse ListQueryLambdaVersions(ctx, workspace, queryLambda).Execute()

List Query Lambda Versions



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
    queryLambda := "queryLambda_example" // string | name of the Query Lambda

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QueryLambdasApi.ListQueryLambdaVersions(context.Background(), workspace, queryLambda).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryLambdasApi.ListQueryLambdaVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListQueryLambdaVersions`: ListQueryLambdaVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `QueryLambdasApi.ListQueryLambdaVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**queryLambda** | **string** | name of the Query Lambda | 

### Other Parameters

Other parameters are passed through a pointer to a apiListQueryLambdaVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListQueryLambdaVersionsResponse**](ListQueryLambdaVersionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQueryLambdasInWorkspace

> ListQueryLambdasResponse ListQueryLambdasInWorkspace(ctx, workspace).Execute()

List Query Lambdas in Workspace



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
    resp, r, err := api_client.QueryLambdasApi.ListQueryLambdasInWorkspace(context.Background(), workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryLambdasApi.ListQueryLambdasInWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListQueryLambdasInWorkspace`: ListQueryLambdasResponse
    fmt.Fprintf(os.Stdout, "Response from `QueryLambdasApi.ListQueryLambdasInWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiListQueryLambdasInWorkspaceRequest struct via the builder pattern


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


## UpdateQueryLambda

> QueryLambdaVersionResponse UpdateQueryLambda(ctx, workspace, queryLambda).Body(body).Create(create).Execute()

Update Query Lambda



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
    queryLambda := "queryLambda_example" // string | name of the Query Lambda
    body := *openapiclient.NewUpdateQueryLambdaRequest() // UpdateQueryLambdaRequest | JSON object
    create := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QueryLambdasApi.UpdateQueryLambda(context.Background(), workspace, queryLambda).Body(body).Create(create).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryLambdasApi.UpdateQueryLambda``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateQueryLambda`: QueryLambdaVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `QueryLambdasApi.UpdateQueryLambda`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**queryLambda** | **string** | name of the Query Lambda | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQueryLambdaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateQueryLambdaRequest**](UpdateQueryLambdaRequest.md) | JSON object | 
 **create** | **bool** |  | 

### Return type

[**QueryLambdaVersionResponse**](QueryLambdaVersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

