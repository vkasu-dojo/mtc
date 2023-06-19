# \AttachmentApi

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAttachmentById**](AttachmentApi.md#GetAttachmentById) | **Get** /attachments/{id} | Get attachment by id
[**GetBlogpostAttachments**](AttachmentApi.md#GetBlogpostAttachments) | **Get** /blogposts/{id}/attachments | Get attachments for blog post
[**GetCustomContentAttachments**](AttachmentApi.md#GetCustomContentAttachments) | **Get** /custom-content/{id}/attachments | Get attachments for custom content
[**GetLabelAttachments**](AttachmentApi.md#GetLabelAttachments) | **Get** /labels/{id}/attachments | Get attachments for label
[**GetPageAttachments**](AttachmentApi.md#GetPageAttachments) | **Get** /pages/{id}/attachments | Get attachments for page



## GetAttachmentById

> Attachment GetAttachmentById(ctx, id).Version(version).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get attachment by id



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
    id := "id_example" // string | The ID of the attachment to be returned. If you don't know the attachment's ID, use Get attachments for page/blogpost/custom content.
    version := int32(56) // int32 | Allows you to retrieve a previously published version. Specify the previous version's number to retrieve its details. (optional)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttachmentApi.GetAttachmentById(context.Background(), id).Version(version).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentApi.GetAttachmentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttachmentById`: Attachment
    fmt.Fprintf(os.Stdout, "Response from `AttachmentApi.GetAttachmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the attachment to be returned. If you don&#39;t know the attachment&#39;s ID, use Get attachments for page/blogpost/custom content. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** | Allows you to retrieve a previously published version. Specify the previous version&#39;s number to retrieve its details. | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**Attachment**](Attachment.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlogpostAttachments

> MultiEntityResultAttachment GetBlogpostAttachments(ctx, id).Sort(sort).Cursor(cursor).MediaType(mediaType).Filename(filename).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get attachments for blog post



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
    id := int64(789) // int64 | The ID of the blog post for which attachments should be returned.
    sort := openapiclient.AttachmentSortOrder("created-date") // AttachmentSortOrder | Used to sort the result by a particular field. (optional)
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    mediaType := "mediaType_example" // string | Filters on the mediaType of attachments. Only one may be specified. (optional)
    filename := "filename_example" // string | Filters on the file-name of attachments. Only one may be specified. (optional)
    limit := int32(56) // int32 | Maximum number of attachments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 50)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttachmentApi.GetBlogpostAttachments(context.Background(), id).Sort(sort).Cursor(cursor).MediaType(mediaType).Filename(filename).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentApi.GetBlogpostAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlogpostAttachments`: MultiEntityResultAttachment
    fmt.Fprintf(os.Stdout, "Response from `AttachmentApi.GetBlogpostAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the blog post for which attachments should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogpostAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | [**AttachmentSortOrder**](AttachmentSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **mediaType** | **string** | Filters on the mediaType of attachments. Only one may be specified. | 
 **filename** | **string** | Filters on the file-name of attachments. Only one may be specified. | 
 **limit** | **int32** | Maximum number of attachments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 50]
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultAttachment**](MultiEntityResultAttachment.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomContentAttachments

> MultiEntityResultAttachment GetCustomContentAttachments(ctx, id).Sort(sort).Cursor(cursor).MediaType(mediaType).Filename(filename).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get attachments for custom content



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
    id := int64(789) // int64 | The ID of the custom content for which attachments should be returned.
    sort := openapiclient.AttachmentSortOrder("created-date") // AttachmentSortOrder | Used to sort the result by a particular field. (optional)
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    mediaType := "mediaType_example" // string | Filters on the mediaType of attachments. Only one may be specified. (optional)
    filename := "filename_example" // string | Filters on the file-name of attachments. Only one may be specified. (optional)
    limit := int32(56) // int32 | Maximum number of attachments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 50)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttachmentApi.GetCustomContentAttachments(context.Background(), id).Sort(sort).Cursor(cursor).MediaType(mediaType).Filename(filename).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentApi.GetCustomContentAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomContentAttachments`: MultiEntityResultAttachment
    fmt.Fprintf(os.Stdout, "Response from `AttachmentApi.GetCustomContentAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the custom content for which attachments should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomContentAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | [**AttachmentSortOrder**](AttachmentSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **mediaType** | **string** | Filters on the mediaType of attachments. Only one may be specified. | 
 **filename** | **string** | Filters on the file-name of attachments. Only one may be specified. | 
 **limit** | **int32** | Maximum number of attachments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 50]
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultAttachment**](MultiEntityResultAttachment.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLabelAttachments

> MultiEntityResultAttachment1 GetLabelAttachments(ctx, id).Sort(sort).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get attachments for label



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
    id := int64(789) // int64 | The ID of the label for which attachments should be returned.
    sort := openapiclient.AttachmentSortOrder("created-date") // AttachmentSortOrder | Used to sort the result by a particular field. (optional)
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttachmentApi.GetLabelAttachments(context.Background(), id).Sort(sort).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentApi.GetLabelAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLabelAttachments`: MultiEntityResultAttachment1
    fmt.Fprintf(os.Stdout, "Response from `AttachmentApi.GetLabelAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the label for which attachments should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | [**AttachmentSortOrder**](AttachmentSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultAttachment1**](MultiEntityResultAttachment1.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageAttachments

> MultiEntityResultAttachment GetPageAttachments(ctx, id).Sort(sort).Cursor(cursor).MediaType(mediaType).Filename(filename).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get attachments for page



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
    id := int64(789) // int64 | The ID of the page for which attachments should be returned.
    sort := openapiclient.AttachmentSortOrder("created-date") // AttachmentSortOrder | Used to sort the result by a particular field. (optional)
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    mediaType := "mediaType_example" // string | Filters on the mediaType of attachments. Only one may be specified. (optional)
    filename := "filename_example" // string | Filters on the file-name of attachments. Only one may be specified. (optional)
    limit := int32(56) // int32 | Maximum number of attachments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 50)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttachmentApi.GetPageAttachments(context.Background(), id).Sort(sort).Cursor(cursor).MediaType(mediaType).Filename(filename).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentApi.GetPageAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPageAttachments`: MultiEntityResultAttachment
    fmt.Fprintf(os.Stdout, "Response from `AttachmentApi.GetPageAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page for which attachments should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | [**AttachmentSortOrder**](AttachmentSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **mediaType** | **string** | Filters on the mediaType of attachments. Only one may be specified. | 
 **filename** | **string** | Filters on the file-name of attachments. Only one may be specified. | 
 **limit** | **int32** | Maximum number of attachments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 50]
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultAttachment**](MultiEntityResultAttachment.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

