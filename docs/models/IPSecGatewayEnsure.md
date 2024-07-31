# IPSecGatewayEnsure

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the IPSecGateway. | |
|**Metadata** | Pointer to **map[string]interface{}** | Metadata | [optional] |
|**Properties** | [**IPSecGateway**](IPSecGateway.md) |  | |

## Methods

### NewIPSecGatewayEnsure

`func NewIPSecGatewayEnsure(id string, properties IPSecGateway, ) *IPSecGatewayEnsure`

NewIPSecGatewayEnsure instantiates a new IPSecGatewayEnsure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecGatewayEnsureWithDefaults

`func NewIPSecGatewayEnsureWithDefaults() *IPSecGatewayEnsure`

NewIPSecGatewayEnsureWithDefaults instantiates a new IPSecGatewayEnsure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPSecGatewayEnsure) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPSecGatewayEnsure) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPSecGatewayEnsure) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *IPSecGatewayEnsure) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IPSecGatewayEnsure) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IPSecGatewayEnsure) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IPSecGatewayEnsure) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *IPSecGatewayEnsure) GetProperties() IPSecGateway`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IPSecGatewayEnsure) GetPropertiesOk() (*IPSecGateway, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IPSecGatewayEnsure) SetProperties(v IPSecGateway)`

SetProperties sets Properties field to given value.



