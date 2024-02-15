# IndexerStatusResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**IndexerId** | Pointer to **int32** |  | [optional] 
**DisabledTill** | Pointer to **NullableTime** |  | [optional] 
**MostRecentFailure** | Pointer to **NullableTime** |  | [optional] 
**InitialFailure** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewIndexerStatusResource

`func NewIndexerStatusResource() *IndexerStatusResource`

NewIndexerStatusResource instantiates a new IndexerStatusResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexerStatusResourceWithDefaults

`func NewIndexerStatusResourceWithDefaults() *IndexerStatusResource`

NewIndexerStatusResourceWithDefaults instantiates a new IndexerStatusResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IndexerStatusResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IndexerStatusResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IndexerStatusResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IndexerStatusResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndexerId

`func (o *IndexerStatusResource) GetIndexerId() int32`

GetIndexerId returns the IndexerId field if non-nil, zero value otherwise.

### GetIndexerIdOk

`func (o *IndexerStatusResource) GetIndexerIdOk() (*int32, bool)`

GetIndexerIdOk returns a tuple with the IndexerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexerId

`func (o *IndexerStatusResource) SetIndexerId(v int32)`

SetIndexerId sets IndexerId field to given value.

### HasIndexerId

`func (o *IndexerStatusResource) HasIndexerId() bool`

HasIndexerId returns a boolean if a field has been set.

### GetDisabledTill

`func (o *IndexerStatusResource) GetDisabledTill() time.Time`

GetDisabledTill returns the DisabledTill field if non-nil, zero value otherwise.

### GetDisabledTillOk

`func (o *IndexerStatusResource) GetDisabledTillOk() (*time.Time, bool)`

GetDisabledTillOk returns a tuple with the DisabledTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledTill

`func (o *IndexerStatusResource) SetDisabledTill(v time.Time)`

SetDisabledTill sets DisabledTill field to given value.

### HasDisabledTill

`func (o *IndexerStatusResource) HasDisabledTill() bool`

HasDisabledTill returns a boolean if a field has been set.

### SetDisabledTillNil

`func (o *IndexerStatusResource) SetDisabledTillNil(b bool)`

 SetDisabledTillNil sets the value for DisabledTill to be an explicit nil

### UnsetDisabledTill
`func (o *IndexerStatusResource) UnsetDisabledTill()`

UnsetDisabledTill ensures that no value is present for DisabledTill, not even an explicit nil
### GetMostRecentFailure

`func (o *IndexerStatusResource) GetMostRecentFailure() time.Time`

GetMostRecentFailure returns the MostRecentFailure field if non-nil, zero value otherwise.

### GetMostRecentFailureOk

`func (o *IndexerStatusResource) GetMostRecentFailureOk() (*time.Time, bool)`

GetMostRecentFailureOk returns a tuple with the MostRecentFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostRecentFailure

`func (o *IndexerStatusResource) SetMostRecentFailure(v time.Time)`

SetMostRecentFailure sets MostRecentFailure field to given value.

### HasMostRecentFailure

`func (o *IndexerStatusResource) HasMostRecentFailure() bool`

HasMostRecentFailure returns a boolean if a field has been set.

### SetMostRecentFailureNil

`func (o *IndexerStatusResource) SetMostRecentFailureNil(b bool)`

 SetMostRecentFailureNil sets the value for MostRecentFailure to be an explicit nil

### UnsetMostRecentFailure
`func (o *IndexerStatusResource) UnsetMostRecentFailure()`

UnsetMostRecentFailure ensures that no value is present for MostRecentFailure, not even an explicit nil
### GetInitialFailure

`func (o *IndexerStatusResource) GetInitialFailure() time.Time`

GetInitialFailure returns the InitialFailure field if non-nil, zero value otherwise.

### GetInitialFailureOk

`func (o *IndexerStatusResource) GetInitialFailureOk() (*time.Time, bool)`

GetInitialFailureOk returns a tuple with the InitialFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFailure

`func (o *IndexerStatusResource) SetInitialFailure(v time.Time)`

SetInitialFailure sets InitialFailure field to given value.

### HasInitialFailure

`func (o *IndexerStatusResource) HasInitialFailure() bool`

HasInitialFailure returns a boolean if a field has been set.

### SetInitialFailureNil

`func (o *IndexerStatusResource) SetInitialFailureNil(b bool)`

 SetInitialFailureNil sets the value for InitialFailure to be an explicit nil

### UnsetInitialFailure
`func (o *IndexerStatusResource) UnsetInitialFailure()`

UnsetInitialFailure ensures that no value is present for InitialFailure, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


