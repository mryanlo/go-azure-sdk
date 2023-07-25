package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplySlotConfigToProductionOperationResponse struct {
	HttpResponse *http.Response
}

// ApplySlotConfigToProduction ...
func (c WebAppsClient) ApplySlotConfigToProduction(ctx context.Context, id SiteId, input CsmSlotEntity) (result ApplySlotConfigToProductionOperationResponse, err error) {
	req, err := c.preparerForApplySlotConfigToProduction(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ApplySlotConfigToProduction", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ApplySlotConfigToProduction", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForApplySlotConfigToProduction(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ApplySlotConfigToProduction", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForApplySlotConfigToProduction prepares the ApplySlotConfigToProduction request.
func (c WebAppsClient) preparerForApplySlotConfigToProduction(ctx context.Context, id SiteId, input CsmSlotEntity) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/applySlotConfig", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForApplySlotConfigToProduction handles the response to the ApplySlotConfigToProduction request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForApplySlotConfigToProduction(resp *http.Response) (result ApplySlotConfigToProductionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}