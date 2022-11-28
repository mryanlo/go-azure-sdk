package batchaccount

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionProperties struct {
	KeySource          *KeySource          `json:"keySource,omitempty"`
	KeyVaultProperties *KeyVaultProperties `json:"keyVaultProperties"`
}
