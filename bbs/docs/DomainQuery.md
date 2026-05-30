

# DomainQuery

内容板块查询条件。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**ids** | **List&lt;String&gt;** | 板块 ID 列表。 |  [optional] |
|**name** | **String** | 板块名称。 |  [optional] |
|**description** | **String** | 板块描述。 |  [optional] |
|**status** | [**StatusEnum**](#StatusEnum) | 板块状态。 |  [optional] |
|**url** | **String** | 板块 URL。 |  [optional] |
|**icon** | **String** | 板块图标。 |  [optional] |
|**isNav** | **Boolean** | 是否在导航中展示。 |  [optional] |



## Enum: StatusEnum

| Name | Value |
|---- | -----|
| DOMAIN_STATUS_UNSPECIFIED | &quot;DOMAIN_STATUS_UNSPECIFIED&quot; |
| DOMAIN_STATUS_NORMAL | &quot;DOMAIN_STATUS_NORMAL&quot; |
| DOMAIN_STATUS_DISABLED | &quot;DOMAIN_STATUS_DISABLED&quot; |



