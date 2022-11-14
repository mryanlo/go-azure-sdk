
## `github.com/hashicorp/go-azure-sdk/resource-manager/batch/2022-01-01/batchaccount` Documentation

The `batchaccount` SDK allows for interaction with the Azure Resource Manager Service `batch` (API Version `2022-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/batch/2022-01-01/batchaccount"
```


### Client Initialization

```go
client := batchaccount.NewBatchAccountClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BatchAccountClient.Create`

```go
ctx := context.TODO()
id := batchaccount.NewBatchAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

payload := batchaccount.BatchAccountCreateParameters{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BatchAccountClient.Delete`

```go
ctx := context.TODO()
id := batchaccount.NewBatchAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `BatchAccountClient.Get`

```go
ctx := context.TODO()
id := batchaccount.NewBatchAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BatchAccountClient.GetKeys`

```go
ctx := context.TODO()
id := batchaccount.NewBatchAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

read, err := client.GetKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BatchAccountClient.List`

```go
ctx := context.TODO()
id := batchaccount.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BatchAccountClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := batchaccount.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BatchAccountClient.ListOutboundNetworkDependenciesEndpoints`

```go
ctx := context.TODO()
id := batchaccount.NewBatchAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

// alternatively `client.ListOutboundNetworkDependenciesEndpoints(ctx, id)` can be used to do batched pagination
items, err := client.ListOutboundNetworkDependenciesEndpointsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BatchAccountClient.RegenerateKey`

```go
ctx := context.TODO()
id := batchaccount.NewBatchAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

payload := batchaccount.BatchAccountRegenerateKeyParameters{
	// ...
}


read, err := client.RegenerateKey(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BatchAccountClient.SynchronizeAutoStorageKeys`

```go
ctx := context.TODO()
id := batchaccount.NewBatchAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

read, err := client.SynchronizeAutoStorageKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BatchAccountClient.Update`

```go
ctx := context.TODO()
id := batchaccount.NewBatchAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

payload := batchaccount.BatchAccountUpdateParameters{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```