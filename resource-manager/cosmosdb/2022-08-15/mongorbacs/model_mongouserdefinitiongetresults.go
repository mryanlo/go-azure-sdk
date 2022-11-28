package mongorbacs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MongoUserDefinitionGetResults struct {
	Id         *string                      `json:"id,omitempty"`
	Name       *string                      `json:"name,omitempty"`
	Properties *MongoUserDefinitionResource `json:"properties"`
	Type       *string                      `json:"type,omitempty"`
}
