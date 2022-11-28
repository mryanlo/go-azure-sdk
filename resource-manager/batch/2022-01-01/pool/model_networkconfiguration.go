package pool

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkConfiguration struct {
	DynamicVNetAssignmentScope   *DynamicVNetAssignmentScope   `json:"dynamicVNetAssignmentScope,omitempty"`
	EndpointConfiguration        *PoolEndpointConfiguration    `json:"endpointConfiguration"`
	PublicIPAddressConfiguration *PublicIPAddressConfiguration `json:"publicIPAddressConfiguration"`
	SubnetId                     *string                       `json:"subnetId,omitempty"`
}
