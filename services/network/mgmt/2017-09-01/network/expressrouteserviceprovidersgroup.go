package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ExpressRouteServiceProvidersGroupClient is the network Client
type ExpressRouteServiceProvidersGroupClient struct {
	BaseClient
}

// NewExpressRouteServiceProvidersGroupClient creates an instance of the ExpressRouteServiceProvidersGroupClient
// client.
func NewExpressRouteServiceProvidersGroupClient(subscriptionID string) ExpressRouteServiceProvidersGroupClient {
	return NewExpressRouteServiceProvidersGroupClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExpressRouteServiceProvidersGroupClientWithBaseURI creates an instance of the
// ExpressRouteServiceProvidersGroupClient client.
func NewExpressRouteServiceProvidersGroupClientWithBaseURI(baseURI string, subscriptionID string) ExpressRouteServiceProvidersGroupClient {
	return ExpressRouteServiceProvidersGroupClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets all the available express route service providers.
func (client ExpressRouteServiceProvidersGroupClient) List(ctx context.Context) (result ExpressRouteServiceProviderListResultPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteServiceProvidersGroupClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ersplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteServiceProvidersGroupClient", "List", resp, "Failure sending request")
		return
	}

	result.ersplr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteServiceProvidersGroupClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ExpressRouteServiceProvidersGroupClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteServiceProviders", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteServiceProvidersGroupClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExpressRouteServiceProvidersGroupClient) ListResponder(resp *http.Response) (result ExpressRouteServiceProviderListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ExpressRouteServiceProvidersGroupClient) listNextResults(lastResults ExpressRouteServiceProviderListResult) (result ExpressRouteServiceProviderListResult, err error) {
	req, err := lastResults.expressRouteServiceProviderListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteServiceProvidersGroupClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteServiceProvidersGroupClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteServiceProvidersGroupClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ExpressRouteServiceProvidersGroupClient) ListComplete(ctx context.Context) (result ExpressRouteServiceProviderListResultIterator, err error) {
	result.page, err = client.List(ctx)
	return
}
