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

// IPSecGateway Properties with all data needed to create a new IPSec Gateway.
type IPSecGateway struct {
	// The human readable name of your IPSecGateway.
	Name *string `json:"name"`
	// Human readable description of the IPSecGateway.
	Description *string `json:"description,omitempty"`
	// Public IP address to be assigned to the gateway. __Note__: This must be an IP address in the same datacenter as the connections.
	GatewayIP *string `json:"gatewayIP"`
	// The network connection for your gateway. __Note__: all connections must belong to the same datacenterId.
	Connections *[]Connection `json:"connections"`
	// The IKE version that is permitted for the VPN tunnels.\\ Options:  - IKEv2
	Version *string `json:"version,omitempty"`
}

// NewIPSecGateway instantiates a new IPSecGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPSecGateway(name string, gatewayIP string, connections []Connection) *IPSecGateway {
	this := IPSecGateway{}

	this.Name = &name
	this.GatewayIP = &gatewayIP
	this.Connections = &connections
	var version string = "IKEv2"
	this.Version = &version

	return &this
}

// NewIPSecGatewayWithDefaults instantiates a new IPSecGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPSecGatewayWithDefaults() *IPSecGateway {
	this := IPSecGateway{}
	var version string = "IKEv2"
	this.Version = &version
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IPSecGateway) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPSecGateway) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *IPSecGateway) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *IPSecGateway) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IPSecGateway) GetDescription() *string {
	if o == nil {
		return nil
	}

	return o.Description

}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPSecGateway) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Description, true
}

// SetDescription sets field value
func (o *IPSecGateway) SetDescription(v string) {

	o.Description = &v

}

// HasDescription returns a boolean if a field has been set.
func (o *IPSecGateway) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// GetGatewayIP returns the GatewayIP field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IPSecGateway) GetGatewayIP() *string {
	if o == nil {
		return nil
	}

	return o.GatewayIP

}

// GetGatewayIPOk returns a tuple with the GatewayIP field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPSecGateway) GetGatewayIPOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.GatewayIP, true
}

// SetGatewayIP sets field value
func (o *IPSecGateway) SetGatewayIP(v string) {

	o.GatewayIP = &v

}

// HasGatewayIP returns a boolean if a field has been set.
func (o *IPSecGateway) HasGatewayIP() bool {
	if o != nil && o.GatewayIP != nil {
		return true
	}

	return false
}

// GetConnections returns the Connections field value
// If the value is explicit nil, the zero value for []Connection will be returned
func (o *IPSecGateway) GetConnections() *[]Connection {
	if o == nil {
		return nil
	}

	return o.Connections

}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPSecGateway) GetConnectionsOk() (*[]Connection, bool) {
	if o == nil {
		return nil, false
	}

	return o.Connections, true
}

// SetConnections sets field value
func (o *IPSecGateway) SetConnections(v []Connection) {

	o.Connections = &v

}

// HasConnections returns a boolean if a field has been set.
func (o *IPSecGateway) HasConnections() bool {
	if o != nil && o.Connections != nil {
		return true
	}

	return false
}

// GetVersion returns the Version field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IPSecGateway) GetVersion() *string {
	if o == nil {
		return nil
	}

	return o.Version

}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPSecGateway) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Version, true
}

// SetVersion sets field value
func (o *IPSecGateway) SetVersion(v string) {

	o.Version = &v

}

// HasVersion returns a boolean if a field has been set.
func (o *IPSecGateway) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

func (o IPSecGateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	if o.GatewayIP != nil {
		toSerialize["gatewayIP"] = o.GatewayIP
	}

	if o.Connections != nil {
		toSerialize["connections"] = o.Connections
	}

	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	return json.Marshal(toSerialize)
}

type NullableIPSecGateway struct {
	value *IPSecGateway
	isSet bool
}

func (v NullableIPSecGateway) Get() *IPSecGateway {
	return v.value
}

func (v *NullableIPSecGateway) Set(val *IPSecGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableIPSecGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableIPSecGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPSecGateway(val *IPSecGateway) *NullableIPSecGateway {
	return &NullableIPSecGateway{value: val, isSet: true}
}

func (v NullableIPSecGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPSecGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
