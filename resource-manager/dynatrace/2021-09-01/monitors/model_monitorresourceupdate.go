package monitors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitorResourceUpdate struct {
	DynatraceEnvironmentProperties *DynatraceEnvironmentProperties `json:"dynatraceEnvironmentProperties"`
	MarketplaceSubscriptionStatus  *MarketplaceSubscriptionStatus  `json:"marketplaceSubscriptionStatus,omitempty"`
	MonitoringStatus               *MonitoringStatus               `json:"monitoringStatus,omitempty"`
	PlanData                       *PlanData                       `json:"planData"`
	Tags                           *map[string]string              `json:"tags,omitempty"`
	UserInfo                       *UserInfo                       `json:"userInfo"`
}
