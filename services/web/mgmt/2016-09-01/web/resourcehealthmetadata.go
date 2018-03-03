package web

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

// ResourceHealthMetadataClient is the webSite Management Client
type ResourceHealthMetadataClient struct {
	BaseClient
}

// NewResourceHealthMetadataClient creates an instance of the ResourceHealthMetadataClient client.
func NewResourceHealthMetadataClient(subscriptionID string) ResourceHealthMetadataClient {
	return NewResourceHealthMetadataClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewResourceHealthMetadataClientWithBaseURI creates an instance of the ResourceHealthMetadataClient client.
func NewResourceHealthMetadataClientWithBaseURI(baseURI string, subscriptionID string) ResourceHealthMetadataClient {
	return ResourceHealthMetadataClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetBySite gets the category of ResourceHealthMetadata to use for the given site
//
// resourceGroupName is name of the resource group to which the resource belongs. name is name of web app
func (client ResourceHealthMetadataClient) GetBySite(ctx context.Context, resourceGroupName string, name string) (result ResourceHealthMetadataCollectionPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.ResourceHealthMetadataClient", "GetBySite", err.Error())
	}

	result.fn = client.getBySiteNextResults
	req, err := client.GetBySitePreparer(ctx, resourceGroupName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "GetBySite", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetBySiteSender(req)
	if err != nil {
		result.rhmc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "GetBySite", resp, "Failure sending request")
		return
	}

	result.rhmc, err = client.GetBySiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "GetBySite", resp, "Failure responding to request")
	}

	return
}

// GetBySitePreparer prepares the GetBySite request.
func (client ResourceHealthMetadataClient) GetBySitePreparer(ctx context.Context, resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/resourceHealthMetadata/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetBySiteSender sends the GetBySite request. The method will close the
// http.Response Body if it receives an error.
func (client ResourceHealthMetadataClient) GetBySiteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetBySiteResponder handles the response to the GetBySite request. The method always
// closes the http.Response Body.
func (client ResourceHealthMetadataClient) GetBySiteResponder(resp *http.Response) (result ResourceHealthMetadataCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getBySiteNextResults retrieves the next set of results, if any.
func (client ResourceHealthMetadataClient) getBySiteNextResults(lastResults ResourceHealthMetadataCollection) (result ResourceHealthMetadataCollection, err error) {
	req, err := lastResults.resourceHealthMetadataCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "getBySiteNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetBySiteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "getBySiteNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetBySiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "getBySiteNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetBySiteComplete enumerates all values, automatically crossing page boundaries as required.
func (client ResourceHealthMetadataClient) GetBySiteComplete(ctx context.Context, resourceGroupName string, name string) (result ResourceHealthMetadataCollectionIterator, err error) {
	result.page, err = client.GetBySite(ctx, resourceGroupName, name)
	return
}

// GetBySiteSlot gets the category of ResourceHealthMetadata to use for the given site
//
// resourceGroupName is name of the resource group to which the resource belongs. name is name of web app slot is
// name of web app slot. If not specified then will default to production slot.
func (client ResourceHealthMetadataClient) GetBySiteSlot(ctx context.Context, resourceGroupName string, name string, slot string) (result ResourceHealthMetadataCollectionPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.ResourceHealthMetadataClient", "GetBySiteSlot", err.Error())
	}

	result.fn = client.getBySiteSlotNextResults
	req, err := client.GetBySiteSlotPreparer(ctx, resourceGroupName, name, slot)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "GetBySiteSlot", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetBySiteSlotSender(req)
	if err != nil {
		result.rhmc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "GetBySiteSlot", resp, "Failure sending request")
		return
	}

	result.rhmc, err = client.GetBySiteSlotResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "GetBySiteSlot", resp, "Failure responding to request")
	}

	return
}

// GetBySiteSlotPreparer prepares the GetBySiteSlot request.
func (client ResourceHealthMetadataClient) GetBySiteSlotPreparer(ctx context.Context, resourceGroupName string, name string, slot string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"slot":              autorest.Encode("path", slot),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/resourceHealthMetadata/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetBySiteSlotSender sends the GetBySiteSlot request. The method will close the
// http.Response Body if it receives an error.
func (client ResourceHealthMetadataClient) GetBySiteSlotSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetBySiteSlotResponder handles the response to the GetBySiteSlot request. The method always
// closes the http.Response Body.
func (client ResourceHealthMetadataClient) GetBySiteSlotResponder(resp *http.Response) (result ResourceHealthMetadataCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getBySiteSlotNextResults retrieves the next set of results, if any.
func (client ResourceHealthMetadataClient) getBySiteSlotNextResults(lastResults ResourceHealthMetadataCollection) (result ResourceHealthMetadataCollection, err error) {
	req, err := lastResults.resourceHealthMetadataCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "getBySiteSlotNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetBySiteSlotSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "getBySiteSlotNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetBySiteSlotResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "getBySiteSlotNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetBySiteSlotComplete enumerates all values, automatically crossing page boundaries as required.
func (client ResourceHealthMetadataClient) GetBySiteSlotComplete(ctx context.Context, resourceGroupName string, name string, slot string) (result ResourceHealthMetadataCollectionIterator, err error) {
	result.page, err = client.GetBySiteSlot(ctx, resourceGroupName, name, slot)
	return
}

// List list all ResourceHealthMetadata for all sites in the subscription.
func (client ResourceHealthMetadataClient) List(ctx context.Context) (result ResourceHealthMetadataCollectionPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rhmc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "List", resp, "Failure sending request")
		return
	}

	result.rhmc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ResourceHealthMetadataClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Web/resourceHealthMetadata", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ResourceHealthMetadataClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ResourceHealthMetadataClient) ListResponder(resp *http.Response) (result ResourceHealthMetadataCollection, err error) {
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
func (client ResourceHealthMetadataClient) listNextResults(lastResults ResourceHealthMetadataCollection) (result ResourceHealthMetadataCollection, err error) {
	req, err := lastResults.resourceHealthMetadataCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ResourceHealthMetadataClient) ListComplete(ctx context.Context) (result ResourceHealthMetadataCollectionIterator, err error) {
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup list all ResourceHealthMetadata for all sites in the resource group in the subscription.
//
// resourceGroupName is name of the resource group to which the resource belongs.
func (client ResourceHealthMetadataClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ResourceHealthMetadataCollectionPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.ResourceHealthMetadataClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.rhmc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.rhmc, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ResourceHealthMetadataClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/resourceHealthMetadata", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ResourceHealthMetadataClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ResourceHealthMetadataClient) ListByResourceGroupResponder(resp *http.Response) (result ResourceHealthMetadataCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ResourceHealthMetadataClient) listByResourceGroupNextResults(lastResults ResourceHealthMetadataCollection) (result ResourceHealthMetadataCollection, err error) {
	req, err := lastResults.resourceHealthMetadataCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ResourceHealthMetadataClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ResourceHealthMetadataCollectionIterator, err error) {
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// ListBySite gets the category of ResourceHealthMetadata to use for the given site as a collection
//
// resourceGroupName is name of the resource group to which the resource belongs. name is name of web app.
func (client ResourceHealthMetadataClient) ListBySite(ctx context.Context, resourceGroupName string, name string) (result ResourceHealthMetadataCollectionPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.ResourceHealthMetadataClient", "ListBySite", err.Error())
	}

	result.fn = client.listBySiteNextResults
	req, err := client.ListBySitePreparer(ctx, resourceGroupName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "ListBySite", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySiteSender(req)
	if err != nil {
		result.rhmc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "ListBySite", resp, "Failure sending request")
		return
	}

	result.rhmc, err = client.ListBySiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "ListBySite", resp, "Failure responding to request")
	}

	return
}

// ListBySitePreparer prepares the ListBySite request.
func (client ResourceHealthMetadataClient) ListBySitePreparer(ctx context.Context, resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/resourceHealthMetadata", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySiteSender sends the ListBySite request. The method will close the
// http.Response Body if it receives an error.
func (client ResourceHealthMetadataClient) ListBySiteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListBySiteResponder handles the response to the ListBySite request. The method always
// closes the http.Response Body.
func (client ResourceHealthMetadataClient) ListBySiteResponder(resp *http.Response) (result ResourceHealthMetadataCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySiteNextResults retrieves the next set of results, if any.
func (client ResourceHealthMetadataClient) listBySiteNextResults(lastResults ResourceHealthMetadataCollection) (result ResourceHealthMetadataCollection, err error) {
	req, err := lastResults.resourceHealthMetadataCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listBySiteNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySiteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listBySiteNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listBySiteNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySiteComplete enumerates all values, automatically crossing page boundaries as required.
func (client ResourceHealthMetadataClient) ListBySiteComplete(ctx context.Context, resourceGroupName string, name string) (result ResourceHealthMetadataCollectionIterator, err error) {
	result.page, err = client.ListBySite(ctx, resourceGroupName, name)
	return
}

// ListBySiteSlot gets the category of ResourceHealthMetadata to use for the given site as a collection
//
// resourceGroupName is name of the resource group to which the resource belongs. name is name of web app. slot is
// name of web app slot. If not specified then will default to production slot.
func (client ResourceHealthMetadataClient) ListBySiteSlot(ctx context.Context, resourceGroupName string, name string, slot string) (result ResourceHealthMetadataCollectionPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.ResourceHealthMetadataClient", "ListBySiteSlot", err.Error())
	}

	result.fn = client.listBySiteSlotNextResults
	req, err := client.ListBySiteSlotPreparer(ctx, resourceGroupName, name, slot)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "ListBySiteSlot", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySiteSlotSender(req)
	if err != nil {
		result.rhmc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "ListBySiteSlot", resp, "Failure sending request")
		return
	}

	result.rhmc, err = client.ListBySiteSlotResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "ListBySiteSlot", resp, "Failure responding to request")
	}

	return
}

// ListBySiteSlotPreparer prepares the ListBySiteSlot request.
func (client ResourceHealthMetadataClient) ListBySiteSlotPreparer(ctx context.Context, resourceGroupName string, name string, slot string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"slot":              autorest.Encode("path", slot),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/resourceHealthMetadata", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySiteSlotSender sends the ListBySiteSlot request. The method will close the
// http.Response Body if it receives an error.
func (client ResourceHealthMetadataClient) ListBySiteSlotSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListBySiteSlotResponder handles the response to the ListBySiteSlot request. The method always
// closes the http.Response Body.
func (client ResourceHealthMetadataClient) ListBySiteSlotResponder(resp *http.Response) (result ResourceHealthMetadataCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySiteSlotNextResults retrieves the next set of results, if any.
func (client ResourceHealthMetadataClient) listBySiteSlotNextResults(lastResults ResourceHealthMetadataCollection) (result ResourceHealthMetadataCollection, err error) {
	req, err := lastResults.resourceHealthMetadataCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listBySiteSlotNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySiteSlotSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listBySiteSlotNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySiteSlotResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listBySiteSlotNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySiteSlotComplete enumerates all values, automatically crossing page boundaries as required.
func (client ResourceHealthMetadataClient) ListBySiteSlotComplete(ctx context.Context, resourceGroupName string, name string, slot string) (result ResourceHealthMetadataCollectionIterator, err error) {
	result.page, err = client.ListBySiteSlot(ctx, resourceGroupName, name, slot)
	return
}
