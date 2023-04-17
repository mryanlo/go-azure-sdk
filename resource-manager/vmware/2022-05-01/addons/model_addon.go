package addons

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Addon struct {
	Id         *string          `json:"id,omitempty"`
	Name       *string          `json:"name,omitempty"`
	Properties *AddonProperties `json:"properties,omitempty"`
	Type       *string          `json:"type,omitempty"`
}

var _ json.Unmarshaler = &Addon{}

func (s *Addon) UnmarshalJSON(bytes []byte) error {
	type alias Addon
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into Addon: %+v", err)
	}

	s.Id = decoded.Id
	s.Name = decoded.Name
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Addon into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["properties"]; ok {
		impl, err := unmarshalAddonPropertiesImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Properties' for 'Addon': %+v", err)
		}
		s.Properties = &impl
	}
	return nil
}
