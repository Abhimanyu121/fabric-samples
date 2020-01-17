/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package main

import (
	//"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a car
type SmartContract struct {
	contractapi.Contract
}

// Car describes basic details of what makes up a car


// QueryResult structure used for handling result of query


// InitLedger adds a base set of cars to the ledger
func (s *SmartContract) PutPrescription(ctx contractapi.TransactionContextInterface, adhaarNumber string, prescription string) error {
	in := []byte(prescription)
	//err := ctx.GetStub().PutState("CAR"+strconv.Itoa(i), carAsBytes)
	err := 	ctx.GetStub().PutState("prescription"+adhaarNumber,in)
	if err != nil {
		return fmt.Errorf("Failed to put to put  prescription. %s", err.Error())
	}
	return nil
}
func (s *SmartContract) QueryPrescription(ctx contractapi.TransactionContextInterface, adhaarNumber string) (string, error) {
	prescription, err := ctx.GetStub().GetState("prescription"+adhaarNumber)

	if err != nil {
		return "", fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if prescription == nil {
		return "", fmt.Errorf("%s does not exist", adhaarNumber)
	}

	prescriptionAsString := string(prescription)

	return prescriptionAsString, nil
}

func (s *SmartContract) PutHistory(ctx contractapi.TransactionContextInterface, adhaarNumber string, history string) error {
	in := []byte(history)
	err := 	ctx.GetStub().PutState("history"+adhaarNumber,in)
	if err != nil {
		return fmt.Errorf("Failed to put to put  prescription. %s", err.Error())
	}
	return nil
}
func (s *SmartContract) QueryHistory(ctx contractapi.TransactionContextInterface, adhaarNumber string) (string, error) {
	history, err := ctx.GetStub().GetState("history"+adhaarNumber)

	if err != nil {
		return "", fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if history == nil {
		return "", fmt.Errorf("%s does not exist", adhaarNumber)
	}

	historyAsString := string(history)

	return historyAsString, nil
}
// CreateCar adds a new car to the world state with given details
// func (s *SmartContract) CreateCar(ctx contractapi.TransactionContextInterface, carNumber string, make string, model string, colour string, owner string) error {
// 	car := Car{
// 		Make:   make,
// 		Model:  model,
// 		Colour: colour,
// 		Owner:  owner,
// 	}

// 	carAsBytes, _ := json.Marshal(car)

// 	return ctx.GetStub().PutState(carNumber, carAsBytes)
// }

// QueryCar returns the car stored in the world state with given id


// QueryAllCars returns all cars found in world state
// func (s *SmartContract) QueryAllCars(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
// 	startKey := "CAR0"
// 	endKey := "CAR99"

// 	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resultsIterator.Close()

// 	results := []QueryResult{}

// 	for resultsIterator.HasNext() {
// 		queryResponse, err := resultsIterator.Next()

// 		if err != nil {
// 			return nil, err
// 		}

// 		car := new(Car)
// 		_ = json.Unmarshal(queryResponse.Value, car)

// 		queryResult := QueryResult{Key: queryResponse.Key, Record: car}
// 		results = append(results, queryResult)
// 	}

// 	return results, nil
// }

// ChangeCarOwner updates the owner field of car with given id in world state
// func (s *SmartContract) ChangeCarOwner(ctx contractapi.TransactionContextInterface, carNumber string, newOwner string) error {
// 	car, err := s.QueryCar(ctx, carNumber)

// 	if err != nil {
// 		return err
// 	}

// 	car.Owner = newOwner

// 	carAsBytes, _ := json.Marshal(car)

// 	return ctx.GetStub().PutState(carNumber, carAsBytes)
// }

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabcar chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabcar chaincode: %s", err.Error())
	}
}
