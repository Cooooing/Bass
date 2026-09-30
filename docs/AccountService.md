# \AccountService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Avatar**](AccountService.md#Avatar) | **Get** /v1/user/account/avatar | 
[**CompleteProfileImageUpload**](AccountService.md#CompleteProfileImageUpload) | **Post** /v1/user/account/complete-profile-image-upload | 
[**GetCurrent**](AccountService.md#GetCurrent) | **Post** /v1/user/account/get-current | 
[**GetProfile**](AccountService.md#GetProfile) | **Post** /v1/user/account/get-profile | 
[**ListFollowers**](AccountService.md#ListFollowers) | **Post** /v1/user/account/list-followers | 
[**ListFollowing**](AccountService.md#ListFollowing) | **Post** /v1/user/account/list-following | 
[**PrepareProfileImageUpload**](AccountService.md#PrepareProfileImageUpload) | **Post** /v1/user/account/prepare-profile-image-upload | 
[**UpdateEmail**](AccountService.md#UpdateEmail) | **Post** /v1/user/account/update-email | 
[**UpdatePassword**](AccountService.md#UpdatePassword) | **Post** /v1/user/account/update-password | 
[**UpdatePhone**](AccountService.md#UpdatePhone) | **Post** /v1/user/account/update-phone | 
[**UpdateProfile**](AccountService.md#UpdateProfile) | **Post** /v1/user/account/update-profile | 



## Avatar

> ImageResp Avatar(ctx).Name(name).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "name_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountService.Avatar(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.Avatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Avatar`: ImageResp
	fmt.Fprintf(os.Stdout, "Response from `AccountService.Avatar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 

### Return type

[**ImageResp**](ImageResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteProfileImageUpload

> CompleteProfileImageUploadAccountResp CompleteProfileImageUpload(ctx).CompleteProfileImageUploadAccountReq(completeProfileImageUploadAccountReq).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	completeProfileImageUploadAccountReq := *openapiclient.NewCompleteProfileImageUploadAccountReq("Purpose_example", "Hash_example") // CompleteProfileImageUploadAccountReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountService.CompleteProfileImageUpload(context.Background()).CompleteProfileImageUploadAccountReq(completeProfileImageUploadAccountReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.CompleteProfileImageUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteProfileImageUpload`: CompleteProfileImageUploadAccountResp
	fmt.Fprintf(os.Stdout, "Response from `AccountService.CompleteProfileImageUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteProfileImageUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **completeProfileImageUploadAccountReq** | [**CompleteProfileImageUploadAccountReq**](CompleteProfileImageUploadAccountReq.md) |  | 

### Return type

[**CompleteProfileImageUploadAccountResp**](CompleteProfileImageUploadAccountResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrent

> GetCurrentAccountResp GetCurrent(ctx).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountService.GetCurrent(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.GetCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrent`: GetCurrentAccountResp
	fmt.Fprintf(os.Stdout, "Response from `AccountService.GetCurrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**GetCurrentAccountResp**](GetCurrentAccountResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfile

> GetProfileResp GetProfile(ctx).GetProfileReq(getProfileReq).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	getProfileReq := *openapiclient.NewGetProfileReq("Name_example") // GetProfileReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountService.GetProfile(context.Background()).GetProfileReq(getProfileReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.GetProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfile`: GetProfileResp
	fmt.Fprintf(os.Stdout, "Response from `AccountService.GetProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getProfileReq** | [**GetProfileReq**](GetProfileReq.md) |  | 

### Return type

[**GetProfileResp**](GetProfileResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFollowers

> ListFollowersResp ListFollowers(ctx).ListFollowersReq(listFollowersReq).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	listFollowersReq := *openapiclient.NewListFollowersReq("Name_example") // ListFollowersReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountService.ListFollowers(context.Background()).ListFollowersReq(listFollowersReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.ListFollowers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFollowers`: ListFollowersResp
	fmt.Fprintf(os.Stdout, "Response from `AccountService.ListFollowers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFollowersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listFollowersReq** | [**ListFollowersReq**](ListFollowersReq.md) |  | 

### Return type

[**ListFollowersResp**](ListFollowersResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFollowing

> ListFollowingResp ListFollowing(ctx).ListFollowingReq(listFollowingReq).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	listFollowingReq := *openapiclient.NewListFollowingReq("Name_example") // ListFollowingReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountService.ListFollowing(context.Background()).ListFollowingReq(listFollowingReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.ListFollowing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFollowing`: ListFollowingResp
	fmt.Fprintf(os.Stdout, "Response from `AccountService.ListFollowing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFollowingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listFollowingReq** | [**ListFollowingReq**](ListFollowingReq.md) |  | 

### Return type

[**ListFollowingResp**](ListFollowingResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrepareProfileImageUpload

> PrepareProfileImageUploadAccountResp PrepareProfileImageUpload(ctx).PrepareProfileImageUploadAccountReq(prepareProfileImageUploadAccountReq).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prepareProfileImageUploadAccountReq := *openapiclient.NewPrepareProfileImageUploadAccountReq("Purpose_example", "Hash_example", "MimeType_example", "Size_example") // PrepareProfileImageUploadAccountReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountService.PrepareProfileImageUpload(context.Background()).PrepareProfileImageUploadAccountReq(prepareProfileImageUploadAccountReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.PrepareProfileImageUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrepareProfileImageUpload`: PrepareProfileImageUploadAccountResp
	fmt.Fprintf(os.Stdout, "Response from `AccountService.PrepareProfileImageUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrepareProfileImageUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prepareProfileImageUploadAccountReq** | [**PrepareProfileImageUploadAccountReq**](PrepareProfileImageUploadAccountReq.md) |  | 

### Return type

[**PrepareProfileImageUploadAccountResp**](PrepareProfileImageUploadAccountResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmail

> map[string]interface{} UpdateEmail(ctx).UpdateEmailAccountReq(updateEmailAccountReq).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateEmailAccountReq := *openapiclient.NewUpdateEmailAccountReq("Email_example", "Code_example") // UpdateEmailAccountReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountService.UpdateEmail(context.Background()).UpdateEmailAccountReq(updateEmailAccountReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.UpdateEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEmail`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountService.UpdateEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateEmailAccountReq** | [**UpdateEmailAccountReq**](UpdateEmailAccountReq.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePassword

> map[string]interface{} UpdatePassword(ctx).UpdatePasswordAccountReq(updatePasswordAccountReq).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updatePasswordAccountReq := *openapiclient.NewUpdatePasswordAccountReq("OldPassword_example", "NewPassword_example") // UpdatePasswordAccountReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountService.UpdatePassword(context.Background()).UpdatePasswordAccountReq(updatePasswordAccountReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.UpdatePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePassword`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountService.UpdatePassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatePasswordAccountReq** | [**UpdatePasswordAccountReq**](UpdatePasswordAccountReq.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePhone

> map[string]interface{} UpdatePhone(ctx).UpdatePhoneAccountReq(updatePhoneAccountReq).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updatePhoneAccountReq := *openapiclient.NewUpdatePhoneAccountReq("Phone_example", "Code_example") // UpdatePhoneAccountReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountService.UpdatePhone(context.Background()).UpdatePhoneAccountReq(updatePhoneAccountReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.UpdatePhone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePhone`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountService.UpdatePhone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePhoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatePhoneAccountReq** | [**UpdatePhoneAccountReq**](UpdatePhoneAccountReq.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProfile

> UpdateProfileAccountResp UpdateProfile(ctx).UpdateProfileAccountReq(updateProfileAccountReq).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateProfileAccountReq := *openapiclient.NewUpdateProfileAccountReq() // UpdateProfileAccountReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountService.UpdateProfile(context.Background()).UpdateProfileAccountReq(updateProfileAccountReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.UpdateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProfile`: UpdateProfileAccountResp
	fmt.Fprintf(os.Stdout, "Response from `AccountService.UpdateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateProfileAccountReq** | [**UpdateProfileAccountReq**](UpdateProfileAccountReq.md) |  | 

### Return type

[**UpdateProfileAccountResp**](UpdateProfileAccountResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

