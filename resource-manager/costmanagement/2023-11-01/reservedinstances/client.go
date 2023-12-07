package reservedinstances

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservedInstancesClient struct {
	Client *resourcemanager.Client
}

func NewReservedInstancesClientWithBaseURI(sdkApi sdkEnv.Api) (*ReservedInstancesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "reservedinstances", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReservedInstancesClient: %+v", err)
	}

	return &ReservedInstancesClient{
		Client: client,
	}, nil
}
