package mbse

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

/*
//	Manage the Assets
*/

// Create a CR BC Asset
func (s *SmartContract) ManageMBSEAssets(ctx contractapi.TransactionContextInterface, assetActionvalue string) (string, error) {

	assetAction := new(AssetAction)

	err := json.Unmarshal([]byte(assetActionvalue), assetAction)
	if err != nil {
		return "", fmt.Errorf("error. %s", err)
	}

	fmt.Println("ActionID : ", assetAction.ActionId)

	switch {

	case (assetAction.ActionId == "C" && assetAction.BCAssetType == "ChangeRequest"):
		fmt.Println("Create CR!!!")
		return s.CreateCR(ctx, *assetAction)

	case (assetAction.ActionId == "U" && assetAction.BCAssetType == "ChangeRequest"):
		fmt.Println("Update CR!!!")
		return s.UpdateCR(ctx, *assetAction)

	case (assetAction.ActionId == "R" && assetAction.BCAssetType == "ChangeRequest"):
		fmt.Println("Read CR!!!")
		return s.ReadCR(ctx, *assetAction)

	case (assetAction.ActionId == "P" && assetAction.BCAssetType == "ChangeRequest"):
		fmt.Println("Insert a CR Decision!!!")
		// return s.InsertACRDecisions(ctx, *assetAction)

	case (assetAction.ActionId == "D" && assetAction.BCAssetType == "ChangeRequest"):
		fmt.Println("Withdraw a CR!!!")
		// return s.WithdrawACR(ctx, *assetAction)

	case (assetAction.ActionId == "P" && assetAction.BCAssetType == "ChangeRequest" && assetAction.ActionName == "InsertACRComment"):
		fmt.Println("Insert a CR comment!!!")
		// return s.InsertACRComment(ctx, assetAction.BCAsset.BCAssetId)

	case (assetAction.ActionId == "C" && assetAction.BCAssetType == "DocumentPackage"):
		fmt.Println("Create Document Package!!!")
		// return s.CreateDP(ctx, *assetAction)

	case (assetAction.ActionId == "PU" && assetAction.BCAssetType == "DocumentPackage"):
		fmt.Println("Update Document Package!!!")
		// return s.InsertDocumentpackagestatus(ctx, *assetAction)

	case (assetAction.ActionId == "R" && assetAction.BCAssetType == "DocumentPackage"):
		fmt.Println("Read Document Package!!!")
		// return s.ReadDP(ctx, assetAction.BCAsset.BCAssetId)

	case (assetAction.ActionId == "C" && assetAction.BCAssetType == "BCDocument"):
		fmt.Println("Create BC Document!!!")
		// return s.CreateBCDocument(ctx, *assetAction)

	case (assetAction.ActionId == "PU" && assetAction.BCAssetType == "BCDocument"):
		fmt.Println("Update BC Document!!!")
		// return s.InsertBCDocumentStatus(ctx, *assetAction)

	case (assetAction.ActionId == "R" && assetAction.BCAssetType == "BCDocument"):
		fmt.Println("Read BC Document!!!")
		// return s.ReadBCDocument(ctx, assetAction.BCAsset.BCAssetId)

	}

	return "", nil
}
