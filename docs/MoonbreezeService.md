# \MoonbreezeService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create**](MoonbreezeService.md#create) | **POST** /v1/content/moonbreeze/create | 
[**page_member**](MoonbreezeService.md#page_member) | **POST** /v1/content/moonbreeze/page-member | 
[**page_public**](MoonbreezeService.md#page_public) | **POST** /v1/content/moonbreeze/page-public | 
[**page_watching**](MoonbreezeService.md#page_watching) | **POST** /v1/content/moonbreeze/page-watching | 



## create

> models::CreateMoonbreezeResp create(create_moonbreeze_req)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**create_moonbreeze_req** | [**CreateMoonbreezeReq**](CreateMoonbreezeReq.md) |  | [required] |

### Return type

[**models::CreateMoonbreezeResp**](CreateMoonbreeze_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## page_member

> models::PageMoonbreezesResp page_member(page_member_moonbreezes_req)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**page_member_moonbreezes_req** | [**PageMemberMoonbreezesReq**](PageMemberMoonbreezesReq.md) |  | [required] |

### Return type

[**models::PageMoonbreezesResp**](PageMoonbreezes_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## page_public

> models::PageMoonbreezesResp page_public(page_public_moonbreezes_req)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**page_public_moonbreezes_req** | [**PagePublicMoonbreezesReq**](PagePublicMoonbreezesReq.md) |  | [required] |

### Return type

[**models::PageMoonbreezesResp**](PageMoonbreezes_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## page_watching

> models::PageMoonbreezesResp page_watching(page_watching_moonbreezes_req)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**page_watching_moonbreezes_req** | [**PageWatchingMoonbreezesReq**](PageWatchingMoonbreezesReq.md) |  | [required] |

### Return type

[**models::PageMoonbreezesResp**](PageMoonbreezes_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

