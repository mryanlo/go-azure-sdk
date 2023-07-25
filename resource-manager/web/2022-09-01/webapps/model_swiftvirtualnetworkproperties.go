package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SwiftVirtualNetworkProperties struct {
	SubnetResourceId *string `json:"subnetResourceId,omitempty"`
	SwiftSupported   *bool   `json:"swiftSupported,omitempty"`
}