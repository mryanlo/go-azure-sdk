package dataconnectors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MCASDataConnectorProperties struct {
	DataTypes *MCASDataConnectorDataTypes `json:"dataTypes,omitempty"`
	TenantId  *string                     `json:"tenantId,omitempty"`
}
