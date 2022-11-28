package metadata

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetadataProperties struct {
	Author                   *MetadataAuthor       `json:"author"`
	Categories               *MetadataCategories   `json:"categories"`
	ContentId                *string               `json:"contentId,omitempty"`
	ContentSchemaVersion     *string               `json:"contentSchemaVersion,omitempty"`
	CustomVersion            *string               `json:"customVersion,omitempty"`
	Dependencies             *MetadataDependencies `json:"dependencies"`
	FirstPublishDate         *string               `json:"firstPublishDate,omitempty"`
	Icon                     *string               `json:"icon,omitempty"`
	Kind                     Kind                  `json:"kind"`
	LastPublishDate          *string               `json:"lastPublishDate,omitempty"`
	ParentId                 string                `json:"parentId"`
	PreviewImages            *[]string             `json:"previewImages,omitempty"`
	PreviewImagesDark        *[]string             `json:"previewImagesDark,omitempty"`
	Providers                *[]string             `json:"providers,omitempty"`
	Source                   *MetadataSource       `json:"source"`
	Support                  *MetadataSupport      `json:"support"`
	ThreatAnalysisTactics    *[]string             `json:"threatAnalysisTactics,omitempty"`
	ThreatAnalysisTechniques *[]string             `json:"threatAnalysisTechniques,omitempty"`
	Version                  *string               `json:"version,omitempty"`
}
