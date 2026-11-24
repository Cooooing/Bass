# \DomainService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_domain**](DomainService.md#create_domain) | **POST** /v1/content/domain/create | 
[**list_domains**](DomainService.md#list_domains) | **POST** /v1/content/domain/list | 
[**update_domain**](DomainService.md#update_domain) | **POST** /v1/content/domain/update | 



## create_domain

> models::CreateDomainResp create_domain(create_domain_req)


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


## list_domains

> models::ListDomainsResp list_domains(list_domains_req)


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


## update_domain

> models::UpdateDomainResp update_domain(update_domain_req)


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

