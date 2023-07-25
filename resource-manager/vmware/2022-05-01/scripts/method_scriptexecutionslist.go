package scripts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScriptExecutionsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ScriptExecution
}

type ScriptExecutionsListCompleteResult struct {
	Items []ScriptExecution
}

// ScriptExecutionsList ...
func (c ScriptsClient) ScriptExecutionsList(ctx context.Context, id PrivateCloudId) (result ScriptExecutionsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/scriptExecutions", id.ID()),
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
		Values *[]ScriptExecution `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ScriptExecutionsListComplete retrieves all the results into a single object
func (c ScriptsClient) ScriptExecutionsListComplete(ctx context.Context, id PrivateCloudId) (ScriptExecutionsListCompleteResult, error) {
	return c.ScriptExecutionsListCompleteMatchingPredicate(ctx, id, ScriptExecutionOperationPredicate{})
}

// ScriptExecutionsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ScriptsClient) ScriptExecutionsListCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate ScriptExecutionOperationPredicate) (result ScriptExecutionsListCompleteResult, err error) {
	items := make([]ScriptExecution, 0)

	resp, err := c.ScriptExecutionsList(ctx, id)
	if err != nil {
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

	result = ScriptExecutionsListCompleteResult{
		Items: items,
	}
	return
}