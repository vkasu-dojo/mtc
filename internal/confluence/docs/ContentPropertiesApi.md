# \ContentPropertiesApi

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAttachmentProperty**](ContentPropertiesApi.md#CreateAttachmentProperty) | **Post** /attachments/{attachment-id}/properties | Create attachment property
[**CreateBlogpostProperty**](ContentPropertiesApi.md#CreateBlogpostProperty) | **Post** /blogposts/{blogpost-id}/properties | Create blog post property
[**CreateCommentProperty**](ContentPropertiesApi.md#CreateCommentProperty) | **Post** /comments/{comment-id}/properties | Create comment property
[**CreateCustomContentProperty**](ContentPropertiesApi.md#CreateCustomContentProperty) | **Post** /custom-content/{custom-content-id}/properties | Create custom content property
[**CreatePageProperty**](ContentPropertiesApi.md#CreatePageProperty) | **Post** /pages/{page-id}/properties | Create page property
[**DeleteAttachmentPropertyById**](ContentPropertiesApi.md#DeleteAttachmentPropertyById) | **Delete** /attachments/{attachment-id}/properties/{property-id} | Delete attachment property by id
[**DeleteBlogpostPropertyById**](ContentPropertiesApi.md#DeleteBlogpostPropertyById) | **Delete** /blogposts/{blogpost-id}/properties/{property-id} | Delete blog post property by id
[**DeleteCommentPropertyById**](ContentPropertiesApi.md#DeleteCommentPropertyById) | **Delete** /comments/{comment-id}/properties/{property-id} | Delete comment property by id
[**DeleteCustomContentPropertyById**](ContentPropertiesApi.md#DeleteCustomContentPropertyById) | **Delete** /custom-content/{custom-content-id}/properties/{property-id} | Delete custom content property by id
[**DeletePagePropertyById**](ContentPropertiesApi.md#DeletePagePropertyById) | **Delete** /pages/{page-id}/properties/{property-id} | Delete page property by id
[**GetAttachmentContentProperties**](ContentPropertiesApi.md#GetAttachmentContentProperties) | **Get** /attachments/{attachment-id}/properties | Get content properties for attachment
[**GetAttachmentContentPropertiesById**](ContentPropertiesApi.md#GetAttachmentContentPropertiesById) | **Get** /attachments/{attachment-id}/properties/{property-id} | Get content property for attachment by id
[**GetBlogpostContentProperties**](ContentPropertiesApi.md#GetBlogpostContentProperties) | **Get** /blogposts/{blogpost-id}/properties | Get content properties for blog post
[**GetBlogpostContentPropertiesById**](ContentPropertiesApi.md#GetBlogpostContentPropertiesById) | **Get** /blogposts/{blogpost-id}/properties/{property-id} | Get content property for blog post by id
[**GetCommentContentProperties**](ContentPropertiesApi.md#GetCommentContentProperties) | **Get** /comments/{comment-id}/properties | Get content properties for comment
[**GetCommentContentPropertiesById**](ContentPropertiesApi.md#GetCommentContentPropertiesById) | **Get** /comments/{comment-id}/properties/{property-id} | Get content property for comment by id
[**GetCustomContentContentProperties**](ContentPropertiesApi.md#GetCustomContentContentProperties) | **Get** /custom-content/{custom-content-id}/properties | Get content properties for custom content
[**GetCustomContentContentPropertiesById**](ContentPropertiesApi.md#GetCustomContentContentPropertiesById) | **Get** /custom-content/{custom-content-id}/properties/{property-id} | Get content property for custom content by id
[**GetPageContentProperties**](ContentPropertiesApi.md#GetPageContentProperties) | **Get** /pages/{page-id}/properties | Get content properties for page
[**GetPageContentPropertiesById**](ContentPropertiesApi.md#GetPageContentPropertiesById) | **Get** /pages/{page-id}/properties/{property-id} | Get content property for page by id
[**UpdateAttachmentPropertyById**](ContentPropertiesApi.md#UpdateAttachmentPropertyById) | **Put** /attachments/{attachment-id}/properties/{property-id} | Update attachment property by id
[**UpdateBlogpostPropertyById**](ContentPropertiesApi.md#UpdateBlogpostPropertyById) | **Put** /blogposts/{blogpost-id}/properties/{property-id} | Update blog post property by id
[**UpdateCommentPropertyById**](ContentPropertiesApi.md#UpdateCommentPropertyById) | **Put** /comments/{comment-id}/properties/{property-id} | Update comment property by id
[**UpdateCustomContentPropertyById**](ContentPropertiesApi.md#UpdateCustomContentPropertyById) | **Put** /custom-content/{custom-content-id}/properties/{property-id} | Update custom content property by id
[**UpdatePagePropertyById**](ContentPropertiesApi.md#UpdatePagePropertyById) | **Put** /pages/{page-id}/properties/{property-id} | Update page property by id



## CreateAttachmentProperty

> ContentProperty CreateAttachmentProperty(ctx, attachmentId).ContentPropertyCreateRequest(contentPropertyCreateRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Create attachment property



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
    attachmentId := "attachmentId_example" // string | The ID of the attachment to create a property for.
    contentPropertyCreateRequest := *openapiclient.NewContentPropertyCreateRequest() // ContentPropertyCreateRequest | The attachment property to be created
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPropertiesApi.CreateAttachmentProperty(context.Background(), attachmentId).ContentPropertyCreateRequest(contentPropertyCreateRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.CreateAttachmentProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAttachmentProperty`: ContentProperty
    fmt.Fprintf(os.Stdout, "Response from `ContentPropertiesApi.CreateAttachmentProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** | The ID of the attachment to create a property for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAttachmentPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentPropertyCreateRequest** | [**ContentPropertyCreateRequest**](ContentPropertyCreateRequest.md) | The attachment property to be created | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**ContentProperty**](ContentProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBlogpostProperty

> ContentProperty CreateBlogpostProperty(ctx, blogpostId).ContentPropertyCreateRequest(contentPropertyCreateRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Create blog post property



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
    blogpostId := int64(789) // int64 | The ID of the blog post to create a property for.
    contentPropertyCreateRequest := *openapiclient.NewContentPropertyCreateRequest() // ContentPropertyCreateRequest | The blog post property to be created
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPropertiesApi.CreateBlogpostProperty(context.Background(), blogpostId).ContentPropertyCreateRequest(contentPropertyCreateRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.CreateBlogpostProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBlogpostProperty`: ContentProperty
    fmt.Fprintf(os.Stdout, "Response from `ContentPropertiesApi.CreateBlogpostProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogpostId** | **int64** | The ID of the blog post to create a property for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlogpostPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentPropertyCreateRequest** | [**ContentPropertyCreateRequest**](ContentPropertyCreateRequest.md) | The blog post property to be created | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**ContentProperty**](ContentProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCommentProperty

> ContentProperty CreateCommentProperty(ctx, commentId).ContentPropertyCreateRequest(contentPropertyCreateRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Create comment property



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
    commentId := int64(789) // int64 | The ID of the comment to create a property for.
    contentPropertyCreateRequest := *openapiclient.NewContentPropertyCreateRequest() // ContentPropertyCreateRequest | The comment property to be created
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPropertiesApi.CreateCommentProperty(context.Background(), commentId).ContentPropertyCreateRequest(contentPropertyCreateRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.CreateCommentProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCommentProperty`: ContentProperty
    fmt.Fprintf(os.Stdout, "Response from `ContentPropertiesApi.CreateCommentProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int64** | The ID of the comment to create a property for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommentPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentPropertyCreateRequest** | [**ContentPropertyCreateRequest**](ContentPropertyCreateRequest.md) | The comment property to be created | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**ContentProperty**](ContentProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomContentProperty

> ContentProperty CreateCustomContentProperty(ctx, customContentId).ContentPropertyCreateRequest(contentPropertyCreateRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Create custom content property



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
    customContentId := int64(789) // int64 | The ID of the custom content to create a property for.
    contentPropertyCreateRequest := *openapiclient.NewContentPropertyCreateRequest() // ContentPropertyCreateRequest | The custom content property to be created
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPropertiesApi.CreateCustomContentProperty(context.Background(), customContentId).ContentPropertyCreateRequest(contentPropertyCreateRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.CreateCustomContentProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomContentProperty`: ContentProperty
    fmt.Fprintf(os.Stdout, "Response from `ContentPropertiesApi.CreateCustomContentProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customContentId** | **int64** | The ID of the custom content to create a property for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomContentPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentPropertyCreateRequest** | [**ContentPropertyCreateRequest**](ContentPropertyCreateRequest.md) | The custom content property to be created | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**ContentProperty**](ContentProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePageProperty

> ContentProperty CreatePageProperty(ctx, pageId).ContentPropertyCreateRequest(contentPropertyCreateRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Create page property



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
    pageId := int64(789) // int64 | The ID of the page to create a property for.
    contentPropertyCreateRequest := *openapiclient.NewContentPropertyCreateRequest() // ContentPropertyCreateRequest | The page property to be created
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPropertiesApi.CreatePageProperty(context.Background(), pageId).ContentPropertyCreateRequest(contentPropertyCreateRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.CreatePageProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePageProperty`: ContentProperty
    fmt.Fprintf(os.Stdout, "Response from `ContentPropertiesApi.CreatePageProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **int64** | The ID of the page to create a property for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePagePropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentPropertyCreateRequest** | [**ContentPropertyCreateRequest**](ContentPropertyCreateRequest.md) | The page property to be created | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**ContentProperty**](ContentProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAttachmentPropertyById

> DeleteAttachmentPropertyById(ctx, attachmentId, propertyId).Execute()

Delete attachment property by id



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
    attachmentId := "attachmentId_example" // string | The ID of the attachment the property belongs to.
    propertyId := int64(789) // int64 | The ID of the property to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContentPropertiesApi.DeleteAttachmentPropertyById(context.Background(), attachmentId, propertyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.DeleteAttachmentPropertyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** | The ID of the attachment the property belongs to. | 
**propertyId** | **int64** | The ID of the property to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttachmentPropertyByIdRequest struct via the builder pattern


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


## DeleteBlogpostPropertyById

> DeleteBlogpostPropertyById(ctx, blogpostId, propertyId).Execute()

Delete blog post property by id



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
    blogpostId := int64(789) // int64 | The ID of the blog post the property belongs to.
    propertyId := int64(789) // int64 | The ID of the property to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContentPropertiesApi.DeleteBlogpostPropertyById(context.Background(), blogpostId, propertyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.DeleteBlogpostPropertyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogpostId** | **int64** | The ID of the blog post the property belongs to. | 
**propertyId** | **int64** | The ID of the property to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlogpostPropertyByIdRequest struct via the builder pattern


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


## DeleteCommentPropertyById

> DeleteCommentPropertyById(ctx, commentId, propertyId).Execute()

Delete comment property by id



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
    commentId := int64(789) // int64 | The ID of the comment the property belongs to.
    propertyId := int64(789) // int64 | The ID of the property to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContentPropertiesApi.DeleteCommentPropertyById(context.Background(), commentId, propertyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.DeleteCommentPropertyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int64** | The ID of the comment the property belongs to. | 
**propertyId** | **int64** | The ID of the property to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommentPropertyByIdRequest struct via the builder pattern


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


## DeleteCustomContentPropertyById

> DeleteCustomContentPropertyById(ctx, customContentId, propertyId).Execute()

Delete custom content property by id



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
    customContentId := int64(789) // int64 | The ID of the custom content the property belongs to.
    propertyId := int64(789) // int64 | The ID of the property to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContentPropertiesApi.DeleteCustomContentPropertyById(context.Background(), customContentId, propertyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.DeleteCustomContentPropertyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customContentId** | **int64** | The ID of the custom content the property belongs to. | 
**propertyId** | **int64** | The ID of the property to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomContentPropertyByIdRequest struct via the builder pattern


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


## DeletePagePropertyById

> DeletePagePropertyById(ctx, pageId, propertyId).Execute()

Delete page property by id



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
    pageId := int64(789) // int64 | The ID of the page the property belongs to.
    propertyId := int64(789) // int64 | The ID of the property to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContentPropertiesApi.DeletePagePropertyById(context.Background(), pageId, propertyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.DeletePagePropertyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **int64** | The ID of the page the property belongs to. | 
**propertyId** | **int64** | The ID of the property to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagePropertyByIdRequest struct via the builder pattern


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


## GetAttachmentContentProperties

> MultiEntityResultContentProperty GetAttachmentContentProperties(ctx, attachmentId).Key(key).Sort(sort).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get content properties for attachment



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
    attachmentId := "attachmentId_example" // string | The ID of the attachment for which content properties should be returned.
    key := "key_example" // string | Filters the response to return a specific content property with matching key (case sensitive). (optional)
    sort := openapiclient.ContentPropertySortOrder("key") // ContentPropertySortOrder | Used to sort the result by a particular field. (optional)
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    limit := int32(56) // int32 | Maximum number of attachments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPropertiesApi.GetAttachmentContentProperties(context.Background(), attachmentId).Key(key).Sort(sort).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.GetAttachmentContentProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttachmentContentProperties`: MultiEntityResultContentProperty
    fmt.Fprintf(os.Stdout, "Response from `ContentPropertiesApi.GetAttachmentContentProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** | The ID of the attachment for which content properties should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentContentPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** | Filters the response to return a specific content property with matching key (case sensitive). | 
 **sort** | [**ContentPropertySortOrder**](ContentPropertySortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of attachments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultContentProperty**](MultiEntityResultContentProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachmentContentPropertiesById

> ContentProperty GetAttachmentContentPropertiesById(ctx, attachmentId, propertyId).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get content property for attachment by id



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
    attachmentId := "attachmentId_example" // string | The ID of the attachment for which content properties should be returned.
    propertyId := int64(789) // int64 | The ID of the content property to be returned
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPropertiesApi.GetAttachmentContentPropertiesById(context.Background(), attachmentId, propertyId).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.GetAttachmentContentPropertiesById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttachmentContentPropertiesById`: ContentProperty
    fmt.Fprintf(os.Stdout, "Response from `ContentPropertiesApi.GetAttachmentContentPropertiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** | The ID of the attachment for which content properties should be returned. | 
**propertyId** | **int64** | The ID of the content property to be returned | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentContentPropertiesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**ContentProperty**](ContentProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlogpostContentProperties

> MultiEntityResultContentProperty GetBlogpostContentProperties(ctx, blogpostId).Key(key).Sort(sort).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get content properties for blog post



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
    blogpostId := int64(789) // int64 | The ID of the blog post for which content properties should be returned.
    key := "key_example" // string | Filters the response to return a specific content property with matching key (case sensitive). (optional)
    sort := openapiclient.ContentPropertySortOrder("key") // ContentPropertySortOrder | Used to sort the result by a particular field. (optional)
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    limit := int32(56) // int32 | Maximum number of attachments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPropertiesApi.GetBlogpostContentProperties(context.Background(), blogpostId).Key(key).Sort(sort).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.GetBlogpostContentProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlogpostContentProperties`: MultiEntityResultContentProperty
    fmt.Fprintf(os.Stdout, "Response from `ContentPropertiesApi.GetBlogpostContentProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogpostId** | **int64** | The ID of the blog post for which content properties should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogpostContentPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** | Filters the response to return a specific content property with matching key (case sensitive). | 
 **sort** | [**ContentPropertySortOrder**](ContentPropertySortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of attachments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultContentProperty**](MultiEntityResultContentProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlogpostContentPropertiesById

> ContentProperty GetBlogpostContentPropertiesById(ctx, blogpostId, propertyId).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get content property for blog post by id



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
    blogpostId := int64(789) // int64 | The ID of the blog post for which content properties should be returned.
    propertyId := int64(789) // int64 | The ID of the property being requested
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPropertiesApi.GetBlogpostContentPropertiesById(context.Background(), blogpostId, propertyId).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.GetBlogpostContentPropertiesById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlogpostContentPropertiesById`: ContentProperty
    fmt.Fprintf(os.Stdout, "Response from `ContentPropertiesApi.GetBlogpostContentPropertiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogpostId** | **int64** | The ID of the blog post for which content properties should be returned. | 
**propertyId** | **int64** | The ID of the property being requested | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogpostContentPropertiesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**ContentProperty**](ContentProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommentContentProperties

> MultiEntityResultContentProperty GetCommentContentProperties(ctx, commentId).Key(key).Sort(sort).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get content properties for comment



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
    commentId := int64(789) // int64 | The ID of the comment for which content properties should be returned.
    key := "key_example" // string | Filters the response to return a specific content property with matching key (case sensitive). (optional)
    sort := openapiclient.ContentPropertySortOrder("key") // ContentPropertySortOrder | Used to sort the result by a particular field. (optional)
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    limit := int32(56) // int32 | Maximum number of attachments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPropertiesApi.GetCommentContentProperties(context.Background(), commentId).Key(key).Sort(sort).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.GetCommentContentProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommentContentProperties`: MultiEntityResultContentProperty
    fmt.Fprintf(os.Stdout, "Response from `ContentPropertiesApi.GetCommentContentProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int64** | The ID of the comment for which content properties should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentContentPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** | Filters the response to return a specific content property with matching key (case sensitive). | 
 **sort** | [**ContentPropertySortOrder**](ContentPropertySortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of attachments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultContentProperty**](MultiEntityResultContentProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommentContentPropertiesById

> ContentProperty GetCommentContentPropertiesById(ctx, commentId, propertyId).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get content property for comment by id



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
    commentId := int64(789) // int64 | The ID of the comment for which content properties should be returned.
    propertyId := int64(789) // int64 | The ID of the content property being requested.
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPropertiesApi.GetCommentContentPropertiesById(context.Background(), commentId, propertyId).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.GetCommentContentPropertiesById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommentContentPropertiesById`: ContentProperty
    fmt.Fprintf(os.Stdout, "Response from `ContentPropertiesApi.GetCommentContentPropertiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int64** | The ID of the comment for which content properties should be returned. | 
**propertyId** | **int64** | The ID of the content property being requested. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentContentPropertiesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**ContentProperty**](ContentProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomContentContentProperties

> MultiEntityResultContentProperty GetCustomContentContentProperties(ctx, customContentId).Key(key).Sort(sort).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get content properties for custom content



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
    customContentId := int64(789) // int64 | The ID of the custom content for which content properties should be returned.
    key := "key_example" // string | Filters the response to return a specific content property with matching key (case sensitive). (optional)
    sort := openapiclient.ContentPropertySortOrder("key") // ContentPropertySortOrder | Used to sort the result by a particular field. (optional)
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    limit := int32(56) // int32 | Maximum number of attachments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPropertiesApi.GetCustomContentContentProperties(context.Background(), customContentId).Key(key).Sort(sort).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.GetCustomContentContentProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomContentContentProperties`: MultiEntityResultContentProperty
    fmt.Fprintf(os.Stdout, "Response from `ContentPropertiesApi.GetCustomContentContentProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customContentId** | **int64** | The ID of the custom content for which content properties should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomContentContentPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** | Filters the response to return a specific content property with matching key (case sensitive). | 
 **sort** | [**ContentPropertySortOrder**](ContentPropertySortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of attachments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultContentProperty**](MultiEntityResultContentProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomContentContentPropertiesById

> ContentProperty GetCustomContentContentPropertiesById(ctx, customContentId, propertyId).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get content property for custom content by id



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
    customContentId := int64(789) // int64 | The ID of the custom content for which content properties should be returned.
    propertyId := int64(789) // int64 | The ID of the content property being requested.
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPropertiesApi.GetCustomContentContentPropertiesById(context.Background(), customContentId, propertyId).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.GetCustomContentContentPropertiesById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomContentContentPropertiesById`: ContentProperty
    fmt.Fprintf(os.Stdout, "Response from `ContentPropertiesApi.GetCustomContentContentPropertiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customContentId** | **int64** | The ID of the custom content for which content properties should be returned. | 
**propertyId** | **int64** | The ID of the content property being requested. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomContentContentPropertiesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**ContentProperty**](ContentProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageContentProperties

> MultiEntityResultContentProperty GetPageContentProperties(ctx, pageId).Key(key).Sort(sort).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get content properties for page



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
    pageId := int64(789) // int64 | The ID of the page for which content properties should be returned.
    key := "key_example" // string | Filters the response to return a specific content property with matching key (case sensitive). (optional)
    sort := openapiclient.ContentPropertySortOrder("key") // ContentPropertySortOrder | Used to sort the result by a particular field. (optional)
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    limit := int32(56) // int32 | Maximum number of attachments per result to return. If more results exist, use the `Link` header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPropertiesApi.GetPageContentProperties(context.Background(), pageId).Key(key).Sort(sort).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.GetPageContentProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPageContentProperties`: MultiEntityResultContentProperty
    fmt.Fprintf(os.Stdout, "Response from `ContentPropertiesApi.GetPageContentProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **int64** | The ID of the page for which content properties should be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageContentPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** | Filters the response to return a specific content property with matching key (case sensitive). | 
 **sort** | [**ContentPropertySortOrder**](ContentPropertySortOrder.md) | Used to sort the result by a particular field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of attachments per result to return. If more results exist, use the &#x60;Link&#x60; header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultContentProperty**](MultiEntityResultContentProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageContentPropertiesById

> ContentProperty GetPageContentPropertiesById(ctx, pageId, propertyId).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get content property for page by id



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
    pageId := int64(789) // int64 | The ID of the page for which content properties should be returned.
    propertyId := int64(789) // int64 | The ID of the content property being requested.
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPropertiesApi.GetPageContentPropertiesById(context.Background(), pageId, propertyId).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.GetPageContentPropertiesById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPageContentPropertiesById`: ContentProperty
    fmt.Fprintf(os.Stdout, "Response from `ContentPropertiesApi.GetPageContentPropertiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **int64** | The ID of the page for which content properties should be returned. | 
**propertyId** | **int64** | The ID of the content property being requested. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageContentPropertiesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**ContentProperty**](ContentProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAttachmentPropertyById

> ContentProperty UpdateAttachmentPropertyById(ctx, attachmentId, propertyId).ContentPropertyUpdateRequest(contentPropertyUpdateRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Update attachment property by id



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
    attachmentId := "attachmentId_example" // string | The ID of the attachment the property belongs to.
    propertyId := int64(789) // int64 | The ID of the property to be updated.
    contentPropertyUpdateRequest := *openapiclient.NewContentPropertyUpdateRequest() // ContentPropertyUpdateRequest | The attachment property to be updated.
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPropertiesApi.UpdateAttachmentPropertyById(context.Background(), attachmentId, propertyId).ContentPropertyUpdateRequest(contentPropertyUpdateRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.UpdateAttachmentPropertyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAttachmentPropertyById`: ContentProperty
    fmt.Fprintf(os.Stdout, "Response from `ContentPropertiesApi.UpdateAttachmentPropertyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** | The ID of the attachment the property belongs to. | 
**propertyId** | **int64** | The ID of the property to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAttachmentPropertyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentPropertyUpdateRequest** | [**ContentPropertyUpdateRequest**](ContentPropertyUpdateRequest.md) | The attachment property to be updated. | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**ContentProperty**](ContentProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBlogpostPropertyById

> ContentProperty UpdateBlogpostPropertyById(ctx, blogpostId, propertyId).ContentPropertyUpdateRequest(contentPropertyUpdateRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Update blog post property by id



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
    blogpostId := int64(789) // int64 | The ID of the blog post the property belongs to.
    propertyId := int64(789) // int64 | The ID of the property to be updated.
    contentPropertyUpdateRequest := *openapiclient.NewContentPropertyUpdateRequest() // ContentPropertyUpdateRequest | The blog post property to be updated.
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPropertiesApi.UpdateBlogpostPropertyById(context.Background(), blogpostId, propertyId).ContentPropertyUpdateRequest(contentPropertyUpdateRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.UpdateBlogpostPropertyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBlogpostPropertyById`: ContentProperty
    fmt.Fprintf(os.Stdout, "Response from `ContentPropertiesApi.UpdateBlogpostPropertyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogpostId** | **int64** | The ID of the blog post the property belongs to. | 
**propertyId** | **int64** | The ID of the property to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlogpostPropertyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentPropertyUpdateRequest** | [**ContentPropertyUpdateRequest**](ContentPropertyUpdateRequest.md) | The blog post property to be updated. | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**ContentProperty**](ContentProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCommentPropertyById

> ContentProperty UpdateCommentPropertyById(ctx, commentId, propertyId).ContentPropertyUpdateRequest(contentPropertyUpdateRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Update comment property by id



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
    commentId := int64(789) // int64 | The ID of the comment the property belongs to.
    propertyId := int64(789) // int64 | The ID of the property to be updated.
    contentPropertyUpdateRequest := *openapiclient.NewContentPropertyUpdateRequest() // ContentPropertyUpdateRequest | The comment property to be updated.
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPropertiesApi.UpdateCommentPropertyById(context.Background(), commentId, propertyId).ContentPropertyUpdateRequest(contentPropertyUpdateRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.UpdateCommentPropertyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCommentPropertyById`: ContentProperty
    fmt.Fprintf(os.Stdout, "Response from `ContentPropertiesApi.UpdateCommentPropertyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int64** | The ID of the comment the property belongs to. | 
**propertyId** | **int64** | The ID of the property to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCommentPropertyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentPropertyUpdateRequest** | [**ContentPropertyUpdateRequest**](ContentPropertyUpdateRequest.md) | The comment property to be updated. | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**ContentProperty**](ContentProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomContentPropertyById

> ContentProperty UpdateCustomContentPropertyById(ctx, customContentId, propertyId).ContentPropertyUpdateRequest(contentPropertyUpdateRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Update custom content property by id



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
    customContentId := int64(789) // int64 | The ID of the custom content the property belongs to.
    propertyId := int64(789) // int64 | The ID of the property to be updated.
    contentPropertyUpdateRequest := *openapiclient.NewContentPropertyUpdateRequest() // ContentPropertyUpdateRequest | The custom content property to be updated.
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPropertiesApi.UpdateCustomContentPropertyById(context.Background(), customContentId, propertyId).ContentPropertyUpdateRequest(contentPropertyUpdateRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.UpdateCustomContentPropertyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomContentPropertyById`: ContentProperty
    fmt.Fprintf(os.Stdout, "Response from `ContentPropertiesApi.UpdateCustomContentPropertyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customContentId** | **int64** | The ID of the custom content the property belongs to. | 
**propertyId** | **int64** | The ID of the property to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomContentPropertyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentPropertyUpdateRequest** | [**ContentPropertyUpdateRequest**](ContentPropertyUpdateRequest.md) | The custom content property to be updated. | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**ContentProperty**](ContentProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePagePropertyById

> ContentProperty UpdatePagePropertyById(ctx, pageId, propertyId).ContentPropertyUpdateRequest(contentPropertyUpdateRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Update page property by id



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
    pageId := int64(789) // int64 | The ID of the page the property belongs to.
    propertyId := int64(789) // int64 | The ID of the property to be updated.
    contentPropertyUpdateRequest := *openapiclient.NewContentPropertyUpdateRequest() // ContentPropertyUpdateRequest | The page property to be updated.
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPropertiesApi.UpdatePagePropertyById(context.Background(), pageId, propertyId).ContentPropertyUpdateRequest(contentPropertyUpdateRequest).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPropertiesApi.UpdatePagePropertyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePagePropertyById`: ContentProperty
    fmt.Fprintf(os.Stdout, "Response from `ContentPropertiesApi.UpdatePagePropertyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **int64** | The ID of the page the property belongs to. | 
**propertyId** | **int64** | The ID of the property to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePagePropertyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentPropertyUpdateRequest** | [**ContentPropertyUpdateRequest**](ContentPropertyUpdateRequest.md) | The page property to be updated. | 
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**ContentProperty**](ContentProperty.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

