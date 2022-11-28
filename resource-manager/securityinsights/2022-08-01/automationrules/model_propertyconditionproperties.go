package automationrules

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AutomationRuleCondition = PropertyConditionProperties{}

type PropertyConditionProperties struct {
	ConditionProperties *AutomationRulePropertyValuesCondition `json:"conditionProperties"`

	// Fields inherited from AutomationRuleCondition
}

var _ json.Marshaler = PropertyConditionProperties{}

func (s PropertyConditionProperties) MarshalJSON() ([]byte, error) {
	type wrapper PropertyConditionProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PropertyConditionProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PropertyConditionProperties: %+v", err)
	}
	decoded["conditionType"] = "Property"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PropertyConditionProperties: %+v", err)
	}

	return encoded, nil
}
