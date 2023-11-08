// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeRelationshipsCreatedBy The user who created the downtime.
type DowntimeRelationshipsCreatedBy struct {
	// Data for the user who created the downtime.
	Data NullableDowntimeRelationshipsCreatedByData `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewDowntimeRelationshipsCreatedBy instantiates a new DowntimeRelationshipsCreatedBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeRelationshipsCreatedBy() *DowntimeRelationshipsCreatedBy {
	this := DowntimeRelationshipsCreatedBy{}
	return &this
}

// NewDowntimeRelationshipsCreatedByWithDefaults instantiates a new DowntimeRelationshipsCreatedBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeRelationshipsCreatedByWithDefaults() *DowntimeRelationshipsCreatedBy {
	this := DowntimeRelationshipsCreatedBy{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DowntimeRelationshipsCreatedBy) GetData() DowntimeRelationshipsCreatedByData {
	if o == nil || o.Data.Get() == nil {
		var ret DowntimeRelationshipsCreatedByData
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DowntimeRelationshipsCreatedBy) GetDataOk() (*DowntimeRelationshipsCreatedByData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *DowntimeRelationshipsCreatedBy) HasData() bool {
	return o != nil && o.Data.IsSet()
}

// SetData gets a reference to the given NullableDowntimeRelationshipsCreatedByData and assigns it to the Data field.
func (o *DowntimeRelationshipsCreatedBy) SetData(v DowntimeRelationshipsCreatedByData) {
	o.Data.Set(&v)
}

// SetDataNil sets the value for Data to be an explicit nil.
func (o *DowntimeRelationshipsCreatedBy) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil.
func (o *DowntimeRelationshipsCreatedBy) UnsetData() {
	o.Data.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeRelationshipsCreatedBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DowntimeRelationshipsCreatedBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data NullableDowntimeRelationshipsCreatedByData `json:"data,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
