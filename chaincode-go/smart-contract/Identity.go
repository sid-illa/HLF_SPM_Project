package mbse

import (
	"fmt"

	"github.com/hyperledger/fabric-ca/lib/attrmgr"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

/*
//	Get Identity Attributes
*/

// Get a attribute from the current user Identity
func (s *SmartContract) GetIdentityAttribute(ctx contractapi.TransactionContextInterface, reqAttr string) string {
	// get the required attribute
	attributeValue, ok, err := cid.GetAttributeValue(ctx.GetStub(), reqAttr)
	if err != nil {
		// There was an error trying to retrieve the attribute
		fmt.Println("error: ", err)
		return ""
	}
	if !ok {
		// The client identity does not possess the attribute
		fmt.Println("Attribute value not present for :", reqAttr)
		return ""
	}

	return attributeValue
}

// Get all attributes from the current user Identity
func (s *SmartContract) GetAllIdentityAttributes(ctx contractapi.TransactionContextInterface) *attrmgr.Attributes {
	//reading certificate from the stub
	cert, err := ctx.GetClientIdentity().GetX509Certificate()
	if err != nil {
		fmt.Print("error")
	}

	//get attributes from the certificate
	value, err := attrmgr.New().GetAttributesFromCert(cert)
	if err != nil {
		fmt.Print("error")
	}

	return value
}
