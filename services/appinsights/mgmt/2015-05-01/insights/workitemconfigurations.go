package insights

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

// WorkItemConfigurationsClient is the composite Swagger for Application Insights Management Client
type WorkItemConfigurationsClient struct {
	BaseClient
}

// NewWorkItemConfigurationsClient creates an instance of the WorkItemConfigurationsClient client.
func NewWorkItemConfigurationsClient(subscriptionID string) WorkItemConfigurationsClient {
	return NewWorkItemConfigurationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkItemConfigurationsClientWithBaseURI creates an instance of the WorkItemConfigurationsClient client.
func NewWorkItemConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) WorkItemConfigurationsClient {
	return WorkItemConfigurationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create a work item configuration for an Application Insights component.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights
// component resource. workItemConfigurationProperties is properties that need to be specified to create a work
// item configuration of a Application Insights component.
func (client WorkItemConfigurationsClient) Create(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigurationProperties WorkItemCreateConfiguration) (result WorkItemConfiguration, err error) {
	req, err := client.CreatePreparer(ctx, resourceGroupName, resourceName, workItemConfigurationProperties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client WorkItemConfigurationsClient) CreatePreparer(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigurationProperties WorkItemCreateConfiguration) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/WorkItemConfigs", pathParameters),
		autorest.WithJSON(workItemConfigurationProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client WorkItemConfigurationsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client WorkItemConfigurationsClient) CreateResponder(resp *http.Response) (result WorkItemConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete an workitem configuration of an Application Insights component.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights
// component resource. workItemConfigID is the unique work item configuration Id. This can be either friendly name
// of connector as defined in connector configuration
func (client WorkItemConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigID string) (result SetObject, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, resourceName, workItemConfigID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client WorkItemConfigurationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workItemConfigId":  autorest.Encode("path", workItemConfigID),
	}

	const APIVersion = "2015-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/WorkItemConfigs/{workItemConfigId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client WorkItemConfigurationsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client WorkItemConfigurationsClient) DeleteResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDefault gets default work item configurations that exist for the application
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights
// component resource.
func (client WorkItemConfigurationsClient) GetDefault(ctx context.Context, resourceGroupName string, resourceName string) (result WorkItemConfiguration, err error) {
	req, err := client.GetDefaultPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "GetDefault", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDefaultSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "GetDefault", resp, "Failure sending request")
		return
	}

	result, err = client.GetDefaultResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "GetDefault", resp, "Failure responding to request")
	}

	return
}

// GetDefaultPreparer prepares the GetDefault request.
func (client WorkItemConfigurationsClient) GetDefaultPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/DefaultWorkItemConfig", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDefaultSender sends the GetDefault request. The method will close the
// http.Response Body if it receives an error.
func (client WorkItemConfigurationsClient) GetDefaultSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetDefaultResponder handles the response to the GetDefault request. The method always
// closes the http.Response Body.
func (client WorkItemConfigurationsClient) GetDefaultResponder(resp *http.Response) (result WorkItemConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets the list work item configurations that exist for the application
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights
// component resource.
func (client WorkItemConfigurationsClient) List(ctx context.Context, resourceGroupName string, resourceName string) (result ListWorkItemConfiguration, err error) {
	req, err := client.ListPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client WorkItemConfigurationsClient) ListPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/WorkItemConfigs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client WorkItemConfigurationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WorkItemConfigurationsClient) ListResponder(resp *http.Response) (result ListWorkItemConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
