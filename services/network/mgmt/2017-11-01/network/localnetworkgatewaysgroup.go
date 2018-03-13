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
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// LocalNetworkGatewaysGroupClient is the network Client
type LocalNetworkGatewaysGroupClient struct {
	BaseClient
}

// NewLocalNetworkGatewaysGroupClient creates an instance of the LocalNetworkGatewaysGroupClient client.
func NewLocalNetworkGatewaysGroupClient(subscriptionID string) LocalNetworkGatewaysGroupClient {
	return NewLocalNetworkGatewaysGroupClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLocalNetworkGatewaysGroupClientWithBaseURI creates an instance of the LocalNetworkGatewaysGroupClient client.
func NewLocalNetworkGatewaysGroupClientWithBaseURI(baseURI string, subscriptionID string) LocalNetworkGatewaysGroupClient {
	return LocalNetworkGatewaysGroupClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a local network gateway in the specified resource group.
//
// resourceGroupName is the name of the resource group. localNetworkGatewayName is the name of the local network
// gateway. parameters is parameters supplied to the create or update local network gateway operation.
func (client LocalNetworkGatewaysGroupClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters LocalNetworkGateway) (result LocalNetworkGatewaysGroupCreateOrUpdateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: localNetworkGatewayName,
			Constraints: []validation.Constraint{{Target: "localNetworkGatewayName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.LocalNetworkGatewayPropertiesFormat", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("network.LocalNetworkGatewaysGroupClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, localNetworkGatewayName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysGroupClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysGroupClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client LocalNetworkGatewaysGroupClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters LocalNetworkGateway) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"localNetworkGatewayName": autorest.Encode("path", localNetworkGatewayName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client LocalNetworkGatewaysGroupClient) CreateOrUpdateSender(req *http.Request) (future LocalNetworkGatewaysGroupCreateOrUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated))
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client LocalNetworkGatewaysGroupClient) CreateOrUpdateResponder(resp *http.Response) (result LocalNetworkGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified local network gateway.
//
// resourceGroupName is the name of the resource group. localNetworkGatewayName is the name of the local network
// gateway.
func (client LocalNetworkGatewaysGroupClient) Delete(ctx context.Context, resourceGroupName string, localNetworkGatewayName string) (result LocalNetworkGatewaysGroupDeleteFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: localNetworkGatewayName,
			Constraints: []validation.Constraint{{Target: "localNetworkGatewayName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("network.LocalNetworkGatewaysGroupClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, localNetworkGatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysGroupClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysGroupClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client LocalNetworkGatewaysGroupClient) DeletePreparer(ctx context.Context, resourceGroupName string, localNetworkGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"localNetworkGatewayName": autorest.Encode("path", localNetworkGatewayName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client LocalNetworkGatewaysGroupClient) DeleteSender(req *http.Request) (future LocalNetworkGatewaysGroupDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client LocalNetworkGatewaysGroupClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified local network gateway in a resource group.
//
// resourceGroupName is the name of the resource group. localNetworkGatewayName is the name of the local network
// gateway.
func (client LocalNetworkGatewaysGroupClient) Get(ctx context.Context, resourceGroupName string, localNetworkGatewayName string) (result LocalNetworkGateway, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: localNetworkGatewayName,
			Constraints: []validation.Constraint{{Target: "localNetworkGatewayName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("network.LocalNetworkGatewaysGroupClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, localNetworkGatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysGroupClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysGroupClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysGroupClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client LocalNetworkGatewaysGroupClient) GetPreparer(ctx context.Context, resourceGroupName string, localNetworkGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"localNetworkGatewayName": autorest.Encode("path", localNetworkGatewayName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LocalNetworkGatewaysGroupClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LocalNetworkGatewaysGroupClient) GetResponder(resp *http.Response) (result LocalNetworkGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all the local network gateways in a resource group.
//
// resourceGroupName is the name of the resource group.
func (client LocalNetworkGatewaysGroupClient) List(ctx context.Context, resourceGroupName string) (result LocalNetworkGatewayListResultPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysGroupClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lnglr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysGroupClient", "List", resp, "Failure sending request")
		return
	}

	result.lnglr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysGroupClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client LocalNetworkGatewaysGroupClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LocalNetworkGatewaysGroupClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LocalNetworkGatewaysGroupClient) ListResponder(resp *http.Response) (result LocalNetworkGatewayListResult, err error) {
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
func (client LocalNetworkGatewaysGroupClient) listNextResults(lastResults LocalNetworkGatewayListResult) (result LocalNetworkGatewayListResult, err error) {
	req, err := lastResults.localNetworkGatewayListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysGroupClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysGroupClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysGroupClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client LocalNetworkGatewaysGroupClient) ListComplete(ctx context.Context, resourceGroupName string) (result LocalNetworkGatewayListResultIterator, err error) {
	result.page, err = client.List(ctx, resourceGroupName)
	return
}

// UpdateTags updates a local network gateway tags.
//
// resourceGroupName is the name of the resource group. localNetworkGatewayName is the name of the local network
// gateway. parameters is parameters supplied to update local network gateway tags.
func (client LocalNetworkGatewaysGroupClient) UpdateTags(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters TagsObject) (result LocalNetworkGatewaysGroupUpdateTagsFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: localNetworkGatewayName,
			Constraints: []validation.Constraint{{Target: "localNetworkGatewayName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("network.LocalNetworkGatewaysGroupClient", "UpdateTags", err.Error())
	}

	req, err := client.UpdateTagsPreparer(ctx, resourceGroupName, localNetworkGatewayName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysGroupClient", "UpdateTags", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateTagsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysGroupClient", "UpdateTags", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdateTagsPreparer prepares the UpdateTags request.
func (client LocalNetworkGatewaysGroupClient) UpdateTagsPreparer(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters TagsObject) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"localNetworkGatewayName": autorest.Encode("path", localNetworkGatewayName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateTagsSender sends the UpdateTags request. The method will close the
// http.Response Body if it receives an error.
func (client LocalNetworkGatewaysGroupClient) UpdateTagsSender(req *http.Request) (future LocalNetworkGatewaysGroupUpdateTagsFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK))
	return
}

// UpdateTagsResponder handles the response to the UpdateTags request. The method always
// closes the http.Response Body.
func (client LocalNetworkGatewaysGroupClient) UpdateTagsResponder(resp *http.Response) (result LocalNetworkGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
