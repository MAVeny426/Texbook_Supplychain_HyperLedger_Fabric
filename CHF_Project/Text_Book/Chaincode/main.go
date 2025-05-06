package main

import (
	"textbook/contracts"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Textbookcontract struct {
	contractapi.Contract
}

func main() {
	textbookcontract := new(contracts.Textbookcontract)

	chaincode, err := contractapi.NewChaincode(textbookcontract)
	// cc, err := contractapi.NewChaincode(new(Textbookcontract))
	if err != nil {
		panic("Error creating chaincode: " + err.Error())
	}

	if err := chaincode.Start(); err != nil {
		panic("Error starting chaincode: " + err.Error())
	}
}
