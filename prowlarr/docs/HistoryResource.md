# HistoryResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**IndexerId** | Pointer to **int32** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**DownloadId** | Pointer to **NullableString** |  | [optional] 
**Successful** | Pointer to **bool** |  | [optional] 
**EventType** | Pointer to [**HistoryEventType**](HistoryEventType.md) |  | [optional] 
**Data** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewHistoryResource

`func NewHistoryResource() *HistoryResource`

NewHistoryResource instantiates a new HistoryResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryResourceWithDefaults

`func NewHistoryResourceWithDefaults() *HistoryResource`

NewHistoryResourceWithDefaults instantiates a new HistoryResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoryResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HistoryResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndexerId

`func (o *HistoryResource) GetIndexerId() int32`

GetIndexerId returns the IndexerId field if non-nil, zero value otherwise.

### GetIndexerIdOk

`func (o *HistoryResource) GetIndexerIdOk() (*int32, bool)`

GetIndexerIdOk returns a tuple with the IndexerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexerId

`func (o *HistoryResource) SetIndexerId(v int32)`

SetIndexerId sets IndexerId field to given value.

### HasIndexerId

`func (o *HistoryResource) HasIndexerId() bool`

HasIndexerId returns a boolean if a field has been set.

### GetDate

`func (o *HistoryResource) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *HistoryResource) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *HistoryResource) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *HistoryResource) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDownloadId

`func (o *HistoryResource) GetDownloadId() string`

GetDownloadId returns the DownloadId field if non-nil, zero value otherwise.

### GetDownloadIdOk

`func (o *HistoryResource) GetDownloadIdOk() (*string, bool)`

GetDownloadIdOk returns a tuple with the DownloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadId

`func (o *HistoryResource) SetDownloadId(v string)`

SetDownloadId sets DownloadId field to given value.

### HasDownloadId

`func (o *HistoryResource) HasDownloadId() bool`

HasDownloadId returns a boolean if a field has been set.

### SetDownloadIdNil

`func (o *HistoryResource) SetDownloadIdNil(b bool)`

 SetDownloadIdNil sets the value for DownloadId to be an explicit nil

### UnsetDownloadId
`func (o *HistoryResource) UnsetDownloadId()`

UnsetDownloadId ensures that no value is present for DownloadId, not even an explicit nil
### GetSuccessful

`func (o *HistoryResource) GetSuccessful() bool`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *HistoryResource) GetSuccessfulOk() (*bool, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *HistoryResource) SetSuccessful(v bool)`

SetSuccessful sets Successful field to given value.

### HasSuccessful

`func (o *HistoryResource) HasSuccessful() bool`

HasSuccessful returns a boolean if a field has been set.

### GetEventType

`func (o *HistoryResource) GetEventType() HistoryEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *HistoryResource) GetEventTypeOk() (*HistoryEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *HistoryResource) SetEventType(v HistoryEventType)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *HistoryResource) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetData

`func (o *HistoryResource) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HistoryResource) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HistoryResource) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *HistoryResource) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *HistoryResource) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *HistoryResource) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


