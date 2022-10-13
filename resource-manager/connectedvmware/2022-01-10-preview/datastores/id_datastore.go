package datastores

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = DataStoreId{}

// DataStoreId is a struct representing the Resource ID for a Data Store
type DataStoreId struct {
	SubscriptionId    string
	ResourceGroupName string
	DatastoreName     string
}

// NewDataStoreID returns a new DataStoreId struct
func NewDataStoreID(subscriptionId string, resourceGroupName string, datastoreName string) DataStoreId {
	return DataStoreId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		DatastoreName:     datastoreName,
	}
}

// ParseDataStoreID parses 'input' into a DataStoreId
func ParseDataStoreID(input string) (*DataStoreId, error) {
	parser := resourceids.NewParserFromResourceIdType(DataStoreId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DataStoreId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.DatastoreName, ok = parsed.Parsed["datastoreName"]; !ok {
		return nil, fmt.Errorf("the segment 'datastoreName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseDataStoreIDInsensitively parses 'input' case-insensitively into a DataStoreId
// note: this method should only be used for API response data and not user input
func ParseDataStoreIDInsensitively(input string) (*DataStoreId, error) {
	parser := resourceids.NewParserFromResourceIdType(DataStoreId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DataStoreId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.DatastoreName, ok = parsed.Parsed["datastoreName"]; !ok {
		return nil, fmt.Errorf("the segment 'datastoreName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateDataStoreID checks that 'input' can be parsed as a Data Store ID
func ValidateDataStoreID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDataStoreID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Data Store ID
func (id DataStoreId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ConnectedVMwarevSphere/dataStores/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DatastoreName)
}

// Segments returns a slice of Resource ID Segments which comprise this Data Store ID
func (id DataStoreId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftConnectedVMwarevSphere", "Microsoft.ConnectedVMwarevSphere", "Microsoft.ConnectedVMwarevSphere"),
		resourceids.StaticSegment("staticDataStores", "dataStores", "dataStores"),
		resourceids.UserSpecifiedSegment("datastoreName", "datastoreValue"),
	}
}

// String returns a human-readable description of this Data Store ID
func (id DataStoreId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Datastore Name: %q", id.DatastoreName),
	}
	return fmt.Sprintf("Data Store (%s)", strings.Join(components, "\n"))
}
