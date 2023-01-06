# HostStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** |  | [optional] 
**NumberOfQueries** | Pointer to **int32** |  | [optional] 
**NumberOfGrabs** | Pointer to **int32** |  | [optional] 

## Methods

### NewHostStatistics

`func NewHostStatistics() *HostStatistics`

NewHostStatistics instantiates a new HostStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostStatisticsWithDefaults

`func NewHostStatisticsWithDefaults() *HostStatistics`

NewHostStatisticsWithDefaults instantiates a new HostStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *HostStatistics) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HostStatistics) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HostStatistics) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *HostStatistics) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *HostStatistics) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *HostStatistics) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetNumberOfQueries

`func (o *HostStatistics) GetNumberOfQueries() int32`

GetNumberOfQueries returns the NumberOfQueries field if non-nil, zero value otherwise.

### GetNumberOfQueriesOk

`func (o *HostStatistics) GetNumberOfQueriesOk() (*int32, bool)`

GetNumberOfQueriesOk returns a tuple with the NumberOfQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfQueries

`func (o *HostStatistics) SetNumberOfQueries(v int32)`

SetNumberOfQueries sets NumberOfQueries field to given value.

### HasNumberOfQueries

`func (o *HostStatistics) HasNumberOfQueries() bool`

HasNumberOfQueries returns a boolean if a field has been set.

### GetNumberOfGrabs

`func (o *HostStatistics) GetNumberOfGrabs() int32`

GetNumberOfGrabs returns the NumberOfGrabs field if non-nil, zero value otherwise.

### GetNumberOfGrabsOk

`func (o *HostStatistics) GetNumberOfGrabsOk() (*int32, bool)`

GetNumberOfGrabsOk returns a tuple with the NumberOfGrabs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfGrabs

`func (o *HostStatistics) SetNumberOfGrabs(v int32)`

SetNumberOfGrabs sets NumberOfGrabs field to given value.

### HasNumberOfGrabs

`func (o *HostStatistics) HasNumberOfGrabs() bool`

HasNumberOfGrabs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


