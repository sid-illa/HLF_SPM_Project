package mbse

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

/*
//	Change Request Policies
*/

// Create Chanage Request Security Policy verifies the request is authorized to create a change request or not.
func (s *SmartContract) CreateCR_SP(ctx contractapi.TransactionContextInterface, crDecision CRDecisionStruct, orgRoles []OrgRolesStruct) bool {
	fmt.Println("Executing CreateCR Policy!!!")

	/*
		Conditions to evaluate
		crDecisions == "{}"
		user.organization == action.newobject.project.leadOrganization
		user.role == "CSE"
	*/
	userOrganization := s.GetIdentityAttribute(ctx, "organization")
	userRole := s.GetIdentityAttribute(ctx, "role")

	return (crDecision == CRDecisionStruct{} && s.Contains(s.GetOrgId(orgRoles, "lead"), userOrganization) && userRole == "cse")
}

// Update Change Request Security Policy verifies the request is authorized to update a Change Request or not.
func (s *SmartContract) UpdateCR_SP(ctx contractapi.TransactionContextInterface, crDecision CRDecisionStruct, orgRoles []OrgRolesStruct) bool {
	fmt.Println("Executing UpdateCR Policy!!!")

	/*
		Conditions to evaluate
		crDecisions == "{"CRDecisionTime": "2022-04-19T20:20:39+00:00", "CRDecisionNum": 123, "CRDecisionStauts": "approve"}"
		user.organization == action.newobject.project.leadOrganization
		user.role == "CSE"
	*/
	userOrganization := s.GetIdentityAttribute(ctx, "organization")
	userRole := s.GetIdentityAttribute(ctx, "role")

	return (crDecision != CRDecisionStruct{} && s.Contains(s.GetOrgId(orgRoles, "lead"), userOrganization) && userRole == "cse")
}

// Insert A Change Request Decision Security Policy verifies the request is authorized to add a CR decision int0 a change request object or not.
func (s *SmartContract) InsertACRDecision_SP(ctx contractapi.TransactionContextInterface, attributesToUpdate []string, orgRoles []OrgRolesStruct) bool {
	fmt.Println("Executing InsertACRDecision Policy!!!")

	// user.organization == action.newobject.project.controlBoard
	// user.abac.bc_orgrole == 'manager'
	// action.attributesToUpdate == 'crdecisions'
	userOrganization := s.GetIdentityAttribute(ctx, "organization")
	userRole := s.GetIdentityAttribute(ctx, "role")

	// Needs implementation
	return (s.Contains(attributesToUpdate, "IsWithdrawn") && s.Contains(s.GetOrgId(orgRoles, "controlBoard"), userOrganization) && userRole == "manager")
}

// Withdraw a CR
func (s *SmartContract) WithdrawACR_SP(ctx contractapi.TransactionContextInterface, attributesToRead []string, orgRoles []OrgRolesStruct) bool {
	fmt.Println("Executing WithdrawACR Policy!!!")

	// user.organization == action.newobject.project.leadOrganization
	// user.abac.bc_orgrole == 'cse'
	// (action.newobject.attributesToUpdate == "IsWithdrawn" && action.newobject.attributesToUpdate == "WithdrawnTime")
	userOrganization := s.GetIdentityAttribute(ctx, "organization")
	userRole := s.GetIdentityAttribute(ctx, "role")

	return (s.Contains(attributesToRead, "IsWithdrawn") && s.Contains(attributesToRead, "WithdrawnTime") && s.Contains(s.GetOrgId(orgRoles, "lead"), userOrganization) && userRole == "cse")
}

// policy for permitting the read of a change request object,Model ,.
func (s *SmartContract) ReadCR_SP(ctx contractapi.TransactionContextInterface, orgRoles []OrgRolesStruct) bool {
	fmt.Println("Executing ReadCR Policy!!!")

	// user.organization == object.project.organizations
	// (user.abac.bc_orgrole == "reviewer" || user.abac.bc_orgrole == "cse"|| user.abac.bc_orgrole == "manager" || user.abac.bc_orgrole == "keyreviewer" || user.abac.bc_orgrole == "se")
	// action.newobject.CRDecisions == object.CRDecisions
	allowedRolesToRead := []string{"reviewer", "cse", "manager", "keyreviewer", "se"}
	userRole := s.GetIdentityAttribute(ctx, "role")
	userOrganization := s.GetIdentityAttribute(ctx, "organization")
	fmt.Println("")

	return (s.Contains(s.GetOrgId(orgRoles, ""), userOrganization) && s.Contains(allowedRolesToRead, userRole))
}

// Insert a CR Comment
func (s *SmartContract) InsertACRComment_SP(ctx contractapi.TransactionContextInterface, attributesToUpdate []string, orgRoles []OrgRolesStruct) bool {
	fmt.Println("Executing InsertACRComment Policy!!!")

	// user.organization == object.project.organizations
	// (user.abac.bc_orgrole == "reviewer" || user.abac.bc_orgrole == "keyreviewer")
	// action.newobject.attributesToUpdate == 'CRComments'
	allowedRolesToRead := []string{"reviewer", "keyreviewer"}
	userRole := s.GetIdentityAttribute(ctx, "role")
	userOrganization := s.GetIdentityAttribute(ctx, "organization")

	return (s.Contains(attributesToUpdate, "CRComments") && s.Contains(s.GetOrgId(orgRoles, ""), userOrganization) && s.Contains(allowedRolesToRead, userRole))
}
