# DevelopmentConfigResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ConsoleLogLevel** | Pointer to **NullableString** |  | [optional] 
**LogSql** | Pointer to **bool** |  | [optional] 
**LogIndexerResponse** | Pointer to **bool** |  | [optional] 
**LogRotate** | Pointer to **int32** |  | [optional] 
**FilterSentryEvents** | Pointer to **bool** |  | [optional] 

## Methods

### NewDevelopmentConfigResource

`func NewDevelopmentConfigResource() *DevelopmentConfigResource`

NewDevelopmentConfigResource instantiates a new DevelopmentConfigResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevelopmentConfigResourceWithDefaults

`func NewDevelopmentConfigResourceWithDefaults() *DevelopmentConfigResource`

NewDevelopmentConfigResourceWithDefaults instantiates a new DevelopmentConfigResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DevelopmentConfigResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DevelopmentConfigResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DevelopmentConfigResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DevelopmentConfigResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConsoleLogLevel

`func (o *DevelopmentConfigResource) GetConsoleLogLevel() string`

GetConsoleLogLevel returns the ConsoleLogLevel field if non-nil, zero value otherwise.

### GetConsoleLogLevelOk

`func (o *DevelopmentConfigResource) GetConsoleLogLevelOk() (*string, bool)`

GetConsoleLogLevelOk returns a tuple with the ConsoleLogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleLogLevel

`func (o *DevelopmentConfigResource) SetConsoleLogLevel(v string)`

SetConsoleLogLevel sets ConsoleLogLevel field to given value.

### HasConsoleLogLevel

`func (o *DevelopmentConfigResource) HasConsoleLogLevel() bool`

HasConsoleLogLevel returns a boolean if a field has been set.

### SetConsoleLogLevelNil

`func (o *DevelopmentConfigResource) SetConsoleLogLevelNil(b bool)`

 SetConsoleLogLevelNil sets the value for ConsoleLogLevel to be an explicit nil

### UnsetConsoleLogLevel
`func (o *DevelopmentConfigResource) UnsetConsoleLogLevel()`

UnsetConsoleLogLevel ensures that no value is present for ConsoleLogLevel, not even an explicit nil
### GetLogSql

`func (o *DevelopmentConfigResource) GetLogSql() bool`

GetLogSql returns the LogSql field if non-nil, zero value otherwise.

### GetLogSqlOk

`func (o *DevelopmentConfigResource) GetLogSqlOk() (*bool, bool)`

GetLogSqlOk returns a tuple with the LogSql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSql

`func (o *DevelopmentConfigResource) SetLogSql(v bool)`

SetLogSql sets LogSql field to given value.

### HasLogSql

`func (o *DevelopmentConfigResource) HasLogSql() bool`

HasLogSql returns a boolean if a field has been set.

### GetLogIndexerResponse

`func (o *DevelopmentConfigResource) GetLogIndexerResponse() bool`

GetLogIndexerResponse returns the LogIndexerResponse field if non-nil, zero value otherwise.

### GetLogIndexerResponseOk

`func (o *DevelopmentConfigResource) GetLogIndexerResponseOk() (*bool, bool)`

GetLogIndexerResponseOk returns a tuple with the LogIndexerResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogIndexerResponse

`func (o *DevelopmentConfigResource) SetLogIndexerResponse(v bool)`

SetLogIndexerResponse sets LogIndexerResponse field to given value.

### HasLogIndexerResponse

`func (o *DevelopmentConfigResource) HasLogIndexerResponse() bool`

HasLogIndexerResponse returns a boolean if a field has been set.

### GetLogRotate

`func (o *DevelopmentConfigResource) GetLogRotate() int32`

GetLogRotate returns the LogRotate field if non-nil, zero value otherwise.

### GetLogRotateOk

`func (o *DevelopmentConfigResource) GetLogRotateOk() (*int32, bool)`

GetLogRotateOk returns a tuple with the LogRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRotate

`func (o *DevelopmentConfigResource) SetLogRotate(v int32)`

SetLogRotate sets LogRotate field to given value.

### HasLogRotate

`func (o *DevelopmentConfigResource) HasLogRotate() bool`

HasLogRotate returns a boolean if a field has been set.

### GetFilterSentryEvents

`func (o *DevelopmentConfigResource) GetFilterSentryEvents() bool`

GetFilterSentryEvents returns the FilterSentryEvents field if non-nil, zero value otherwise.

### GetFilterSentryEventsOk

`func (o *DevelopmentConfigResource) GetFilterSentryEventsOk() (*bool, bool)`

GetFilterSentryEventsOk returns a tuple with the FilterSentryEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSentryEvents

`func (o *DevelopmentConfigResource) SetFilterSentryEvents(v bool)`

SetFilterSentryEvents sets FilterSentryEvents field to given value.

### HasFilterSentryEvents

`func (o *DevelopmentConfigResource) HasFilterSentryEvents() bool`

HasFilterSentryEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


