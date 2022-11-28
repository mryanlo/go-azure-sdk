package diskpools

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiskPool struct {
	Id                *string                `json:"id,omitempty"`
	Location          string                 `json:"location"`
	ManagedBy         *string                `json:"managedBy,omitempty"`
	ManagedByExtended *[]string              `json:"managedByExtended,omitempty"`
	Name              *string                `json:"name,omitempty"`
	Properties        DiskPoolProperties     `json:"properties"`
	Sku               *Sku                   `json:"sku"`
	SystemData        *systemdata.SystemData `json:"systemData,omitempty"`
	Tags              *map[string]string     `json:"tags,omitempty"`
	Type              *string                `json:"type,omitempty"`
}
