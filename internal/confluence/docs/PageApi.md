# \PageApi

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePage**](PageApi.md#CreatePage) | **Post** /pages | Create page
[**DeletePage**](PageApi.md#DeletePage) | **Delete** /pages/{id} | Delete page
[**GetLabelPages**](PageApi.md#GetLabelPages) | **Get** /labels/{id}/pages | Get pages for label
[**GetPageById**](PageApi.md#GetPageById) | **Get** /pages/{id} | Get page by id
[**GetPages**](PageApi.md#GetPages) | **Get** /pages | Get pages
[**GetPagesInSpace**](PageApi.md#GetPagesInSpace) | **Get** /spaces/{id}/pages | Get pages in space
[**UpdatePage**](PageApi.md#UpdatePage) | **Put** /pages/{id} | Update page



## CreatePage

> Page CreatePage(ctx).CreatePageRequest(createPageRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Create page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "internal/confluence"
)

func main() {
    createPageRequest := *openapiclient.NewCreatePageRequest("SpaceId_example") // CreatePageRequest | 
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PageApi.CreatePage(context.Background()).CreatePageRequest(createPageRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageApi.CreatePage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePage`: Page
    fmt.Fprintf(os.Stdout, "Response from `PageApi.CreatePage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPageRequest** | [**CreatePageRequest**](CreatePageRequest.md) |  | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**Page**](Page.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePage

> DeletePage(ctx, id).Execute()

Delete page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "internal/confluence"
)

func main() {
    id := int64(789) // int64 | The ID of the page to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PageApi.DeletePage(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageApi.DeletePage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLabelPages

> MultiEntityResultPage GetLabelPages(ctx, id).BodyFormat(bodyFormat).Sort(sort).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get pages for label



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "internal/confluence"
)

func main() {
    id := int64(789) // int64 | The ID of the label for which pages should be returned.
    bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
    sort := openapiclient.PageSortOrder("id") // PageSortOrder | Used to sort the result by a particular field. (optional)
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PageApi.GetLabelPages(context.Background(), id).BodyFormat(bodyFormat).Sort(sort).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageApi.GetLabelPages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLabelPages`: MultiEntityResultPage
    fmt.Fprintf(os.Stdout, "Response from `PageApi.GetLabelPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the label for which pages should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **sort** | [**PageSortOrder**](PageSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultPage**](MultiEntityResultPage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageById

> Page GetPageById(ctx, id).BodyFormat(bodyFormat).GetDraft(getDraft).Version(version).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get page by id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "internal/confluence"
)

func main() {
    id := int64(789) // int64 | The ID of the page to be returned. If you don't know the page ID, use Get pages and filter the results.
    bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
    getDraft := true // bool | Retrieve the draft version of this page. (optional) (default to false)
    version := int32(56) // int32 | Allows you to retrieve a previously published version. Specify the previous version's number to retrieve its details. (optional)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PageApi.GetPageById(context.Background(), id).BodyFormat(bodyFormat).GetDraft(getDraft).Version(version).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageApi.GetPageById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPageById`: Page
    fmt.Fprintf(os.Stdout, "Response from `PageApi.GetPageById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page to be returned. If you don&#39;t know the page ID, use Get pages and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **getDraft** | **bool** | Retrieve the draft version of this page. | [default to false]
 **version** | **int32** | Allows you to retrieve a previously published version. Specify the previous version&#39;s number to retrieve its details. | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**Page**](Page.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPages

> MultiEntityResultPage GetPages(ctx).Id(id).Status(status).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get pages



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "internal/confluence"
)

func main() {
    id := []int64{int64(123)} // []int64 | Filter the results based on page ids. Multiple page ids can be specified as a comma-separated list. (optional)
    status := "status_example" // string | Filter the results to pages based on their status. (optional)
    bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PageApi.GetPages(context.Background()).Id(id).Status(status).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageApi.GetPages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPages`: MultiEntityResultPage
    fmt.Fprintf(os.Stdout, "Response from `PageApi.GetPages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]int64** | Filter the results based on page ids. Multiple page ids can be specified as a comma-separated list. | 
 **status** | **string** | Filter the results to pages based on their status. | 
 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultPage**](MultiEntityResultPage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesInSpace

> MultiEntityResultPage GetPagesInSpace(ctx, id).Status(status).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get pages in space



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "internal/confluence"
)

func main() {
    id := int64(789) // int64 | The ID of the space for which pages should be returned.
    status := "status_example" // string | Filter the results to pages based on their status. (optional)
    bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PageApi.GetPagesInSpace(context.Background(), id).Status(status).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageApi.GetPagesInSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesInSpace`: MultiEntityResultPage
    fmt.Fprintf(os.Stdout, "Response from `PageApi.GetPagesInSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the space for which pages should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesInSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter the results to pages based on their status. | 
 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultPage**](MultiEntityResultPage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePage

> Page UpdatePage(ctx, id).UpdatePageRequest(updatePageRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Update page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "internal/confluence"
)

func main() {
    id := int64(789) // int64 | The ID of the page to be updated. If you don't know the page ID, use Get Pages and filter the results.
    updatePageRequest := *openapiclient.NewUpdatePageRequest("Id_example", "Status_example", "Title_example", "SpaceId_example", openapiclient.createPage_request_body{PageBodyWrite: openapiclient.NewPageBodyWrite()}, *openapiclient.NewUpdateBlogPostRequestVersion()) // UpdatePageRequest | 
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PageApi.UpdatePage(context.Background(), id).UpdatePageRequest(updatePageRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageApi.UpdatePage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePage`: Page
    fmt.Fprintf(os.Stdout, "Response from `PageApi.UpdatePage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page to be updated. If you don&#39;t know the page ID, use Get Pages and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePageRequest** | [**UpdatePageRequest**](UpdatePageRequest.md) |  | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**Page**](Page.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

