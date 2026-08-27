# \CharacterService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create**](CharacterService.md#create) | **POST** /v1/game-idle/character/create | 
[**list**](CharacterService.md#list) | **POST** /v1/game-idle/character/list | 



## create

> models::CreateCharacterResp create(create_character_req)


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


## list

> models::ListCharacterResp list(body)


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

