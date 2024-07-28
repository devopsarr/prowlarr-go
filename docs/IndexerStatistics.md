# IndexerStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexerId** | Pointer to **int32** |  | [optional] 
**IndexerName** | Pointer to **NullableString** |  | [optional] 
**AverageResponseTime** | Pointer to **int32** |  | [optional] 
**AverageGrabResponseTime** | Pointer to **int32** |  | [optional] 
**NumberOfQueries** | Pointer to **int32** |  | [optional] 
**NumberOfGrabs** | Pointer to **int32** |  | [optional] 
**NumberOfRssQueries** | Pointer to **int32** |  | [optional] 
**NumberOfAuthQueries** | Pointer to **int32** |  | [optional] 
**NumberOfFailedQueries** | Pointer to **int32** |  | [optional] 
**NumberOfFailedGrabs** | Pointer to **int32** |  | [optional] 
**NumberOfFailedRssQueries** | Pointer to **int32** |  | [optional] 
**NumberOfFailedAuthQueries** | Pointer to **int32** |  | [optional] 

## Methods

### NewIndexerStatistics

`func NewIndexerStatistics() *IndexerStatistics`

NewIndexerStatistics instantiates a new IndexerStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexerStatisticsWithDefaults

`func NewIndexerStatisticsWithDefaults() *IndexerStatistics`

NewIndexerStatisticsWithDefaults instantiates a new IndexerStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexerId

`func (o *IndexerStatistics) GetIndexerId() int32`

GetIndexerId returns the IndexerId field if non-nil, zero value otherwise.

### GetIndexerIdOk

`func (o *IndexerStatistics) GetIndexerIdOk() (*int32, bool)`

GetIndexerIdOk returns a tuple with the IndexerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexerId

`func (o *IndexerStatistics) SetIndexerId(v int32)`

SetIndexerId sets IndexerId field to given value.

### HasIndexerId

`func (o *IndexerStatistics) HasIndexerId() bool`

HasIndexerId returns a boolean if a field has been set.

### GetIndexerName

`func (o *IndexerStatistics) GetIndexerName() string`

GetIndexerName returns the IndexerName field if non-nil, zero value otherwise.

### GetIndexerNameOk

`func (o *IndexerStatistics) GetIndexerNameOk() (*string, bool)`

GetIndexerNameOk returns a tuple with the IndexerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexerName

`func (o *IndexerStatistics) SetIndexerName(v string)`

SetIndexerName sets IndexerName field to given value.

### HasIndexerName

`func (o *IndexerStatistics) HasIndexerName() bool`

HasIndexerName returns a boolean if a field has been set.

### SetIndexerNameNil

`func (o *IndexerStatistics) SetIndexerNameNil(b bool)`

 SetIndexerNameNil sets the value for IndexerName to be an explicit nil

### UnsetIndexerName
`func (o *IndexerStatistics) UnsetIndexerName()`

UnsetIndexerName ensures that no value is present for IndexerName, not even an explicit nil
### GetAverageResponseTime

`func (o *IndexerStatistics) GetAverageResponseTime() int32`

GetAverageResponseTime returns the AverageResponseTime field if non-nil, zero value otherwise.

### GetAverageResponseTimeOk

`func (o *IndexerStatistics) GetAverageResponseTimeOk() (*int32, bool)`

GetAverageResponseTimeOk returns a tuple with the AverageResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageResponseTime

`func (o *IndexerStatistics) SetAverageResponseTime(v int32)`

SetAverageResponseTime sets AverageResponseTime field to given value.

### HasAverageResponseTime

`func (o *IndexerStatistics) HasAverageResponseTime() bool`

HasAverageResponseTime returns a boolean if a field has been set.

### GetAverageGrabResponseTime

`func (o *IndexerStatistics) GetAverageGrabResponseTime() int32`

GetAverageGrabResponseTime returns the AverageGrabResponseTime field if non-nil, zero value otherwise.

### GetAverageGrabResponseTimeOk

`func (o *IndexerStatistics) GetAverageGrabResponseTimeOk() (*int32, bool)`

GetAverageGrabResponseTimeOk returns a tuple with the AverageGrabResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageGrabResponseTime

`func (o *IndexerStatistics) SetAverageGrabResponseTime(v int32)`

SetAverageGrabResponseTime sets AverageGrabResponseTime field to given value.

### HasAverageGrabResponseTime

`func (o *IndexerStatistics) HasAverageGrabResponseTime() bool`

HasAverageGrabResponseTime returns a boolean if a field has been set.

### GetNumberOfQueries

`func (o *IndexerStatistics) GetNumberOfQueries() int32`

GetNumberOfQueries returns the NumberOfQueries field if non-nil, zero value otherwise.

### GetNumberOfQueriesOk

`func (o *IndexerStatistics) GetNumberOfQueriesOk() (*int32, bool)`

GetNumberOfQueriesOk returns a tuple with the NumberOfQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfQueries

`func (o *IndexerStatistics) SetNumberOfQueries(v int32)`

SetNumberOfQueries sets NumberOfQueries field to given value.

### HasNumberOfQueries

`func (o *IndexerStatistics) HasNumberOfQueries() bool`

HasNumberOfQueries returns a boolean if a field has been set.

### GetNumberOfGrabs

`func (o *IndexerStatistics) GetNumberOfGrabs() int32`

GetNumberOfGrabs returns the NumberOfGrabs field if non-nil, zero value otherwise.

### GetNumberOfGrabsOk

`func (o *IndexerStatistics) GetNumberOfGrabsOk() (*int32, bool)`

GetNumberOfGrabsOk returns a tuple with the NumberOfGrabs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfGrabs

`func (o *IndexerStatistics) SetNumberOfGrabs(v int32)`

SetNumberOfGrabs sets NumberOfGrabs field to given value.

### HasNumberOfGrabs

`func (o *IndexerStatistics) HasNumberOfGrabs() bool`

HasNumberOfGrabs returns a boolean if a field has been set.

### GetNumberOfRssQueries

`func (o *IndexerStatistics) GetNumberOfRssQueries() int32`

GetNumberOfRssQueries returns the NumberOfRssQueries field if non-nil, zero value otherwise.

### GetNumberOfRssQueriesOk

`func (o *IndexerStatistics) GetNumberOfRssQueriesOk() (*int32, bool)`

GetNumberOfRssQueriesOk returns a tuple with the NumberOfRssQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfRssQueries

`func (o *IndexerStatistics) SetNumberOfRssQueries(v int32)`

SetNumberOfRssQueries sets NumberOfRssQueries field to given value.

### HasNumberOfRssQueries

`func (o *IndexerStatistics) HasNumberOfRssQueries() bool`

HasNumberOfRssQueries returns a boolean if a field has been set.

### GetNumberOfAuthQueries

`func (o *IndexerStatistics) GetNumberOfAuthQueries() int32`

GetNumberOfAuthQueries returns the NumberOfAuthQueries field if non-nil, zero value otherwise.

### GetNumberOfAuthQueriesOk

`func (o *IndexerStatistics) GetNumberOfAuthQueriesOk() (*int32, bool)`

GetNumberOfAuthQueriesOk returns a tuple with the NumberOfAuthQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAuthQueries

`func (o *IndexerStatistics) SetNumberOfAuthQueries(v int32)`

SetNumberOfAuthQueries sets NumberOfAuthQueries field to given value.

### HasNumberOfAuthQueries

`func (o *IndexerStatistics) HasNumberOfAuthQueries() bool`

HasNumberOfAuthQueries returns a boolean if a field has been set.

### GetNumberOfFailedQueries

`func (o *IndexerStatistics) GetNumberOfFailedQueries() int32`

GetNumberOfFailedQueries returns the NumberOfFailedQueries field if non-nil, zero value otherwise.

### GetNumberOfFailedQueriesOk

`func (o *IndexerStatistics) GetNumberOfFailedQueriesOk() (*int32, bool)`

GetNumberOfFailedQueriesOk returns a tuple with the NumberOfFailedQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFailedQueries

`func (o *IndexerStatistics) SetNumberOfFailedQueries(v int32)`

SetNumberOfFailedQueries sets NumberOfFailedQueries field to given value.

### HasNumberOfFailedQueries

`func (o *IndexerStatistics) HasNumberOfFailedQueries() bool`

HasNumberOfFailedQueries returns a boolean if a field has been set.

### GetNumberOfFailedGrabs

`func (o *IndexerStatistics) GetNumberOfFailedGrabs() int32`

GetNumberOfFailedGrabs returns the NumberOfFailedGrabs field if non-nil, zero value otherwise.

### GetNumberOfFailedGrabsOk

`func (o *IndexerStatistics) GetNumberOfFailedGrabsOk() (*int32, bool)`

GetNumberOfFailedGrabsOk returns a tuple with the NumberOfFailedGrabs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFailedGrabs

`func (o *IndexerStatistics) SetNumberOfFailedGrabs(v int32)`

SetNumberOfFailedGrabs sets NumberOfFailedGrabs field to given value.

### HasNumberOfFailedGrabs

`func (o *IndexerStatistics) HasNumberOfFailedGrabs() bool`

HasNumberOfFailedGrabs returns a boolean if a field has been set.

### GetNumberOfFailedRssQueries

`func (o *IndexerStatistics) GetNumberOfFailedRssQueries() int32`

GetNumberOfFailedRssQueries returns the NumberOfFailedRssQueries field if non-nil, zero value otherwise.

### GetNumberOfFailedRssQueriesOk

`func (o *IndexerStatistics) GetNumberOfFailedRssQueriesOk() (*int32, bool)`

GetNumberOfFailedRssQueriesOk returns a tuple with the NumberOfFailedRssQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFailedRssQueries

`func (o *IndexerStatistics) SetNumberOfFailedRssQueries(v int32)`

SetNumberOfFailedRssQueries sets NumberOfFailedRssQueries field to given value.

### HasNumberOfFailedRssQueries

`func (o *IndexerStatistics) HasNumberOfFailedRssQueries() bool`

HasNumberOfFailedRssQueries returns a boolean if a field has been set.

### GetNumberOfFailedAuthQueries

`func (o *IndexerStatistics) GetNumberOfFailedAuthQueries() int32`

GetNumberOfFailedAuthQueries returns the NumberOfFailedAuthQueries field if non-nil, zero value otherwise.

### GetNumberOfFailedAuthQueriesOk

`func (o *IndexerStatistics) GetNumberOfFailedAuthQueriesOk() (*int32, bool)`

GetNumberOfFailedAuthQueriesOk returns a tuple with the NumberOfFailedAuthQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFailedAuthQueries

`func (o *IndexerStatistics) SetNumberOfFailedAuthQueries(v int32)`

SetNumberOfFailedAuthQueries sets NumberOfFailedAuthQueries field to given value.

### HasNumberOfFailedAuthQueries

`func (o *IndexerStatistics) HasNumberOfFailedAuthQueries() bool`

HasNumberOfFailedAuthQueries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


