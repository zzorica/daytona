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

// checks if the CreateProjectDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProjectDTO{}

// CreateProjectDTO struct for CreateProjectDTO
type CreateProjectDTO struct {
	ExistingConfig *ExistingConfigDTO      `json:"existingConfig,omitempty"`
	NewConfig      *CreateProjectConfigDTO `json:"newConfig,omitempty"`
}

// NewCreateProjectDTO instantiates a new CreateProjectDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProjectDTO() *CreateProjectDTO {
	this := CreateProjectDTO{}
	return &this
}

// NewCreateProjectDTOWithDefaults instantiates a new CreateProjectDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectDTOWithDefaults() *CreateProjectDTO {
	this := CreateProjectDTO{}
	return &this
}

// GetExistingConfig returns the ExistingConfig field value if set, zero value otherwise.
func (o *CreateProjectDTO) GetExistingConfig() ExistingConfigDTO {
	if o == nil || IsNil(o.ExistingConfig) {
		var ret ExistingConfigDTO
		return ret
	}
	return *o.ExistingConfig
}

// GetExistingConfigOk returns a tuple with the ExistingConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectDTO) GetExistingConfigOk() (*ExistingConfigDTO, bool) {
	if o == nil || IsNil(o.ExistingConfig) {
		return nil, false
	}
	return o.ExistingConfig, true
}

// HasExistingConfig returns a boolean if a field has been set.
func (o *CreateProjectDTO) HasExistingConfig() bool {
	if o != nil && !IsNil(o.ExistingConfig) {
		return true
	}

	return false
}

// SetExistingConfig gets a reference to the given ExistingConfigDTO and assigns it to the ExistingConfig field.
func (o *CreateProjectDTO) SetExistingConfig(v ExistingConfigDTO) {
	o.ExistingConfig = &v
}

// GetNewConfig returns the NewConfig field value if set, zero value otherwise.
func (o *CreateProjectDTO) GetNewConfig() CreateProjectConfigDTO {
	if o == nil || IsNil(o.NewConfig) {
		var ret CreateProjectConfigDTO
		return ret
	}
	return *o.NewConfig
}

// GetNewConfigOk returns a tuple with the NewConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectDTO) GetNewConfigOk() (*CreateProjectConfigDTO, bool) {
	if o == nil || IsNil(o.NewConfig) {
		return nil, false
	}
	return o.NewConfig, true
}

// HasNewConfig returns a boolean if a field has been set.
func (o *CreateProjectDTO) HasNewConfig() bool {
	if o != nil && !IsNil(o.NewConfig) {
		return true
	}

	return false
}

// SetNewConfig gets a reference to the given CreateProjectConfigDTO and assigns it to the NewConfig field.
func (o *CreateProjectDTO) SetNewConfig(v CreateProjectConfigDTO) {
	o.NewConfig = &v
}

func (o CreateProjectDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProjectDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExistingConfig) {
		toSerialize["existingConfig"] = o.ExistingConfig
	}
	if !IsNil(o.NewConfig) {
		toSerialize["newConfig"] = o.NewConfig
	}
	return toSerialize, nil
}

type NullableCreateProjectDTO struct {
	value *CreateProjectDTO
	isSet bool
}

func (v NullableCreateProjectDTO) Get() *CreateProjectDTO {
	return v.value
}

func (v *NullableCreateProjectDTO) Set(val *CreateProjectDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProjectDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProjectDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProjectDTO(val *CreateProjectDTO) *NullableCreateProjectDTO {
	return &NullableCreateProjectDTO{value: val, isSet: true}
}

func (v NullableCreateProjectDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProjectDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
