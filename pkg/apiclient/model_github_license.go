/*
Daytona Server API

Daytona Server API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// checks if the GithubLicense type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GithubLicense{}

// GithubLicense struct for GithubLicense
type GithubLicense struct {
	Body           *string  `json:"body,omitempty"`
	Conditions     []string `json:"conditions,omitempty"`
	Description    *string  `json:"description,omitempty"`
	Featured       *bool    `json:"featured,omitempty"`
	HtmlUrl        *string  `json:"html_url,omitempty"`
	Implementation *string  `json:"implementation,omitempty"`
	Key            *string  `json:"key,omitempty"`
	Limitations    []string `json:"limitations,omitempty"`
	Name           *string  `json:"name,omitempty"`
	Permissions    []string `json:"permissions,omitempty"`
	SpdxId         *string  `json:"spdx_id,omitempty"`
	Url            *string  `json:"url,omitempty"`
}

// NewGithubLicense instantiates a new GithubLicense object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubLicense() *GithubLicense {
	this := GithubLicense{}
	return &this
}

// NewGithubLicenseWithDefaults instantiates a new GithubLicense object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubLicenseWithDefaults() *GithubLicense {
	this := GithubLicense{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *GithubLicense) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubLicense) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *GithubLicense) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *GithubLicense) SetBody(v string) {
	o.Body = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *GithubLicense) GetConditions() []string {
	if o == nil || IsNil(o.Conditions) {
		var ret []string
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubLicense) GetConditionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *GithubLicense) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []string and assigns it to the Conditions field.
func (o *GithubLicense) SetConditions(v []string) {
	o.Conditions = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GithubLicense) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubLicense) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GithubLicense) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GithubLicense) SetDescription(v string) {
	o.Description = &v
}

// GetFeatured returns the Featured field value if set, zero value otherwise.
func (o *GithubLicense) GetFeatured() bool {
	if o == nil || IsNil(o.Featured) {
		var ret bool
		return ret
	}
	return *o.Featured
}

// GetFeaturedOk returns a tuple with the Featured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubLicense) GetFeaturedOk() (*bool, bool) {
	if o == nil || IsNil(o.Featured) {
		return nil, false
	}
	return o.Featured, true
}

// HasFeatured returns a boolean if a field has been set.
func (o *GithubLicense) HasFeatured() bool {
	if o != nil && !IsNil(o.Featured) {
		return true
	}

	return false
}

// SetFeatured gets a reference to the given bool and assigns it to the Featured field.
func (o *GithubLicense) SetFeatured(v bool) {
	o.Featured = &v
}

// GetHtmlUrl returns the HtmlUrl field value if set, zero value otherwise.
func (o *GithubLicense) GetHtmlUrl() string {
	if o == nil || IsNil(o.HtmlUrl) {
		var ret string
		return ret
	}
	return *o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubLicense) GetHtmlUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HtmlUrl) {
		return nil, false
	}
	return o.HtmlUrl, true
}

// HasHtmlUrl returns a boolean if a field has been set.
func (o *GithubLicense) HasHtmlUrl() bool {
	if o != nil && !IsNil(o.HtmlUrl) {
		return true
	}

	return false
}

// SetHtmlUrl gets a reference to the given string and assigns it to the HtmlUrl field.
func (o *GithubLicense) SetHtmlUrl(v string) {
	o.HtmlUrl = &v
}

// GetImplementation returns the Implementation field value if set, zero value otherwise.
func (o *GithubLicense) GetImplementation() string {
	if o == nil || IsNil(o.Implementation) {
		var ret string
		return ret
	}
	return *o.Implementation
}

// GetImplementationOk returns a tuple with the Implementation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubLicense) GetImplementationOk() (*string, bool) {
	if o == nil || IsNil(o.Implementation) {
		return nil, false
	}
	return o.Implementation, true
}

// HasImplementation returns a boolean if a field has been set.
func (o *GithubLicense) HasImplementation() bool {
	if o != nil && !IsNil(o.Implementation) {
		return true
	}

	return false
}

// SetImplementation gets a reference to the given string and assigns it to the Implementation field.
func (o *GithubLicense) SetImplementation(v string) {
	o.Implementation = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *GithubLicense) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubLicense) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *GithubLicense) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *GithubLicense) SetKey(v string) {
	o.Key = &v
}

// GetLimitations returns the Limitations field value if set, zero value otherwise.
func (o *GithubLicense) GetLimitations() []string {
	if o == nil || IsNil(o.Limitations) {
		var ret []string
		return ret
	}
	return o.Limitations
}

// GetLimitationsOk returns a tuple with the Limitations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubLicense) GetLimitationsOk() ([]string, bool) {
	if o == nil || IsNil(o.Limitations) {
		return nil, false
	}
	return o.Limitations, true
}

// HasLimitations returns a boolean if a field has been set.
func (o *GithubLicense) HasLimitations() bool {
	if o != nil && !IsNil(o.Limitations) {
		return true
	}

	return false
}

// SetLimitations gets a reference to the given []string and assigns it to the Limitations field.
func (o *GithubLicense) SetLimitations(v []string) {
	o.Limitations = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GithubLicense) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubLicense) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GithubLicense) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GithubLicense) SetName(v string) {
	o.Name = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *GithubLicense) GetPermissions() []string {
	if o == nil || IsNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubLicense) GetPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *GithubLicense) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *GithubLicense) SetPermissions(v []string) {
	o.Permissions = v
}

// GetSpdxId returns the SpdxId field value if set, zero value otherwise.
func (o *GithubLicense) GetSpdxId() string {
	if o == nil || IsNil(o.SpdxId) {
		var ret string
		return ret
	}
	return *o.SpdxId
}

// GetSpdxIdOk returns a tuple with the SpdxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubLicense) GetSpdxIdOk() (*string, bool) {
	if o == nil || IsNil(o.SpdxId) {
		return nil, false
	}
	return o.SpdxId, true
}

// HasSpdxId returns a boolean if a field has been set.
func (o *GithubLicense) HasSpdxId() bool {
	if o != nil && !IsNil(o.SpdxId) {
		return true
	}

	return false
}

// SetSpdxId gets a reference to the given string and assigns it to the SpdxId field.
func (o *GithubLicense) SetSpdxId(v string) {
	o.SpdxId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GithubLicense) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubLicense) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GithubLicense) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GithubLicense) SetUrl(v string) {
	o.Url = &v
}

func (o GithubLicense) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GithubLicense) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Featured) {
		toSerialize["featured"] = o.Featured
	}
	if !IsNil(o.HtmlUrl) {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if !IsNil(o.Implementation) {
		toSerialize["implementation"] = o.Implementation
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Limitations) {
		toSerialize["limitations"] = o.Limitations
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.SpdxId) {
		toSerialize["spdx_id"] = o.SpdxId
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableGithubLicense struct {
	value *GithubLicense
	isSet bool
}

func (v NullableGithubLicense) Get() *GithubLicense {
	return v.value
}

func (v *NullableGithubLicense) Set(val *GithubLicense) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubLicense) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubLicense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubLicense(val *GithubLicense) *NullableGithubLicense {
	return &NullableGithubLicense{value: val, isSet: true}
}

func (v NullableGithubLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubLicense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}