# \CharacterService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_character**](CharacterService.md#create_character) | **POST** /v1/game-idle/character/create | 
[**list_character**](CharacterService.md#list_character) | **POST** /v1/game-idle/character/list | 



## create_character

> models::CreateCharacterResp create_character(create_character_req)


创建角色。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**create_character_req** | [**CreateCharacterReq**](CreateCharacterReq.md) |  | [required] |

### Return type

[**models::CreateCharacterResp**](CreateCharacter_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_character

> models::ListCharacterResp list_character(body)


查询当前账号角色列表。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::ListCharacterResp**](ListCharacter_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

