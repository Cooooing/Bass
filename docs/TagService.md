# \TagService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**bind_article_tags**](TagService.md#bind_article_tags) | **POST** /v1/content/tag/bind-article | 
[**create_tag**](TagService.md#create_tag) | **POST** /v1/content/tag/create | 
[**list_article_tags**](TagService.md#list_article_tags) | **POST** /v1/content/tag/list-article-tags | 
[**list_tags**](TagService.md#list_tags) | **POST** /v1/content/tag/list | 
[**unbind_article_tags**](TagService.md#unbind_article_tags) | **POST** /v1/content/tag/unbind-article | 
[**update_tag**](TagService.md#update_tag) | **POST** /v1/content/tag/update | 



## bind_article_tags

> serde_json::Value bind_article_tags(bind_article_tags_req)


绑定文章标签。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**bind_article_tags_req** | [**BindArticleTagsReq**](BindArticleTagsReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## create_tag

> models::CreateTagResp create_tag(create_tag_req)


创建标签。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**create_tag_req** | [**CreateTagReq**](CreateTagReq.md) |  | [required] |

### Return type

[**models::CreateTagResp**](CreateTag_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_article_tags

> models::ListArticleTagsResp list_article_tags(list_article_tags_req)


查询文章标签列表。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_article_tags_req** | [**ListArticleTagsReq**](ListArticleTagsReq.md) |  | [required] |

### Return type

[**models::ListArticleTagsResp**](ListArticleTags_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_tags

> models::ListTagsResp list_tags(list_tags_req)


查询标签列表。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_tags_req** | [**ListTagsReq**](ListTagsReq.md) |  | [required] |

### Return type

[**models::ListTagsResp**](ListTags_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## unbind_article_tags

> serde_json::Value unbind_article_tags(unbind_article_tags_req)


解绑文章标签。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**unbind_article_tags_req** | [**UnbindArticleTagsReq**](UnbindArticleTagsReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## update_tag

> models::UpdateTagResp update_tag(update_tag_req)


更新标签。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**update_tag_req** | [**UpdateTagReq**](UpdateTagReq.md) |  | [required] |

### Return type

[**models::UpdateTagResp**](UpdateTag_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

