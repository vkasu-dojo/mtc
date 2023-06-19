# \UserApi

All URIs are relative to *https://no-default/wiki/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckAccessByEmail**](UserApi.md#CheckAccessByEmail) | **Post** /user/access/check-access-by-email | Check site access for a list of emails
[**InviteByEmail**](UserApi.md#InviteByEmail) | **Post** /user/access/invite-by-email | Invite a list of emails to the site



## CheckAccessByEmail

> CheckAccessByEmail200Response CheckAccessByEmail(ctx).CheckAccessByEmailRequest(checkAccessByEmailRequest).Execute()

Check site access for a list of emails



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
    checkAccessByEmailRequest := *openapiclient.NewCheckAccessByEmailRequest([]string{"Emails_example"}) // CheckAccessByEmailRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.CheckAccessByEmail(context.Background()).CheckAccessByEmailRequest(checkAccessByEmailRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.CheckAccessByEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckAccessByEmail`: CheckAccessByEmail200Response
    fmt.Fprintf(os.Stdout, "Response from `UserApi.CheckAccessByEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckAccessByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkAccessByEmailRequest** | [**CheckAccessByEmailRequest**](CheckAccessByEmailRequest.md) |  | 

### Return type

[**CheckAccessByEmail200Response**](CheckAccessByEmail200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteByEmail

> InviteByEmail200Response InviteByEmail(ctx).CheckAccessByEmailRequest(checkAccessByEmailRequest).Execute()

Invite a list of emails to the site



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
    checkAccessByEmailRequest := *openapiclient.NewCheckAccessByEmailRequest([]string{"Emails_example"}) // CheckAccessByEmailRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.InviteByEmail(context.Background()).CheckAccessByEmailRequest(checkAccessByEmailRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.InviteByEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InviteByEmail`: InviteByEmail200Response
    fmt.Fprintf(os.Stdout, "Response from `UserApi.InviteByEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInviteByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkAccessByEmailRequest** | [**CheckAccessByEmailRequest**](CheckAccessByEmailRequest.md) |  | 

### Return type

[**InviteByEmail200Response**](InviteByEmail200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth), [oAuthDefinitions](../README.md#oAuthDefinitions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

