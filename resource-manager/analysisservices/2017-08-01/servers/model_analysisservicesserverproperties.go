package servers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AnalysisServicesServerProperties struct {
	AsAdministrators        *ServerAdministrators `json:"asAdministrators"`
	BackupBlobContainerUri  *string               `json:"backupBlobContainerUri,omitempty"`
	GatewayDetails          *GatewayDetails       `json:"gatewayDetails"`
	IPV4FirewallSettings    *IPv4FirewallSettings `json:"ipV4FirewallSettings"`
	ManagedMode             *ManagedMode          `json:"managedMode,omitempty"`
	ProvisioningState       *ProvisioningState    `json:"provisioningState,omitempty"`
	QuerypoolConnectionMode *ConnectionMode       `json:"querypoolConnectionMode,omitempty"`
	ServerFullName          *string               `json:"serverFullName,omitempty"`
	ServerMonitorMode       *ServerMonitorMode    `json:"serverMonitorMode,omitempty"`
	Sku                     *ResourceSku          `json:"sku"`
	State                   *State                `json:"state,omitempty"`
}
