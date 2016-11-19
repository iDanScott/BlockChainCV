package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type BlockCV struct{}

func (bc *BlockCV) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	return nil, nil
}

func (bc *BlockCV) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query is running " + function)
	fmt.Print("Args: ")
	for i := 0; i < len(args); i++ {
		fmt.Print(args[i])
	}

	return nil, errors.New("Function was not understood")
}

func (bc *BlockCV) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Invoke is running " + function)
	fmt.Print("Args: ")
	for i := 0; i < len(args); i++ {
		fmt.Print(args[i])
	}

	return nil, errors.New("Function was not understood")
}

func main() {
	err := shim.Start(new(BlockCV))
	if err != nil {
		fmt.Printf("Error starting BlockCV: %s", err)
	}
}
