package batchaccount

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BatchAccountCreateProperties struct {
	AllowedAuthenticationModes *[]AuthenticationMode      `json:"allowedAuthenticationModes,omitempty"`
	AutoStorage                *AutoStorageBaseProperties `json:"autoStorage"`
	Encryption                 *EncryptionProperties      `json:"encryption"`
	KeyVaultReference          *KeyVaultReference         `json:"keyVaultReference"`
	PoolAllocationMode         *PoolAllocationMode        `json:"poolAllocationMode,omitempty"`
	PublicNetworkAccess        *PublicNetworkAccessType   `json:"publicNetworkAccess,omitempty"`
}
