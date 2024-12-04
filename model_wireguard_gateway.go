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

// WireguardGateway Properties with all data needed to create a new WireGuard Gateway.
type WireguardGateway struct {
	// The human readable name of your WireguardGateway.
	Name *string `json:"name"`
	// Human readable description of the WireguardGateway.
	Description *string `json:"description,omitempty"`
	// Public IP address to be assigned to the gateway. __Note__: This must be an IP address in the same datacenter as the connections.
	GatewayIP *string `json:"gatewayIP"`
	// Describes a range of IP V4 addresses in CIDR notation.
	InterfaceIPv4CIDR *string `json:"interfaceIPv4CIDR,omitempty"`
	// Describes a range of IP V6 addresses in CIDR notation.
	InterfaceIPv6CIDR *string `json:"interfaceIPv6CIDR,omitempty"`
	// The network connection for your gateway. __Note__: all connections must belong to the same datacenterId. There is a limit to the total number of connections. Please refer to product documentation.
	Connections *[]Connection `json:"connections"`
	// PrivateKey used for WireGuard Server
	PrivateKey *string `json:"privateKey"`
	// IP port number
	ListenPort *int32 `json:"listenPort,omitempty"`
	// Gateway performance options.  See product documentation for full details.\\ Options: - STANDARD - STANDARD_HA - ENHANCED - ENHANCED_HA - PREMIUM - PREMIUM_HA
	Tier              *string            `json:"tier,omitempty"`
	MaintenanceWindow *MaintenanceWindow `json:"maintenanceWindow,omitempty"`
}

// NewWireguardGateway instantiates a new WireguardGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireguardGateway(name string, gatewayIP string, connections []Connection, privateKey string) *WireguardGateway {
	this := WireguardGateway{}

	this.Name = &name
	this.GatewayIP = &gatewayIP
	this.Connections = &connections
	this.PrivateKey = &privateKey
	var listenPort int32 = 51820
	this.ListenPort = &listenPort
	var tier string = "STANDARD"
	this.Tier = &tier

	return &this
}

// NewWireguardGatewayWithDefaults instantiates a new WireguardGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireguardGatewayWithDefaults() *WireguardGateway {
	this := WireguardGateway{}
	var listenPort int32 = 51820
	this.ListenPort = &listenPort
	var tier string = "STANDARD"
	this.Tier = &tier
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WireguardGateway) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardGateway) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *WireguardGateway) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *WireguardGateway) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WireguardGateway) GetDescription() *string {
	if o == nil {
		return nil
	}

	return o.Description

}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardGateway) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Description, true
}

// SetDescription sets field value
func (o *WireguardGateway) SetDescription(v string) {

	o.Description = &v

}

// HasDescription returns a boolean if a field has been set.
func (o *WireguardGateway) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// GetGatewayIP returns the GatewayIP field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WireguardGateway) GetGatewayIP() *string {
	if o == nil {
		return nil
	}

	return o.GatewayIP

}

// GetGatewayIPOk returns a tuple with the GatewayIP field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardGateway) GetGatewayIPOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.GatewayIP, true
}

// SetGatewayIP sets field value
func (o *WireguardGateway) SetGatewayIP(v string) {

	o.GatewayIP = &v

}

// HasGatewayIP returns a boolean if a field has been set.
func (o *WireguardGateway) HasGatewayIP() bool {
	if o != nil && o.GatewayIP != nil {
		return true
	}

	return false
}

// GetInterfaceIPv4CIDR returns the InterfaceIPv4CIDR field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WireguardGateway) GetInterfaceIPv4CIDR() *string {
	if o == nil {
		return nil
	}

	return o.InterfaceIPv4CIDR

}

// GetInterfaceIPv4CIDROk returns a tuple with the InterfaceIPv4CIDR field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardGateway) GetInterfaceIPv4CIDROk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.InterfaceIPv4CIDR, true
}

// SetInterfaceIPv4CIDR sets field value
func (o *WireguardGateway) SetInterfaceIPv4CIDR(v string) {

	o.InterfaceIPv4CIDR = &v

}

// HasInterfaceIPv4CIDR returns a boolean if a field has been set.
func (o *WireguardGateway) HasInterfaceIPv4CIDR() bool {
	if o != nil && o.InterfaceIPv4CIDR != nil {
		return true
	}

	return false
}

// GetInterfaceIPv6CIDR returns the InterfaceIPv6CIDR field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WireguardGateway) GetInterfaceIPv6CIDR() *string {
	if o == nil {
		return nil
	}

	return o.InterfaceIPv6CIDR

}

// GetInterfaceIPv6CIDROk returns a tuple with the InterfaceIPv6CIDR field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardGateway) GetInterfaceIPv6CIDROk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.InterfaceIPv6CIDR, true
}

// SetInterfaceIPv6CIDR sets field value
func (o *WireguardGateway) SetInterfaceIPv6CIDR(v string) {

	o.InterfaceIPv6CIDR = &v

}

// HasInterfaceIPv6CIDR returns a boolean if a field has been set.
func (o *WireguardGateway) HasInterfaceIPv6CIDR() bool {
	if o != nil && o.InterfaceIPv6CIDR != nil {
		return true
	}

	return false
}

// GetConnections returns the Connections field value
// If the value is explicit nil, the zero value for []Connection will be returned
func (o *WireguardGateway) GetConnections() *[]Connection {
	if o == nil {
		return nil
	}

	return o.Connections

}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardGateway) GetConnectionsOk() (*[]Connection, bool) {
	if o == nil {
		return nil, false
	}

	return o.Connections, true
}

// SetConnections sets field value
func (o *WireguardGateway) SetConnections(v []Connection) {

	o.Connections = &v

}

// HasConnections returns a boolean if a field has been set.
func (o *WireguardGateway) HasConnections() bool {
	if o != nil && o.Connections != nil {
		return true
	}

	return false
}

// GetPrivateKey returns the PrivateKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WireguardGateway) GetPrivateKey() *string {
	if o == nil {
		return nil
	}

	return o.PrivateKey

}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardGateway) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.PrivateKey, true
}

// SetPrivateKey sets field value
func (o *WireguardGateway) SetPrivateKey(v string) {

	o.PrivateKey = &v

}

// HasPrivateKey returns a boolean if a field has been set.
func (o *WireguardGateway) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// GetListenPort returns the ListenPort field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *WireguardGateway) GetListenPort() *int32 {
	if o == nil {
		return nil
	}

	return o.ListenPort

}

// GetListenPortOk returns a tuple with the ListenPort field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardGateway) GetListenPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.ListenPort, true
}

// SetListenPort sets field value
func (o *WireguardGateway) SetListenPort(v int32) {

	o.ListenPort = &v

}

// HasListenPort returns a boolean if a field has been set.
func (o *WireguardGateway) HasListenPort() bool {
	if o != nil && o.ListenPort != nil {
		return true
	}

	return false
}

// GetTier returns the Tier field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WireguardGateway) GetTier() *string {
	if o == nil {
		return nil
	}

	return o.Tier

}

// GetTierOk returns a tuple with the Tier field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardGateway) GetTierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Tier, true
}

// SetTier sets field value
func (o *WireguardGateway) SetTier(v string) {

	o.Tier = &v

}

// HasTier returns a boolean if a field has been set.
func (o *WireguardGateway) HasTier() bool {
	if o != nil && o.Tier != nil {
		return true
	}

	return false
}

// GetMaintenanceWindow returns the MaintenanceWindow field value
// If the value is explicit nil, the zero value for MaintenanceWindow will be returned
func (o *WireguardGateway) GetMaintenanceWindow() *MaintenanceWindow {
	if o == nil {
		return nil
	}

	return o.MaintenanceWindow

}

// GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardGateway) GetMaintenanceWindowOk() (*MaintenanceWindow, bool) {
	if o == nil {
		return nil, false
	}

	return o.MaintenanceWindow, true
}

// SetMaintenanceWindow sets field value
func (o *WireguardGateway) SetMaintenanceWindow(v MaintenanceWindow) {

	o.MaintenanceWindow = &v

}

// HasMaintenanceWindow returns a boolean if a field has been set.
func (o *WireguardGateway) HasMaintenanceWindow() bool {
	if o != nil && o.MaintenanceWindow != nil {
		return true
	}

	return false
}

func (o WireguardGateway) MarshalJSON() ([]byte, error) {
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

	if o.InterfaceIPv4CIDR != nil {
		toSerialize["interfaceIPv4CIDR"] = o.InterfaceIPv4CIDR
	}

	if o.InterfaceIPv6CIDR != nil {
		toSerialize["interfaceIPv6CIDR"] = o.InterfaceIPv6CIDR
	}

	if o.Connections != nil {
		toSerialize["connections"] = o.Connections
	}

	if o.PrivateKey != nil {
		toSerialize["privateKey"] = o.PrivateKey
	}

	if o.ListenPort != nil {
		toSerialize["listenPort"] = o.ListenPort
	}

	if o.Tier != nil {
		toSerialize["tier"] = o.Tier
	}

	if o.MaintenanceWindow != nil {
		toSerialize["maintenanceWindow"] = o.MaintenanceWindow
	}

	return json.Marshal(toSerialize)
}

type NullableWireguardGateway struct {
	value *WireguardGateway
	isSet bool
}

func (v NullableWireguardGateway) Get() *WireguardGateway {
	return v.value
}

func (v *NullableWireguardGateway) Set(val *WireguardGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableWireguardGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableWireguardGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireguardGateway(val *WireguardGateway) *NullableWireguardGateway {
	return &NullableWireguardGateway{value: val, isSet: true}
}

func (v NullableWireguardGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireguardGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
