# ListNotificationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PageRequest**](PageRequest.md) | 分页参数。 | [optional] 

## Methods

### NewListNotificationsRequest

`func NewListNotificationsRequest() *ListNotificationsRequest`

NewListNotificationsRequest instantiates a new ListNotificationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNotificationsRequestWithDefaults

`func NewListNotificationsRequestWithDefaults() *ListNotificationsRequest`

NewListNotificationsRequestWithDefaults instantiates a new ListNotificationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListNotificationsRequest) GetPage() PageRequest`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListNotificationsRequest) GetPageOk() (*PageRequest, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListNotificationsRequest) SetPage(v PageRequest)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListNotificationsRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


