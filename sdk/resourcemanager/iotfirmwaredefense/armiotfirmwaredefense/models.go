//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armiotfirmwaredefense

import "time"

// BinaryHardening - Binary hardening of a firmware.
type BinaryHardening struct {
	// The architecture of the uploaded firmware.
	Architecture *string

	// ID for the binary hardening result.
	BinaryHardeningID *string

	// class for binary hardening.
	Class *string

	// Binary hardening features.
	Features *BinaryHardeningFeatures

	// path for binary hardening.
	Path *string

	// The rpath of the uploaded firmware.
	Rpath *string

	// The runpath of the uploaded firmware.
	Runpath *string
}

// BinaryHardeningFeatures - Binary hardening features.
type BinaryHardeningFeatures struct {
	// Canary flag.
	Canary *CanaryFlag

	// NX flag.
	Nx *NxFlag

	// PIE flag.
	Pie *PieFlag

	// RELRO flag.
	Relro *RelroFlag

	// Stripped flag.
	Stripped *StrippedFlag
}

// BinaryHardeningList - List result for binary hardening
type BinaryHardeningList struct {
	// The uri to fetch the next page of asset.
	NextLink *string

	// READ-ONLY; The list of binary hardening results.
	Value []*BinaryHardening
}

// BinaryHardeningSummary - Binary hardening summary percentages.
type BinaryHardeningSummary struct {
	// Canary summary percentage
	Canary *int32

	// NX summary percentage
	Nx *int32

	// PIE summary percentage
	Pie *int32

	// RELRO summary percentage
	Relro *int32

	// Stripped summary percentage
	Stripped *int32

	// Total number of binaries that were analyzed
	TotalFiles *int64
}

// Component of a firmware.
type Component struct {
	// ID for the component.
	ComponentID *string

	// Name for the component.
	ComponentName *string

	// Flag if new update is available for the component.
	IsUpdateAvailable *IsUpdateAvailable

	// License for the component.
	License *string

	// Paths of the component.
	Paths []*string

	// Release date for the component.
	ReleaseDate *time.Time

	// Version for the component.
	Version *string
}

// ComponentList - List result for components
type ComponentList struct {
	// The uri to fetch the next page of asset.
	NextLink *string

	// READ-ONLY; The list of components.
	Value []*Component
}

// CryptoCertificate - Crypto certificate properties
type CryptoCertificate struct {
	// ID for the certificate.
	CryptoCertID *string

	// Encoding used for the certificate.
	Encoding *string

	// Expiration date for the certificate.
	ExpirationDate *time.Time

	// Fingerprint of the certificate.
	Fingerprint *string

	// Indicates if the certificate is expired.
	IsExpired *IsExpired

	// Indicates if the certificate was self-signed.
	IsSelfSigned *IsSelfSigned

	// Indicates the certificate's key size is considered too small to be secure for the key algorithm.
	IsShortKeySize *IsShortKeySize

	// Indicates the signature algorithm used is insecure.
	IsWeakSignature *IsWeakSignature

	// Issue date for the certificate.
	IssuedDate *time.Time

	// Issuer information of the certificate.
	Issuer *CryptoCertificateEntity

	// Key algorithm used in the certificate.
	KeyAlgorithm *string

	// Size of the certificate's key in bits
	KeySize *int64

	// Name of the certificate.
	Name *string

	// A matching paired private key.
	PairedKey *PairedKey

	// Role of the certificate (Root CA, etc)
	Role *string

	// Serial number of the certificate.
	SerialNumber *string

	// The signature algorithm used in the certificate.
	SignatureAlgorithm *string

	// Subject information of the certificate.
	Subject *CryptoCertificateEntity

	// List of functions the certificate can fulfill.
	Usage []*string

	// READ-ONLY; List of files paths for this certificate
	FilePaths []*string
}

// CryptoCertificateEntity - Information on an entity (distinguished name) in a cryptographic certificate.
type CryptoCertificateEntity struct {
	// Common name of the certificate entity.
	CommonName *string

	// Country code of the certificate entity.
	Country *string

	// Organization of the certificate entity.
	Organization *string

	// The organizational unit of the certificate entity.
	OrganizationalUnit *string

	// Geographical state or province of the certificate entity.
	State *string
}

// CryptoCertificateList - Crypto certificates list
type CryptoCertificateList struct {
	// The uri to fetch the next page of asset.
	NextLink *string

	// READ-ONLY; Crypto certificates list
	Value []*CryptoCertificate
}

// CryptoCertificateSummary - Cryptographic certificate summary values.
type CryptoCertificateSummary struct {
	// Total number of expired certificates found.
	Expired *int64

	// Total number of nearly expired certificates found.
	ExpiringSoon *int64

	// Total number of paired private keys found for the certificates.
	PairedKeys *int64

	// Total number of certificates found that are self-signed.
	SelfSigned *int64

	// Total number of certificates found that have an insecure key size for the key algorithm.
	ShortKeySize *int64

	// Total number of certificates found.
	TotalCertificates *int64

	// Total number of certificates found using a weak signature algorithm.
	WeakSignature *int64
}

// CryptoKey - Crypto key properties
type CryptoKey struct {
	// ID for the key.
	CryptoKeyID *string

	// Indicates the key size is considered too small to be secure for the algorithm.
	IsShortKeySize *IsShortKeySize

	// Key algorithm name.
	KeyAlgorithm *string

	// Size of the key in bits.
	KeySize *int64

	// Type of the key (public or private).
	KeyType *string

	// A matching paired key or certificate.
	PairedKey *PairedKey

	// Functions the key can fulfill.
	Usage []*string

	// READ-ONLY; List of files paths for this key.
	FilePaths []*string
}

// CryptoKeyList - Crypto keys list
type CryptoKeyList struct {
	// The uri to fetch the next page of asset.
	NextLink *string

	// READ-ONLY; Crypto keys list
	Value []*CryptoKey
}

// CryptoKeySummary - Cryptographic key summary values.
type CryptoKeySummary struct {
	// Total number of keys found that have a matching paired key or certificate.
	PairedKeys *int64

	// Total number of private keys found.
	PrivateKeys *int64

	// Total number of (non-certificate) public keys found.
	PublicKeys *int64

	// Total number of keys found that have an insecure key size for the algorithm.
	ShortKeySize *int64

	// Total number of cryptographic keys found.
	TotalKeys *int64
}

// Cve - Known CVEs of a firmware.
type Cve struct {
	// Component of CVE
	Component any

	// ID of CVE
	CveID *string

	// A single CVSS score to represent the CVE. If a V3 score is specified, then it will use the V3 score. Otherwise if the V2
	// score is specified it will be the V2 score
	CvssScore *string

	// Cvss V2 score of CVE
	CvssV2Score *string

	// Cvss V3 score of CVE
	CvssV3Score *string

	// Cvss version of CVE
	CvssVersion *string

	// Description of CVE
	Description *string

	// Name of CVE
	Name *string

	// Publish date of CVE
	PublishDate *time.Time

	// Severity of CVE
	Severity *string

	// Updated date of CVE
	UpdatedDate *time.Time

	// READ-ONLY; The list of CVE links.
	Links []*CveLink
}

// CveLink - Link for CVE
type CveLink struct {
	// Href of CVE link
	Href *string

	// Label of CVE link
	Label *string
}

// CveList - List result for CVE
type CveList struct {
	// The uri to fetch the next page of asset.
	NextLink *string

	// READ-ONLY; The list of CVE results.
	Value []*Cve
}

// CveSummary - CVE summary values.
type CveSummary struct {
	// The total number of critical severity CVEs detected
	Critical *int64

	// The total number of high severity CVEs detected
	High *int64

	// The total number of low severity CVEs detected
	Low *int64

	// The total number of medium severity CVEs detected
	Medium *int64

	// The total number of undefined severity CVEs detected
	Undefined *int64

	// The total number of unknown severity CVEs detected
	Unknown *int64
}

// Firmware definition
type Firmware struct {
	// The properties of a firmware
	Properties *FirmwareProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// FirmwareClientCreateOptions contains the optional parameters for the FirmwareClient.Create method.
type FirmwareClientCreateOptions struct {
	// placeholder for future optional parameters
}

// FirmwareClientDeleteOptions contains the optional parameters for the FirmwareClient.Delete method.
type FirmwareClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// FirmwareClientGenerateBinaryHardeningDetailsOptions contains the optional parameters for the FirmwareClient.GenerateBinaryHardeningDetails
// method.
type FirmwareClientGenerateBinaryHardeningDetailsOptions struct {
	// placeholder for future optional parameters
}

// FirmwareClientGenerateBinaryHardeningSummaryOptions contains the optional parameters for the FirmwareClient.GenerateBinaryHardeningSummary
// method.
type FirmwareClientGenerateBinaryHardeningSummaryOptions struct {
	// placeholder for future optional parameters
}

// FirmwareClientGenerateComponentDetailsOptions contains the optional parameters for the FirmwareClient.GenerateComponentDetails
// method.
type FirmwareClientGenerateComponentDetailsOptions struct {
	// placeholder for future optional parameters
}

// FirmwareClientGenerateCryptoCertificateSummaryOptions contains the optional parameters for the FirmwareClient.GenerateCryptoCertificateSummary
// method.
type FirmwareClientGenerateCryptoCertificateSummaryOptions struct {
	// placeholder for future optional parameters
}

// FirmwareClientGenerateCryptoKeySummaryOptions contains the optional parameters for the FirmwareClient.GenerateCryptoKeySummary
// method.
type FirmwareClientGenerateCryptoKeySummaryOptions struct {
	// placeholder for future optional parameters
}

// FirmwareClientGenerateCveSummaryOptions contains the optional parameters for the FirmwareClient.GenerateCveSummary method.
type FirmwareClientGenerateCveSummaryOptions struct {
	// placeholder for future optional parameters
}

// FirmwareClientGenerateDownloadURLOptions contains the optional parameters for the FirmwareClient.GenerateDownloadURL method.
type FirmwareClientGenerateDownloadURLOptions struct {
	// placeholder for future optional parameters
}

// FirmwareClientGenerateFilesystemDownloadURLOptions contains the optional parameters for the FirmwareClient.GenerateFilesystemDownloadURL
// method.
type FirmwareClientGenerateFilesystemDownloadURLOptions struct {
	// placeholder for future optional parameters
}

// FirmwareClientGenerateSummaryOptions contains the optional parameters for the FirmwareClient.GenerateSummary method.
type FirmwareClientGenerateSummaryOptions struct {
	// placeholder for future optional parameters
}

// FirmwareClientGetOptions contains the optional parameters for the FirmwareClient.Get method.
type FirmwareClientGetOptions struct {
	// placeholder for future optional parameters
}

// FirmwareClientListByWorkspaceOptions contains the optional parameters for the FirmwareClient.NewListByWorkspacePager method.
type FirmwareClientListByWorkspaceOptions struct {
	// placeholder for future optional parameters
}

// FirmwareClientListGenerateBinaryHardeningListOptions contains the optional parameters for the FirmwareClient.NewListGenerateBinaryHardeningListPager
// method.
type FirmwareClientListGenerateBinaryHardeningListOptions struct {
	// placeholder for future optional parameters
}

// FirmwareClientListGenerateComponentListOptions contains the optional parameters for the FirmwareClient.NewListGenerateComponentListPager
// method.
type FirmwareClientListGenerateComponentListOptions struct {
	// placeholder for future optional parameters
}

// FirmwareClientListGenerateCryptoCertificateListOptions contains the optional parameters for the FirmwareClient.NewListGenerateCryptoCertificateListPager
// method.
type FirmwareClientListGenerateCryptoCertificateListOptions struct {
	// placeholder for future optional parameters
}

// FirmwareClientListGenerateCryptoKeyListOptions contains the optional parameters for the FirmwareClient.NewListGenerateCryptoKeyListPager
// method.
type FirmwareClientListGenerateCryptoKeyListOptions struct {
	// placeholder for future optional parameters
}

// FirmwareClientListGenerateCveListOptions contains the optional parameters for the FirmwareClient.NewListGenerateCveListPager
// method.
type FirmwareClientListGenerateCveListOptions struct {
	// placeholder for future optional parameters
}

// FirmwareClientListGeneratePasswordHashListOptions contains the optional parameters for the FirmwareClient.NewListGeneratePasswordHashListPager
// method.
type FirmwareClientListGeneratePasswordHashListOptions struct {
	// placeholder for future optional parameters
}

// FirmwareClientUpdateOptions contains the optional parameters for the FirmwareClient.Update method.
type FirmwareClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// FirmwareList - List of firmwares
type FirmwareList struct {
	// The uri to fetch the next page of asset.
	NextLink *string

	// READ-ONLY; The list of firmwares.
	Value []*Firmware
}

// FirmwareProperties - Firmware properties.
type FirmwareProperties struct {
	// User-specified description of the firmware.
	Description *string

	// File name for a firmware that user uploaded.
	FileName *string

	// File size of the uploaded firmware image.
	FileSize *int64

	// Firmware model.
	Model *string

	// The status of firmware scan.
	Status *Status

	// A list of errors or other messages generated during firmware analysis
	StatusMessages []any

	// Firmware vendor.
	Vendor *string

	// Firmware version.
	Version *string

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState
}

// FirmwareSummary - Summary result after scanning the firmware.
type FirmwareSummary struct {
	// Time used for analysis
	AnalysisTimeSeconds *int64

	// Binary count
	BinaryCount *int64

	// Components count.
	ComponentCount *int64

	// Extracted file count.
	ExtractedFileCount *int64

	// Total extracted size of the firmware in bytes.
	ExtractedSize *int64

	// Firmware file size in bytes.
	FileSize *int64

	// The number of root file systems found.
	RootFileSystems *int64
}

// FirmwareUpdateDefinition - Firmware definition
type FirmwareUpdateDefinition struct {
	// The editable properties of a firmware
	Properties *FirmwareProperties
}

// GenerateUploadURLRequest - Properties for generating an upload URL
type GenerateUploadURLRequest struct {
	// A unique ID for the firmware to be uploaded.
	FirmwareID *string
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PairedKey - Details of a matching paired key or certificate.
type PairedKey struct {
	// Additional paired key properties
	AdditionalProperties any

	// ID of the paired key or certificate.
	ID *string

	// The type indicating whether the paired object is a key or certificate.
	Type *string
}

// PasswordHash - Password hash properties
type PasswordHash struct {
	// Algorithm of the password hash
	Algorithm *string

	// Context of password hash
	Context *string

	// File path of the password hash
	FilePath *string

	// Hash of the password
	Hash *string

	// ID for password hash
	PasswordHashID *string

	// Salt of the password hash
	Salt *string

	// User name of password hash
	Username *string
}

// PasswordHashList - Password hashes list
type PasswordHashList struct {
	// The uri to fetch the next page of asset.
	NextLink *string

	// READ-ONLY; Password hashes list
	Value []*PasswordHash
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// URLToken - Url data for creating or accessing a blob file.
type URLToken struct {
	// READ-ONLY; SAS URL for creating or accessing a blob file.
	URL *string

	// READ-ONLY; SAS URL for file uploading. Kept for backwards compatibility
	UploadURL *string
}

// Workspace - Firmware analysis workspace.
type Workspace struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Workspace properties.
	Properties *WorkspaceProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// WorkspaceList - Return a list of firmware analysis workspaces.
type WorkspaceList struct {
	// The uri to fetch the next page of asset.
	NextLink *string

	// READ-ONLY; The list of firmware analysis workspaces.
	Value []*Workspace
}

// WorkspaceProperties - Workspace properties.
type WorkspaceProperties struct {
	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState
}

// WorkspaceUpdateDefinition - Firmware analysis workspace.
type WorkspaceUpdateDefinition struct {
	// The editable workspace properties.
	Properties *WorkspaceProperties
}

// WorkspacesClientCreateOptions contains the optional parameters for the WorkspacesClient.Create method.
type WorkspacesClientCreateOptions struct {
	// placeholder for future optional parameters
}

// WorkspacesClientDeleteOptions contains the optional parameters for the WorkspacesClient.Delete method.
type WorkspacesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// WorkspacesClientGenerateUploadURLOptions contains the optional parameters for the WorkspacesClient.GenerateUploadURL method.
type WorkspacesClientGenerateUploadURLOptions struct {
	// placeholder for future optional parameters
}

// WorkspacesClientGetOptions contains the optional parameters for the WorkspacesClient.Get method.
type WorkspacesClientGetOptions struct {
	// placeholder for future optional parameters
}

// WorkspacesClientListByResourceGroupOptions contains the optional parameters for the WorkspacesClient.NewListByResourceGroupPager
// method.
type WorkspacesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// WorkspacesClientListBySubscriptionOptions contains the optional parameters for the WorkspacesClient.NewListBySubscriptionPager
// method.
type WorkspacesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// WorkspacesClientUpdateOptions contains the optional parameters for the WorkspacesClient.Update method.
type WorkspacesClientUpdateOptions struct {
	// placeholder for future optional parameters
}
