package mbse

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

/*
//	Manage the Change Requests
*/

// Create a CR BC Asset
func (s *SmartContract) CreateCR(ctx contractapi.TransactionContextInterface, assetAction AssetAction) (string, error) {

	isAuthorized := s.CreateCR_SP(ctx, assetAction.BCAsset.CRDecision, assetAction.BCAsset.Project.OrgRoles)
	if !isAuthorized {
		return "", fmt.Errorf("not authorized to create a Change request")
	}

	exists, err := s.AssetExists(ctx, assetAction.BCAsset.BCAssetId)
	if err != nil {
		return "Error ", fmt.Errorf("could not read from world state. %s", err)
	} else if exists {
		return "Error ", fmt.Errorf("the asset %s already exists", assetAction.BCAsset.BCAssetId)
	}

	bytes, _ := json.Marshal(assetAction.BCAsset)

	return "Successfully created Change Request", ctx.GetStub().PutState(assetAction.BCAsset.BCAssetId, bytes)
}

// Update a CR BC Asset
func (s *SmartContract) UpdateCR(ctx contractapi.TransactionContextInterface, assetAction AssetAction) (string, error) {

	isAuthorized := s.UpdateCR_SP(ctx, assetAction.AttributesToUpdate.CRDecision, assetAction.BCAsset.Project.OrgRoles)
	if !isAuthorized {
		return "Error ", fmt.Errorf("not authorized to update a Change request")
	}

	asset, err := s.ReadCR(ctx, assetAction)
	if err != nil {
		return "", err
	}

	bcAsset := new(BCAsset)

	err = json.Unmarshal([]byte(asset), bcAsset)
	if err != nil {
		return "Error ", fmt.Errorf("error. %s", err)
	}

	bcAsset.CRSubmissionTime = assetAction.AttributesToUpdate.CRSubmissionTime
	bcAsset.CRDecision = assetAction.AttributesToUpdate.CRDecision

	bytes, _ := json.Marshal(bcAsset)

	return "Successfully Updated CR Decision ", ctx.GetStub().PutState(bcAsset.BCAssetId, bytes)
}

// ReadAsset returns the asset stored in the world state with given id.
// func (s *SmartContract) ReadCR(ctx contractapi.TransactionContextInterface, id string) (*BCAsset, error) {
func (s *SmartContract) ReadCR(ctx contractapi.TransactionContextInterface, assetAction AssetAction) (string, error) {

	isAuthorized := s.ReadCR_SP(ctx, assetAction.BCAsset.Project.OrgRoles)
	if !isAuthorized {
		return "Error ", fmt.Errorf("not authorized to Read this Change request")
	}

	assetActionJSON, err := ctx.GetStub().GetState(assetAction.BCAsset.BCAssetId)
	if err != nil {
		return "Error ", fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetActionJSON == nil {
		return "Error ", fmt.Errorf("the asset %s does not exist", assetAction.BCAsset.BCAssetId)
	}

	// return &bcAsset, nil
	return string(assetActionJSON[:]), nil
}

// Insert A Change Request Decision Security Policy verifies the request is authorized to add a CR decision int0 a change request object or not.
func (s *SmartContract) InsertACRPolicy(ctx contractapi.TransactionContextInterface, assetActionvalue string) error {

	assetAction := new(AssetAction)

	err := json.Unmarshal([]byte(assetActionvalue), assetAction)
	if err != nil {
		return fmt.Errorf("error. %s", err)
	}

	isAuthorized := s.CreateCR_SP(ctx, assetAction.BCAsset.CRDecision, assetAction.BCAsset.Project.OrgRoles)
	if !isAuthorized {
		return fmt.Errorf("not authorized to create a Change request")
	}

	exists, err := s.AssetExists(ctx, assetAction.BCAsset.BCAssetId)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if exists {
		return fmt.Errorf("the asset %s already exists", assetAction.BCAsset.BCAssetId)
	}

	bytes, _ := json.Marshal(assetAction.BCAsset)

	return ctx.GetStub().PutState(assetAction.BCAsset.BCAssetId, bytes)

}

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {

	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// GetSubmittingClientIdentity returns the name and issuer of the identity that
// invokes the smart contract. This function base64 decodes the identity string
// before returning the value to the client or smart contract.
func (s *SmartContract) GetSubmittingClientIdentity(ctx contractapi.TransactionContextInterface) (string, error) {

	b64ID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return "", fmt.Errorf("failed to read clientID: %v", err)
	}
	decodeID, err := base64.StdEncoding.DecodeString(b64ID)
	if err != nil {
		return "", fmt.Errorf("failed to base64 decode clientID: %v", err)
	}
	return string(decodeID), nil
}
