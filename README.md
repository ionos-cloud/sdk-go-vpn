# Go API client for ionoscloud

POC Docs for VPN gateway as service

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: v1.0.1
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://docs.ionos.com/support/general-information/contact-information](https://docs.ionos.com/support/general-information/contact-information)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import ionoscloud "github.com/ionos-cloud/sdk-go-vpn"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), ionoscloud.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), ionoscloud.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

## Documentation for API Endpoints

All URIs are relative to *https://vpn.de-fra.ionos.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*IPSecGatewaysApi* | [**IpsecgatewaysDelete**](docs/api/IPSecGatewaysApi.md#ipsecgatewaysdelete) | **Delete** /ipsecgateways/{gatewayId} | Delete IPSecGateway
*IPSecGatewaysApi* | [**IpsecgatewaysFindById**](docs/api/IPSecGatewaysApi.md#ipsecgatewaysfindbyid) | **Get** /ipsecgateways/{gatewayId} | Retrieve IPSecGateway
*IPSecGatewaysApi* | [**IpsecgatewaysGet**](docs/api/IPSecGatewaysApi.md#ipsecgatewaysget) | **Get** /ipsecgateways | Retrieve all IPSecGateways
*IPSecGatewaysApi* | [**IpsecgatewaysPost**](docs/api/IPSecGatewaysApi.md#ipsecgatewayspost) | **Post** /ipsecgateways | Create IPSecGateway
*IPSecGatewaysApi* | [**IpsecgatewaysPut**](docs/api/IPSecGatewaysApi.md#ipsecgatewaysput) | **Put** /ipsecgateways/{gatewayId} | Ensure IPSecGateway
*IPSecTunnelsApi* | [**IpsecgatewaysTunnelsDelete**](docs/api/IPSecTunnelsApi.md#ipsecgatewaystunnelsdelete) | **Delete** /ipsecgateways/{gatewayId}/tunnels/{tunnelId} | Delete IPSecTunnel
*IPSecTunnelsApi* | [**IpsecgatewaysTunnelsFindById**](docs/api/IPSecTunnelsApi.md#ipsecgatewaystunnelsfindbyid) | **Get** /ipsecgateways/{gatewayId}/tunnels/{tunnelId} | Retrieve IPSecTunnel
*IPSecTunnelsApi* | [**IpsecgatewaysTunnelsGet**](docs/api/IPSecTunnelsApi.md#ipsecgatewaystunnelsget) | **Get** /ipsecgateways/{gatewayId}/tunnels | Retrieve all IPSecTunnels
*IPSecTunnelsApi* | [**IpsecgatewaysTunnelsPost**](docs/api/IPSecTunnelsApi.md#ipsecgatewaystunnelspost) | **Post** /ipsecgateways/{gatewayId}/tunnels | Create IPSecTunnel
*IPSecTunnelsApi* | [**IpsecgatewaysTunnelsPut**](docs/api/IPSecTunnelsApi.md#ipsecgatewaystunnelsput) | **Put** /ipsecgateways/{gatewayId}/tunnels/{tunnelId} | Ensure IPSecTunnel
*WireguardGatewaysApi* | [**WireguardgatewaysDelete**](docs/api/WireguardGatewaysApi.md#wireguardgatewaysdelete) | **Delete** /wireguardgateways/{gatewayId} | Delete WireguardGateway
*WireguardGatewaysApi* | [**WireguardgatewaysFindById**](docs/api/WireguardGatewaysApi.md#wireguardgatewaysfindbyid) | **Get** /wireguardgateways/{gatewayId} | Retrieve WireguardGateway
*WireguardGatewaysApi* | [**WireguardgatewaysGet**](docs/api/WireguardGatewaysApi.md#wireguardgatewaysget) | **Get** /wireguardgateways | Retrieve all WireguardGateways
*WireguardGatewaysApi* | [**WireguardgatewaysPost**](docs/api/WireguardGatewaysApi.md#wireguardgatewayspost) | **Post** /wireguardgateways | Create WireguardGateway
*WireguardGatewaysApi* | [**WireguardgatewaysPut**](docs/api/WireguardGatewaysApi.md#wireguardgatewaysput) | **Put** /wireguardgateways/{gatewayId} | Ensure WireguardGateway
*WireguardPeersApi* | [**WireguardgatewaysPeersDelete**](docs/api/WireguardPeersApi.md#wireguardgatewayspeersdelete) | **Delete** /wireguardgateways/{gatewayId}/peers/{peerId} | Delete WireguardPeer
*WireguardPeersApi* | [**WireguardgatewaysPeersFindById**](docs/api/WireguardPeersApi.md#wireguardgatewayspeersfindbyid) | **Get** /wireguardgateways/{gatewayId}/peers/{peerId} | Retrieve WireguardPeer
*WireguardPeersApi* | [**WireguardgatewaysPeersGet**](docs/api/WireguardPeersApi.md#wireguardgatewayspeersget) | **Get** /wireguardgateways/{gatewayId}/peers | Retrieve all WireguardPeers
*WireguardPeersApi* | [**WireguardgatewaysPeersPost**](docs/api/WireguardPeersApi.md#wireguardgatewayspeerspost) | **Post** /wireguardgateways/{gatewayId}/peers | Create WireguardPeer
*WireguardPeersApi* | [**WireguardgatewaysPeersPut**](docs/api/WireguardPeersApi.md#wireguardgatewayspeersput) | **Put** /wireguardgateways/{gatewayId}/peers/{peerId} | Ensure WireguardPeer


## Documentation For Models

 - [Connection](docs/models/Connection.md)
 - [ESPEncryption](docs/models/ESPEncryption.md)
 - [Error](docs/models/Error.md)
 - [ErrorMessages](docs/models/ErrorMessages.md)
 - [IKEEncryption](docs/models/IKEEncryption.md)
 - [IPSecGateway](docs/models/IPSecGateway.md)
 - [IPSecGatewayCreate](docs/models/IPSecGatewayCreate.md)
 - [IPSecGatewayEnsure](docs/models/IPSecGatewayEnsure.md)
 - [IPSecGatewayMetadata](docs/models/IPSecGatewayMetadata.md)
 - [IPSecGatewayRead](docs/models/IPSecGatewayRead.md)
 - [IPSecGatewayReadList](docs/models/IPSecGatewayReadList.md)
 - [IPSecGatewayReadListAllOf](docs/models/IPSecGatewayReadListAllOf.md)
 - [IPSecPSK](docs/models/IPSecPSK.md)
 - [IPSecTunnel](docs/models/IPSecTunnel.md)
 - [IPSecTunnelAuth](docs/models/IPSecTunnelAuth.md)
 - [IPSecTunnelCreate](docs/models/IPSecTunnelCreate.md)
 - [IPSecTunnelEnsure](docs/models/IPSecTunnelEnsure.md)
 - [IPSecTunnelMetadata](docs/models/IPSecTunnelMetadata.md)
 - [IPSecTunnelRead](docs/models/IPSecTunnelRead.md)
 - [IPSecTunnelReadList](docs/models/IPSecTunnelReadList.md)
 - [IPSecTunnelReadListAllOf](docs/models/IPSecTunnelReadListAllOf.md)
 - [Links](docs/models/Links.md)
 - [Metadata](docs/models/Metadata.md)
 - [Pagination](docs/models/Pagination.md)
 - [ResourceStatus](docs/models/ResourceStatus.md)
 - [WireguardEndpoint](docs/models/WireguardEndpoint.md)
 - [WireguardGateway](docs/models/WireguardGateway.md)
 - [WireguardGatewayCreate](docs/models/WireguardGatewayCreate.md)
 - [WireguardGatewayEnsure](docs/models/WireguardGatewayEnsure.md)
 - [WireguardGatewayMetadata](docs/models/WireguardGatewayMetadata.md)
 - [WireguardGatewayMetadataAllOf](docs/models/WireguardGatewayMetadataAllOf.md)
 - [WireguardGatewayRead](docs/models/WireguardGatewayRead.md)
 - [WireguardGatewayReadList](docs/models/WireguardGatewayReadList.md)
 - [WireguardGatewayReadListAllOf](docs/models/WireguardGatewayReadListAllOf.md)
 - [WireguardPeer](docs/models/WireguardPeer.md)
 - [WireguardPeerCreate](docs/models/WireguardPeerCreate.md)
 - [WireguardPeerEnsure](docs/models/WireguardPeerEnsure.md)
 - [WireguardPeerMetadata](docs/models/WireguardPeerMetadata.md)
 - [WireguardPeerRead](docs/models/WireguardPeerRead.md)
 - [WireguardPeerReadList](docs/models/WireguardPeerReadList.md)
 - [WireguardPeerReadListAllOf](docs/models/WireguardPeerReadListAllOf.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### tokenAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@cloud.ionos.com

