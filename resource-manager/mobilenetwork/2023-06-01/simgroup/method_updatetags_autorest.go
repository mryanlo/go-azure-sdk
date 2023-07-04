package simgroup

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateTagsOperationResponse struct {
	HttpResponse *http.Response
	Model        *SimGroup
}

// UpdateTags ...
func (c SIMGroupClient) UpdateTags(ctx context.Context, id SimGroupId, input IdentityAndTagsObject) (result UpdateTagsOperationResponse, err error) {
	req, err := c.preparerForUpdateTags(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "simgroup.SIMGroupClient", "UpdateTags", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "simgroup.SIMGroupClient", "UpdateTags", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateTags(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "simgroup.SIMGroupClient", "UpdateTags", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateTags prepares the UpdateTags request.
func (c SIMGroupClient) preparerForUpdateTags(ctx context.Context, id SimGroupId, input IdentityAndTagsObject) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateTags handles the response to the UpdateTags request. The method always
// closes the http.Response Body.
func (c SIMGroupClient) responderForUpdateTags(resp *http.Response) (result UpdateTagsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}