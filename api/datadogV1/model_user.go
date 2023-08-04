// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// User Create, edit, and disable users.
type User struct {
	// The access role of the user. Options are **st** (standard user), **adm** (admin user), or **ro** (read-only user).
	AccessRole NullableAccessRole `json:"access_role,omitempty"`
	// The new disabled status of the user.
	Disabled *bool `json:"disabled,omitempty"`
	// The new email of the user.
	Email *string `json:"email,omitempty"`
	// The user handle, must be a valid email.
	Handle *string `json:"handle,omitempty"`
	// Gravatar icon associated to the user.
	Icon *string `json:"icon,omitempty"`
	// The name of the user.
	Name *string `json:"name,omitempty"`
	// Whether or not the user logged in Datadog at least once.
	Verified *bool `json:"verified,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUser instantiates a new User object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetAccessRole returns the AccessRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetAccessRole() AccessRole {
	if o == nil || o.AccessRole.Get() == nil {
		var ret AccessRole
		return ret
	}
	return *o.AccessRole.Get()
}

// GetAccessRoleOk returns a tuple with the AccessRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *User) GetAccessRoleOk() (*AccessRole, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessRole.Get(), o.AccessRole.IsSet()
}

// HasAccessRole returns a boolean if a field has been set.
func (o *User) HasAccessRole() bool {
	return o != nil && o.AccessRole.IsSet()
}

// SetAccessRole gets a reference to the given NullableAccessRole and assigns it to the AccessRole field.
func (o *User) SetAccessRole(v AccessRole) {
	o.AccessRole.Set(&v)
}

// SetAccessRoleNil sets the value for AccessRole to be an explicit nil.
func (o *User) SetAccessRoleNil() {
	o.AccessRole.Set(nil)
}

// UnsetAccessRole ensures that no value is present for AccessRole, not even an explicit nil.
func (o *User) UnsetAccessRole() {
	o.AccessRole.Unset()
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *User) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *User) HasDisabled() bool {
	return o != nil && o.Disabled != nil
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *User) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *User) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *User) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *User) SetEmail(v string) {
	o.Email = &v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *User) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *User) HasHandle() bool {
	return o != nil && o.Handle != nil
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *User) SetHandle(v string) {
	o.Handle = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *User) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *User) HasIcon() bool {
	return o != nil && o.Icon != nil
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *User) SetIcon(v string) {
	o.Icon = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *User) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *User) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *User) SetName(v string) {
	o.Name = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *User) GetVerified() bool {
	if o == nil || o.Verified == nil {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetVerifiedOk() (*bool, bool) {
	if o == nil || o.Verified == nil {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *User) HasVerified() bool {
	return o != nil && o.Verified != nil
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *User) SetVerified(v bool) {
	o.Verified = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.AccessRole.IsSet() {
		toSerialize["access_role"] = o.AccessRole.Get()
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Verified != nil {
		toSerialize["verified"] = o.Verified
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *User) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		AccessRole NullableAccessRole `json:"access_role,omitempty"`
		Disabled   *bool              `json:"disabled,omitempty"`
		Email      *string            `json:"email,omitempty"`
		Handle     *string            `json:"handle,omitempty"`
		Icon       *string            `json:"icon,omitempty"`
		Name       *string            `json:"name,omitempty"`
		Verified   *bool              `json:"verified,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"access_role", "disabled", "email", "handle", "icon", "name", "verified"})
	} else {
		return err
	}

	hasInvalidField := false
	if v := all.AccessRole; v.Get() != nil && !v.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.AccessRole = all.AccessRole
	}
	o.Disabled = all.Disabled
	o.Email = all.Email
	o.Handle = all.Handle
	o.Icon = all.Icon
	o.Name = all.Name
	o.Verified = all.Verified

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}

	return nil
}
