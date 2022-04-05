//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcognitiveservices

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// AccountsClientCreatePollerResponse contains the response from method AccountsClient.Create.
type AccountsClientCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *AccountsClientCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l AccountsClientCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (AccountsClientCreateResponse, error) {
	respType := AccountsClientCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Account)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a AccountsClientCreatePollerResponse from the provided client and resume token.
func (l *AccountsClientCreatePollerResponse) Resume(ctx context.Context, client *AccountsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("AccountsClient.Create", token, client.pl)
	if err != nil {
		return err
	}
	poller := &AccountsClientCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// AccountsClientCreateResponse contains the response from method AccountsClient.Create.
type AccountsClientCreateResponse struct {
	AccountsClientCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsClientCreateResult contains the result from method AccountsClient.Create.
type AccountsClientCreateResult struct {
	Account
}

// AccountsClientDeletePollerResponse contains the response from method AccountsClient.Delete.
type AccountsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *AccountsClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l AccountsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (AccountsClientDeleteResponse, error) {
	respType := AccountsClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a AccountsClientDeletePollerResponse from the provided client and resume token.
func (l *AccountsClientDeletePollerResponse) Resume(ctx context.Context, client *AccountsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("AccountsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &AccountsClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// AccountsClientDeleteResponse contains the response from method AccountsClient.Delete.
type AccountsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsClientGetResponse contains the response from method AccountsClient.Get.
type AccountsClientGetResponse struct {
	AccountsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsClientGetResult contains the result from method AccountsClient.Get.
type AccountsClientGetResult struct {
	Account
}

// AccountsClientListByResourceGroupResponse contains the response from method AccountsClient.ListByResourceGroup.
type AccountsClientListByResourceGroupResponse struct {
	AccountsClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsClientListByResourceGroupResult contains the result from method AccountsClient.ListByResourceGroup.
type AccountsClientListByResourceGroupResult struct {
	AccountListResult
}

// AccountsClientListKeysResponse contains the response from method AccountsClient.ListKeys.
type AccountsClientListKeysResponse struct {
	AccountsClientListKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsClientListKeysResult contains the result from method AccountsClient.ListKeys.
type AccountsClientListKeysResult struct {
	APIKeys
}

// AccountsClientListModelsResponse contains the response from method AccountsClient.ListModels.
type AccountsClientListModelsResponse struct {
	AccountsClientListModelsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsClientListModelsResult contains the result from method AccountsClient.ListModels.
type AccountsClientListModelsResult struct {
	AccountModelListResult
}

// AccountsClientListResponse contains the response from method AccountsClient.List.
type AccountsClientListResponse struct {
	AccountsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsClientListResult contains the result from method AccountsClient.List.
type AccountsClientListResult struct {
	AccountListResult
}

// AccountsClientListSKUsResponse contains the response from method AccountsClient.ListSKUs.
type AccountsClientListSKUsResponse struct {
	AccountsClientListSKUsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsClientListSKUsResult contains the result from method AccountsClient.ListSKUs.
type AccountsClientListSKUsResult struct {
	AccountSKUListResult
}

// AccountsClientListUsagesResponse contains the response from method AccountsClient.ListUsages.
type AccountsClientListUsagesResponse struct {
	AccountsClientListUsagesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsClientListUsagesResult contains the result from method AccountsClient.ListUsages.
type AccountsClientListUsagesResult struct {
	UsageListResult
}

// AccountsClientRegenerateKeyResponse contains the response from method AccountsClient.RegenerateKey.
type AccountsClientRegenerateKeyResponse struct {
	AccountsClientRegenerateKeyResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsClientRegenerateKeyResult contains the result from method AccountsClient.RegenerateKey.
type AccountsClientRegenerateKeyResult struct {
	APIKeys
}

// AccountsClientUpdatePollerResponse contains the response from method AccountsClient.Update.
type AccountsClientUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *AccountsClientUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l AccountsClientUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (AccountsClientUpdateResponse, error) {
	respType := AccountsClientUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Account)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a AccountsClientUpdatePollerResponse from the provided client and resume token.
func (l *AccountsClientUpdatePollerResponse) Resume(ctx context.Context, client *AccountsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("AccountsClient.Update", token, client.pl)
	if err != nil {
		return err
	}
	poller := &AccountsClientUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// AccountsClientUpdateResponse contains the response from method AccountsClient.Update.
type AccountsClientUpdateResponse struct {
	AccountsClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsClientUpdateResult contains the result from method AccountsClient.Update.
type AccountsClientUpdateResult struct {
	Account
}

// CommitmentPlansClientCreateOrUpdateResponse contains the response from method CommitmentPlansClient.CreateOrUpdate.
type CommitmentPlansClientCreateOrUpdateResponse struct {
	CommitmentPlansClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CommitmentPlansClientCreateOrUpdateResult contains the result from method CommitmentPlansClient.CreateOrUpdate.
type CommitmentPlansClientCreateOrUpdateResult struct {
	CommitmentPlan
}

// CommitmentPlansClientDeletePollerResponse contains the response from method CommitmentPlansClient.Delete.
type CommitmentPlansClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *CommitmentPlansClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l CommitmentPlansClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (CommitmentPlansClientDeleteResponse, error) {
	respType := CommitmentPlansClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a CommitmentPlansClientDeletePollerResponse from the provided client and resume token.
func (l *CommitmentPlansClientDeletePollerResponse) Resume(ctx context.Context, client *CommitmentPlansClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("CommitmentPlansClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &CommitmentPlansClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// CommitmentPlansClientDeleteResponse contains the response from method CommitmentPlansClient.Delete.
type CommitmentPlansClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CommitmentPlansClientGetResponse contains the response from method CommitmentPlansClient.Get.
type CommitmentPlansClientGetResponse struct {
	CommitmentPlansClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CommitmentPlansClientGetResult contains the result from method CommitmentPlansClient.Get.
type CommitmentPlansClientGetResult struct {
	CommitmentPlan
}

// CommitmentPlansClientListResponse contains the response from method CommitmentPlansClient.List.
type CommitmentPlansClientListResponse struct {
	CommitmentPlansClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CommitmentPlansClientListResult contains the result from method CommitmentPlansClient.List.
type CommitmentPlansClientListResult struct {
	CommitmentPlanListResult
}

// CommitmentTiersClientListResponse contains the response from method CommitmentTiersClient.List.
type CommitmentTiersClientListResponse struct {
	CommitmentTiersClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CommitmentTiersClientListResult contains the result from method CommitmentTiersClient.List.
type CommitmentTiersClientListResult struct {
	CommitmentTierListResult
}

// DeletedAccountsClientGetResponse contains the response from method DeletedAccountsClient.Get.
type DeletedAccountsClientGetResponse struct {
	DeletedAccountsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DeletedAccountsClientGetResult contains the result from method DeletedAccountsClient.Get.
type DeletedAccountsClientGetResult struct {
	Account
}

// DeletedAccountsClientListResponse contains the response from method DeletedAccountsClient.List.
type DeletedAccountsClientListResponse struct {
	DeletedAccountsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DeletedAccountsClientListResult contains the result from method DeletedAccountsClient.List.
type DeletedAccountsClientListResult struct {
	AccountListResult
}

// DeletedAccountsClientPurgePollerResponse contains the response from method DeletedAccountsClient.Purge.
type DeletedAccountsClientPurgePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *DeletedAccountsClientPurgePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l DeletedAccountsClientPurgePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (DeletedAccountsClientPurgeResponse, error) {
	respType := DeletedAccountsClientPurgeResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a DeletedAccountsClientPurgePollerResponse from the provided client and resume token.
func (l *DeletedAccountsClientPurgePollerResponse) Resume(ctx context.Context, client *DeletedAccountsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("DeletedAccountsClient.Purge", token, client.pl)
	if err != nil {
		return err
	}
	poller := &DeletedAccountsClientPurgePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// DeletedAccountsClientPurgeResponse contains the response from method DeletedAccountsClient.Purge.
type DeletedAccountsClientPurgeResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DeploymentsClientCreateOrUpdatePollerResponse contains the response from method DeploymentsClient.CreateOrUpdate.
type DeploymentsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *DeploymentsClientCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l DeploymentsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (DeploymentsClientCreateOrUpdateResponse, error) {
	respType := DeploymentsClientCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Deployment)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a DeploymentsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *DeploymentsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *DeploymentsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("DeploymentsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &DeploymentsClientCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// DeploymentsClientCreateOrUpdateResponse contains the response from method DeploymentsClient.CreateOrUpdate.
type DeploymentsClientCreateOrUpdateResponse struct {
	DeploymentsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DeploymentsClientCreateOrUpdateResult contains the result from method DeploymentsClient.CreateOrUpdate.
type DeploymentsClientCreateOrUpdateResult struct {
	Deployment
}

// DeploymentsClientDeletePollerResponse contains the response from method DeploymentsClient.Delete.
type DeploymentsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *DeploymentsClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l DeploymentsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (DeploymentsClientDeleteResponse, error) {
	respType := DeploymentsClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a DeploymentsClientDeletePollerResponse from the provided client and resume token.
func (l *DeploymentsClientDeletePollerResponse) Resume(ctx context.Context, client *DeploymentsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("DeploymentsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &DeploymentsClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// DeploymentsClientDeleteResponse contains the response from method DeploymentsClient.Delete.
type DeploymentsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DeploymentsClientGetResponse contains the response from method DeploymentsClient.Get.
type DeploymentsClientGetResponse struct {
	DeploymentsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DeploymentsClientGetResult contains the result from method DeploymentsClient.Get.
type DeploymentsClientGetResult struct {
	Deployment
}

// DeploymentsClientListResponse contains the response from method DeploymentsClient.List.
type DeploymentsClientListResponse struct {
	DeploymentsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DeploymentsClientListResult contains the result from method DeploymentsClient.List.
type DeploymentsClientListResult struct {
	DeploymentListResult
}

// ManagementClientCheckDomainAvailabilityResponse contains the response from method ManagementClient.CheckDomainAvailability.
type ManagementClientCheckDomainAvailabilityResponse struct {
	ManagementClientCheckDomainAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ManagementClientCheckDomainAvailabilityResult contains the result from method ManagementClient.CheckDomainAvailability.
type ManagementClientCheckDomainAvailabilityResult struct {
	DomainAvailability
}

// ManagementClientCheckSKUAvailabilityResponse contains the response from method ManagementClient.CheckSKUAvailability.
type ManagementClientCheckSKUAvailabilityResponse struct {
	ManagementClientCheckSKUAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ManagementClientCheckSKUAvailabilityResult contains the result from method ManagementClient.CheckSKUAvailability.
type ManagementClientCheckSKUAvailabilityResult struct {
	SKUAvailabilityListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientListResult contains the result from method OperationsClient.List.
type OperationsClientListResult struct {
	OperationListResult
}

// PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse contains the response from method PrivateEndpointConnectionsClient.CreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateEndpointConnectionsClientCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionsClientCreateOrUpdateResponse, error) {
	respType := PrivateEndpointConnectionsClientCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.PrivateEndpointConnection)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *PrivateEndpointConnectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateEndpointConnectionsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &PrivateEndpointConnectionsClientCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.CreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	PrivateEndpointConnectionsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsClientCreateOrUpdateResult contains the result from method PrivateEndpointConnectionsClient.CreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeletePollerResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateEndpointConnectionsClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PrivateEndpointConnectionsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionsClientDeleteResponse, error) {
	respType := PrivateEndpointConnectionsClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateEndpointConnectionsClientDeletePollerResponse from the provided client and resume token.
func (l *PrivateEndpointConnectionsClientDeletePollerResponse) Resume(ctx context.Context, client *PrivateEndpointConnectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateEndpointConnectionsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &PrivateEndpointConnectionsClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnectionsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsClientGetResult contains the result from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListResponse contains the response from method PrivateEndpointConnectionsClient.List.
type PrivateEndpointConnectionsClientListResponse struct {
	PrivateEndpointConnectionsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsClientListResult contains the result from method PrivateEndpointConnectionsClient.List.
type PrivateEndpointConnectionsClientListResult struct {
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientListResponse contains the response from method PrivateLinkResourcesClient.List.
type PrivateLinkResourcesClientListResponse struct {
	PrivateLinkResourcesClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourcesClientListResult contains the result from method PrivateLinkResourcesClient.List.
type PrivateLinkResourcesClientListResult struct {
	PrivateLinkResourceListResult
}

// ResourceSKUsClientListResponse contains the response from method ResourceSKUsClient.List.
type ResourceSKUsClientListResponse struct {
	ResourceSKUsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ResourceSKUsClientListResult contains the result from method ResourceSKUsClient.List.
type ResourceSKUsClientListResult struct {
	ResourceSKUListResult
}
