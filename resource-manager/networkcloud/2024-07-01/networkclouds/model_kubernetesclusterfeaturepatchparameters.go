package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KubernetesClusterFeaturePatchParameters struct {
	Properties *KubernetesClusterFeaturePatchProperties `json:"properties,omitempty"`
	Tags       *map[string]string                       `json:"tags,omitempty"`
}
