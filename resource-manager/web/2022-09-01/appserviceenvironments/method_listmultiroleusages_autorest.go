package appserviceenvironments

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMultiRoleUsagesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Usage

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListMultiRoleUsagesOperationResponse, error)
}

type ListMultiRoleUsagesCompleteResult struct {
	Items []Usage
}

func (r ListMultiRoleUsagesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListMultiRoleUsagesOperationResponse) LoadMore(ctx context.Context) (resp ListMultiRoleUsagesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListMultiRoleUsages ...
func (c AppServiceEnvironmentsClient) ListMultiRoleUsages(ctx context.Context, id HostingEnvironmentId) (resp ListMultiRoleUsagesOperationResponse, err error) {
	req, err := c.preparerForListMultiRoleUsages(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRoleUsages", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRoleUsages", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListMultiRoleUsages(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRoleUsages", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListMultiRoleUsages prepares the ListMultiRoleUsages request.
func (c AppServiceEnvironmentsClient) preparerForListMultiRoleUsages(ctx context.Context, id HostingEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/multiRolePools/default/usages", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListMultiRoleUsagesWithNextLink prepares the ListMultiRoleUsages request with the given nextLink token.
func (c AppServiceEnvironmentsClient) preparerForListMultiRoleUsagesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListMultiRoleUsages handles the response to the ListMultiRoleUsages request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForListMultiRoleUsages(resp *http.Response) (result ListMultiRoleUsagesOperationResponse, err error) {
	type page struct {
		Values   []Usage `json:"value"`
		NextLink *string `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListMultiRoleUsagesOperationResponse, err error) {
			req, err := c.preparerForListMultiRoleUsagesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRoleUsages", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRoleUsages", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListMultiRoleUsages(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRoleUsages", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListMultiRoleUsagesComplete retrieves all of the results into a single object
func (c AppServiceEnvironmentsClient) ListMultiRoleUsagesComplete(ctx context.Context, id HostingEnvironmentId) (ListMultiRoleUsagesCompleteResult, error) {
	return c.ListMultiRoleUsagesCompleteMatchingPredicate(ctx, id, UsageOperationPredicate{})
}

// ListMultiRoleUsagesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppServiceEnvironmentsClient) ListMultiRoleUsagesCompleteMatchingPredicate(ctx context.Context, id HostingEnvironmentId, predicate UsageOperationPredicate) (resp ListMultiRoleUsagesCompleteResult, err error) {
	items := make([]Usage, 0)

	page, err := c.ListMultiRoleUsages(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := ListMultiRoleUsagesCompleteResult{
		Items: items,
	}
	return out, nil
}