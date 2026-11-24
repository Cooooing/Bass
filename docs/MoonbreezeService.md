# \MoonbreezeService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_moonbreeze**](MoonbreezeService.md#create_moonbreeze) | **POST** /v1/content/moonbreeze/create | 
[**page_member_moonbreezes**](MoonbreezeService.md#page_member_moonbreezes) | **POST** /v1/content/moonbreeze/page-member | 
[**page_public_moonbreezes**](MoonbreezeService.md#page_public_moonbreezes) | **POST** /v1/content/moonbreeze/page-public | 
[**page_watching_moonbreezes**](MoonbreezeService.md#page_watching_moonbreezes) | **POST** /v1/content/moonbreeze/page-watching | 



## create_moonbreeze

> models::CreateMoonbreezeResp create_moonbreeze(create_moonbreeze_req)


服务方法。

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


## page_member_moonbreezes

> models::PageMemberMoonbreezesResp page_member_moonbreezes(page_member_moonbreezes_req)


服务方法。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**page_member_moonbreezes_req** | [**PageMemberMoonbreezesReq**](PageMemberMoonbreezesReq.md) |  | [required] |

### Return type

[**models::PageMemberMoonbreezesResp**](PageMemberMoonbreezes_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## page_public_moonbreezes

> models::PagePublicMoonbreezesResp page_public_moonbreezes(page_public_moonbreezes_req)


服务方法。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**page_public_moonbreezes_req** | [**PagePublicMoonbreezesReq**](PagePublicMoonbreezesReq.md) |  | [required] |

### Return type

[**models::PagePublicMoonbreezesResp**](PagePublicMoonbreezes_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## page_watching_moonbreezes

> models::PageWatchingMoonbreezesResp page_watching_moonbreezes(page_watching_moonbreezes_req)


服务方法。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**page_watching_moonbreezes_req** | [**PageWatchingMoonbreezesReq**](PageWatchingMoonbreezesReq.md) |  | [required] |

### Return type

[**models::PageWatchingMoonbreezesResp**](PageWatchingMoonbreezes_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

