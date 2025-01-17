package privatelink

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionsListByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateEndpointConnection
}

type PrivateEndpointConnectionsListByWorkspaceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PrivateEndpointConnection
}

type PrivateEndpointConnectionsListByWorkspaceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *PrivateEndpointConnectionsListByWorkspaceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// PrivateEndpointConnectionsListByWorkspace ...
func (c PrivateLinkClient) PrivateEndpointConnectionsListByWorkspace(ctx context.Context, id WorkspaceId) (result PrivateEndpointConnectionsListByWorkspaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &PrivateEndpointConnectionsListByWorkspaceCustomPager{},
		Path:       fmt.Sprintf("%s/privateEndpointConnections", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]PrivateEndpointConnection `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// PrivateEndpointConnectionsListByWorkspaceComplete retrieves all the results into a single object
func (c PrivateLinkClient) PrivateEndpointConnectionsListByWorkspaceComplete(ctx context.Context, id WorkspaceId) (PrivateEndpointConnectionsListByWorkspaceCompleteResult, error) {
	return c.PrivateEndpointConnectionsListByWorkspaceCompleteMatchingPredicate(ctx, id, PrivateEndpointConnectionOperationPredicate{})
}

// PrivateEndpointConnectionsListByWorkspaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateLinkClient) PrivateEndpointConnectionsListByWorkspaceCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate PrivateEndpointConnectionOperationPredicate) (result PrivateEndpointConnectionsListByWorkspaceCompleteResult, err error) {
	items := make([]PrivateEndpointConnection, 0)

	resp, err := c.PrivateEndpointConnectionsListByWorkspace(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = PrivateEndpointConnectionsListByWorkspaceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
