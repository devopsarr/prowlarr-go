/*
Prowlarr

Prowlarr API docs

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
	"time"
)

// SystemResource struct for SystemResource
type SystemResource struct {
	AppName NullableString `json:"appName,omitempty"`
	InstanceName NullableString `json:"instanceName,omitempty"`
	Version NullableString `json:"version,omitempty"`
	BuildTime *time.Time `json:"buildTime,omitempty"`
	IsDebug *bool `json:"isDebug,omitempty"`
	IsProduction *bool `json:"isProduction,omitempty"`
	IsAdmin *bool `json:"isAdmin,omitempty"`
	IsUserInteractive *bool `json:"isUserInteractive,omitempty"`
	StartupPath NullableString `json:"startupPath,omitempty"`
	AppData NullableString `json:"appData,omitempty"`
	OsName NullableString `json:"osName,omitempty"`
	OsVersion NullableString `json:"osVersion,omitempty"`
	IsNetCore *bool `json:"isNetCore,omitempty"`
	IsLinux *bool `json:"isLinux,omitempty"`
	IsOsx *bool `json:"isOsx,omitempty"`
	IsWindows *bool `json:"isWindows,omitempty"`
	IsDocker *bool `json:"isDocker,omitempty"`
	Mode *RuntimeMode `json:"mode,omitempty"`
	Branch NullableString `json:"branch,omitempty"`
	DatabaseType *DatabaseType `json:"databaseType,omitempty"`
	DatabaseVersion *string `json:"databaseVersion,omitempty"`
	Authentication *AuthenticationType `json:"authentication,omitempty"`
	MigrationVersion *int32 `json:"migrationVersion,omitempty"`
	UrlBase NullableString `json:"urlBase,omitempty"`
	RuntimeVersion *string `json:"runtimeVersion,omitempty"`
	RuntimeName NullableString `json:"runtimeName,omitempty"`
	StartTime *time.Time `json:"startTime,omitempty"`
	PackageVersion NullableString `json:"packageVersion,omitempty"`
	PackageAuthor NullableString `json:"packageAuthor,omitempty"`
	PackageUpdateMechanism *UpdateMechanism `json:"packageUpdateMechanism,omitempty"`
	PackageUpdateMechanismMessage NullableString `json:"packageUpdateMechanismMessage,omitempty"`
}

// NewSystemResource instantiates a new SystemResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemResource() *SystemResource {
	this := SystemResource{}
	return &this
}

// NewSystemResourceWithDefaults instantiates a new SystemResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemResourceWithDefaults() *SystemResource {
	this := SystemResource{}
	return &this
}

// GetAppName returns the AppName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemResource) GetAppName() string {
	if o == nil || isNil(o.AppName.Get()) {
		var ret string
		return ret
	}
	return *o.AppName.Get()
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemResource) GetAppNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AppName.Get(), o.AppName.IsSet()
}

// HasAppName returns a boolean if a field has been set.
func (o *SystemResource) HasAppName() bool {
	if o != nil && o.AppName.IsSet() {
		return true
	}

	return false
}

// SetAppName gets a reference to the given NullableString and assigns it to the AppName field.
func (o *SystemResource) SetAppName(v string) {
	o.AppName.Set(&v)
}
// SetAppNameNil sets the value for AppName to be an explicit nil
func (o *SystemResource) SetAppNameNil() {
	o.AppName.Set(nil)
}

// UnsetAppName ensures that no value is present for AppName, not even an explicit nil
func (o *SystemResource) UnsetAppName() {
	o.AppName.Unset()
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemResource) GetInstanceName() string {
	if o == nil || isNil(o.InstanceName.Get()) {
		var ret string
		return ret
	}
	return *o.InstanceName.Get()
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemResource) GetInstanceNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.InstanceName.Get(), o.InstanceName.IsSet()
}

// HasInstanceName returns a boolean if a field has been set.
func (o *SystemResource) HasInstanceName() bool {
	if o != nil && o.InstanceName.IsSet() {
		return true
	}

	return false
}

// SetInstanceName gets a reference to the given NullableString and assigns it to the InstanceName field.
func (o *SystemResource) SetInstanceName(v string) {
	o.InstanceName.Set(&v)
}
// SetInstanceNameNil sets the value for InstanceName to be an explicit nil
func (o *SystemResource) SetInstanceNameNil() {
	o.InstanceName.Set(nil)
}

// UnsetInstanceName ensures that no value is present for InstanceName, not even an explicit nil
func (o *SystemResource) UnsetInstanceName() {
	o.InstanceName.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemResource) GetVersion() string {
	if o == nil || isNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemResource) GetVersionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *SystemResource) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *SystemResource) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *SystemResource) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *SystemResource) UnsetVersion() {
	o.Version.Unset()
}

// GetBuildTime returns the BuildTime field value if set, zero value otherwise.
func (o *SystemResource) GetBuildTime() time.Time {
	if o == nil || isNil(o.BuildTime) {
		var ret time.Time
		return ret
	}
	return *o.BuildTime
}

// GetBuildTimeOk returns a tuple with the BuildTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResource) GetBuildTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.BuildTime) {
    return nil, false
	}
	return o.BuildTime, true
}

// HasBuildTime returns a boolean if a field has been set.
func (o *SystemResource) HasBuildTime() bool {
	if o != nil && !isNil(o.BuildTime) {
		return true
	}

	return false
}

// SetBuildTime gets a reference to the given time.Time and assigns it to the BuildTime field.
func (o *SystemResource) SetBuildTime(v time.Time) {
	o.BuildTime = &v
}

// GetIsDebug returns the IsDebug field value if set, zero value otherwise.
func (o *SystemResource) GetIsDebug() bool {
	if o == nil || isNil(o.IsDebug) {
		var ret bool
		return ret
	}
	return *o.IsDebug
}

// GetIsDebugOk returns a tuple with the IsDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResource) GetIsDebugOk() (*bool, bool) {
	if o == nil || isNil(o.IsDebug) {
    return nil, false
	}
	return o.IsDebug, true
}

// HasIsDebug returns a boolean if a field has been set.
func (o *SystemResource) HasIsDebug() bool {
	if o != nil && !isNil(o.IsDebug) {
		return true
	}

	return false
}

// SetIsDebug gets a reference to the given bool and assigns it to the IsDebug field.
func (o *SystemResource) SetIsDebug(v bool) {
	o.IsDebug = &v
}

// GetIsProduction returns the IsProduction field value if set, zero value otherwise.
func (o *SystemResource) GetIsProduction() bool {
	if o == nil || isNil(o.IsProduction) {
		var ret bool
		return ret
	}
	return *o.IsProduction
}

// GetIsProductionOk returns a tuple with the IsProduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResource) GetIsProductionOk() (*bool, bool) {
	if o == nil || isNil(o.IsProduction) {
    return nil, false
	}
	return o.IsProduction, true
}

// HasIsProduction returns a boolean if a field has been set.
func (o *SystemResource) HasIsProduction() bool {
	if o != nil && !isNil(o.IsProduction) {
		return true
	}

	return false
}

// SetIsProduction gets a reference to the given bool and assigns it to the IsProduction field.
func (o *SystemResource) SetIsProduction(v bool) {
	o.IsProduction = &v
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise.
func (o *SystemResource) GetIsAdmin() bool {
	if o == nil || isNil(o.IsAdmin) {
		var ret bool
		return ret
	}
	return *o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResource) GetIsAdminOk() (*bool, bool) {
	if o == nil || isNil(o.IsAdmin) {
    return nil, false
	}
	return o.IsAdmin, true
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *SystemResource) HasIsAdmin() bool {
	if o != nil && !isNil(o.IsAdmin) {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.
func (o *SystemResource) SetIsAdmin(v bool) {
	o.IsAdmin = &v
}

// GetIsUserInteractive returns the IsUserInteractive field value if set, zero value otherwise.
func (o *SystemResource) GetIsUserInteractive() bool {
	if o == nil || isNil(o.IsUserInteractive) {
		var ret bool
		return ret
	}
	return *o.IsUserInteractive
}

// GetIsUserInteractiveOk returns a tuple with the IsUserInteractive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResource) GetIsUserInteractiveOk() (*bool, bool) {
	if o == nil || isNil(o.IsUserInteractive) {
    return nil, false
	}
	return o.IsUserInteractive, true
}

// HasIsUserInteractive returns a boolean if a field has been set.
func (o *SystemResource) HasIsUserInteractive() bool {
	if o != nil && !isNil(o.IsUserInteractive) {
		return true
	}

	return false
}

// SetIsUserInteractive gets a reference to the given bool and assigns it to the IsUserInteractive field.
func (o *SystemResource) SetIsUserInteractive(v bool) {
	o.IsUserInteractive = &v
}

// GetStartupPath returns the StartupPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemResource) GetStartupPath() string {
	if o == nil || isNil(o.StartupPath.Get()) {
		var ret string
		return ret
	}
	return *o.StartupPath.Get()
}

// GetStartupPathOk returns a tuple with the StartupPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemResource) GetStartupPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.StartupPath.Get(), o.StartupPath.IsSet()
}

// HasStartupPath returns a boolean if a field has been set.
func (o *SystemResource) HasStartupPath() bool {
	if o != nil && o.StartupPath.IsSet() {
		return true
	}

	return false
}

// SetStartupPath gets a reference to the given NullableString and assigns it to the StartupPath field.
func (o *SystemResource) SetStartupPath(v string) {
	o.StartupPath.Set(&v)
}
// SetStartupPathNil sets the value for StartupPath to be an explicit nil
func (o *SystemResource) SetStartupPathNil() {
	o.StartupPath.Set(nil)
}

// UnsetStartupPath ensures that no value is present for StartupPath, not even an explicit nil
func (o *SystemResource) UnsetStartupPath() {
	o.StartupPath.Unset()
}

// GetAppData returns the AppData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemResource) GetAppData() string {
	if o == nil || isNil(o.AppData.Get()) {
		var ret string
		return ret
	}
	return *o.AppData.Get()
}

// GetAppDataOk returns a tuple with the AppData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemResource) GetAppDataOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AppData.Get(), o.AppData.IsSet()
}

// HasAppData returns a boolean if a field has been set.
func (o *SystemResource) HasAppData() bool {
	if o != nil && o.AppData.IsSet() {
		return true
	}

	return false
}

// SetAppData gets a reference to the given NullableString and assigns it to the AppData field.
func (o *SystemResource) SetAppData(v string) {
	o.AppData.Set(&v)
}
// SetAppDataNil sets the value for AppData to be an explicit nil
func (o *SystemResource) SetAppDataNil() {
	o.AppData.Set(nil)
}

// UnsetAppData ensures that no value is present for AppData, not even an explicit nil
func (o *SystemResource) UnsetAppData() {
	o.AppData.Unset()
}

// GetOsName returns the OsName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemResource) GetOsName() string {
	if o == nil || isNil(o.OsName.Get()) {
		var ret string
		return ret
	}
	return *o.OsName.Get()
}

// GetOsNameOk returns a tuple with the OsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemResource) GetOsNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.OsName.Get(), o.OsName.IsSet()
}

// HasOsName returns a boolean if a field has been set.
func (o *SystemResource) HasOsName() bool {
	if o != nil && o.OsName.IsSet() {
		return true
	}

	return false
}

// SetOsName gets a reference to the given NullableString and assigns it to the OsName field.
func (o *SystemResource) SetOsName(v string) {
	o.OsName.Set(&v)
}
// SetOsNameNil sets the value for OsName to be an explicit nil
func (o *SystemResource) SetOsNameNil() {
	o.OsName.Set(nil)
}

// UnsetOsName ensures that no value is present for OsName, not even an explicit nil
func (o *SystemResource) UnsetOsName() {
	o.OsName.Unset()
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemResource) GetOsVersion() string {
	if o == nil || isNil(o.OsVersion.Get()) {
		var ret string
		return ret
	}
	return *o.OsVersion.Get()
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemResource) GetOsVersionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.OsVersion.Get(), o.OsVersion.IsSet()
}

// HasOsVersion returns a boolean if a field has been set.
func (o *SystemResource) HasOsVersion() bool {
	if o != nil && o.OsVersion.IsSet() {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given NullableString and assigns it to the OsVersion field.
func (o *SystemResource) SetOsVersion(v string) {
	o.OsVersion.Set(&v)
}
// SetOsVersionNil sets the value for OsVersion to be an explicit nil
func (o *SystemResource) SetOsVersionNil() {
	o.OsVersion.Set(nil)
}

// UnsetOsVersion ensures that no value is present for OsVersion, not even an explicit nil
func (o *SystemResource) UnsetOsVersion() {
	o.OsVersion.Unset()
}

// GetIsNetCore returns the IsNetCore field value if set, zero value otherwise.
func (o *SystemResource) GetIsNetCore() bool {
	if o == nil || isNil(o.IsNetCore) {
		var ret bool
		return ret
	}
	return *o.IsNetCore
}

// GetIsNetCoreOk returns a tuple with the IsNetCore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResource) GetIsNetCoreOk() (*bool, bool) {
	if o == nil || isNil(o.IsNetCore) {
    return nil, false
	}
	return o.IsNetCore, true
}

// HasIsNetCore returns a boolean if a field has been set.
func (o *SystemResource) HasIsNetCore() bool {
	if o != nil && !isNil(o.IsNetCore) {
		return true
	}

	return false
}

// SetIsNetCore gets a reference to the given bool and assigns it to the IsNetCore field.
func (o *SystemResource) SetIsNetCore(v bool) {
	o.IsNetCore = &v
}

// GetIsLinux returns the IsLinux field value if set, zero value otherwise.
func (o *SystemResource) GetIsLinux() bool {
	if o == nil || isNil(o.IsLinux) {
		var ret bool
		return ret
	}
	return *o.IsLinux
}

// GetIsLinuxOk returns a tuple with the IsLinux field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResource) GetIsLinuxOk() (*bool, bool) {
	if o == nil || isNil(o.IsLinux) {
    return nil, false
	}
	return o.IsLinux, true
}

// HasIsLinux returns a boolean if a field has been set.
func (o *SystemResource) HasIsLinux() bool {
	if o != nil && !isNil(o.IsLinux) {
		return true
	}

	return false
}

// SetIsLinux gets a reference to the given bool and assigns it to the IsLinux field.
func (o *SystemResource) SetIsLinux(v bool) {
	o.IsLinux = &v
}

// GetIsOsx returns the IsOsx field value if set, zero value otherwise.
func (o *SystemResource) GetIsOsx() bool {
	if o == nil || isNil(o.IsOsx) {
		var ret bool
		return ret
	}
	return *o.IsOsx
}

// GetIsOsxOk returns a tuple with the IsOsx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResource) GetIsOsxOk() (*bool, bool) {
	if o == nil || isNil(o.IsOsx) {
    return nil, false
	}
	return o.IsOsx, true
}

// HasIsOsx returns a boolean if a field has been set.
func (o *SystemResource) HasIsOsx() bool {
	if o != nil && !isNil(o.IsOsx) {
		return true
	}

	return false
}

// SetIsOsx gets a reference to the given bool and assigns it to the IsOsx field.
func (o *SystemResource) SetIsOsx(v bool) {
	o.IsOsx = &v
}

// GetIsWindows returns the IsWindows field value if set, zero value otherwise.
func (o *SystemResource) GetIsWindows() bool {
	if o == nil || isNil(o.IsWindows) {
		var ret bool
		return ret
	}
	return *o.IsWindows
}

// GetIsWindowsOk returns a tuple with the IsWindows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResource) GetIsWindowsOk() (*bool, bool) {
	if o == nil || isNil(o.IsWindows) {
    return nil, false
	}
	return o.IsWindows, true
}

// HasIsWindows returns a boolean if a field has been set.
func (o *SystemResource) HasIsWindows() bool {
	if o != nil && !isNil(o.IsWindows) {
		return true
	}

	return false
}

// SetIsWindows gets a reference to the given bool and assigns it to the IsWindows field.
func (o *SystemResource) SetIsWindows(v bool) {
	o.IsWindows = &v
}

// GetIsDocker returns the IsDocker field value if set, zero value otherwise.
func (o *SystemResource) GetIsDocker() bool {
	if o == nil || isNil(o.IsDocker) {
		var ret bool
		return ret
	}
	return *o.IsDocker
}

// GetIsDockerOk returns a tuple with the IsDocker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResource) GetIsDockerOk() (*bool, bool) {
	if o == nil || isNil(o.IsDocker) {
    return nil, false
	}
	return o.IsDocker, true
}

// HasIsDocker returns a boolean if a field has been set.
func (o *SystemResource) HasIsDocker() bool {
	if o != nil && !isNil(o.IsDocker) {
		return true
	}

	return false
}

// SetIsDocker gets a reference to the given bool and assigns it to the IsDocker field.
func (o *SystemResource) SetIsDocker(v bool) {
	o.IsDocker = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *SystemResource) GetMode() RuntimeMode {
	if o == nil || isNil(o.Mode) {
		var ret RuntimeMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResource) GetModeOk() (*RuntimeMode, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *SystemResource) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given RuntimeMode and assigns it to the Mode field.
func (o *SystemResource) SetMode(v RuntimeMode) {
	o.Mode = &v
}

// GetBranch returns the Branch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemResource) GetBranch() string {
	if o == nil || isNil(o.Branch.Get()) {
		var ret string
		return ret
	}
	return *o.Branch.Get()
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemResource) GetBranchOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Branch.Get(), o.Branch.IsSet()
}

// HasBranch returns a boolean if a field has been set.
func (o *SystemResource) HasBranch() bool {
	if o != nil && o.Branch.IsSet() {
		return true
	}

	return false
}

// SetBranch gets a reference to the given NullableString and assigns it to the Branch field.
func (o *SystemResource) SetBranch(v string) {
	o.Branch.Set(&v)
}
// SetBranchNil sets the value for Branch to be an explicit nil
func (o *SystemResource) SetBranchNil() {
	o.Branch.Set(nil)
}

// UnsetBranch ensures that no value is present for Branch, not even an explicit nil
func (o *SystemResource) UnsetBranch() {
	o.Branch.Unset()
}

// GetDatabaseType returns the DatabaseType field value if set, zero value otherwise.
func (o *SystemResource) GetDatabaseType() DatabaseType {
	if o == nil || isNil(o.DatabaseType) {
		var ret DatabaseType
		return ret
	}
	return *o.DatabaseType
}

// GetDatabaseTypeOk returns a tuple with the DatabaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResource) GetDatabaseTypeOk() (*DatabaseType, bool) {
	if o == nil || isNil(o.DatabaseType) {
    return nil, false
	}
	return o.DatabaseType, true
}

// HasDatabaseType returns a boolean if a field has been set.
func (o *SystemResource) HasDatabaseType() bool {
	if o != nil && !isNil(o.DatabaseType) {
		return true
	}

	return false
}

// SetDatabaseType gets a reference to the given DatabaseType and assigns it to the DatabaseType field.
func (o *SystemResource) SetDatabaseType(v DatabaseType) {
	o.DatabaseType = &v
}

// GetDatabaseVersion returns the DatabaseVersion field value if set, zero value otherwise.
func (o *SystemResource) GetDatabaseVersion() string {
	if o == nil || isNil(o.DatabaseVersion) {
		var ret string
		return ret
	}
	return *o.DatabaseVersion
}

// GetDatabaseVersionOk returns a tuple with the DatabaseVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResource) GetDatabaseVersionOk() (*string, bool) {
	if o == nil || isNil(o.DatabaseVersion) {
    return nil, false
	}
	return o.DatabaseVersion, true
}

// HasDatabaseVersion returns a boolean if a field has been set.
func (o *SystemResource) HasDatabaseVersion() bool {
	if o != nil && !isNil(o.DatabaseVersion) {
		return true
	}

	return false
}

// SetDatabaseVersion gets a reference to the given string and assigns it to the DatabaseVersion field.
func (o *SystemResource) SetDatabaseVersion(v string) {
	o.DatabaseVersion = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *SystemResource) GetAuthentication() AuthenticationType {
	if o == nil || isNil(o.Authentication) {
		var ret AuthenticationType
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResource) GetAuthenticationOk() (*AuthenticationType, bool) {
	if o == nil || isNil(o.Authentication) {
    return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *SystemResource) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given AuthenticationType and assigns it to the Authentication field.
func (o *SystemResource) SetAuthentication(v AuthenticationType) {
	o.Authentication = &v
}

// GetMigrationVersion returns the MigrationVersion field value if set, zero value otherwise.
func (o *SystemResource) GetMigrationVersion() int32 {
	if o == nil || isNil(o.MigrationVersion) {
		var ret int32
		return ret
	}
	return *o.MigrationVersion
}

// GetMigrationVersionOk returns a tuple with the MigrationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResource) GetMigrationVersionOk() (*int32, bool) {
	if o == nil || isNil(o.MigrationVersion) {
    return nil, false
	}
	return o.MigrationVersion, true
}

// HasMigrationVersion returns a boolean if a field has been set.
func (o *SystemResource) HasMigrationVersion() bool {
	if o != nil && !isNil(o.MigrationVersion) {
		return true
	}

	return false
}

// SetMigrationVersion gets a reference to the given int32 and assigns it to the MigrationVersion field.
func (o *SystemResource) SetMigrationVersion(v int32) {
	o.MigrationVersion = &v
}

// GetUrlBase returns the UrlBase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemResource) GetUrlBase() string {
	if o == nil || isNil(o.UrlBase.Get()) {
		var ret string
		return ret
	}
	return *o.UrlBase.Get()
}

// GetUrlBaseOk returns a tuple with the UrlBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemResource) GetUrlBaseOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.UrlBase.Get(), o.UrlBase.IsSet()
}

// HasUrlBase returns a boolean if a field has been set.
func (o *SystemResource) HasUrlBase() bool {
	if o != nil && o.UrlBase.IsSet() {
		return true
	}

	return false
}

// SetUrlBase gets a reference to the given NullableString and assigns it to the UrlBase field.
func (o *SystemResource) SetUrlBase(v string) {
	o.UrlBase.Set(&v)
}
// SetUrlBaseNil sets the value for UrlBase to be an explicit nil
func (o *SystemResource) SetUrlBaseNil() {
	o.UrlBase.Set(nil)
}

// UnsetUrlBase ensures that no value is present for UrlBase, not even an explicit nil
func (o *SystemResource) UnsetUrlBase() {
	o.UrlBase.Unset()
}

// GetRuntimeVersion returns the RuntimeVersion field value if set, zero value otherwise.
func (o *SystemResource) GetRuntimeVersion() string {
	if o == nil || isNil(o.RuntimeVersion) {
		var ret string
		return ret
	}
	return *o.RuntimeVersion
}

// GetRuntimeVersionOk returns a tuple with the RuntimeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResource) GetRuntimeVersionOk() (*string, bool) {
	if o == nil || isNil(o.RuntimeVersion) {
    return nil, false
	}
	return o.RuntimeVersion, true
}

// HasRuntimeVersion returns a boolean if a field has been set.
func (o *SystemResource) HasRuntimeVersion() bool {
	if o != nil && !isNil(o.RuntimeVersion) {
		return true
	}

	return false
}

// SetRuntimeVersion gets a reference to the given string and assigns it to the RuntimeVersion field.
func (o *SystemResource) SetRuntimeVersion(v string) {
	o.RuntimeVersion = &v
}

// GetRuntimeName returns the RuntimeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemResource) GetRuntimeName() string {
	if o == nil || isNil(o.RuntimeName.Get()) {
		var ret string
		return ret
	}
	return *o.RuntimeName.Get()
}

// GetRuntimeNameOk returns a tuple with the RuntimeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemResource) GetRuntimeNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RuntimeName.Get(), o.RuntimeName.IsSet()
}

// HasRuntimeName returns a boolean if a field has been set.
func (o *SystemResource) HasRuntimeName() bool {
	if o != nil && o.RuntimeName.IsSet() {
		return true
	}

	return false
}

// SetRuntimeName gets a reference to the given NullableString and assigns it to the RuntimeName field.
func (o *SystemResource) SetRuntimeName(v string) {
	o.RuntimeName.Set(&v)
}
// SetRuntimeNameNil sets the value for RuntimeName to be an explicit nil
func (o *SystemResource) SetRuntimeNameNil() {
	o.RuntimeName.Set(nil)
}

// UnsetRuntimeName ensures that no value is present for RuntimeName, not even an explicit nil
func (o *SystemResource) UnsetRuntimeName() {
	o.RuntimeName.Unset()
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *SystemResource) GetStartTime() time.Time {
	if o == nil || isNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResource) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTime) {
    return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *SystemResource) HasStartTime() bool {
	if o != nil && !isNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *SystemResource) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetPackageVersion returns the PackageVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemResource) GetPackageVersion() string {
	if o == nil || isNil(o.PackageVersion.Get()) {
		var ret string
		return ret
	}
	return *o.PackageVersion.Get()
}

// GetPackageVersionOk returns a tuple with the PackageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemResource) GetPackageVersionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.PackageVersion.Get(), o.PackageVersion.IsSet()
}

// HasPackageVersion returns a boolean if a field has been set.
func (o *SystemResource) HasPackageVersion() bool {
	if o != nil && o.PackageVersion.IsSet() {
		return true
	}

	return false
}

// SetPackageVersion gets a reference to the given NullableString and assigns it to the PackageVersion field.
func (o *SystemResource) SetPackageVersion(v string) {
	o.PackageVersion.Set(&v)
}
// SetPackageVersionNil sets the value for PackageVersion to be an explicit nil
func (o *SystemResource) SetPackageVersionNil() {
	o.PackageVersion.Set(nil)
}

// UnsetPackageVersion ensures that no value is present for PackageVersion, not even an explicit nil
func (o *SystemResource) UnsetPackageVersion() {
	o.PackageVersion.Unset()
}

// GetPackageAuthor returns the PackageAuthor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemResource) GetPackageAuthor() string {
	if o == nil || isNil(o.PackageAuthor.Get()) {
		var ret string
		return ret
	}
	return *o.PackageAuthor.Get()
}

// GetPackageAuthorOk returns a tuple with the PackageAuthor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemResource) GetPackageAuthorOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.PackageAuthor.Get(), o.PackageAuthor.IsSet()
}

// HasPackageAuthor returns a boolean if a field has been set.
func (o *SystemResource) HasPackageAuthor() bool {
	if o != nil && o.PackageAuthor.IsSet() {
		return true
	}

	return false
}

// SetPackageAuthor gets a reference to the given NullableString and assigns it to the PackageAuthor field.
func (o *SystemResource) SetPackageAuthor(v string) {
	o.PackageAuthor.Set(&v)
}
// SetPackageAuthorNil sets the value for PackageAuthor to be an explicit nil
func (o *SystemResource) SetPackageAuthorNil() {
	o.PackageAuthor.Set(nil)
}

// UnsetPackageAuthor ensures that no value is present for PackageAuthor, not even an explicit nil
func (o *SystemResource) UnsetPackageAuthor() {
	o.PackageAuthor.Unset()
}

// GetPackageUpdateMechanism returns the PackageUpdateMechanism field value if set, zero value otherwise.
func (o *SystemResource) GetPackageUpdateMechanism() UpdateMechanism {
	if o == nil || isNil(o.PackageUpdateMechanism) {
		var ret UpdateMechanism
		return ret
	}
	return *o.PackageUpdateMechanism
}

// GetPackageUpdateMechanismOk returns a tuple with the PackageUpdateMechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResource) GetPackageUpdateMechanismOk() (*UpdateMechanism, bool) {
	if o == nil || isNil(o.PackageUpdateMechanism) {
    return nil, false
	}
	return o.PackageUpdateMechanism, true
}

// HasPackageUpdateMechanism returns a boolean if a field has been set.
func (o *SystemResource) HasPackageUpdateMechanism() bool {
	if o != nil && !isNil(o.PackageUpdateMechanism) {
		return true
	}

	return false
}

// SetPackageUpdateMechanism gets a reference to the given UpdateMechanism and assigns it to the PackageUpdateMechanism field.
func (o *SystemResource) SetPackageUpdateMechanism(v UpdateMechanism) {
	o.PackageUpdateMechanism = &v
}

// GetPackageUpdateMechanismMessage returns the PackageUpdateMechanismMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemResource) GetPackageUpdateMechanismMessage() string {
	if o == nil || isNil(o.PackageUpdateMechanismMessage.Get()) {
		var ret string
		return ret
	}
	return *o.PackageUpdateMechanismMessage.Get()
}

// GetPackageUpdateMechanismMessageOk returns a tuple with the PackageUpdateMechanismMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemResource) GetPackageUpdateMechanismMessageOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.PackageUpdateMechanismMessage.Get(), o.PackageUpdateMechanismMessage.IsSet()
}

// HasPackageUpdateMechanismMessage returns a boolean if a field has been set.
func (o *SystemResource) HasPackageUpdateMechanismMessage() bool {
	if o != nil && o.PackageUpdateMechanismMessage.IsSet() {
		return true
	}

	return false
}

// SetPackageUpdateMechanismMessage gets a reference to the given NullableString and assigns it to the PackageUpdateMechanismMessage field.
func (o *SystemResource) SetPackageUpdateMechanismMessage(v string) {
	o.PackageUpdateMechanismMessage.Set(&v)
}
// SetPackageUpdateMechanismMessageNil sets the value for PackageUpdateMechanismMessage to be an explicit nil
func (o *SystemResource) SetPackageUpdateMechanismMessageNil() {
	o.PackageUpdateMechanismMessage.Set(nil)
}

// UnsetPackageUpdateMechanismMessage ensures that no value is present for PackageUpdateMechanismMessage, not even an explicit nil
func (o *SystemResource) UnsetPackageUpdateMechanismMessage() {
	o.PackageUpdateMechanismMessage.Unset()
}

func (o SystemResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppName.IsSet() {
		toSerialize["appName"] = o.AppName.Get()
	}
	if o.InstanceName.IsSet() {
		toSerialize["instanceName"] = o.InstanceName.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if !isNil(o.BuildTime) {
		toSerialize["buildTime"] = o.BuildTime
	}
	if !isNil(o.IsDebug) {
		toSerialize["isDebug"] = o.IsDebug
	}
	if !isNil(o.IsProduction) {
		toSerialize["isProduction"] = o.IsProduction
	}
	if !isNil(o.IsAdmin) {
		toSerialize["isAdmin"] = o.IsAdmin
	}
	if !isNil(o.IsUserInteractive) {
		toSerialize["isUserInteractive"] = o.IsUserInteractive
	}
	if o.StartupPath.IsSet() {
		toSerialize["startupPath"] = o.StartupPath.Get()
	}
	if o.AppData.IsSet() {
		toSerialize["appData"] = o.AppData.Get()
	}
	if o.OsName.IsSet() {
		toSerialize["osName"] = o.OsName.Get()
	}
	if o.OsVersion.IsSet() {
		toSerialize["osVersion"] = o.OsVersion.Get()
	}
	if !isNil(o.IsNetCore) {
		toSerialize["isNetCore"] = o.IsNetCore
	}
	if !isNil(o.IsLinux) {
		toSerialize["isLinux"] = o.IsLinux
	}
	if !isNil(o.IsOsx) {
		toSerialize["isOsx"] = o.IsOsx
	}
	if !isNil(o.IsWindows) {
		toSerialize["isWindows"] = o.IsWindows
	}
	if !isNil(o.IsDocker) {
		toSerialize["isDocker"] = o.IsDocker
	}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if o.Branch.IsSet() {
		toSerialize["branch"] = o.Branch.Get()
	}
	if !isNil(o.DatabaseType) {
		toSerialize["databaseType"] = o.DatabaseType
	}
	if !isNil(o.DatabaseVersion) {
		toSerialize["databaseVersion"] = o.DatabaseVersion
	}
	if !isNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	if !isNil(o.MigrationVersion) {
		toSerialize["migrationVersion"] = o.MigrationVersion
	}
	if o.UrlBase.IsSet() {
		toSerialize["urlBase"] = o.UrlBase.Get()
	}
	if !isNil(o.RuntimeVersion) {
		toSerialize["runtimeVersion"] = o.RuntimeVersion
	}
	if o.RuntimeName.IsSet() {
		toSerialize["runtimeName"] = o.RuntimeName.Get()
	}
	if !isNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if o.PackageVersion.IsSet() {
		toSerialize["packageVersion"] = o.PackageVersion.Get()
	}
	if o.PackageAuthor.IsSet() {
		toSerialize["packageAuthor"] = o.PackageAuthor.Get()
	}
	if !isNil(o.PackageUpdateMechanism) {
		toSerialize["packageUpdateMechanism"] = o.PackageUpdateMechanism
	}
	if o.PackageUpdateMechanismMessage.IsSet() {
		toSerialize["packageUpdateMechanismMessage"] = o.PackageUpdateMechanismMessage.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSystemResource struct {
	value *SystemResource
	isSet bool
}

func (v NullableSystemResource) Get() *SystemResource {
	return v.value
}

func (v *NullableSystemResource) Set(val *SystemResource) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemResource) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemResource(val *SystemResource) *NullableSystemResource {
	return &NullableSystemResource{value: val, isSet: true}
}

func (v NullableSystemResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

