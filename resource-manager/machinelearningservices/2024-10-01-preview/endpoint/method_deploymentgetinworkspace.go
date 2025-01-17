package endpoint

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentGetInWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]EndpointDeploymentResourcePropertiesBasicResource
}

type DeploymentGetInWorkspaceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []EndpointDeploymentResourcePropertiesBasicResource
}

type DeploymentGetInWorkspaceOperationOptions struct {
	EndpointType *EndpointType
	Skip         *string
}

func DefaultDeploymentGetInWorkspaceOperationOptions() DeploymentGetInWorkspaceOperationOptions {
	return DeploymentGetInWorkspaceOperationOptions{}
}

func (o DeploymentGetInWorkspaceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DeploymentGetInWorkspaceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeploymentGetInWorkspaceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.EndpointType != nil {
		out.Append("endpointType", fmt.Sprintf("%v", *o.EndpointType))
	}
	if o.Skip != nil {
		out.Append("$skip", fmt.Sprintf("%v", *o.Skip))
	}
	return &out
}

type DeploymentGetInWorkspaceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DeploymentGetInWorkspaceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DeploymentGetInWorkspace ...
func (c EndpointClient) DeploymentGetInWorkspace(ctx context.Context, id WorkspaceId, options DeploymentGetInWorkspaceOperationOptions) (result DeploymentGetInWorkspaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &DeploymentGetInWorkspaceCustomPager{},
		Path:          fmt.Sprintf("%s/deployments", id.ID()),
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
		Values *[]EndpointDeploymentResourcePropertiesBasicResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// DeploymentGetInWorkspaceComplete retrieves all the results into a single object
func (c EndpointClient) DeploymentGetInWorkspaceComplete(ctx context.Context, id WorkspaceId, options DeploymentGetInWorkspaceOperationOptions) (DeploymentGetInWorkspaceCompleteResult, error) {
	return c.DeploymentGetInWorkspaceCompleteMatchingPredicate(ctx, id, options, EndpointDeploymentResourcePropertiesBasicResourceOperationPredicate{})
}

// DeploymentGetInWorkspaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EndpointClient) DeploymentGetInWorkspaceCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, options DeploymentGetInWorkspaceOperationOptions, predicate EndpointDeploymentResourcePropertiesBasicResourceOperationPredicate) (result DeploymentGetInWorkspaceCompleteResult, err error) {
	items := make([]EndpointDeploymentResourcePropertiesBasicResource, 0)

	resp, err := c.DeploymentGetInWorkspace(ctx, id, options)
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

	result = DeploymentGetInWorkspaceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
