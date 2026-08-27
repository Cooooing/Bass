

# RespRecord


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** |  |  [optional] |
|**transactionNo** | **String** |  |  [optional] |
|**recordType** | [**RecordTypeEnum**](#RecordTypeEnum) |  |  [optional] |
|**direction** | [**DirectionEnum**](#DirectionEnum) |  |  [optional] |
|**amount** | **String** |  |  [optional] |
|**balanceBefore** | **String** |  |  [optional] |
|**balanceAfter** | **String** |  |  [optional] |
|**remark** | **String** |  |  [optional] |
|**createdAt** | **OffsetDateTime** |  |  [optional] |



## Enum: RecordTypeEnum

| Name | Value |
|---- | -----|
| ECONOMY_RECORD_TYPE_UNSPECIFIED | &quot;ECONOMY_RECORD_TYPE_UNSPECIFIED&quot; |
| ECONOMY_RECORD_TYPE_SIGN_IN_REWARD | &quot;ECONOMY_RECORD_TYPE_SIGN_IN_REWARD&quot; |
| ECONOMY_RECORD_TYPE_ARTICLE_THANK_REWARD | &quot;ECONOMY_RECORD_TYPE_ARTICLE_THANK_REWARD&quot; |
| ECONOMY_RECORD_TYPE_COMMENT_THANK_REWARD | &quot;ECONOMY_RECORD_TYPE_COMMENT_THANK_REWARD&quot; |
| ECONOMY_RECORD_TYPE_ARTICLE_REWARD_OUT | &quot;ECONOMY_RECORD_TYPE_ARTICLE_REWARD_OUT&quot; |
| ECONOMY_RECORD_TYPE_ARTICLE_REWARD_IN | &quot;ECONOMY_RECORD_TYPE_ARTICLE_REWARD_IN&quot; |
| ECONOMY_RECORD_TYPE_ADMIN_ADD | &quot;ECONOMY_RECORD_TYPE_ADMIN_ADD&quot; |
| ECONOMY_RECORD_TYPE_ADMIN_DEDUCT | &quot;ECONOMY_RECORD_TYPE_ADMIN_DEDUCT&quot; |



## Enum: DirectionEnum

| Name | Value |
|---- | -----|
| ECONOMY_RECORD_DIRECTION_UNSPECIFIED | &quot;ECONOMY_RECORD_DIRECTION_UNSPECIFIED&quot; |
| ECONOMY_RECORD_DIRECTION_INCOME | &quot;ECONOMY_RECORD_DIRECTION_INCOME&quot; |
| ECONOMY_RECORD_DIRECTION_EXPENSE | &quot;ECONOMY_RECORD_DIRECTION_EXPENSE&quot; |



