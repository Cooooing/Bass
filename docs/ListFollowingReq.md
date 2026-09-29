# ListFollowingReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Page** | Pointer to [**PageReq**](PageReq.md) |  | [optional] 

## Methods

### NewListFollowingReq

`func NewListFollowingReq(name string, ) *ListFollowingReq`

NewListFollowingReq instantiates a new ListFollowingReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFollowingReqWithDefaults

`func NewListFollowingReqWithDefaults() *ListFollowingReq`

NewListFollowingReqWithDefaults instantiates a new ListFollowingReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ListFollowingReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListFollowingReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListFollowingReq) SetName(v string)`

SetName sets Name field to given value.


### GetPage

`func (o *ListFollowingReq) GetPage() PageReq`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListFollowingReq) GetPageOk() (*PageReq, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListFollowingReq) SetPage(v PageReq)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListFollowingReq) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


