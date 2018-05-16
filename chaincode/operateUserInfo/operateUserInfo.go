package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// SmartContract structure
type SmartContract struct {
}

/*
type MaxNumber struct {
	MaxUserID string `json:"maxUserID"`
}
*/

// User structure
type User struct {
	//UserID   string `json:"userID"`
	UserInfo string `json:"userInfo"`
}

//Init method is called as a result of deployment "operateUserInfo"
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	/*
		var maxNumber = MaxNumber{
			MaxUserID: "0",
		}
		maxNumberAsBytes, _ := json.Marshal(maxNumber)
		APIstub.PutState("maxUserID", maxNumberAsBytes)
	*/
	return shim.Success(nil)
}

//Invoke method is called as a result of an application request to run the Smart Contract "borrowApplication"
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "createUser" {
		return s.createUser(APIstub, args)
	} else if function == "queryUserInfo" {
		return s.queryUserInfo(APIstub, args)
	} else if function == "queryAllUserInfo" {
		return s.queryAllUserInfo(APIstub)
	} else if function == "changeUserInfo" {
		return s.changeUserInfo(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) createUser(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	/*
		//get maxNumber of application
		maxNumberAsBytes, _ := APIstub.GetState("maxUserID")

		maxNumber := MaxNumber{}
		json.Unmarshal(maxNumberAsBytes, &maxNumber)

		var tmpNumber int
		tmpNumber, _ = strconv.Atoi(maxNumber.MaxUserID)
		tmpNumber = tmpNumber + 1

		maxNumber.MaxUserID = strconv.Itoa(tmpNumber)
	*/
	var user = User{
		//	UserID:   args[0],
		UserInfo: args[1],
	}

	//maxNumberAsBytes, _ = json.Marshal(maxNumber)
	userAsBytes, _ := json.Marshal(user)

	//APIstub.PutState("maxUserID", maxNumberAsBytes)
	APIstub.PutState(args[0], userAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) changeUserInfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	/*
		changeUserInfoAsBytes, _ := APIstub.GetState(args[0])
		user := User{}

		json.Unmarshal(changeUserInfoAsBytes, &user)
		user.UserInfo = args[1]

		changeUserInfoAsBytes, _ = json.Marshal(user)
	*/
	APIstub.PutState(args[0], args[1])

	return shim.Success(nil)
}

func (s *SmartContract) queryUserInfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	userAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(userAsBytes)
}

func (s *SmartContract) queryExchangeApplication(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	exchangeApplicationAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(exchangeApplicationAsBytes)
}

func (s *SmartContract) queryAllUserInfo(APIstub shim.ChaincodeStubInterface) sc.Response {
	/*
		startKey := "0"
		endKey := "999"
	*/
	resultsIterator, err := APIstub.GetStateByRange("", "")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllUserInfo:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {
	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
