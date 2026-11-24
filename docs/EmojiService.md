# \EmojiService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_emoji_item**](EmojiService.md#add_emoji_item) | **POST** /v1/im/emoji/items | 
[**create_emoji_group**](EmojiService.md#create_emoji_group) | **POST** /v1/im/emoji/groups | 
[**create_emoji_share**](EmojiService.md#create_emoji_share) | **POST** /v1/im/emoji/shares | 
[**delete_emoji_group**](EmojiService.md#delete_emoji_group) | **POST** /v1/im/emoji/groups/delete | 
[**delete_emoji_item**](EmojiService.md#delete_emoji_item) | **POST** /v1/im/emoji/items/delete | 
[**import_emoji_share**](EmojiService.md#import_emoji_share) | **POST** /v1/im/emoji/shares/import | 
[**list_emoji_groups**](EmojiService.md#list_emoji_groups) | **POST** /v1/im/emoji/groups/list | 
[**move_emoji_item**](EmojiService.md#move_emoji_item) | **POST** /v1/im/emoji/items/move | 
[**rename_emoji_group**](EmojiService.md#rename_emoji_group) | **POST** /v1/im/emoji/groups/rename | 
[**rename_emoji_item**](EmojiService.md#rename_emoji_item) | **POST** /v1/im/emoji/items/rename | 
[**reorder_emoji_groups**](EmojiService.md#reorder_emoji_groups) | **POST** /v1/im/emoji/groups/reorder | 
[**reorder_emoji_items**](EmojiService.md#reorder_emoji_items) | **POST** /v1/im/emoji/items/reorder | 



## add_emoji_item

> models::AddEmojiItemResp add_emoji_item(add_emoji_item_req)


向分组加入已上传图片。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**add_emoji_item_req** | [**AddEmojiItemReq**](AddEmojiItemReq.md) |  | [required] |

### Return type

[**models::AddEmojiItemResp**](AddEmojiItem_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## create_emoji_group

> models::CreateEmojiGroupResp create_emoji_group(create_emoji_group_req)


创建一个表情分组。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**create_emoji_group_req** | [**CreateEmojiGroupReq**](CreateEmojiGroupReq.md) |  | [required] |

### Return type

[**models::CreateEmojiGroupResp**](CreateEmojiGroup_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## create_emoji_share

> models::CreateEmojiShareResp create_emoji_share(create_emoji_share_req)


生成七天有效的分组分享码。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**create_emoji_share_req** | [**CreateEmojiShareReq**](CreateEmojiShareReq.md) |  | [required] |

### Return type

[**models::CreateEmojiShareResp**](CreateEmojiShare_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## delete_emoji_group

> serde_json::Value delete_emoji_group(delete_emoji_group_req)


删除一个表情分组。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**delete_emoji_group_req** | [**DeleteEmojiGroupReq**](DeleteEmojiGroupReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## delete_emoji_item

> serde_json::Value delete_emoji_item(delete_emoji_item_req)


从分组移除表情。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**delete_emoji_item_req** | [**DeleteEmojiItemReq**](DeleteEmojiItemReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## import_emoji_share

> models::ImportEmojiShareResp import_emoji_share(import_emoji_share_req)


把分享码导入为一个新分组。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**import_emoji_share_req** | [**ImportEmojiShareReq**](ImportEmojiShareReq.md) |  | [required] |

### Return type

[**models::ImportEmojiShareResp**](ImportEmojiShare_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_emoji_groups

> models::ListEmojiGroupsResp list_emoji_groups(list_emoji_groups_req)


查询当前账号的表情分组。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_emoji_groups_req** | [**ListEmojiGroupsReq**](ListEmojiGroupsReq.md) |  | [required] |

### Return type

[**models::ListEmojiGroupsResp**](ListEmojiGroups_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## move_emoji_item

> models::MoveEmojiItemResp move_emoji_item(move_emoji_item_req)


将表情移动到另一个当前账号分组。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**move_emoji_item_req** | [**MoveEmojiItemReq**](MoveEmojiItemReq.md) |  | [required] |

### Return type

[**models::MoveEmojiItemResp**](MoveEmojiItem_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rename_emoji_group

> models::RenameEmojiGroupResp rename_emoji_group(rename_emoji_group_req)


重命名一个表情分组。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**rename_emoji_group_req** | [**RenameEmojiGroupReq**](RenameEmojiGroupReq.md) |  | [required] |

### Return type

[**models::RenameEmojiGroupResp**](RenameEmojiGroup_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## rename_emoji_item

> models::RenameEmojiItemResp rename_emoji_item(rename_emoji_item_req)


重命名一个表情别名。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**rename_emoji_item_req** | [**RenameEmojiItemReq**](RenameEmojiItemReq.md) |  | [required] |

### Return type

[**models::RenameEmojiItemResp**](RenameEmojiItem_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## reorder_emoji_groups

> serde_json::Value reorder_emoji_groups(reorder_emoji_groups_req)


重排表情分组。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**reorder_emoji_groups_req** | [**ReorderEmojiGroupsReq**](ReorderEmojiGroupsReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## reorder_emoji_items

> serde_json::Value reorder_emoji_items(reorder_emoji_items_req)


重排同一分组内的表情。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**reorder_emoji_items_req** | [**ReorderEmojiItemsReq**](ReorderEmojiItemsReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

