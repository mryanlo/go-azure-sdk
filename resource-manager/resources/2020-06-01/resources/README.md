
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-06-01/resources` Documentation

The `resources` SDK allows for interaction with the Azure Resource Manager Service `resources` (API Version `2020-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-06-01/resources"
```


### Client Initialization

```go
client := resources.NewResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `ResourcesClient.CheckExistence`

```go
ctx := context.TODO()
id := resources.NewResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceProviderNamespaceValue", "parentResourcePathValue", "resourceTypeValue", "resourceValue")
read, err := client.CheckExistence(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourcesClient.CheckExistenceById`

```go
ctx := context.TODO()
id := resources.NewScopeID()
read, err := client.CheckExistenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourcesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := resources.NewResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceProviderNamespaceValue", "parentResourcePathValue", "resourceTypeValue", "resourceValue")

payload := resources.GenericResource{
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


### Example Usage: `ResourcesClient.CreateOrUpdateById`

```go
ctx := context.TODO()
id := resources.NewScopeID()

payload := resources.GenericResource{
	// ...
}

future, err := client.CreateOrUpdateById(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `ResourcesClient.Delete`

```go
ctx := context.TODO()
id := resources.NewResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceProviderNamespaceValue", "parentResourcePathValue", "resourceTypeValue", "resourceValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `ResourcesClient.DeleteById`

```go
ctx := context.TODO()
id := resources.NewScopeID()
future, err := client.DeleteById(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `ResourcesClient.Get`

```go
ctx := context.TODO()
id := resources.NewResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceProviderNamespaceValue", "parentResourcePathValue", "resourceTypeValue", "resourceValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourcesClient.GetById`

```go
ctx := context.TODO()
id := resources.NewScopeID()
read, err := client.GetById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourcesClient.List`

```go
ctx := context.TODO()
id := resources.NewSubscriptionID()
// alternatively `client.List(ctx, id, resources.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, resources.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ResourcesClient.MoveResources`

```go
ctx := context.TODO()
id := resources.NewSubscriptionResourceGroupID("12345678-1234-9876-4563-123456789012", "sourceResourceGroupValue")

payload := resources.ResourcesMoveInfo{
	// ...
}

future, err := client.MoveResources(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `ResourcesClient.Update`

```go
ctx := context.TODO()
id := resources.NewResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceProviderNamespaceValue", "parentResourcePathValue", "resourceTypeValue", "resourceValue")

payload := resources.GenericResource{
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


### Example Usage: `ResourcesClient.UpdateById`

```go
ctx := context.TODO()
id := resources.NewScopeID()

payload := resources.GenericResource{
	// ...
}

future, err := client.UpdateById(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `ResourcesClient.ValidateMoveResources`

```go
ctx := context.TODO()
id := resources.NewSubscriptionResourceGroupID("12345678-1234-9876-4563-123456789012", "sourceResourceGroupValue")

payload := resources.ResourcesMoveInfo{
	// ...
}

future, err := client.ValidateMoveResources(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```