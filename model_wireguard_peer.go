/*
 * VPN Gateways
 *
 * POC Docs for VPN gateway as service
 *
 * API version: 0.0.1
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// WireguardPeer Properties with all data needed to create a new WireGuard Gateway Peer.\\ __Note__: there is a limit of 20 peers per WireGuard Gateway.
type WireguardPeer struct {
	// The human readable name of your WireguardGateway Peer.
	Name *string `json:"name"`
	// Human readable description of the WireguardGateway Peer.
	Description *string            `json:"description,omitempty"`
	Endpoint    *WireguardEndpoint `json:"endpoint,omitempty"`
	// The subnet CIDRs that are allowed to connect to the WireGuard Gateway.  Specify \"a.b.c.d/32\" for an individual IP address.  Specify \"0.0.0.0/0\" or \"::/0\" for all addresses.
	AllowedIPs *[]string `json:"allowedIPs"`
	// WireGuard public key of the connecting peer
	PublicKey *string `json:"publicKey"`
}

// NewWireguardPeer instantiates a new WireguardPeer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireguardPeer(name string, allowedIPs []string, publicKey string) *WireguardPeer {
	this := WireguardPeer{}

	this.Name = &name
	this.AllowedIPs = &allowedIPs
	this.PublicKey = &publicKey

	return &this
}

// NewWireguardPeerWithDefaults instantiates a new WireguardPeer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireguardPeerWithDefaults() *WireguardPeer {
	this := WireguardPeer{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WireguardPeer) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardPeer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *WireguardPeer) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *WireguardPeer) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WireguardPeer) GetDescription() *string {
	if o == nil {
		return nil
	}

	return o.Description

}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardPeer) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Description, true
}

// SetDescription sets field value
func (o *WireguardPeer) SetDescription(v string) {

	o.Description = &v

}

// HasDescription returns a boolean if a field has been set.
func (o *WireguardPeer) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// GetEndpoint returns the Endpoint field value
// If the value is explicit nil, the zero value for WireguardEndpoint will be returned
func (o *WireguardPeer) GetEndpoint() *WireguardEndpoint {
	if o == nil {
		return nil
	}

	return o.Endpoint

}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardPeer) GetEndpointOk() (*WireguardEndpoint, bool) {
	if o == nil {
		return nil, false
	}

	return o.Endpoint, true
}

// SetEndpoint sets field value
func (o *WireguardPeer) SetEndpoint(v WireguardEndpoint) {

	o.Endpoint = &v

}

// HasEndpoint returns a boolean if a field has been set.
func (o *WireguardPeer) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// GetAllowedIPs returns the AllowedIPs field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *WireguardPeer) GetAllowedIPs() *[]string {
	if o == nil {
		return nil
	}

	return o.AllowedIPs

}

// GetAllowedIPsOk returns a tuple with the AllowedIPs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardPeer) GetAllowedIPsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.AllowedIPs, true
}

// SetAllowedIPs sets field value
func (o *WireguardPeer) SetAllowedIPs(v []string) {

	o.AllowedIPs = &v

}

// HasAllowedIPs returns a boolean if a field has been set.
func (o *WireguardPeer) HasAllowedIPs() bool {
	if o != nil && o.AllowedIPs != nil {
		return true
	}

	return false
}

// GetPublicKey returns the PublicKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WireguardPeer) GetPublicKey() *string {
	if o == nil {
		return nil
	}

	return o.PublicKey

}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardPeer) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.PublicKey, true
}

// SetPublicKey sets field value
func (o *WireguardPeer) SetPublicKey(v string) {

	o.PublicKey = &v

}

// HasPublicKey returns a boolean if a field has been set.
func (o *WireguardPeer) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

func (o WireguardPeer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}

	if o.AllowedIPs != nil {
		toSerialize["allowedIPs"] = o.AllowedIPs
	}

	if o.PublicKey != nil {
		toSerialize["publicKey"] = o.PublicKey
	}

	return json.Marshal(toSerialize)
}

type NullableWireguardPeer struct {
	value *WireguardPeer
	isSet bool
}

func (v NullableWireguardPeer) Get() *WireguardPeer {
	return v.value
}

func (v *NullableWireguardPeer) Set(val *WireguardPeer) {
	v.value = val
	v.isSet = true
}

func (v NullableWireguardPeer) IsSet() bool {
	return v.isSet
}

func (v *NullableWireguardPeer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireguardPeer(val *WireguardPeer) *NullableWireguardPeer {
	return &NullableWireguardPeer{value: val, isSet: true}
}

func (v NullableWireguardPeer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireguardPeer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
