

# TagQuery

标签查询条件。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**ids** | **List&lt;String&gt;** | 标签 ID 列表。 |  [optional] |
|**name** | **String** | 标签名称。 |  [optional] |
|**names** | **List&lt;String&gt;** | 标签名称列表。 |  [optional] |
|**description** | **String** | 标签描述。 |  [optional] |
|**status** | [**StatusEnum**](#StatusEnum) | 标签启停状态。 |  [optional] |
|**domainId** | **String** | 所属板块 ID。 |  [optional] |



## Enum: StatusEnum

| Name | Value |
|---- | -----|
| TAG_STATUS_UNSPECIFIED | &quot;TAG_STATUS_UNSPECIFIED&quot; |
| TAG_STATUS_ENABLED | &quot;TAG_STATUS_ENABLED&quot; |
| TAG_STATUS_DISABLED | &quot;TAG_STATUS_DISABLED&quot; |



