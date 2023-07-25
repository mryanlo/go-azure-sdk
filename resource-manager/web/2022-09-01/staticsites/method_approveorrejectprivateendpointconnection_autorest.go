package staticsites

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApproveOrRejectPrivateEndpointConnectionOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ApproveOrRejectPrivateEndpointConnection ...
func (c StaticSitesClient) ApproveOrRejectPrivateEndpointConnection(ctx context.Context, id StaticSitePrivateEndpointConnectionId, input PrivateLinkConnectionApprovalRequestResource) (result ApproveOrRejectPrivateEndpointConnectionOperationResponse, err error) {
	req, err := c.preparerForApproveOrRejectPrivateEndpointConnection(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ApproveOrRejectPrivateEndpointConnection", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForApproveOrRejectPrivateEndpointConnection(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ApproveOrRejectPrivateEndpointConnection", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ApproveOrRejectPrivateEndpointConnectionThenPoll performs ApproveOrRejectPrivateEndpointConnection then polls until it's completed
func (c StaticSitesClient) ApproveOrRejectPrivateEndpointConnectionThenPoll(ctx context.Context, id StaticSitePrivateEndpointConnectionId, input PrivateLinkConnectionApprovalRequestResource) error {
	result, err := c.ApproveOrRejectPrivateEndpointConnection(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ApproveOrRejectPrivateEndpointConnection: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ApproveOrRejectPrivateEndpointConnection: %+v", err)
	}

	return nil
}

// preparerForApproveOrRejectPrivateEndpointConnection prepares the ApproveOrRejectPrivateEndpointConnection request.
func (c StaticSitesClient) preparerForApproveOrRejectPrivateEndpointConnection(ctx context.Context, id StaticSitePrivateEndpointConnectionId, input PrivateLinkConnectionApprovalRequestResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForApproveOrRejectPrivateEndpointConnection sends the ApproveOrRejectPrivateEndpointConnection request. The method will close the
// http.Response Body if it receives an error.
func (c StaticSitesClient) senderForApproveOrRejectPrivateEndpointConnection(ctx context.Context, req *http.Request) (future ApproveOrRejectPrivateEndpointConnectionOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}