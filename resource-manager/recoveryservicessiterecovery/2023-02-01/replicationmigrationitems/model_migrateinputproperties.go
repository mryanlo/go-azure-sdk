package replicationmigrationitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrateInputProperties struct {
	ProviderSpecificDetails MigrateProviderSpecificInput `json:"providerSpecificDetails"`
}

var _ json.Unmarshaler = &MigrateInputProperties{}

func (s *MigrateInputProperties) UnmarshalJSON(bytes []byte) error {

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MigrateInputProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["providerSpecificDetails"]; ok {
		impl, err := unmarshalMigrateProviderSpecificInputImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ProviderSpecificDetails' for 'MigrateInputProperties': %+v", err)
		}
		s.ProviderSpecificDetails = &impl
	}
	return nil
}
