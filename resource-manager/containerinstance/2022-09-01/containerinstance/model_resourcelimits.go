package containerinstance

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceLimits struct {
	Cpu        *float64     `json:"cpu,omitempty"`
	Gpu        *GpuResource `json:"gpu"`
	MemoryInGB *float64     `json:"memoryInGB,omitempty"`
}
