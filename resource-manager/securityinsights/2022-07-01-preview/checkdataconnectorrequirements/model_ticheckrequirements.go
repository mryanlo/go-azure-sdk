package checkdataconnectorrequirements

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataConnectorsCheckRequirements = TICheckRequirements{}

type TICheckRequirements struct {
	Properties *DataConnectorTenantId `json:"properties"`

	// Fields inherited from DataConnectorsCheckRequirements
}

var _ json.Marshaler = TICheckRequirements{}

func (s TICheckRequirements) MarshalJSON() ([]byte, error) {
	type wrapper TICheckRequirements
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TICheckRequirements: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TICheckRequirements: %+v", err)
	}
	decoded["kind"] = "ThreatIntelligence"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TICheckRequirements: %+v", err)
	}

	return encoded, nil
}
