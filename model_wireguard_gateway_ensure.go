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

// WireguardGatewayEnsure struct for WireguardGatewayEnsure
type WireguardGatewayEnsure struct {
	// The ID (UUID) of the WireguardGateway.
	Id *string `json:"id"`
	// Metadata
	Metadata   *map[string]interface{} `json:"metadata,omitempty"`
	Properties *WireguardGateway       `json:"properties"`
}

// NewWireguardGatewayEnsure instantiates a new WireguardGatewayEnsure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireguardGatewayEnsure(id string, properties WireguardGateway) *WireguardGatewayEnsure {
	this := WireguardGatewayEnsure{}

	this.Id = &id
	this.Properties = &properties

	return &this
}

// NewWireguardGatewayEnsureWithDefaults instantiates a new WireguardGatewayEnsure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireguardGatewayEnsureWithDefaults() *WireguardGatewayEnsure {
	this := WireguardGatewayEnsure{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WireguardGatewayEnsure) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardGatewayEnsure) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *WireguardGatewayEnsure) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *WireguardGatewayEnsure) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *WireguardGatewayEnsure) GetMetadata() *map[string]interface{} {
	if o == nil {
		return nil
	}

	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardGatewayEnsure) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}

	return o.Metadata, true
}

// SetMetadata sets field value
func (o *WireguardGatewayEnsure) SetMetadata(v map[string]interface{}) {

	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *WireguardGatewayEnsure) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for WireguardGateway will be returned
func (o *WireguardGatewayEnsure) GetProperties() *WireguardGateway {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardGatewayEnsure) GetPropertiesOk() (*WireguardGateway, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *WireguardGatewayEnsure) SetProperties(v WireguardGateway) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *WireguardGatewayEnsure) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

func (o WireguardGatewayEnsure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	return json.Marshal(toSerialize)
}

type NullableWireguardGatewayEnsure struct {
	value *WireguardGatewayEnsure
	isSet bool
}

func (v NullableWireguardGatewayEnsure) Get() *WireguardGatewayEnsure {
	return v.value
}

func (v *NullableWireguardGatewayEnsure) Set(val *WireguardGatewayEnsure) {
	v.value = val
	v.isSet = true
}

func (v NullableWireguardGatewayEnsure) IsSet() bool {
	return v.isSet
}

func (v *NullableWireguardGatewayEnsure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireguardGatewayEnsure(val *WireguardGatewayEnsure) *NullableWireguardGatewayEnsure {
	return &NullableWireguardGatewayEnsure{value: val, isSet: true}
}

func (v NullableWireguardGatewayEnsure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireguardGatewayEnsure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
