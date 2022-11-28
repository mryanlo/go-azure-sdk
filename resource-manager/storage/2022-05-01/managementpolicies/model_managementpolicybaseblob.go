package managementpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementPolicyBaseBlob struct {
	Delete                      *DateAfterModification `json:"delete"`
	EnableAutoTierToHotFromCool *bool                  `json:"enableAutoTierToHotFromCool,omitempty"`
	TierToArchive               *DateAfterModification `json:"tierToArchive"`
	TierToCool                  *DateAfterModification `json:"tierToCool"`
}
