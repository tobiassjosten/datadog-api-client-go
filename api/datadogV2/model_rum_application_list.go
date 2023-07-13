// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMApplicationList RUM application list.
type RUMApplicationList struct {
	// RUM application list attributes.
	Attributes RUMApplicationListAttributes `json:"attributes"`
	// RUM application ID.
	Id *string `json:"id,omitempty"`
	// RUM application list type.
	Type RUMApplicationListType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewRUMApplicationList instantiates a new RUMApplicationList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMApplicationList(attributes RUMApplicationListAttributes, typeVar RUMApplicationListType) *RUMApplicationList {
	this := RUMApplicationList{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewRUMApplicationListWithDefaults instantiates a new RUMApplicationList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMApplicationListWithDefaults() *RUMApplicationList {
	this := RUMApplicationList{}
	var typeVar RUMApplicationListType = RUMAPPLICATIONLISTTYPE_RUM_APPLICATION
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *RUMApplicationList) GetAttributes() RUMApplicationListAttributes {
	if o == nil {
		var ret RUMApplicationListAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *RUMApplicationList) GetAttributesOk() (*RUMApplicationListAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *RUMApplicationList) SetAttributes(v RUMApplicationListAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RUMApplicationList) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationList) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RUMApplicationList) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RUMApplicationList) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *RUMApplicationList) GetType() RUMApplicationListType {
	if o == nil {
		var ret RUMApplicationListType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RUMApplicationList) GetTypeOk() (*RUMApplicationListType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RUMApplicationList) SetType(v RUMApplicationListType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMApplicationList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMApplicationList) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Attributes *RUMApplicationListAttributes `json:"attributes"`
		Id         *string                       `json:"id,omitempty"`
		Type       *RUMApplicationListType       `json:"type"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}
	if v := all.Type; !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Attributes = *all.Attributes
	o.Id = all.Id
	o.Type = *all.Type
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
