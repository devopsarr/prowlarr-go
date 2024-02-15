# UserAgentStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserAgent** | Pointer to **NullableString** |  | [optional] 
**NumberOfQueries** | Pointer to **int32** |  | [optional] 
**NumberOfGrabs** | Pointer to **int32** |  | [optional] 

## Methods

### NewUserAgentStatistics

`func NewUserAgentStatistics() *UserAgentStatistics`

NewUserAgentStatistics instantiates a new UserAgentStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAgentStatisticsWithDefaults

`func NewUserAgentStatisticsWithDefaults() *UserAgentStatistics`

NewUserAgentStatisticsWithDefaults instantiates a new UserAgentStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserAgent

`func (o *UserAgentStatistics) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *UserAgentStatistics) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *UserAgentStatistics) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *UserAgentStatistics) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### SetUserAgentNil

`func (o *UserAgentStatistics) SetUserAgentNil(b bool)`

 SetUserAgentNil sets the value for UserAgent to be an explicit nil

### UnsetUserAgent
`func (o *UserAgentStatistics) UnsetUserAgent()`

UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil
### GetNumberOfQueries

`func (o *UserAgentStatistics) GetNumberOfQueries() int32`

GetNumberOfQueries returns the NumberOfQueries field if non-nil, zero value otherwise.

### GetNumberOfQueriesOk

`func (o *UserAgentStatistics) GetNumberOfQueriesOk() (*int32, bool)`

GetNumberOfQueriesOk returns a tuple with the NumberOfQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfQueries

`func (o *UserAgentStatistics) SetNumberOfQueries(v int32)`

SetNumberOfQueries sets NumberOfQueries field to given value.

### HasNumberOfQueries

`func (o *UserAgentStatistics) HasNumberOfQueries() bool`

HasNumberOfQueries returns a boolean if a field has been set.

### GetNumberOfGrabs

`func (o *UserAgentStatistics) GetNumberOfGrabs() int32`

GetNumberOfGrabs returns the NumberOfGrabs field if non-nil, zero value otherwise.

### GetNumberOfGrabsOk

`func (o *UserAgentStatistics) GetNumberOfGrabsOk() (*int32, bool)`

GetNumberOfGrabsOk returns a tuple with the NumberOfGrabs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfGrabs

`func (o *UserAgentStatistics) SetNumberOfGrabs(v int32)`

SetNumberOfGrabs sets NumberOfGrabs field to given value.

### HasNumberOfGrabs

`func (o *UserAgentStatistics) HasNumberOfGrabs() bool`

HasNumberOfGrabs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


