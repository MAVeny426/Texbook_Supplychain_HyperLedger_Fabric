package contracts

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Textbookcontract contract for managing CRUD operations for Textbook
type Textbookcontract struct {
	contractapi.Contract
}

type Textbook struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
	Price  string `json:"price"`
}

// TextbookExists checks whether a textbook exists
func (t *Textbookcontract) TextbookExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	data, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return data != nil, nil
}

// CreateTextbook can be called only by Manufacturer
func (t *Textbookcontract) CreateTextbook(ctx contractapi.TransactionContextInterface, id, title, author, year, price string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	if clientOrgID != "manufacturer-textbook-com" {
		return "", fmt.Errorf("user under following MSPID: %v can't perform this action", clientOrgID)
	}

	exists, err := t.TextbookExists(ctx, id)
	if err != nil {
		return "", err
	}
	if exists {
		return "", fmt.Errorf("the textbook %s already exists", id)
	}

	textbook := Textbook{
		ID:     id,
		Title:  title,
		Author: author,
		Year:   year,
		Price:  price,
	}

	bytes, _ := json.Marshal(textbook)
	if err := ctx.GetStub().PutState(id, bytes); err != nil {
		return "", err
	}
	return fmt.Sprintf("Successfully added textbook %v", id), nil
}

// ReadTextbook is accessible by all
func (t *Textbookcontract) ReadTextbook(ctx contractapi.TransactionContextInterface, id string) (*Textbook, error) {
	bytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if bytes == nil {
		return nil, fmt.Errorf("the textbook %s does not exist", id)
	}

	var textbook Textbook
	if err := json.Unmarshal(bytes, &textbook); err != nil {
		return nil, fmt.Errorf("could not unmarshal world state data to type Textbook")
	}
	return &textbook, nil
}

// DeleteTextbook can be called only by Manufacturer
func (t *Textbookcontract) DeleteTextbook(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}
	if clientOrgID != "manufacturer-textbook-com" {
		return "", fmt.Errorf("user under following MSPID: %v can't perform this action", clientOrgID)
	}

	exists, err := t.TextbookExists(ctx, id)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", fmt.Errorf("the textbook %s does not exist", id)
	}

	if err := ctx.GetStub().DelState(id); err != nil {
		return "", err
	}
	return fmt.Sprintf("Textbook with id %v has been deleted from the world state", id), nil
}

// UpdateTextbook: Manufacturer can update all fields, Dealer can update price only
func (t *Textbookcontract) UpdateTextbook(ctx contractapi.TransactionContextInterface, id, title, author, year, price string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	exists, err := t.TextbookExists(ctx, id)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", fmt.Errorf("the textbook %s does not exist", id)
	}

	existingBytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		return "", err
	}

	var existingTextbook Textbook
	if err := json.Unmarshal(existingBytes, &existingTextbook); err != nil {
		return "", err
	}

	if clientOrgID == "manufacturer-textbook-com" {
		// Manufacturer can update everything
		existingTextbook.Title = title
		existingTextbook.Author = author
		existingTextbook.Year = year
		existingTextbook.Price = price

	} else if clientOrgID == "dealer-textbook-com" {
		// Dealer can only update the price
		existingTextbook.Price = price

	} else {
		return "", fmt.Errorf("user under following MSPID: %v can't perform this action", clientOrgID)
	}

	updatedBytes, _ := json.Marshal(existingTextbook)
	if err := ctx.GetStub().PutState(id, updatedBytes); err != nil {
		return "", err
	}
	return fmt.Sprintf("Textbook with id %v has been updated", id), nil
}
