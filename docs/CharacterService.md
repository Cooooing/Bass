# \CharacterService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCharacter**](CharacterService.md#CreateCharacter) | **Post** /v1/game-idle/character/create | 
[**ListCharacter**](CharacterService.md#ListCharacter) | **Post** /v1/game-idle/character/list | 



## CreateCharacter

> CreateCharacterResp CreateCharacter(ctx).CreateCharacterReq(createCharacterReq).Execute()





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
	createCharacterReq := *openapiclient.NewCreateCharacterReq() // CreateCharacterReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CharacterService.CreateCharacter(context.Background()).CreateCharacterReq(createCharacterReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CharacterService.CreateCharacter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCharacter`: CreateCharacterResp
	fmt.Fprintf(os.Stdout, "Response from `CharacterService.CreateCharacter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCharacterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCharacterReq** | [**CreateCharacterReq**](CreateCharacterReq.md) |  | 

### Return type

[**CreateCharacterResp**](CreateCharacterResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCharacter

> ListCharacterResp ListCharacter(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.CharacterService.ListCharacter(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CharacterService.ListCharacter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCharacter`: ListCharacterResp
	fmt.Fprintf(os.Stdout, "Response from `CharacterService.ListCharacter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCharacterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**ListCharacterResp**](ListCharacterResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

