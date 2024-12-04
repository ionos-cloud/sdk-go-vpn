/*
 * IONOS Cloud VPN Gateway API
 *
 * The Managed VPN Gateway service provides secure and scalable connectivity, enabling encrypted communication between your IONOS cloud resources in a VDC and remote networks (on-premises, multi-cloud, private LANs in other VDCs etc).
 *
 * API version: 1.0.0
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// IPSecPSK Properties with all data needed to define IPSec Authentication PSK. This is required if the method is PSK.
type IPSecPSK struct {
	// The Pre-Shared Key used for IPSec Authentication.
	Key *string `json:"key"`
}

// NewIPSecPSK instantiates a new IPSecPSK object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPSecPSK(key string) *IPSecPSK {
	this := IPSecPSK{}

	this.Key = &key

	return &this
}

// NewIPSecPSKWithDefaults instantiates a new IPSecPSK object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPSecPSKWithDefaults() *IPSecPSK {
	this := IPSecPSK{}
	return &this
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IPSecPSK) GetKey() *string {
	if o == nil {
		return nil
	}

	return o.Key

}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPSecPSK) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Key, true
}

// SetKey sets field value
func (o *IPSecPSK) SetKey(v string) {

	o.Key = &v

}

// HasKey returns a boolean if a field has been set.
func (o *IPSecPSK) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

func (o IPSecPSK) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}

	return json.Marshal(toSerialize)
}

type NullableIPSecPSK struct {
	value *IPSecPSK
	isSet bool
}

func (v NullableIPSecPSK) Get() *IPSecPSK {
	return v.value
}

func (v *NullableIPSecPSK) Set(val *IPSecPSK) {
	v.value = val
	v.isSet = true
}

func (v NullableIPSecPSK) IsSet() bool {
	return v.isSet
}

func (v *NullableIPSecPSK) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPSecPSK(val *IPSecPSK) *NullableIPSecPSK {
	return &NullableIPSecPSK{value: val, isSet: true}
}

func (v NullableIPSecPSK) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPSecPSK) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
