# SchedulePublishArticleReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArticleId** | **string** |  | 
**PublishAt** | **time.Time** |  | 

## Methods

### NewSchedulePublishArticleReq

`func NewSchedulePublishArticleReq(articleId string, publishAt time.Time, ) *SchedulePublishArticleReq`

NewSchedulePublishArticleReq instantiates a new SchedulePublishArticleReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePublishArticleReqWithDefaults

`func NewSchedulePublishArticleReqWithDefaults() *SchedulePublishArticleReq`

NewSchedulePublishArticleReqWithDefaults instantiates a new SchedulePublishArticleReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticleId

`func (o *SchedulePublishArticleReq) GetArticleId() string`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *SchedulePublishArticleReq) GetArticleIdOk() (*string, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *SchedulePublishArticleReq) SetArticleId(v string)`

SetArticleId sets ArticleId field to given value.


### GetPublishAt

`func (o *SchedulePublishArticleReq) GetPublishAt() time.Time`

GetPublishAt returns the PublishAt field if non-nil, zero value otherwise.

### GetPublishAtOk

`func (o *SchedulePublishArticleReq) GetPublishAtOk() (*time.Time, bool)`

GetPublishAtOk returns a tuple with the PublishAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishAt

`func (o *SchedulePublishArticleReq) SetPublishAt(v time.Time)`

SetPublishAt sets PublishAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


