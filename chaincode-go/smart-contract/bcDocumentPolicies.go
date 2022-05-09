package mbse

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

/*
//	BC Document Policies
*/

// policy for permitting the creation of a BCDocument object.
func (s *SmartContract) CreateBCDocument_SP(ctx contractapi.TransactionContextInterface, bCDocumentStatus []BCDocumentStatusStruct, orgRoles []OrgRolesStruct) bool {
	fmt.Println("Executing CreateBCDocument Policy!!!")

	// user.organization == action.newobject.project.leadOrganization
	// user.role == "cse"
	// stringBagSize(action.newobject.BCDocumentStatus) == 0
	userOrganization := s.GetIdentityAttribute(ctx, "organization")
	userRole := s.GetIdentityAttribute(ctx, "role")

	return (len(bCDocumentStatus) == 0 && s.Contains(s.GetOrgId(orgRoles, "lead"), userOrganization) && userRole == "cse")
}

// Update Chanage Request Security Policy verifies the request is authorized to update a Change Request or not.
func (s *SmartContract) ReadBCDocument_SP(ctx contractapi.TransactionContextInterface, orgRoles []OrgRolesStruct) bool {
	fmt.Println("Executing ReadBCDocument Policy!!!")

	// crDecisions == "{"CRDecisionTime": "2022-04-19T20:20:39+00:00", "CRDecisionNum": 123, "CRDecisionStauts": "approve"}"
	// user.organization == object.project.organizations
	// user.role == "CSE"
	userOrganization := s.GetIdentityAttribute(ctx, "organization")
	allowedRolesToRead := []string{"reviewer", "cse", "manager", "keyreviewer", "se"}
	userRole := s.GetIdentityAttribute(ctx, "role")

	return (s.Contains(s.GetOrgId(orgRoles, "lead"), userOrganization) && s.Contains(allowedRolesToRead, userRole))
}

// Insert A Change Request Decision Security Policy verifies the request is authorized to add a CR decision int0 a change request object or not.
func (s *SmartContract) CanInsertBCDocumentStatus_SP(ctx contractapi.TransactionContextInterface, attributesToUpdate []string, orgRoles []OrgRolesStruct) bool {
	fmt.Println("Executing CanInsertBCDocumentStatus Policy!!!")

	// user.organization == action.newobject.project.controlBoard
	// user.abac.bc_orgrole == 'manager'
	// action.attributesToUpdate == 'BCDocumentStatus'
	userOrganization := s.GetIdentityAttribute(ctx, "organization")
	userRole := s.GetIdentityAttribute(ctx, "role")

	// Needs implementation
	return (s.Contains(attributesToUpdate, "BCDocumentStatus") && s.Contains(s.GetOrgId(orgRoles, "controlBoard"), userOrganization) && userRole == "manager")
}
