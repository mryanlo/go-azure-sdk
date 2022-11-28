package appliances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SupportedVersionCatalogVersion struct {
	Data      *SupportedVersionCatalogVersionData `json:"data"`
	Name      *string                             `json:"name,omitempty"`
	Namespace *string                             `json:"namespace,omitempty"`
}
