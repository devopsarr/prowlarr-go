# Go API client for prowlarr

Prowlarr API docs

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1.27.0.4852
- Package version: 1.1.1 <!--- x-release-please-version -->
- Generator version: 7.9.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import prowlarr "github.com/devopsarr/prowlarr-go/prowlarr"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `prowlarr.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), prowlarr.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `prowlarr.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), prowlarr.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `prowlarr.ContextOperationServerIndices` and `prowlarr.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), prowlarr.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), prowlarr.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:9696*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApiInfoAPI* | [**GetApi**](docs/ApiInfoAPI.md#getapi) | **Get** /api | 
*AppProfileAPI* | [**CreateAppProfile**](docs/AppProfileAPI.md#createappprofile) | **Post** /api/v1/appprofile | 
*AppProfileAPI* | [**DeleteAppProfile**](docs/AppProfileAPI.md#deleteappprofile) | **Delete** /api/v1/appprofile/{id} | 
*AppProfileAPI* | [**GetAppProfileById**](docs/AppProfileAPI.md#getappprofilebyid) | **Get** /api/v1/appprofile/{id} | 
*AppProfileAPI* | [**GetAppProfileSchema**](docs/AppProfileAPI.md#getappprofileschema) | **Get** /api/v1/appprofile/schema | 
*AppProfileAPI* | [**ListAppProfile**](docs/AppProfileAPI.md#listappprofile) | **Get** /api/v1/appprofile | 
*AppProfileAPI* | [**UpdateAppProfile**](docs/AppProfileAPI.md#updateappprofile) | **Put** /api/v1/appprofile/{id} | 
*ApplicationAPI* | [**CreateApplications**](docs/ApplicationAPI.md#createapplications) | **Post** /api/v1/applications | 
*ApplicationAPI* | [**CreateApplicationsActionByName**](docs/ApplicationAPI.md#createapplicationsactionbyname) | **Post** /api/v1/applications/action/{name} | 
*ApplicationAPI* | [**DeleteApplications**](docs/ApplicationAPI.md#deleteapplications) | **Delete** /api/v1/applications/{id} | 
*ApplicationAPI* | [**DeleteApplicationsBulk**](docs/ApplicationAPI.md#deleteapplicationsbulk) | **Delete** /api/v1/applications/bulk | 
*ApplicationAPI* | [**GetApplicationsById**](docs/ApplicationAPI.md#getapplicationsbyid) | **Get** /api/v1/applications/{id} | 
*ApplicationAPI* | [**ListApplications**](docs/ApplicationAPI.md#listapplications) | **Get** /api/v1/applications | 
*ApplicationAPI* | [**ListApplicationsSchema**](docs/ApplicationAPI.md#listapplicationsschema) | **Get** /api/v1/applications/schema | 
*ApplicationAPI* | [**PutApplicationsBulk**](docs/ApplicationAPI.md#putapplicationsbulk) | **Put** /api/v1/applications/bulk | 
*ApplicationAPI* | [**TestApplications**](docs/ApplicationAPI.md#testapplications) | **Post** /api/v1/applications/test | 
*ApplicationAPI* | [**TestallApplications**](docs/ApplicationAPI.md#testallapplications) | **Post** /api/v1/applications/testall | 
*ApplicationAPI* | [**UpdateApplications**](docs/ApplicationAPI.md#updateapplications) | **Put** /api/v1/applications/{id} | 
*AuthenticationAPI* | [**CreateLogin**](docs/AuthenticationAPI.md#createlogin) | **Post** /login | 
*AuthenticationAPI* | [**GetLogout**](docs/AuthenticationAPI.md#getlogout) | **Get** /logout | 
*BackupAPI* | [**CreateSystemBackupRestoreById**](docs/BackupAPI.md#createsystembackuprestorebyid) | **Post** /api/v1/system/backup/restore/{id} | 
*BackupAPI* | [**CreateSystemBackupRestoreUpload**](docs/BackupAPI.md#createsystembackuprestoreupload) | **Post** /api/v1/system/backup/restore/upload | 
*BackupAPI* | [**DeleteSystemBackup**](docs/BackupAPI.md#deletesystembackup) | **Delete** /api/v1/system/backup/{id} | 
*BackupAPI* | [**ListSystemBackup**](docs/BackupAPI.md#listsystembackup) | **Get** /api/v1/system/backup | 
*CommandAPI* | [**CreateCommand**](docs/CommandAPI.md#createcommand) | **Post** /api/v1/command | 
*CommandAPI* | [**DeleteCommand**](docs/CommandAPI.md#deletecommand) | **Delete** /api/v1/command/{id} | 
*CommandAPI* | [**GetCommandById**](docs/CommandAPI.md#getcommandbyid) | **Get** /api/v1/command/{id} | 
*CommandAPI* | [**ListCommand**](docs/CommandAPI.md#listcommand) | **Get** /api/v1/command | 
*CustomFilterAPI* | [**CreateCustomFilter**](docs/CustomFilterAPI.md#createcustomfilter) | **Post** /api/v1/customfilter | 
*CustomFilterAPI* | [**DeleteCustomFilter**](docs/CustomFilterAPI.md#deletecustomfilter) | **Delete** /api/v1/customfilter/{id} | 
*CustomFilterAPI* | [**GetCustomFilterById**](docs/CustomFilterAPI.md#getcustomfilterbyid) | **Get** /api/v1/customfilter/{id} | 
*CustomFilterAPI* | [**ListCustomFilter**](docs/CustomFilterAPI.md#listcustomfilter) | **Get** /api/v1/customfilter | 
*CustomFilterAPI* | [**UpdateCustomFilter**](docs/CustomFilterAPI.md#updatecustomfilter) | **Put** /api/v1/customfilter/{id} | 
*DevelopmentConfigAPI* | [**GetDevelopmentConfig**](docs/DevelopmentConfigAPI.md#getdevelopmentconfig) | **Get** /api/v1/config/development | 
*DevelopmentConfigAPI* | [**GetDevelopmentConfigById**](docs/DevelopmentConfigAPI.md#getdevelopmentconfigbyid) | **Get** /api/v1/config/development/{id} | 
*DevelopmentConfigAPI* | [**UpdateDevelopmentConfig**](docs/DevelopmentConfigAPI.md#updatedevelopmentconfig) | **Put** /api/v1/config/development/{id} | 
*DownloadClientAPI* | [**CreateDownloadClient**](docs/DownloadClientAPI.md#createdownloadclient) | **Post** /api/v1/downloadclient | 
*DownloadClientAPI* | [**CreateDownloadClientActionByName**](docs/DownloadClientAPI.md#createdownloadclientactionbyname) | **Post** /api/v1/downloadclient/action/{name} | 
*DownloadClientAPI* | [**DeleteDownloadClient**](docs/DownloadClientAPI.md#deletedownloadclient) | **Delete** /api/v1/downloadclient/{id} | 
*DownloadClientAPI* | [**DeleteDownloadClientBulk**](docs/DownloadClientAPI.md#deletedownloadclientbulk) | **Delete** /api/v1/downloadclient/bulk | 
*DownloadClientAPI* | [**GetDownloadClientById**](docs/DownloadClientAPI.md#getdownloadclientbyid) | **Get** /api/v1/downloadclient/{id} | 
*DownloadClientAPI* | [**ListDownloadClient**](docs/DownloadClientAPI.md#listdownloadclient) | **Get** /api/v1/downloadclient | 
*DownloadClientAPI* | [**ListDownloadClientSchema**](docs/DownloadClientAPI.md#listdownloadclientschema) | **Get** /api/v1/downloadclient/schema | 
*DownloadClientAPI* | [**PutDownloadClientBulk**](docs/DownloadClientAPI.md#putdownloadclientbulk) | **Put** /api/v1/downloadclient/bulk | 
*DownloadClientAPI* | [**TestDownloadClient**](docs/DownloadClientAPI.md#testdownloadclient) | **Post** /api/v1/downloadclient/test | 
*DownloadClientAPI* | [**TestallDownloadClient**](docs/DownloadClientAPI.md#testalldownloadclient) | **Post** /api/v1/downloadclient/testall | 
*DownloadClientAPI* | [**UpdateDownloadClient**](docs/DownloadClientAPI.md#updatedownloadclient) | **Put** /api/v1/downloadclient/{id} | 
*DownloadClientConfigAPI* | [**GetDownloadClientConfig**](docs/DownloadClientConfigAPI.md#getdownloadclientconfig) | **Get** /api/v1/config/downloadclient | 
*DownloadClientConfigAPI* | [**GetDownloadClientConfigById**](docs/DownloadClientConfigAPI.md#getdownloadclientconfigbyid) | **Get** /api/v1/config/downloadclient/{id} | 
*DownloadClientConfigAPI* | [**UpdateDownloadClientConfig**](docs/DownloadClientConfigAPI.md#updatedownloadclientconfig) | **Put** /api/v1/config/downloadclient/{id} | 
*FileSystemAPI* | [**GetFileSystem**](docs/FileSystemAPI.md#getfilesystem) | **Get** /api/v1/filesystem | 
*FileSystemAPI* | [**GetFileSystemType**](docs/FileSystemAPI.md#getfilesystemtype) | **Get** /api/v1/filesystem/type | 
*HealthAPI* | [**ListHealth**](docs/HealthAPI.md#listhealth) | **Get** /api/v1/health | 
*HistoryAPI* | [**GetHistory**](docs/HistoryAPI.md#gethistory) | **Get** /api/v1/history | 
*HistoryAPI* | [**ListHistoryIndexer**](docs/HistoryAPI.md#listhistoryindexer) | **Get** /api/v1/history/indexer | 
*HistoryAPI* | [**ListHistorySince**](docs/HistoryAPI.md#listhistorysince) | **Get** /api/v1/history/since | 
*HostConfigAPI* | [**GetHostConfig**](docs/HostConfigAPI.md#gethostconfig) | **Get** /api/v1/config/host | 
*HostConfigAPI* | [**GetHostConfigById**](docs/HostConfigAPI.md#gethostconfigbyid) | **Get** /api/v1/config/host/{id} | 
*HostConfigAPI* | [**UpdateHostConfig**](docs/HostConfigAPI.md#updatehostconfig) | **Put** /api/v1/config/host/{id} | 
*IndexerAPI* | [**CreateIndexer**](docs/IndexerAPI.md#createindexer) | **Post** /api/v1/indexer | 
*IndexerAPI* | [**CreateIndexerActionByName**](docs/IndexerAPI.md#createindexeractionbyname) | **Post** /api/v1/indexer/action/{name} | 
*IndexerAPI* | [**DeleteIndexer**](docs/IndexerAPI.md#deleteindexer) | **Delete** /api/v1/indexer/{id} | 
*IndexerAPI* | [**DeleteIndexerBulk**](docs/IndexerAPI.md#deleteindexerbulk) | **Delete** /api/v1/indexer/bulk | 
*IndexerAPI* | [**GetIndexerById**](docs/IndexerAPI.md#getindexerbyid) | **Get** /api/v1/indexer/{id} | 
*IndexerAPI* | [**ListIndexer**](docs/IndexerAPI.md#listindexer) | **Get** /api/v1/indexer | 
*IndexerAPI* | [**ListIndexerSchema**](docs/IndexerAPI.md#listindexerschema) | **Get** /api/v1/indexer/schema | 
*IndexerAPI* | [**PutIndexerBulk**](docs/IndexerAPI.md#putindexerbulk) | **Put** /api/v1/indexer/bulk | 
*IndexerAPI* | [**TestIndexer**](docs/IndexerAPI.md#testindexer) | **Post** /api/v1/indexer/test | 
*IndexerAPI* | [**TestallIndexer**](docs/IndexerAPI.md#testallindexer) | **Post** /api/v1/indexer/testall | 
*IndexerAPI* | [**UpdateIndexer**](docs/IndexerAPI.md#updateindexer) | **Put** /api/v1/indexer/{id} | 
*IndexerDefaultCategoriesAPI* | [**ListIndexerCategories**](docs/IndexerDefaultCategoriesAPI.md#listindexercategories) | **Get** /api/v1/indexer/categories | 
*IndexerProxyAPI* | [**CreateIndexerProxy**](docs/IndexerProxyAPI.md#createindexerproxy) | **Post** /api/v1/indexerproxy | 
*IndexerProxyAPI* | [**CreateIndexerProxyActionByName**](docs/IndexerProxyAPI.md#createindexerproxyactionbyname) | **Post** /api/v1/indexerproxy/action/{name} | 
*IndexerProxyAPI* | [**DeleteIndexerProxy**](docs/IndexerProxyAPI.md#deleteindexerproxy) | **Delete** /api/v1/indexerproxy/{id} | 
*IndexerProxyAPI* | [**GetIndexerProxyById**](docs/IndexerProxyAPI.md#getindexerproxybyid) | **Get** /api/v1/indexerproxy/{id} | 
*IndexerProxyAPI* | [**ListIndexerProxy**](docs/IndexerProxyAPI.md#listindexerproxy) | **Get** /api/v1/indexerproxy | 
*IndexerProxyAPI* | [**ListIndexerProxySchema**](docs/IndexerProxyAPI.md#listindexerproxyschema) | **Get** /api/v1/indexerproxy/schema | 
*IndexerProxyAPI* | [**TestIndexerProxy**](docs/IndexerProxyAPI.md#testindexerproxy) | **Post** /api/v1/indexerproxy/test | 
*IndexerProxyAPI* | [**TestallIndexerProxy**](docs/IndexerProxyAPI.md#testallindexerproxy) | **Post** /api/v1/indexerproxy/testall | 
*IndexerProxyAPI* | [**UpdateIndexerProxy**](docs/IndexerProxyAPI.md#updateindexerproxy) | **Put** /api/v1/indexerproxy/{id} | 
*IndexerStatsAPI* | [**GetIndexerStats**](docs/IndexerStatsAPI.md#getindexerstats) | **Get** /api/v1/indexerstats | 
*IndexerStatusAPI* | [**ListIndexerStatus**](docs/IndexerStatusAPI.md#listindexerstatus) | **Get** /api/v1/indexerstatus | 
*LocalizationAPI* | [**GetLocalization**](docs/LocalizationAPI.md#getlocalization) | **Get** /api/v1/localization | 
*LocalizationAPI* | [**ListLocalizationOptions**](docs/LocalizationAPI.md#listlocalizationoptions) | **Get** /api/v1/localization/options | 
*LogAPI* | [**GetLog**](docs/LogAPI.md#getlog) | **Get** /api/v1/log | 
*LogFileAPI* | [**GetLogFileByFilename**](docs/LogFileAPI.md#getlogfilebyfilename) | **Get** /api/v1/log/file/{filename} | 
*LogFileAPI* | [**ListLogFile**](docs/LogFileAPI.md#listlogfile) | **Get** /api/v1/log/file | 
*NewznabAPI* | [**GetIndexerDownload**](docs/NewznabAPI.md#getindexerdownload) | **Get** /api/v1/indexer/{id}/download | 
*NewznabAPI* | [**GetIndexerNewznab**](docs/NewznabAPI.md#getindexernewznab) | **Get** /api/v1/indexer/{id}/newznab | 
*NotificationAPI* | [**CreateNotification**](docs/NotificationAPI.md#createnotification) | **Post** /api/v1/notification | 
*NotificationAPI* | [**CreateNotificationActionByName**](docs/NotificationAPI.md#createnotificationactionbyname) | **Post** /api/v1/notification/action/{name} | 
*NotificationAPI* | [**DeleteNotification**](docs/NotificationAPI.md#deletenotification) | **Delete** /api/v1/notification/{id} | 
*NotificationAPI* | [**GetNotificationById**](docs/NotificationAPI.md#getnotificationbyid) | **Get** /api/v1/notification/{id} | 
*NotificationAPI* | [**ListNotification**](docs/NotificationAPI.md#listnotification) | **Get** /api/v1/notification | 
*NotificationAPI* | [**ListNotificationSchema**](docs/NotificationAPI.md#listnotificationschema) | **Get** /api/v1/notification/schema | 
*NotificationAPI* | [**TestNotification**](docs/NotificationAPI.md#testnotification) | **Post** /api/v1/notification/test | 
*NotificationAPI* | [**TestallNotification**](docs/NotificationAPI.md#testallnotification) | **Post** /api/v1/notification/testall | 
*NotificationAPI* | [**UpdateNotification**](docs/NotificationAPI.md#updatenotification) | **Put** /api/v1/notification/{id} | 
*PingAPI* | [**GetPing**](docs/PingAPI.md#getping) | **Get** /ping | 
*PingAPI* | [**HeadPing**](docs/PingAPI.md#headping) | **Head** /ping | 
*SearchAPI* | [**CreateSearch**](docs/SearchAPI.md#createsearch) | **Post** /api/v1/search | 
*SearchAPI* | [**CreateSearchBulk**](docs/SearchAPI.md#createsearchbulk) | **Post** /api/v1/search/bulk | 
*SearchAPI* | [**ListSearch**](docs/SearchAPI.md#listsearch) | **Get** /api/v1/search | 
*StaticResourceAPI* | [**Get**](docs/StaticResourceAPI.md#get) | **Get** / | 
*StaticResourceAPI* | [**GetByPath**](docs/StaticResourceAPI.md#getbypath) | **Get** /{path} | 
*StaticResourceAPI* | [**GetContentByPath**](docs/StaticResourceAPI.md#getcontentbypath) | **Get** /content/{path} | 
*StaticResourceAPI* | [**GetLogin**](docs/StaticResourceAPI.md#getlogin) | **Get** /login | 
*SystemAPI* | [**CreateSystemRestart**](docs/SystemAPI.md#createsystemrestart) | **Post** /api/v1/system/restart | 
*SystemAPI* | [**CreateSystemShutdown**](docs/SystemAPI.md#createsystemshutdown) | **Post** /api/v1/system/shutdown | 
*SystemAPI* | [**GetSystemRoutes**](docs/SystemAPI.md#getsystemroutes) | **Get** /api/v1/system/routes | 
*SystemAPI* | [**GetSystemRoutesDuplicate**](docs/SystemAPI.md#getsystemroutesduplicate) | **Get** /api/v1/system/routes/duplicate | 
*SystemAPI* | [**GetSystemStatus**](docs/SystemAPI.md#getsystemstatus) | **Get** /api/v1/system/status | 
*TagAPI* | [**CreateTag**](docs/TagAPI.md#createtag) | **Post** /api/v1/tag | 
*TagAPI* | [**DeleteTag**](docs/TagAPI.md#deletetag) | **Delete** /api/v1/tag/{id} | 
*TagAPI* | [**GetTagById**](docs/TagAPI.md#gettagbyid) | **Get** /api/v1/tag/{id} | 
*TagAPI* | [**ListTag**](docs/TagAPI.md#listtag) | **Get** /api/v1/tag | 
*TagAPI* | [**UpdateTag**](docs/TagAPI.md#updatetag) | **Put** /api/v1/tag/{id} | 
*TagDetailsAPI* | [**GetTagDetailById**](docs/TagDetailsAPI.md#gettagdetailbyid) | **Get** /api/v1/tag/detail/{id} | 
*TagDetailsAPI* | [**ListTagDetail**](docs/TagDetailsAPI.md#listtagdetail) | **Get** /api/v1/tag/detail | 
*TaskAPI* | [**GetSystemTaskById**](docs/TaskAPI.md#getsystemtaskbyid) | **Get** /api/v1/system/task/{id} | 
*TaskAPI* | [**ListSystemTask**](docs/TaskAPI.md#listsystemtask) | **Get** /api/v1/system/task | 
*UiConfigAPI* | [**GetUiConfig**](docs/UiConfigAPI.md#getuiconfig) | **Get** /api/v1/config/ui | 
*UiConfigAPI* | [**GetUiConfigById**](docs/UiConfigAPI.md#getuiconfigbyid) | **Get** /api/v1/config/ui/{id} | 
*UiConfigAPI* | [**UpdateUiConfig**](docs/UiConfigAPI.md#updateuiconfig) | **Put** /api/v1/config/ui/{id} | 
*UpdateAPI* | [**ListUpdate**](docs/UpdateAPI.md#listupdate) | **Get** /api/v1/update | 
*UpdateLogFileAPI* | [**GetLogFileUpdateByFilename**](docs/UpdateLogFileAPI.md#getlogfileupdatebyfilename) | **Get** /api/v1/log/file/update/{filename} | 
*UpdateLogFileAPI* | [**ListLogFileUpdate**](docs/UpdateLogFileAPI.md#listlogfileupdate) | **Get** /api/v1/log/file/update | 


## Documentation For Models

 - [ApiInfoResource](docs/ApiInfoResource.md)
 - [AppProfileResource](docs/AppProfileResource.md)
 - [ApplicationBulkResource](docs/ApplicationBulkResource.md)
 - [ApplicationResource](docs/ApplicationResource.md)
 - [ApplicationSyncLevel](docs/ApplicationSyncLevel.md)
 - [ApplyTags](docs/ApplyTags.md)
 - [AuthenticationRequiredType](docs/AuthenticationRequiredType.md)
 - [AuthenticationType](docs/AuthenticationType.md)
 - [BackupResource](docs/BackupResource.md)
 - [BackupType](docs/BackupType.md)
 - [BookSearchParam](docs/BookSearchParam.md)
 - [CertificateValidationType](docs/CertificateValidationType.md)
 - [Command](docs/Command.md)
 - [CommandPriority](docs/CommandPriority.md)
 - [CommandResource](docs/CommandResource.md)
 - [CommandStatus](docs/CommandStatus.md)
 - [CommandTrigger](docs/CommandTrigger.md)
 - [CustomFilterResource](docs/CustomFilterResource.md)
 - [DatabaseType](docs/DatabaseType.md)
 - [DevelopmentConfigResource](docs/DevelopmentConfigResource.md)
 - [DownloadClientBulkResource](docs/DownloadClientBulkResource.md)
 - [DownloadClientCategory](docs/DownloadClientCategory.md)
 - [DownloadClientConfigResource](docs/DownloadClientConfigResource.md)
 - [DownloadClientResource](docs/DownloadClientResource.md)
 - [DownloadProtocol](docs/DownloadProtocol.md)
 - [Field](docs/Field.md)
 - [HealthCheckResult](docs/HealthCheckResult.md)
 - [HealthResource](docs/HealthResource.md)
 - [HistoryEventType](docs/HistoryEventType.md)
 - [HistoryResource](docs/HistoryResource.md)
 - [HistoryResourcePagingResource](docs/HistoryResourcePagingResource.md)
 - [HostConfigResource](docs/HostConfigResource.md)
 - [HostStatistics](docs/HostStatistics.md)
 - [IndexerBulkResource](docs/IndexerBulkResource.md)
 - [IndexerCapabilityResource](docs/IndexerCapabilityResource.md)
 - [IndexerCategory](docs/IndexerCategory.md)
 - [IndexerPrivacy](docs/IndexerPrivacy.md)
 - [IndexerProxyResource](docs/IndexerProxyResource.md)
 - [IndexerResource](docs/IndexerResource.md)
 - [IndexerStatistics](docs/IndexerStatistics.md)
 - [IndexerStatsResource](docs/IndexerStatsResource.md)
 - [IndexerStatusResource](docs/IndexerStatusResource.md)
 - [LocalizationOption](docs/LocalizationOption.md)
 - [LogFileResource](docs/LogFileResource.md)
 - [LogResource](docs/LogResource.md)
 - [LogResourcePagingResource](docs/LogResourcePagingResource.md)
 - [MovieSearchParam](docs/MovieSearchParam.md)
 - [MusicSearchParam](docs/MusicSearchParam.md)
 - [NotificationResource](docs/NotificationResource.md)
 - [PingResource](docs/PingResource.md)
 - [PrivacyLevel](docs/PrivacyLevel.md)
 - [ProviderMessage](docs/ProviderMessage.md)
 - [ProviderMessageType](docs/ProviderMessageType.md)
 - [ProxyType](docs/ProxyType.md)
 - [ReleaseResource](docs/ReleaseResource.md)
 - [RuntimeMode](docs/RuntimeMode.md)
 - [SearchParam](docs/SearchParam.md)
 - [SelectOption](docs/SelectOption.md)
 - [SortDirection](docs/SortDirection.md)
 - [SystemResource](docs/SystemResource.md)
 - [TagDetailsResource](docs/TagDetailsResource.md)
 - [TagResource](docs/TagResource.md)
 - [TaskResource](docs/TaskResource.md)
 - [TvSearchParam](docs/TvSearchParam.md)
 - [UiConfigResource](docs/UiConfigResource.md)
 - [UpdateChanges](docs/UpdateChanges.md)
 - [UpdateMechanism](docs/UpdateMechanism.md)
 - [UpdateResource](docs/UpdateResource.md)
 - [UserAgentStatistics](docs/UserAgentStatistics.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### X-Api-Key

- **Type**: API key
- **API key parameter name**: X-Api-Key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-Api-Key and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		prowlarr.ContextAPIKeys,
		map[string]prowlarr.APIKey{
			"X-Api-Key": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### apikey

- **Type**: API key
- **API key parameter name**: apikey
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: apikey and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		prowlarr.ContextAPIKeys,
		map[string]prowlarr.APIKey{
			"apikey": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



