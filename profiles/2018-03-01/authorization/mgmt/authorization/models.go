// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package authorization

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type ClassicAdministrator = original.ClassicAdministrator
type ClassicAdministratorListResult = original.ClassicAdministratorListResult
type ClassicAdministratorListResultIterator = original.ClassicAdministratorListResultIterator
type ClassicAdministratorListResultPage = original.ClassicAdministratorListResultPage
type ClassicAdministratorProperties = original.ClassicAdministratorProperties
type ClassicAdministratorsClient = original.ClassicAdministratorsClient
type Permission = original.Permission
type PermissionGetResult = original.PermissionGetResult
type PermissionGetResultIterator = original.PermissionGetResultIterator
type PermissionGetResultPage = original.PermissionGetResultPage
type PermissionsClient = original.PermissionsClient
type ProviderOperation = original.ProviderOperation
type ProviderOperationsMetadata = original.ProviderOperationsMetadata
type ProviderOperationsMetadataClient = original.ProviderOperationsMetadataClient
type ProviderOperationsMetadataListResult = original.ProviderOperationsMetadataListResult
type ProviderOperationsMetadataListResultIterator = original.ProviderOperationsMetadataListResultIterator
type ProviderOperationsMetadataListResultPage = original.ProviderOperationsMetadataListResultPage
type ResourceType = original.ResourceType
type RoleAssignment = original.RoleAssignment
type RoleAssignmentCreateParameters = original.RoleAssignmentCreateParameters
type RoleAssignmentFilter = original.RoleAssignmentFilter
type RoleAssignmentListResult = original.RoleAssignmentListResult
type RoleAssignmentListResultIterator = original.RoleAssignmentListResultIterator
type RoleAssignmentListResultPage = original.RoleAssignmentListResultPage
type RoleAssignmentProperties = original.RoleAssignmentProperties
type RoleAssignmentPropertiesWithScope = original.RoleAssignmentPropertiesWithScope
type RoleAssignmentsClient = original.RoleAssignmentsClient
type RoleDefinition = original.RoleDefinition
type RoleDefinitionFilter = original.RoleDefinitionFilter
type RoleDefinitionListResult = original.RoleDefinitionListResult
type RoleDefinitionListResultIterator = original.RoleDefinitionListResultIterator
type RoleDefinitionListResultPage = original.RoleDefinitionListResultPage
type RoleDefinitionProperties = original.RoleDefinitionProperties
type RoleDefinitionsClient = original.RoleDefinitionsClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClassicAdministratorListResultIterator(page ClassicAdministratorListResultPage) ClassicAdministratorListResultIterator {
	return original.NewClassicAdministratorListResultIterator(page)
}
func NewClassicAdministratorListResultPage(getNextPage func(context.Context, ClassicAdministratorListResult) (ClassicAdministratorListResult, error)) ClassicAdministratorListResultPage {
	return original.NewClassicAdministratorListResultPage(getNextPage)
}
func NewClassicAdministratorsClient(subscriptionID string) ClassicAdministratorsClient {
	return original.NewClassicAdministratorsClient(subscriptionID)
}
func NewClassicAdministratorsClientWithBaseURI(baseURI string, subscriptionID string) ClassicAdministratorsClient {
	return original.NewClassicAdministratorsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPermissionGetResultIterator(page PermissionGetResultPage) PermissionGetResultIterator {
	return original.NewPermissionGetResultIterator(page)
}
func NewPermissionGetResultPage(getNextPage func(context.Context, PermissionGetResult) (PermissionGetResult, error)) PermissionGetResultPage {
	return original.NewPermissionGetResultPage(getNextPage)
}
func NewPermissionsClient(subscriptionID string) PermissionsClient {
	return original.NewPermissionsClient(subscriptionID)
}
func NewPermissionsClientWithBaseURI(baseURI string, subscriptionID string) PermissionsClient {
	return original.NewPermissionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProviderOperationsMetadataClient(subscriptionID string) ProviderOperationsMetadataClient {
	return original.NewProviderOperationsMetadataClient(subscriptionID)
}
func NewProviderOperationsMetadataClientWithBaseURI(baseURI string, subscriptionID string) ProviderOperationsMetadataClient {
	return original.NewProviderOperationsMetadataClientWithBaseURI(baseURI, subscriptionID)
}
func NewProviderOperationsMetadataListResultIterator(page ProviderOperationsMetadataListResultPage) ProviderOperationsMetadataListResultIterator {
	return original.NewProviderOperationsMetadataListResultIterator(page)
}
func NewProviderOperationsMetadataListResultPage(getNextPage func(context.Context, ProviderOperationsMetadataListResult) (ProviderOperationsMetadataListResult, error)) ProviderOperationsMetadataListResultPage {
	return original.NewProviderOperationsMetadataListResultPage(getNextPage)
}
func NewRoleAssignmentListResultIterator(page RoleAssignmentListResultPage) RoleAssignmentListResultIterator {
	return original.NewRoleAssignmentListResultIterator(page)
}
func NewRoleAssignmentListResultPage(getNextPage func(context.Context, RoleAssignmentListResult) (RoleAssignmentListResult, error)) RoleAssignmentListResultPage {
	return original.NewRoleAssignmentListResultPage(getNextPage)
}
func NewRoleAssignmentsClient(subscriptionID string) RoleAssignmentsClient {
	return original.NewRoleAssignmentsClient(subscriptionID)
}
func NewRoleAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) RoleAssignmentsClient {
	return original.NewRoleAssignmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRoleDefinitionListResultIterator(page RoleDefinitionListResultPage) RoleDefinitionListResultIterator {
	return original.NewRoleDefinitionListResultIterator(page)
}
func NewRoleDefinitionListResultPage(getNextPage func(context.Context, RoleDefinitionListResult) (RoleDefinitionListResult, error)) RoleDefinitionListResultPage {
	return original.NewRoleDefinitionListResultPage(getNextPage)
}
func NewRoleDefinitionsClient(subscriptionID string) RoleDefinitionsClient {
	return original.NewRoleDefinitionsClient(subscriptionID)
}
func NewRoleDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) RoleDefinitionsClient {
	return original.NewRoleDefinitionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/2018-03-01"
}
func Version() string {
	return original.Version()
}
