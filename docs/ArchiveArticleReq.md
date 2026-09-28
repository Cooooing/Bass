# ArchiveArticleReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArticleId** | **string** |  | 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewArchiveArticleReq

`func NewArchiveArticleReq(articleId string, ) *ArchiveArticleReq`

NewArchiveArticleReq instantiates a new ArchiveArticleReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveArticleReqWithDefaults

`func NewArchiveArticleReqWithDefaults() *ArchiveArticleReq`

NewArchiveArticleReqWithDefaults instantiates a new ArchiveArticleReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticleId

`func (o *ArchiveArticleReq) GetArticleId() string`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *ArchiveArticleReq) GetArticleIdOk() (*string, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *ArchiveArticleReq) SetArticleId(v string)`

SetArticleId sets ArticleId field to given value.


### GetReason

`func (o *ArchiveArticleReq) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ArchiveArticleReq) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ArchiveArticleReq) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ArchiveArticleReq) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


