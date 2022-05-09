/*
 * SPDX-License-Identifier: Apache-2.0
 */

package mbse

type AssetAction struct {
	ActionId           string
	ActionName         string
	BCAssetId          string
	BCAssetType        string
	BCAsset            BCAsset
	AttributesToRead   []string
	AttributesToUpdate AttributesToUpdateStruct
}

type BCAsset struct {
	BCAssetId        string
	BCAssetType      string
	ProjectBCAssetId string
	BCAssetName      string
	Description      string
	CRSubmissionTime string
	IsWithdrawn      bool
	WithdrawnTime    string
	CRDecision       CRDecisionStruct
	CRComments       CRCommentStruct
	Project          ProjectStruct
}

type CRDecisionStruct struct {
	CRDecisionTime   string
	CRDecisionNum    int
	CRDecisionStatus string
}

type CRCommentStruct struct {
	CommentTime string
	Comment     string
	CommenterId string
}

type AttributesToUpdateStruct struct {
	CRSubmissionTime string
	CRDecision       CRDecisionStruct
}

// type BCAsset struct {
// 	BCAssetId                  string
// 	BCAssetType                string
// 	ProjectBCAssetId           string
// 	BCAssetName                string
// 	Description                string
// 	Version                    VersionStruct
// 	BCTransaction              BCTransactionStruct
// 	Payloads                   []PayloadStruct
// 	InSituBCAssetStatuses      []InSituBCAssetStatusStruct
// 	ExternalInfo               ExternalInfoStruct
// 	ABAC                       string
// 	BCAssetJsonSchema          string
// 	StatusesForBCAssetId       string
// 	BCAssetStatuses            []BCAssetStatusesStruct
// 	ForBCAssetType             string
// 	BCAssetTypeConstraints     []BCAssetTypeConstraintsStruct
// 	AttributeProperties        AttributePropertiesStruct
// 	BCDocumentDesc             BCDocumentDescStruct
// 	Payload                    PayloadStruct
// 	BCDocStatuses              []BCDocumentStatusStruct
// 	FromCRId                   string
// 	CRSubmissionTime           string
// 	IsWithdrawn                bool
// 	WithdrawnTime              string
// 	CRDecision                 CRDecisionStruct
// 	CRComments                 CRCommentStruct
// 	BCDocumentIds              string
// 	MBSEBaseModelDesc          string
// 	VariantType                string
// 	BaseModelDescId            string
// 	MBSEModelDesc              string
// 	VariantModelDescId         string
// 	ParentMBSEModelId          string
// 	ChildMBSEModelIds          []string
// 	MBSEModelToRelationships   MBSEModelRelationshipStruct
// 	MBSEModelFromRelationships MBSEModelRelationshipStruct
// 	Project                    ProjectStruct
// 	VCN_SubmitTime             string
// 	VCN_ApprovalStatus         string
// 	VCN_ApprovalTime           string
// 	VerificationId             string
// 	VerificationTitle          string
// 	VerDesc                    string
// 	RequirementId              string
// }

type VersionStruct struct {
	Version     string
	Subversion  string
	StartTime   string
	EndTime     string
	Description string
	Status      string
}

type BCTransactionStruct struct {
	FabOrgMSPId   string
	FabUserId     string
	TransactionId string
	CreationTime  string
}

type ExternalInfoStruct struct {
	SEEId       string
	SEEFilename string
	SEEVersion  string
	SEELink     string
}

type AttributePropertiesStruct struct {
	Required  string
	Optional  string
	Immutable string
}

type BCDocumentDescStruct struct {
	BCDocumentType    string
	GenerationMethod  string
	SourceMBSEModelId string
}

type PayloadStruct struct {
	IsEncrypted   bool
	UseIPFS       bool
	PayloadName   string
	PayloadType   string
	PayloadDesc   string
	EncrypMethod  string
	EncrypKey     string
	PayloadRaw    string
	IPFSCid       string
	IPFS_HashHead string
	IPFS_Name     string
}

type BCDocumentStatusStruct struct {
	DocStatusCode string
	SinceTime     string
	CRId          string
	CRDecisionNum int
}

type InSituBCAssetStatusStruct struct {
	BCAssetStatusName  string
	BCAssetStatusValue string
	SinceTime          string
}

type BCAssetStatusesStruct struct {
	BCAssetStatusName  string
	BCAssetStatusValue string
	SinceTime          string
}

type BCAssetTypeConstraintsStruct struct {
	BCAssetTypeConstraintType         string
	BCAssetTypeConstraintName         string
	MustValidate                      bool
	FK_FromAssetType                  string
	FK_FromAssetTypeAttr              string
	FK_FromAssetTypeAttrIsArray       bool
	FK_ToAssetType                    string
	FK_ToAssetTypeAttr                string
	FK_ToAssetTypeAttrIsArray         bool
	FK_ReverseReferenceConstraintName string
	CheckExpression                   string
}

/////////////////////

type MBSEModelRelationshipStruct struct {
	RelatedMBSEModelId        string
	MBSEModelRelationshipType string
}

type ProjectStruct struct {
	IsTopLevel           bool
	PORTType             []string
	OrgRoles             []OrgRolesStruct
	OrgUserRoleTypes     []OrgUserRoleTypesStruct
	OrgUserRoles         []OrgUserRolesStruct
	BCAssetTypeRoleTypes []BCAssetTypeRoleTypesStruct
	BCAssetTypeRoles     []BCAssetTypeRolesStruct
	ParentProjectId      string
}

type OrgUserRolesStruct struct {
	OrgId     string
	UserId    string
	POURTTYpe string
}

type OrgRolesStruct struct {
	OrgId    string
	PORTType string
}

type OrgUserRoleTypesStruct struct {
	POURTTYpe   string
	Description string
}

type BCAssetTypeRoleTypesStruct struct {
	PBTRTType   string
	Description string
}

type BCAssetTypeRolesStruct struct {
	BCAssetType string
	PBTRTType   string
}
