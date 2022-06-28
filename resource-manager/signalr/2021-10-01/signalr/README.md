
## `github.com/hashicorp/go-azure-sdk/resource-manager/signalr/2021-10-01/signalr` Documentation

The `signalr` SDK allows for interaction with the Azure Resource Manager Service `signalr` (API Version `2021-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/signalr/2021-10-01/signalr"
```


### Client Initialization

```go
client := signalr.NewSignalRClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `SignalRClient.CheckNameAvailability`

```go
ctx := context.TODO()
id := signalr.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := signalr.NameAvailabilityParameters{
	// ...
}

read, err := client.CheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SignalRClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := signalr.NewSignalRID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue")

payload := signalr.SignalRResource{
	// ...
}

future, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `SignalRClient.Delete`

```go
ctx := context.TODO()
id := signalr.NewSignalRID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `SignalRClient.Get`

```go
ctx := context.TODO()
id := signalr.NewSignalRID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SignalRClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := signalr.NewResourceGroupID()
// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SignalRClient.ListBySubscription`

```go
ctx := context.TODO()
id := signalr.NewSubscriptionID()
// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SignalRClient.ListKeys`

```go
ctx := context.TODO()
id := signalr.NewSignalRID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue")
read, err := client.ListKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SignalRClient.ListSkus`

```go
ctx := context.TODO()
id := signalr.NewSignalRID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue")
read, err := client.ListSkus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SignalRClient.PrivateEndpointConnectionsDelete`

```go
ctx := context.TODO()
id := signalr.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue", "privateEndpointConnectionValue")
future, err := client.PrivateEndpointConnectionsDelete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `SignalRClient.PrivateEndpointConnectionsGet`

```go
ctx := context.TODO()
id := signalr.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue", "privateEndpointConnectionValue")
read, err := client.PrivateEndpointConnectionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SignalRClient.PrivateEndpointConnectionsList`

```go
ctx := context.TODO()
id := signalr.NewSignalRID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue")
// alternatively `client.PrivateEndpointConnectionsList(ctx, id)` can be used to do batched pagination
items, err := client.PrivateEndpointConnectionsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SignalRClient.PrivateEndpointConnectionsUpdate`

```go
ctx := context.TODO()
id := signalr.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue", "privateEndpointConnectionValue")

payload := signalr.PrivateEndpointConnection{
	// ...
}

read, err := client.PrivateEndpointConnectionsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SignalRClient.PrivateLinkResourcesList`

```go
ctx := context.TODO()
id := signalr.NewSignalRID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue")
// alternatively `client.PrivateLinkResourcesList(ctx, id)` can be used to do batched pagination
items, err := client.PrivateLinkResourcesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SignalRClient.RegenerateKey`

```go
ctx := context.TODO()
id := signalr.NewSignalRID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue")

payload := signalr.RegenerateKeyParameters{
	// ...
}

future, err := client.RegenerateKey(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `SignalRClient.Restart`

```go
ctx := context.TODO()
id := signalr.NewSignalRID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue")
future, err := client.Restart(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `SignalRClient.SharedPrivateLinkResourcesCreateOrUpdate`

```go
ctx := context.TODO()
id := signalr.NewSharedPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue", "sharedPrivateLinkResourceValue")

payload := signalr.SharedPrivateLinkResource{
	// ...
}

future, err := client.SharedPrivateLinkResourcesCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `SignalRClient.SharedPrivateLinkResourcesDelete`

```go
ctx := context.TODO()
id := signalr.NewSharedPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue", "sharedPrivateLinkResourceValue")
future, err := client.SharedPrivateLinkResourcesDelete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `SignalRClient.SharedPrivateLinkResourcesGet`

```go
ctx := context.TODO()
id := signalr.NewSharedPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue", "sharedPrivateLinkResourceValue")
read, err := client.SharedPrivateLinkResourcesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SignalRClient.SharedPrivateLinkResourcesList`

```go
ctx := context.TODO()
id := signalr.NewSignalRID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue")
// alternatively `client.SharedPrivateLinkResourcesList(ctx, id)` can be used to do batched pagination
items, err := client.SharedPrivateLinkResourcesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SignalRClient.Update`

```go
ctx := context.TODO()
id := signalr.NewSignalRID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue")

payload := signalr.SignalRResource{
	// ...
}

future, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `SignalRClient.UsagesList`

```go
ctx := context.TODO()
id := signalr.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")
// alternatively `client.UsagesList(ctx, id)` can be used to do batched pagination
items, err := client.UsagesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```