# \SpaceApi

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSpaceById**](SpaceApi.md#GetSpaceById) | **Get** /spaces/{id} | Get space by id
[**GetSpaces**](SpaceApi.md#GetSpaces) | **Get** /spaces | Get spaces



## GetSpaceById

> Space GetSpaceById(ctx, id).DescriptionFormat(descriptionFormat).Execute()

Get space by id



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
    id := int64(789) // int64 | The ID of the space to be returned.
    descriptionFormat := openapiclient.SpaceDescriptionBodyRepresentation("plain") // SpaceDescriptionBodyRepresentation | The content format type to be returned in the `description` field of the response. If available, the representation will be available under a response field of the same name under the `description` field. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpaceApi.GetSpaceById(context.Background(), id).DescriptionFormat(descriptionFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpaceApi.GetSpaceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpaceById`: Space
    fmt.Fprintf(os.Stdout, "Response from `SpaceApi.GetSpaceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the space to be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **descriptionFormat** | [**SpaceDescriptionBodyRepresentation**](SpaceDescriptionBodyRepresentation.md) | The content format type to be returned in the &#x60;description&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;description&#x60; field. | 

### Return type

[**Space**](Space.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpaces

> MultiEntityResultSpace GetSpaces(ctx).Ids(ids).Keys(keys).Type_(type_).Status(status).Labels(labels).Sort(sort).DescriptionFormat(descriptionFormat).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()

Get spaces



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
    ids := []int64{int64(123)} // []int64 | Filter the results to spaces based on their IDs. Multiple IDs can be specified as a comma-separated list. (optional)
    keys := []string{"Inner_example"} // []string | Filter the results to spaces based on their keys. Multiple keys can be specified as a comma-separated list. (optional)
    type_ := "type__example" // string | Filter the results to spaces based on their type. (optional)
    status := "status_example" // string | Filter the results to spaces based on their status. (optional)
    labels := []string{"Inner_example"} // []string | Filter the results to spaces based on their labels. Multiple labels can be specified as a comma-separated list. (optional)
    sort := openapiclient.SpaceSortOrder("id") // SpaceSortOrder | Used to sort the result by a particular field. (optional)
    descriptionFormat := openapiclient.SpaceDescriptionBodyRepresentation("plain") // SpaceDescriptionBodyRepresentation | The content format type to be returned in the `description` field of the response. If available, the representation will be available under a response field of the same name under the `description` field. (optional)
    cursor := "cursor_example" // string | Used for pagination, this opaque cursor will be returned in the `next` URL in the `Link` response header. Use the relative URL in the `Link` header to retrieve the `next` set of results. (optional)
    limit := int32(56) // int32 | Maximum number of spaces per result to return. If more results exist, use the `Link` response header to retrieve a relative URL that will return the next set of results. (optional) (default to 25)
    serializeIdsAsStrings := true // bool | Due to JavaScript's max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpaceApi.GetSpaces(context.Background()).Ids(ids).Keys(keys).Type_(type_).Status(status).Labels(labels).Sort(sort).DescriptionFormat(descriptionFormat).Cursor(cursor).Limit(limit).SerializeIdsAsStrings(serializeIdsAsStrings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpaceApi.GetSpaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpaces`: MultiEntityResultSpace
    fmt.Fprintf(os.Stdout, "Response from `SpaceApi.GetSpaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSpacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | Filter the results to spaces based on their IDs. Multiple IDs can be specified as a comma-separated list. | 
 **keys** | **[]string** | Filter the results to spaces based on their keys. Multiple keys can be specified as a comma-separated list. | 
 **type_** | **string** | Filter the results to spaces based on their type. | 
 **status** | **string** | Filter the results to spaces based on their status. | 
 **labels** | **[]string** | Filter the results to spaces based on their labels. Multiple labels can be specified as a comma-separated list. | 
 **sort** | [**SpaceSortOrder**](SpaceSortOrder.md) | Used to sort the result by a particular field. | 
 **descriptionFormat** | [**SpaceDescriptionBodyRepresentation**](SpaceDescriptionBodyRepresentation.md) | The content format type to be returned in the &#x60;description&#x60; field of the response. If available, the representation will be available under a response field of the same name under the &#x60;description&#x60; field. | 
 **cursor** | **string** | Used for pagination, this opaque cursor will be returned in the &#x60;next&#x60; URL in the &#x60;Link&#x60; response header. Use the relative URL in the &#x60;Link&#x60; header to retrieve the &#x60;next&#x60; set of results. | 
 **limit** | **int32** | Maximum number of spaces per result to return. If more results exist, use the &#x60;Link&#x60; response header to retrieve a relative URL that will return the next set of results. | [default to 25]
 **serializeIdsAsStrings** | **bool** | Due to JavaScript&#39;s max integer representation of 2^53-1, the type of any IDs returned in the response body for this endpoint will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, this query param can be passed to this endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail. | [default to false]

### Return type

[**MultiEntityResultSpace**](MultiEntityResultSpace.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

