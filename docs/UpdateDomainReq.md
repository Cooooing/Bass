# UpdateDomainReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | **string** |  | 
**Domain** | [**ReqDomain**](ReqDomain.md) |  | 

## Methods

### NewUpdateDomainReq

`func NewUpdateDomainReq(domainId string, domain ReqDomain, ) *UpdateDomainReq`

NewUpdateDomainReq instantiates a new UpdateDomainReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDomainReqWithDefaults

`func NewUpdateDomainReqWithDefaults() *UpdateDomainReq`

NewUpdateDomainReqWithDefaults instantiates a new UpdateDomainReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *UpdateDomainReq) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *UpdateDomainReq) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *UpdateDomainReq) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetDomain

`func (o *UpdateDomainReq) GetDomain() ReqDomain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UpdateDomainReq) GetDomainOk() (*ReqDomain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UpdateDomainReq) SetDomain(v ReqDomain)`

SetDomain sets Domain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


