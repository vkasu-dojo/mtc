# \BlogPostApi

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlogPost**](BlogPostApi.md#CreateBlogPost) | **Post** /blogposts | Create blog post
[**DeleteBlogPost**](BlogPostApi.md#DeleteBlogPost) | **Delete** /blogposts/{id} | Delete blog post
[**GetBlogPostById**](BlogPostApi.md#GetBlogPostById) | **Get** /blogposts/{id} | Get blog post by id
[**GetBlogPosts**](BlogPostApi.md#GetBlogPosts) | **Get** /blogposts | Get blog posts
[**GetBlogPostsInSpace**](BlogPostApi.md#GetBlogPostsInSpace) | **Get** /spaces/{id}/blogposts | Get blog posts in space
[**GetLabelBlogPosts**](BlogPostApi.md#GetLabelBlogPosts) | **Get** /labels/{id}/blogposts | Get blog posts for label
[**UpdateBlogPost**](BlogPostApi.md#UpdateBlogPost) | **Put** /blogposts/{id} | Update blog post



## CreateBlogPost

> BlogPost CreateBlogPost(ctx).CreateBlogPostRequest(createBlogPostRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Create blog post



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
    createBlogPostRequest := *openapiclient.NewCreateBlogPostRequest("SpaceId_example") // CreateBlogPostRequest | 
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostApi.CreateBlogPost(context.Background()).CreateBlogPostRequest(createBlogPostRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostApi.CreateBlogPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBlogPost`: BlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostApi.CreateBlogPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlogPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBlogPostRequest** | [**CreateBlogPostRequest**](CreateBlogPostRequest.md) |  | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlogPost

> DeleteBlogPost(ctx, id).Execute()

Delete blog post



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
    id := int64(789) // int64 | The ID of the blog post to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BlogPostApi.DeleteBlogPost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostApi.DeleteBlogPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the blog post to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlogPostRequest struct via the builder pattern


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


## GetBlogPostById

> BlogPost GetBlogPostById(ctx, id).BodyFormat(bodyFormat).GetDraft(getDraft).Version(version).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get blog post by id



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
    id := int64(789) // int64 | The ID of the blog post to be returned. If you don't know the blog post ID, use Get blog posts and filter the results.
    bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
    getDraft := true // bool | Retrieve the draft version of this blog post. (optional) (default to false)
    version := int32(56) // int32 | Allows you to retrieve a previously published version. Specify the previous version's number to retrieve its details. (optional)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostApi.GetBlogPostById(context.Background(), id).BodyFormat(bodyFormat).GetDraft(getDraft).Version(version).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostApi.GetBlogPostById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlogPostById`: BlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostApi.GetBlogPostById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the blog post to be returned. If you don&#39;t know the blog post ID, use Get blog posts and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogPostByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **getDraft** | **bool** | Retrieve the draft version of this blog post. | [default to false]
 **version** | **int32** | Allows you to retrieve a previously published version. Specify the previous version&#39;s number to retrieve its details. | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlogPosts

> MultiEntityResultBlogPost GetBlogPosts(ctx).Id(id).Status(status).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get blog posts



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
    id := []int64{int64(123)} // []int64 | Filter the results based on blog post ids. Multiple blog post ids can be specified as a comma-separated list. (optional)
    status := "status_example" // string | Filter the results to blog posts based on their status. (optional)
    bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    limit := int32(56) // int32 | Maximum number of blog posts per result to return. If more results exist, use the `Link` response header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostApi.GetBlogPosts(context.Background()).Id(id).Status(status).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostApi.GetBlogPosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlogPosts`: MultiEntityResultBlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostApi.GetBlogPosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]int64** | Filter the results based on blog post ids. Multiple blog post ids can be specified as a comma-separated list. | 
 **status** | **string** | Filter the results to blog posts based on their status. | 
 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of blog posts per result to return. If more results exist, use the &#x60;Link&#x60; response header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultBlogPost**](MultiEntityResultBlogPost.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlogPostsInSpace

> MultiEntityResultBlogPost GetBlogPostsInSpace(ctx, id).Status(status).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get blog posts in space



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
    id := int64(789) // int64 | The ID of the space for which blog posts should be returned.
    status := "status_example" // string | Filter the results to blog posts based on their status. (optional)
    bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    limit := int32(56) // int32 | Maximum number of blog posts per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostApi.GetBlogPostsInSpace(context.Background(), id).Status(status).BodyFormat(bodyFormat).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostApi.GetBlogPostsInSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlogPostsInSpace`: MultiEntityResultBlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostApi.GetBlogPostsInSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the space for which blog posts should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogPostsInSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter the results to blog posts based on their status. | 
 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of blog posts per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultBlogPost**](MultiEntityResultBlogPost.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLabelBlogPosts

> MultiEntityResultBlogPost GetLabelBlogPosts(ctx, id).BodyFormat(bodyFormat).Sort(sort).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get blog posts for label



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
    id := int64(789) // int64 | The ID of the label for which blog posts should be returned.
    bodyFormat := openapiclient.PrimaryBodyRepresentation("storage") // PrimaryBodyRepresentation | The content format types to be returned in the `body` field of the response. If available, the representation will be available under a response field of the same name under the `body` field. (optional)
    sort := openapiclient.BlogPostSortOrder("id") // BlogPostSortOrder | Used to sort the result by a particular field. (optional)
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    limit := int32(56) // int32 | Maximum number of blog posts per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostApi.GetLabelBlogPosts(context.Background(), id).BodyFormat(bodyFormat).Sort(sort).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostApi.GetLabelBlogPosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLabelBlogPosts`: MultiEntityResultBlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostApi.GetLabelBlogPosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the label for which blog posts should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelBlogPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyFormat** | [**PrimaryBodyRepresentation**](PrimaryBodyRepresentation.md) | The content format types to be returned in the &#x60;body&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;body&#x60; field. | 
 **sort** | [**BlogPostSortOrder**](BlogPostSortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of blog posts per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultBlogPost**](MultiEntityResultBlogPost.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBlogPost

> BlogPost UpdateBlogPost(ctx, id).UpdateBlogPostRequest(updateBlogPostRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Update blog post



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
    id := int64(789) // int64 | The ID of the blog post to be updated. If you don't know the blog post ID, use Get Blog Posts and filter the results.
    updateBlogPostRequest := *openapiclient.NewUpdateBlogPostRequest("Id_example", "Status_example", "Title_example", "SpaceId_example", openapiclient.createBlogPost_request_body{BlogPostBodyWrite: openapiclient.NewBlogPostBodyWrite()}, *openapiclient.NewUpdateBlogPostRequestVersion()) // UpdateBlogPostRequest | 
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlogPostApi.UpdateBlogPost(context.Background(), id).UpdateBlogPostRequest(updateBlogPostRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlogPostApi.UpdateBlogPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBlogPost`: BlogPost
    fmt.Fprintf(os.Stdout, "Response from `BlogPostApi.UpdateBlogPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the blog post to be updated. If you don&#39;t know the blog post ID, use Get Blog Posts and filter the results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlogPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBlogPostRequest** | [**UpdateBlogPostRequest**](UpdateBlogPostRequest.md) |  | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

