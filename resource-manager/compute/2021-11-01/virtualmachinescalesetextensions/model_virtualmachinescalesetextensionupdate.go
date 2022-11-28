package virtualmachinescalesetextensions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetExtensionUpdate struct {
	Id         *string                                    `json:"id,omitempty"`
	Name       *string                                    `json:"name,omitempty"`
	Properties *VirtualMachineScaleSetExtensionProperties `json:"properties"`
	Type       *string                                    `json:"type,omitempty"`
}
