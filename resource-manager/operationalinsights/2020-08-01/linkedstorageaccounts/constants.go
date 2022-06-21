package linkedstorageaccounts

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSourceType string

const (
	DataSourceTypeAlerts      DataSourceType = "Alerts"
	DataSourceTypeAzureWatson DataSourceType = "AzureWatson"
	DataSourceTypeCustomLogs  DataSourceType = "CustomLogs"
	DataSourceTypeIngestion   DataSourceType = "Ingestion"
	DataSourceTypeQuery       DataSourceType = "Query"
)

func PossibleValuesForDataSourceType() []string {
	return []string{
		string(DataSourceTypeAlerts),
		string(DataSourceTypeAzureWatson),
		string(DataSourceTypeCustomLogs),
		string(DataSourceTypeIngestion),
		string(DataSourceTypeQuery),
	}
}

func parseDataSourceType(input string) (*DataSourceType, error) {
	vals := map[string]DataSourceType{
		"alerts":      DataSourceTypeAlerts,
		"azurewatson": DataSourceTypeAzureWatson,
		"customlogs":  DataSourceTypeCustomLogs,
		"ingestion":   DataSourceTypeIngestion,
		"query":       DataSourceTypeQuery,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataSourceType(input)
	return &out, nil
}
