package alertruletemplates

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AlertRuleTemplate = ThreatIntelligenceAlertRuleTemplate{}

type ThreatIntelligenceAlertRuleTemplate struct {
	Properties *ThreatIntelligenceAlertRuleTemplateProperties `json:"properties"`

	// Fields inherited from AlertRuleTemplate
	Id         *string                `json:"id,omitempty"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

var _ json.Marshaler = ThreatIntelligenceAlertRuleTemplate{}

func (s ThreatIntelligenceAlertRuleTemplate) MarshalJSON() ([]byte, error) {
	type wrapper ThreatIntelligenceAlertRuleTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ThreatIntelligenceAlertRuleTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ThreatIntelligenceAlertRuleTemplate: %+v", err)
	}
	decoded["kind"] = "ThreatIntelligence"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ThreatIntelligenceAlertRuleTemplate: %+v", err)
	}

	return encoded, nil
}
