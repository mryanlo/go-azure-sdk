package operationalizationclusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HDInsightProperties struct {
	Address              *string                       `json:"address,omitempty"`
	AdministratorAccount *VirtualMachineSshCredentials `json:"administratorAccount"`
	SshPort              *int64                        `json:"sshPort,omitempty"`
}
