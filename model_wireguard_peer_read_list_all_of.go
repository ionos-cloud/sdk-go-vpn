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

// WireguardPeerReadListAllOf struct for WireguardPeerReadListAllOf
type WireguardPeerReadListAllOf struct {
	// ID of the list of WireguardPeer resources.
	Id *string `json:"id"`
	// The type of the resource.
	Type *string `json:"type"`
	// The URL of the list of WireguardPeer resources.
	Href *string `json:"href"`
	// The list of WireguardPeer resources.
	Items *[]WireguardPeerRead `json:"items,omitempty"`
}

// NewWireguardPeerReadListAllOf instantiates a new WireguardPeerReadListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireguardPeerReadListAllOf(id string, type_ string, href string) *WireguardPeerReadListAllOf {
	this := WireguardPeerReadListAllOf{}

	this.Id = &id
	this.Type = &type_
	this.Href = &href

	return &this
}

// NewWireguardPeerReadListAllOfWithDefaults instantiates a new WireguardPeerReadListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireguardPeerReadListAllOfWithDefaults() *WireguardPeerReadListAllOf {
	this := WireguardPeerReadListAllOf{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WireguardPeerReadListAllOf) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardPeerReadListAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *WireguardPeerReadListAllOf) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *WireguardPeerReadListAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WireguardPeerReadListAllOf) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardPeerReadListAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *WireguardPeerReadListAllOf) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *WireguardPeerReadListAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WireguardPeerReadListAllOf) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardPeerReadListAllOf) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *WireguardPeerReadListAllOf) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *WireguardPeerReadListAllOf) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []WireguardPeerRead will be returned
func (o *WireguardPeerReadListAllOf) GetItems() *[]WireguardPeerRead {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireguardPeerReadListAllOf) GetItemsOk() (*[]WireguardPeerRead, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *WireguardPeerReadListAllOf) SetItems(v []WireguardPeerRead) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *WireguardPeerReadListAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

func (o WireguardPeerReadListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	return json.Marshal(toSerialize)
}

type NullableWireguardPeerReadListAllOf struct {
	value *WireguardPeerReadListAllOf
	isSet bool
}

func (v NullableWireguardPeerReadListAllOf) Get() *WireguardPeerReadListAllOf {
	return v.value
}

func (v *NullableWireguardPeerReadListAllOf) Set(val *WireguardPeerReadListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWireguardPeerReadListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWireguardPeerReadListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireguardPeerReadListAllOf(val *WireguardPeerReadListAllOf) *NullableWireguardPeerReadListAllOf {
	return &NullableWireguardPeerReadListAllOf{value: val, isSet: true}
}

func (v NullableWireguardPeerReadListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireguardPeerReadListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
