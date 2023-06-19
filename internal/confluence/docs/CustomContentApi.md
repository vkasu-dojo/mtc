# \CustomContentApi

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomContent**](CustomContentApi.md#CreateCustomContent) | **Post** /custom-content | Create custom content
[**DeleteCustomContent**](CustomContentApi.md#DeleteCustomContent) | **Delete** /custom-content/{id} | Delete custom content
[**GetCustomContentById**](CustomContentApi.md#GetCustomContentById) | **Get** /custom-content/{id} | Get custom content by id
[**GetCustomContentByType**](CustomContentApi.md#GetCustomContentByType) | **Get** /custom-content | Get custom content by type
[**GetCustomContentByTypeInBlogPost**](CustomContentApi.md#GetCustomContentByTypeInBlogPost) | **Get** /blogposts/{id}/custom-content | Get custom content by type in blog post
[**GetCustomContentByTypeInPage**](CustomContentApi.md#GetCustomContentByTypeInPage) | **Get** /pages/{id}/custom-content | Get custom content by type in page
[**GetCustomContentByTypeInSpace**](CustomContentApi.md#GetCustomContentByTypeInSpace) | **Get** /spaces/{id}/custom-content | Get custom content by type in space
[**UpdateCustomContent**](CustomContentApi.md#UpdateCustomContent) | **Put** /custom-content/{id} | Update custom content



## CreateCustomContent

> CustomContent CreateCustomContent(ctx).CreateCustomContentRequest(createCustomContentRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Create custom content



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
    createCustomContentRequest := *openapiclient.NewCreateCustomContentRequest("Type_example", "Title_example", openapiclient.createCustomContent_request_body{CustomContentBodyWrite: openapiclient.NewCustomContentBodyWrite()}) // CreateCustomContentRequest | 
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomContentApi.CreateCustomContent(context.Background()).CreateCustomContentRequest(createCustomContentRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomContentApi.CreateCustomContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomContent`: CustomContent
    fmt.Fprintf(os.Stdout, "Response from `CustomContentApi.CreateCustomContent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomContentRequest** | [**CreateCustomContentRequest**](CreateCustomContentRequest.md) |  | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**CustomContent**](CustomContent.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomContent

> DeleteCustomContent(ctx, id).Execute()

Delete custom content



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
    id := int64(789) // int64 | The ID of the custom content to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomContentApi.DeleteCustomContent(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomContentApi.DeleteCustomContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the custom content to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomContentRequest struct via the builder pattern


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


## GetCustomContentById

> CustomContent GetCustomContentById(ctx, id).BodyFormat(bodyFormat).Version(version).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get custom content by id



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
    id := int64(789) // int64 | The ID of the custom content to be returned. If you don't know the custom content ID, use Get Custom Content by Type and filter the results.
    bodyFormat := openapiclient.CustomContentBodyRepresentation("raw") // CustomContentBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field.  Note: If the custom content body type is `storage`, the `storage` and `atlas_doc_format` body formats are able to be returned. If the custom content body type is `raw`, only the `raw` body format is able to be returned. (optional)
    version := int32(56) // int32 | Allows you to retrieve a previously published version. Specify the previous version's number to retrieve its details. (optional)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomContentApi.GetCustomContentById(context.Background(), id).BodyFormat(bodyFormat).Version(version).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomContentApi.GetCustomContentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomContentById`: CustomContent
    fmt.Fprintf(os.Stdout, "Response from `CustomContentApi.GetCustomContentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the custom content to be returned. If you don&#39;t know the custom content ID, use Get Custom Content by Type and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomContentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**CustomContentBodyRepresentation**](CustomContentBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field.  Note: If the custom content body type is &#x60;storage&#x60;, the &#x60;storage&#x60; and &#x60;atlas_doc_format&#x60; body formats are able to be returned. If the custom content body type is &#x60;raw&#x60;, only the &#x60;raw&#x60; body format is able to be returned. | 
 **version** | **int32** | Allows you to retrieve a previously published version. Specify the previous version&#39;s number to retrieve its details. | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**CustomContent**](CustomContent.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomContentByType

> MultiEntityResultCustomContent GetCustomContentByType(ctx).Type_(type_).Id(id).Cursor(cursor).Limit(limit).BodyFormat(bodyFormat).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get custom content by type



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
    type_ := "type__example" // string | The type of custom content being requested. See: https://developer.atlassian.com/cloud/confluence/custom-content/ for additional details on custom content.
    id := []int64{int64(123)} // []int64 | Filter the results based on custom content ids. Multiple custom content ids can be specified as a comma-separated list. (optional)
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
    bodyFormat := openapiclient.CustomContentBodyRepresentation("raw") // CustomContentBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field.  Note: If the custom content body type is `storage`, the `storage` and `atlas_doc_format` body formats are able to be returned. If the custom content body type is `raw`, only the `raw` body format is able to be returned. (optional)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomContentApi.GetCustomContentByType(context.Background()).Type_(type_).Id(id).Cursor(cursor).Limit(limit).BodyFormat(bodyFormat).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomContentApi.GetCustomContentByType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomContentByType`: MultiEntityResultCustomContent
    fmt.Fprintf(os.Stdout, "Response from `CustomContentApi.GetCustomContentByType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomContentByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | The type of custom content being requested. See: https://developer.atlassian.com/cloud/confluence/custom-content/ for additional details on custom content. | 
 **id** | **[]int64** | Filter the results based on custom content ids. Multiple custom content ids can be specified as a comma-separated list. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **bodyFormat** | [**CustomContentBodyRepresentation**](CustomContentBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field.  Note: If the custom content body type is &#x60;storage&#x60;, the &#x60;storage&#x60; and &#x60;atlas_doc_format&#x60; body formats are able to be returned. If the custom content body type is &#x60;raw&#x60;, only the &#x60;raw&#x60; body format is able to be returned. | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultCustomContent**](MultiEntityResultCustomContent.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomContentByTypeInBlogPost

> MultiEntityResultCustomContent GetCustomContentByTypeInBlogPost(ctx, id).Type_(type_).Cursor(cursor).Limit(limit).BodyFormat(bodyFormat).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get custom content by type in blog post



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
    id := int64(789) // int64 | The ID of the blog post for which custom content should be returned.
    type_ := "type__example" // string | The type of custom content being requested. See: https://developer.atlassian.com/cloud/confluence/custom-content/ for additional details on custom content.
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
    bodyFormat := openapiclient.CustomContentBodyRepresentation("raw") // CustomContentBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field.  Note: If the custom content body type is `storage`, the `storage` and `atlas_doc_format` body formats are able to be returned. If the custom content body type is `raw`, only the `raw` body format is able to be returned. (optional)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomContentApi.GetCustomContentByTypeInBlogPost(context.Background(), id).Type_(type_).Cursor(cursor).Limit(limit).BodyFormat(bodyFormat).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomContentApi.GetCustomContentByTypeInBlogPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomContentByTypeInBlogPost`: MultiEntityResultCustomContent
    fmt.Fprintf(os.Stdout, "Response from `CustomContentApi.GetCustomContentByTypeInBlogPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the blog post for which custom content should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomContentByTypeInBlogPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | The type of custom content being requested. See: https://developer.atlassian.com/cloud/confluence/custom-content/ for additional details on custom content. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **bodyFormat** | [**CustomContentBodyRepresentation**](CustomContentBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field.  Note: If the custom content body type is &#x60;storage&#x60;, the &#x60;storage&#x60; and &#x60;atlas_doc_format&#x60; body formats are able to be returned. If the custom content body type is &#x60;raw&#x60;, only the &#x60;raw&#x60; body format is able to be returned. | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultCustomContent**](MultiEntityResultCustomContent.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomContentByTypeInPage

> MultiEntityResultCustomContent GetCustomContentByTypeInPage(ctx, id).Type_(type_).Cursor(cursor).Limit(limit).BodyFormat(bodyFormat).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get custom content by type in page



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
    id := int64(789) // int64 | The ID of the page for which custom content should be returned.
    type_ := "type__example" // string | The type of custom content being requested. See: https://developer.atlassian.com/cloud/confluence/custom-content/ for additional details on custom content.
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
    bodyFormat := openapiclient.CustomContentBodyRepresentation("raw") // CustomContentBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field.  Note: If the custom content body type is `storage`, the `storage` and `atlas_doc_format` body formats are able to be returned. If the custom content body type is `raw`, only the `raw` body format is able to be returned. (optional)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomContentApi.GetCustomContentByTypeInPage(context.Background(), id).Type_(type_).Cursor(cursor).Limit(limit).BodyFormat(bodyFormat).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomContentApi.GetCustomContentByTypeInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomContentByTypeInPage`: MultiEntityResultCustomContent
    fmt.Fprintf(os.Stdout, "Response from `CustomContentApi.GetCustomContentByTypeInPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the page for which custom content should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomContentByTypeInPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | The type of custom content being requested. See: https://developer.atlassian.com/cloud/confluence/custom-content/ for additional details on custom content. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **bodyFormat** | [**CustomContentBodyRepresentation**](CustomContentBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field.  Note: If the custom content body type is &#x60;storage&#x60;, the &#x60;storage&#x60; and &#x60;atlas_doc_format&#x60; body formats are able to be returned. If the custom content body type is &#x60;raw&#x60;, only the &#x60;raw&#x60; body format is able to be returned. | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultCustomContent**](MultiEntityResultCustomContent.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomContentByTypeInSpace

> MultiEntityResultCustomContent GetCustomContentByTypeInSpace(ctx, id).Type_(type_).Cursor(cursor).Limit(limit).BodyFormat(bodyFormat).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get custom content by type in space



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
    id := int64(789) // int64 | The ID of the space for which custom content should be returned.
    type_ := "type__example" // string | The type of custom content being requested. See: https://developer.atlassian.com/cloud/confluence/custom-content/ for additional details on custom content.
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    limit := int32(56) // int32 | Maximum number of pages per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
    bodyFormat := openapiclient.CustomContentBodyRepresentation("raw") // CustomContentBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field.  Note: If the custom content body type is `storage`, the `storage` and `atlas_doc_format` body formats are able to be returned. If the custom content body type is `raw`, only the `raw` body format is able to be returned. (optional)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomContentApi.GetCustomContentByTypeInSpace(context.Background(), id).Type_(type_).Cursor(cursor).Limit(limit).BodyFormat(bodyFormat).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomContentApi.GetCustomContentByTypeInSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomContentByTypeInSpace`: MultiEntityResultCustomContent
    fmt.Fprintf(os.Stdout, "Response from `CustomContentApi.GetCustomContentByTypeInSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the space for which custom content should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomContentByTypeInSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | The type of custom content being requested. See: https://developer.atlassian.com/cloud/confluence/custom-content/ for additional details on custom content. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of pages per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **bodyFormat** | [**CustomContentBodyRepresentation**](CustomContentBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field.  Note: If the custom content body type is &#x60;storage&#x60;, the &#x60;storage&#x60; and &#x60;atlas_doc_format&#x60; body formats are able to be returned. If the custom content body type is &#x60;raw&#x60;, only the &#x60;raw&#x60; body format is able to be returned. | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultCustomContent**](MultiEntityResultCustomContent.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomContent

> CustomContent UpdateCustomContent(ctx, id).UpdateCustomContentRequest(updateCustomContentRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Update custom content



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
    id := int64(789) // int64 | The ID of the custom content to be updated. If you don't know the custom content ID, use Get Custom Content by Type and filter the results.
    updateCustomContentRequest := *openapiclient.NewUpdateCustomContentRequest("Id_example", "Type_example", "Status_example", "Title_example", openapiclient.createCustomContent_request_body{CustomContentBodyWrite: openapiclient.NewCustomContentBodyWrite()}, *openapiclient.NewUpdateBlogPostRequestVersion()) // UpdateCustomContentRequest | 
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomContentApi.UpdateCustomContent(context.Background(), id).UpdateCustomContentRequest(updateCustomContentRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomContentApi.UpdateCustomContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomContent`: CustomContent
    fmt.Fprintf(os.Stdout, "Response from `CustomContentApi.UpdateCustomContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the custom content to be updated. If you don&#39;t know the custom content ID, use Get Custom Content by Type and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomContentRequest** | [**UpdateCustomContentRequest**](UpdateCustomContentRequest.md) |  | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**CustomContent**](CustomContent.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

