// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentMeta The definition of `DeploymentMeta` object.
type DeploymentMeta struct {
	// The `DeploymentMeta` `created_at`.
	CreatedAt *string `json:"created_at,omitempty"`
	// The `DeploymentMeta` `user_id`.
	UserId *int64 `json:"user_id,omitempty"`
	// The `DeploymentMeta` `user_name`.
	UserName *string `json:"user_name,omitempty"`
	// The `DeploymentMeta` `user_uuid`.
	UserUuid *string `json:"user_uuid,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentMeta instantiates a new DeploymentMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentMeta() *DeploymentMeta {
	this := DeploymentMeta{}
	return &this
}

// NewDeploymentMetaWithDefaults instantiates a new DeploymentMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentMetaWithDefaults() *DeploymentMeta {
	this := DeploymentMeta{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DeploymentMeta) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentMeta) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DeploymentMeta) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *DeploymentMeta) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *DeploymentMeta) GetUserId() int64 {
	if o == nil || o.UserId == nil {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentMeta) GetUserIdOk() (*int64, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *DeploymentMeta) HasUserId() bool {
	return o != nil && o.UserId != nil
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *DeploymentMeta) SetUserId(v int64) {
	o.UserId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *DeploymentMeta) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentMeta) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *DeploymentMeta) HasUserName() bool {
	return o != nil && o.UserName != nil
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *DeploymentMeta) SetUserName(v string) {
	o.UserName = &v
}

// GetUserUuid returns the UserUuid field value if set, zero value otherwise.
func (o *DeploymentMeta) GetUserUuid() string {
	if o == nil || o.UserUuid == nil {
		var ret string
		return ret
	}
	return *o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentMeta) GetUserUuidOk() (*string, bool) {
	if o == nil || o.UserUuid == nil {
		return nil, false
	}
	return o.UserUuid, true
}

// HasUserUuid returns a boolean if a field has been set.
func (o *DeploymentMeta) HasUserUuid() bool {
	return o != nil && o.UserUuid != nil
}

// SetUserUuid gets a reference to the given string and assigns it to the UserUuid field.
func (o *DeploymentMeta) SetUserUuid(v string) {
	o.UserUuid = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	if o.UserName != nil {
		toSerialize["user_name"] = o.UserName
	}
	if o.UserUuid != nil {
		toSerialize["user_uuid"] = o.UserUuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt *string `json:"created_at,omitempty"`
		UserId    *int64  `json:"user_id,omitempty"`
		UserName  *string `json:"user_name,omitempty"`
		UserUuid  *string `json:"user_uuid,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "user_id", "user_name", "user_uuid"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.UserId = all.UserId
	o.UserName = all.UserName
	o.UserUuid = all.UserUuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
