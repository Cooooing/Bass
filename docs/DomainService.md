# \DomainService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create**](DomainService.md#create) | **POST** /v1/content/domain/create | 
[**list**](DomainService.md#list) | **POST** /v1/content/domain/list | 
[**update**](DomainService.md#update) | **POST** /v1/content/domain/update | 



## create

> models::CreateDomainResp create(create_domain_req)


创建领域。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**create_domain_req** | [**CreateDomainReq**](CreateDomainReq.md) |  | [required] |

### Return type

[**models::CreateDomainResp**](CreateDomain_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list

> models::ListDomainsResp list(list_domains_req)


查询领域列表。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_domains_req** | [**ListDomainsReq**](ListDomainsReq.md) |  | [required] |

### Return type

[**models::ListDomainsResp**](ListDomains_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## update

> models::UpdateDomainResp update(update_domain_req)


更新领域。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**update_domain_req** | [**UpdateDomainReq**](UpdateDomainReq.md) |  | [required] |

### Return type

[**models::UpdateDomainResp**](UpdateDomain_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

