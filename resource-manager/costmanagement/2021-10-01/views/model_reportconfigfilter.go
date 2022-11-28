package views

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReportConfigFilter struct {
	And        *[]ReportConfigFilter             `json:"and,omitempty"`
	Dimensions *ReportConfigComparisonExpression `json:"dimensions"`
	Or         *[]ReportConfigFilter             `json:"or,omitempty"`
	Tags       *ReportConfigComparisonExpression `json:"tags"`
}
