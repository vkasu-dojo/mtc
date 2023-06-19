# \ChildrenApi

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChildCustomContent**](ChildrenApi.md#GetChildCustomContent) | **Get** /custom-content/{id}/children | Get child custom content
[**GetChildPages**](ChildrenApi.md#GetChildPages) | **Get** /pages/{id}/children | Get child pages



## GetChildCustomContent

> MultiEntityResultChildCustomContent GetChildCustomContent(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get child custom content



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
    id := int64(789) // int64 | The ID of the parent custom content. If you don't know the custom content ID, use Get custom-content and filter the results.
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
    sort := []openapiclient.ChildCustomContentSortOrder{openapiclient.ChildCustomContentSortOrder("created-date")} // []ChildCustomContentSortOrder | Used to sort the result by a particular field. (optional)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChildrenApi.GetChildCustomContent(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChildrenApi.GetChildCustomContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChildCustomContent`: MultiEntityResultChildCustomContent
    fmt.Fprintf(os.Stdout, "Response from `ChildrenApi.GetChildCustomContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the parent custom content. If you don&#39;t know the custom content ID, use Get custom-content and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChildCustomContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **sort** | [**[]ChildCustomContentSortOrder**](ChildCustomContentSortOrder.md) | Used to sort the result by a particular field. | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultChildCustomContent**](MultiEntityResultChildCustomContent.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChildPages

> MultiEntityResultChildPage GetChildPages(ctx, id).Cursor(cursor).Limit(limit).Sort(sort).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get child pages



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
    id := int64(789) // int64 | The ID of the parent page. If you don't know the page ID, use Get pages and filter the results.
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
    sort := []openapiclient.ChildPageSortOrder{openapiclient.ChildPageSortOrder("created-date")} // []ChildPageSortOrder | Used to sort the result by a particular field. (optional)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChildrenApi.GetChildPages(context.Background(), id).Cursor(cursor).Limit(limit).Sort(sort).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChildrenApi.GetChildPages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChildPages`: MultiEntityResultChildPage
    fmt.Fprintf(os.Stdout, "Response from `ChildrenApi.GetChildPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the parent page. If you don&#39;t know the page ID, use Get pages and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChildPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **sort** | [**[]ChildPageSortOrder**](ChildPageSortOrder.md) | Used to sort the result by a particular field. | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultChildPage**](MultiEntityResultChildPage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

