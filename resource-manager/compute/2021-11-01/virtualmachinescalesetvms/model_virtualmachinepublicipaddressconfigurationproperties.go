package virtualmachinescalesetvms

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachinePublicIPAddressConfigurationProperties struct {
	DeleteOption             *DeleteOptions                                         `json:"deleteOption,omitempty"`
	DnsSettings              *VirtualMachinePublicIPAddressDnsSettingsConfiguration `json:"dnsSettings,omitempty"`
	IdleTimeoutInMinutes     *int64                                                 `json:"idleTimeoutInMinutes,omitempty"`
	IpTags                   *[]VirtualMachineIpTag                                 `json:"ipTags,omitempty"`
	PublicIPAddressVersion   *IPVersions                                            `json:"publicIPAddressVersion,omitempty"`
	PublicIPAllocationMethod *PublicIPAllocationMethod                              `json:"publicIPAllocationMethod,omitempty"`
	PublicIPPrefix           *SubResource                                           `json:"publicIPPrefix,omitempty"`
}