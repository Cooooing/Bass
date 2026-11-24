# \AccountService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AvatarAccount**](AccountService.md#AvatarAccount) | **Get** /v1/user/account/avatar | 
[**CompleteProfileImageUploadAccount**](AccountService.md#CompleteProfileImageUploadAccount) | **Post** /v1/user/account/complete-profile-image-upload | 
[**GetAccountEconomy**](AccountService.md#GetAccountEconomy) | **Post** /v1/user/account/economy/get | 
[**GetCurrentAccount**](AccountService.md#GetCurrentAccount) | **Post** /v1/user/account/get-current | 
[**GetProfile**](AccountService.md#GetProfile) | **Post** /v1/user/account/get-profile | 
[**ListFollowers**](AccountService.md#ListFollowers) | **Post** /v1/user/account/list-followers | 
[**ListFollowing**](AccountService.md#ListFollowing) | **Post** /v1/user/account/list-following | 
[**ListPointRecords**](AccountService.md#ListPointRecords) | **Post** /v1/user/account/list-point-records | 
[**PrepareProfileImageUploadAccount**](AccountService.md#PrepareProfileImageUploadAccount) | **Post** /v1/user/account/prepare-profile-image-upload | 
[**SearchMentionAccounts**](AccountService.md#SearchMentionAccounts) | **Post** /v1/user/account/search-mention | 
[**TransferAccountPoints**](AccountService.md#TransferAccountPoints) | **Post** /v1/user/account/economy/transfer | 
[**UpdateEmailAccount**](AccountService.md#UpdateEmailAccount) | **Post** /v1/user/account/update-email | 
[**UpdatePasswordAccount**](AccountService.md#UpdatePasswordAccount) | **Post** /v1/user/account/update-password | 
[**UpdatePhoneAccount**](AccountService.md#UpdatePhoneAccount) | **Post** /v1/user/account/update-phone | 
[**UpdateProfileAccount**](AccountService.md#UpdateProfileAccount) | **Post** /v1/user/account/update-profile | 



## AvatarAccount

> AvatarAccountResp AvatarAccount(ctx).Name(name).Execute()





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
	name := "name_example" // string | 账号名，必填 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountService.AvatarAccount(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.AvatarAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvatarAccount`: AvatarAccountResp
	fmt.Fprintf(os.Stdout, "Response from `AccountService.AvatarAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAvatarAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | 账号名，必填 | 

### Return type

[**AvatarAccountResp**](AvatarAccountResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteProfileImageUploadAccount

> CompleteProfileImageUploadAccountResp CompleteProfileImageUploadAccount(ctx).CompleteProfileImageUploadAccountReq(completeProfileImageUploadAccountReq).Execute()





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
	resp, r, err := apiClient.AccountService.CompleteProfileImageUploadAccount(context.Background()).CompleteProfileImageUploadAccountReq(completeProfileImageUploadAccountReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.CompleteProfileImageUploadAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteProfileImageUploadAccount`: CompleteProfileImageUploadAccountResp
	fmt.Fprintf(os.Stdout, "Response from `AccountService.CompleteProfileImageUploadAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteProfileImageUploadAccountRequest struct via the builder pattern


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


## GetAccountEconomy

> GetAccountEconomyResp GetAccountEconomy(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.AccountService.GetAccountEconomy(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.GetAccountEconomy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountEconomy`: GetAccountEconomyResp
	fmt.Fprintf(os.Stdout, "Response from `AccountService.GetAccountEconomy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountEconomyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**GetAccountEconomyResp**](GetAccountEconomyResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentAccount

> GetCurrentAccountResp GetCurrentAccount(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.AccountService.GetCurrentAccount(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.GetCurrentAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentAccount`: GetCurrentAccountResp
	fmt.Fprintf(os.Stdout, "Response from `AccountService.GetCurrentAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentAccountRequest struct via the builder pattern


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


## ListPointRecords

> ListPointRecordsResp ListPointRecords(ctx).ListPointRecordsReq(listPointRecordsReq).Execute()





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
	listPointRecordsReq := *openapiclient.NewListPointRecordsReq("Name_example") // ListPointRecordsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountService.ListPointRecords(context.Background()).ListPointRecordsReq(listPointRecordsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.ListPointRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPointRecords`: ListPointRecordsResp
	fmt.Fprintf(os.Stdout, "Response from `AccountService.ListPointRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPointRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listPointRecordsReq** | [**ListPointRecordsReq**](ListPointRecordsReq.md) |  | 

### Return type

[**ListPointRecordsResp**](ListPointRecordsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrepareProfileImageUploadAccount

> PrepareProfileImageUploadAccountResp PrepareProfileImageUploadAccount(ctx).PrepareProfileImageUploadAccountReq(prepareProfileImageUploadAccountReq).Execute()





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
	resp, r, err := apiClient.AccountService.PrepareProfileImageUploadAccount(context.Background()).PrepareProfileImageUploadAccountReq(prepareProfileImageUploadAccountReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.PrepareProfileImageUploadAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrepareProfileImageUploadAccount`: PrepareProfileImageUploadAccountResp
	fmt.Fprintf(os.Stdout, "Response from `AccountService.PrepareProfileImageUploadAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrepareProfileImageUploadAccountRequest struct via the builder pattern


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


## SearchMentionAccounts

> SearchMentionAccountsResp SearchMentionAccounts(ctx).SearchMentionAccountsReq(searchMentionAccountsReq).Execute()





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
	searchMentionAccountsReq := *openapiclient.NewSearchMentionAccountsReq() // SearchMentionAccountsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountService.SearchMentionAccounts(context.Background()).SearchMentionAccountsReq(searchMentionAccountsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.SearchMentionAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchMentionAccounts`: SearchMentionAccountsResp
	fmt.Fprintf(os.Stdout, "Response from `AccountService.SearchMentionAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchMentionAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchMentionAccountsReq** | [**SearchMentionAccountsReq**](SearchMentionAccountsReq.md) |  | 

### Return type

[**SearchMentionAccountsResp**](SearchMentionAccountsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferAccountPoints

> TransferAccountPointsResp TransferAccountPoints(ctx).TransferAccountPointsReq(transferAccountPointsReq).Execute()





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
	transferAccountPointsReq := *openapiclient.NewTransferAccountPointsReq("TargetName_example", "Amount_example", "IdempotencyKey_example") // TransferAccountPointsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountService.TransferAccountPoints(context.Background()).TransferAccountPointsReq(transferAccountPointsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.TransferAccountPoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferAccountPoints`: TransferAccountPointsResp
	fmt.Fprintf(os.Stdout, "Response from `AccountService.TransferAccountPoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferAccountPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferAccountPointsReq** | [**TransferAccountPointsReq**](TransferAccountPointsReq.md) |  | 

### Return type

[**TransferAccountPointsResp**](TransferAccountPointsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmailAccount

> map[string]interface{} UpdateEmailAccount(ctx).UpdateEmailAccountReq(updateEmailAccountReq).Execute()





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
	resp, r, err := apiClient.AccountService.UpdateEmailAccount(context.Background()).UpdateEmailAccountReq(updateEmailAccountReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.UpdateEmailAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEmailAccount`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountService.UpdateEmailAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailAccountRequest struct via the builder pattern


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


## UpdatePasswordAccount

> map[string]interface{} UpdatePasswordAccount(ctx).UpdatePasswordAccountReq(updatePasswordAccountReq).Execute()





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
	resp, r, err := apiClient.AccountService.UpdatePasswordAccount(context.Background()).UpdatePasswordAccountReq(updatePasswordAccountReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.UpdatePasswordAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePasswordAccount`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountService.UpdatePasswordAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordAccountRequest struct via the builder pattern


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


## UpdatePhoneAccount

> map[string]interface{} UpdatePhoneAccount(ctx).UpdatePhoneAccountReq(updatePhoneAccountReq).Execute()





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
	resp, r, err := apiClient.AccountService.UpdatePhoneAccount(context.Background()).UpdatePhoneAccountReq(updatePhoneAccountReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.UpdatePhoneAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePhoneAccount`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountService.UpdatePhoneAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePhoneAccountRequest struct via the builder pattern


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


## UpdateProfileAccount

> UpdateProfileAccountResp UpdateProfileAccount(ctx).UpdateProfileAccountReq(updateProfileAccountReq).Execute()





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
	resp, r, err := apiClient.AccountService.UpdateProfileAccount(context.Background()).UpdateProfileAccountReq(updateProfileAccountReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.UpdateProfileAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProfileAccount`: UpdateProfileAccountResp
	fmt.Fprintf(os.Stdout, "Response from `AccountService.UpdateProfileAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileAccountRequest struct via the builder pattern


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

