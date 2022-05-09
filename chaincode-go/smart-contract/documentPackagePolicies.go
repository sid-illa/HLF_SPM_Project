package mbse

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

/*
//	Document Package Policies
*/

// policy for permitting the creation of a BCDocument object.
func (s *SmartContract) CreateDocumentPackage_SP(ctx contractapi.TransactionContextInterface, orgRoles []OrgRolesStruct) bool {
	fmt.Println("Executing CreateDocumentPackage Policy!!!")

	// stringBagSize(action.newobject.DocumentPackageStatus) == 0 **** check check TODO
	// user.organization == action.newobject.project.leadOrganization
	// user.role == "cse"
	userOrganization := s.GetIdentityAttribute(ctx, "organization")
	userRole := s.GetIdentityAttribute(ctx, "role")

	return (s.Contains(s.GetOrgId(orgRoles, "lead"), userOrganization) && userRole == "cse")
}

// policy for permitting the reading of a BCDocument object.
func (s *SmartContract) ReadDocumentPackage_SP(ctx contractapi.TransactionContextInterface, crDecision CRDecisionStruct, orgRoles []OrgRolesStruct) bool {
	fmt.Println("Executing ReadDocumentPackage Policy!!!")

	// user.organization == object.project.organizations
	// (user.abac.bc_orgrole == "reviewer" || user.abac.bc_orgrole == "cse"|| user.abac.bc_orgrole == "manager" || user.abac.bc_orgrole == "keyreviewer" || user.abac.bc_orgrole == "se")
	userOrganization := s.GetIdentityAttribute(ctx, "organization")
	allowedRolesToRead := []string{"reviewer", "cse", "manager", "keyreviewer", "se"}
	userRole := s.GetIdentityAttribute(ctx, "role")

	return (s.Contains(s.GetOrgId(orgRoles, "lead"), userOrganization) && s.Contains(allowedRolesToRead, userRole))
}

// policy to check if user can insert a document package status on a BCDocument object or not
func (s *SmartContract) CanInsertDocumentPackageStatus_SP(ctx contractapi.TransactionContextInterface, attributesToUpdate []string, orgRoles []OrgRolesStruct) bool {
	fmt.Println("Executing CanInsertDocumentPackageStatus Policy!!!")

	// user.organization == action.newobject.project.controlBoard
	// user.abac.bc_orgrole == 'manager'
	// action.attributesToUpdate == 'DocumentPackagestatus'
	userOrganization := s.GetIdentityAttribute(ctx, "organization")
	userRole := s.GetIdentityAttribute(ctx, "role")

	// Needs implementation
	return (s.Contains(attributesToUpdate, "DocumentPackagestatus") && s.Contains(s.GetOrgId(orgRoles, "controlBoard"), userOrganization) && userRole == "manager")
}
