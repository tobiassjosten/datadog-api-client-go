// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardTemplateVariablePreset Template variables saved views.
type DashboardTemplateVariablePreset struct {
	// The name of the variable.
	Name *string `json:"name,omitempty"`
	// List of variables.
	TemplateVariables []DashboardTemplateVariablePresetValue `json:"template_variables,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewDashboardTemplateVariablePreset instantiates a new DashboardTemplateVariablePreset object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardTemplateVariablePreset() *DashboardTemplateVariablePreset {
	this := DashboardTemplateVariablePreset{}
	return &this
}

// NewDashboardTemplateVariablePresetWithDefaults instantiates a new DashboardTemplateVariablePreset object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardTemplateVariablePresetWithDefaults() *DashboardTemplateVariablePreset {
	this := DashboardTemplateVariablePreset{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DashboardTemplateVariablePreset) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardTemplateVariablePreset) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DashboardTemplateVariablePreset) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DashboardTemplateVariablePreset) SetName(v string) {
	o.Name = &v
}

// GetTemplateVariables returns the TemplateVariables field value if set, zero value otherwise.
func (o *DashboardTemplateVariablePreset) GetTemplateVariables() []DashboardTemplateVariablePresetValue {
	if o == nil || o.TemplateVariables == nil {
		var ret []DashboardTemplateVariablePresetValue
		return ret
	}
	return o.TemplateVariables
}

// GetTemplateVariablesOk returns a tuple with the TemplateVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardTemplateVariablePreset) GetTemplateVariablesOk() (*[]DashboardTemplateVariablePresetValue, bool) {
	if o == nil || o.TemplateVariables == nil {
		return nil, false
	}
	return &o.TemplateVariables, true
}

// HasTemplateVariables returns a boolean if a field has been set.
func (o *DashboardTemplateVariablePreset) HasTemplateVariables() bool {
	return o != nil && o.TemplateVariables != nil
}

// SetTemplateVariables gets a reference to the given []DashboardTemplateVariablePresetValue and assigns it to the TemplateVariables field.
func (o *DashboardTemplateVariablePreset) SetTemplateVariables(v []DashboardTemplateVariablePresetValue) {
	o.TemplateVariables = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardTemplateVariablePreset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TemplateVariables != nil {
		toSerialize["template_variables"] = o.TemplateVariables
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardTemplateVariablePreset) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name              *string                                `json:"name,omitempty"`
		TemplateVariables []DashboardTemplateVariablePresetValue `json:"template_variables,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "template_variables"})
	} else {
		return err
	}
	o.Name = all.Name
	o.TemplateVariables = all.TemplateVariables

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
