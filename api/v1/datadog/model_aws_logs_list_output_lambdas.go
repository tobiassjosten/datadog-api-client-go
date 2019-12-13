/*
 * (C) Datadog, Inc. 2019
 * All rights reserved
 * Licensed under a 3-clause BSD style license (see LICENSE)
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// AwsLogsListOutputLambdas struct for AwsLogsListOutputLambdas
type AwsLogsListOutputLambdas struct {
	Arn *string `json:"arn,omitempty"`
}

// GetArn returns the Arn field value if set, zero value otherwise.
func (o *AwsLogsListOutputLambdas) GetArn() string {
	if o == nil || o.Arn == nil {
		var ret string
		return ret
	}
	return *o.Arn
}

// GetArnOk returns a tuple with the Arn field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AwsLogsListOutputLambdas) GetArnOk() (string, bool) {
	if o == nil || o.Arn == nil {
		var ret string
		return ret, false
	}
	return *o.Arn, true
}

// HasArn returns a boolean if a field has been set.
func (o *AwsLogsListOutputLambdas) HasArn() bool {
	if o != nil && o.Arn != nil {
		return true
	}

	return false
}

// SetArn gets a reference to the given string and assigns it to the Arn field.
func (o *AwsLogsListOutputLambdas) SetArn(v string) {
	o.Arn = &v
}

type NullableAwsLogsListOutputLambdas struct {
	Value        AwsLogsListOutputLambdas
	ExplicitNull bool
}

func (v NullableAwsLogsListOutputLambdas) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAwsLogsListOutputLambdas) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
