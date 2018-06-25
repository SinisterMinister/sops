package operationsmanagement

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
	"github.com/Azure/go-autorest/autorest"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// ArmTemplateParameter parameter to pass to ARM template
type ArmTemplateParameter struct {
	// Name - name of the parameter.
	Name *string `json:"name,omitempty"`
	// Value - value for the parameter. In Jtoken
	Value *string `json:"value,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// CodeMessageError the error body contract.
type CodeMessageError struct {
	// Error - The error details for a failed request.
	Error *CodeMessageErrorError `json:"error,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// CodeMessageErrorError the error details for a failed request.
type CodeMessageErrorError struct {
	// Code - The error type.
	Code *string `json:"code,omitempty"`
	// Message - The error message.
	Message *string `json:"message,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// ManagementAssociation the container for solution.
type ManagementAssociation struct {
	autorest.Response `json:"-"`
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Properties - Properties for ManagementAssociation object supported by the OperationsManagement resource provider.
	Properties *ManagementAssociationProperties `json:"properties,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// ManagementAssociationProperties managementAssociation properties supported by the OperationsManagement resource
// provider.
type ManagementAssociationProperties struct {
	// ApplicationID - The applicationId of the appliance for this association.
	ApplicationID *string `json:"applicationId,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// ManagementAssociationPropertiesList the list of ManagementAssociation response
type ManagementAssociationPropertiesList struct {
	autorest.Response `json:"-"`
	// Value - List of Management Association properites within the subscription.
	Value *[]ManagementAssociation `json:"value,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// ManagementConfiguration the container for solution.
type ManagementConfiguration struct {
	autorest.Response `json:"-"`
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Properties - Properties for ManagementConfiguration object supported by the OperationsManagement resource provider.
	Properties *ManagementConfigurationProperties `json:"properties,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// ManagementConfigurationProperties managementConfiguration properties supported by the OperationsManagement
// resource provider.
type ManagementConfigurationProperties struct {
	// ApplicationID - The applicationId of the appliance for this Management.
	ApplicationID *string `json:"applicationId,omitempty"`
	// ParentResourceType - The type of the parent resource.
	ParentResourceType *string `json:"parentResourceType,omitempty"`
	// Parameters - Parameters to run the ARM template
	Parameters *[]ArmTemplateParameter `json:"parameters,omitempty"`
	// ProvisioningState - The provisioning state for the ManagementConfiguration.
	ProvisioningState *string `json:"provisioningState,omitempty"`
	// Template - The Json object containing the ARM template to deploy
	Template interface{} `json:"template,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// ManagementConfigurationPropertiesList the list of ManagementConfiguration response
type ManagementConfigurationPropertiesList struct {
	autorest.Response `json:"-"`
	// Value - List of Management Configuration properites within the subscription.
	Value *[]ManagementConfiguration `json:"value,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// Operation supported operation of OperationsManagement resource provider.
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - Display metadata associated with the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// OperationDisplay display metadata associated with the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft OperationsManagement.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Type of operation: get, read, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// OperationListResult result of the request to list solution operations.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of solution operations supported by the OperationsManagement resource provider.
	Value *[]Operation `json:"value,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// Solution the container for solution.
type Solution struct {
	autorest.Response `json:"-"`
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Plan - Plan for solution object supported by the OperationsManagement resource provider.
	Plan *SolutionPlan `json:"plan,omitempty"`
	// Properties - Properties for solution object supported by the OperationsManagement resource provider.
	Properties *SolutionProperties `json:"properties,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// SolutionPlan plan for solution object supported by the OperationsManagement resource provider.
type SolutionPlan struct {
	// Name - name of the solution to be created. For Microsoft published solution it should be in the format of solutionType(workspaceName). SolutionType part is case sensitive. For third party solution, it can be anything.
	Name *string `json:"name,omitempty"`
	// Publisher - Publisher name. For gallery solution, it is Microsoft.
	Publisher *string `json:"publisher,omitempty"`
	// PromotionCode - promotionCode, Not really used now, can you left as empty
	PromotionCode *string `json:"promotionCode,omitempty"`
	// Product - name of the solution to enabled/add. For Microsoft published gallery solution it should be in the format of OMSGallery/<solutionType>. This is case sensitive
	Product *string `json:"product,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// SolutionProperties solution properties supported by the OperationsManagement resource provider.
type SolutionProperties struct {
	// WorkspaceResourceID - The azure resourceId for the workspace where the solution will be deployed/enabled.
	WorkspaceResourceID *string `json:"workspaceResourceId,omitempty"`
	// ProvisioningState - The provisioning state for the solution.
	ProvisioningState *string `json:"provisioningState,omitempty"`
	// ContainedResources - The azure resources that will be contained within the solutions. They will be locked and gets deleted automatically when the solution is deleted.
	ContainedResources *[]string `json:"containedResources,omitempty"`
	// ReferencedResources - The resources that will be referenced from this solution. Deleting any of those solution out of band will break the solution.
	ReferencedResources *[]string `json:"referencedResources,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// SolutionPropertiesList the list of solution response
type SolutionPropertiesList struct {
	autorest.Response `json:"-"`
	// Value - List of solution properites within the subscription.
	Value *[]Solution `json:"value,omitempty"`
}