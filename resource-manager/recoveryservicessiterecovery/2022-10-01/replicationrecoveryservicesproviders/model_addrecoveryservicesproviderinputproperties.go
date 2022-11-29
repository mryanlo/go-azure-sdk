package replicationrecoveryservicesproviders

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddRecoveryServicesProviderInputProperties struct {
	AuthenticationIdentityInput          IdentityProviderInput  `json:"authenticationIdentityInput"`
	BiosId                               *string                `json:"biosId,omitempty"`
	DataPlaneAuthenticationIdentityInput *IdentityProviderInput `json:"dataPlaneAuthenticationIdentityInput"`
	MachineId                            *string                `json:"machineId,omitempty"`
	MachineName                          string                 `json:"machineName"`
	ResourceAccessIdentityInput          IdentityProviderInput  `json:"resourceAccessIdentityInput"`
}