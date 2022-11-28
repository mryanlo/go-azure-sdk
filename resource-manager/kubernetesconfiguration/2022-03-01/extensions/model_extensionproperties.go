package extensions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtensionProperties struct {
	AksAssignedIdentity            *ExtensionPropertiesAksAssignedIdentity `json:"aksAssignedIdentity"`
	AutoUpgradeMinorVersion        *bool                                   `json:"autoUpgradeMinorVersion,omitempty"`
	ConfigurationProtectedSettings *map[string]string                      `json:"configurationProtectedSettings,omitempty"`
	ConfigurationSettings          *map[string]string                      `json:"configurationSettings,omitempty"`
	CustomLocationSettings         *map[string]string                      `json:"customLocationSettings,omitempty"`
	ErrorInfo                      *ErrorDetail                            `json:"errorInfo"`
	ExtensionType                  *string                                 `json:"extensionType,omitempty"`
	InstalledVersion               *string                                 `json:"installedVersion,omitempty"`
	PackageUri                     *string                                 `json:"packageUri,omitempty"`
	ProvisioningState              *ProvisioningState                      `json:"provisioningState,omitempty"`
	ReleaseTrain                   *string                                 `json:"releaseTrain,omitempty"`
	Scope                          *Scope                                  `json:"scope"`
	Statuses                       *[]ExtensionStatus                      `json:"statuses,omitempty"`
	Version                        *string                                 `json:"version,omitempty"`
}
