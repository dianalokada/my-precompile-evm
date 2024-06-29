// Code generated
// This file is a generated precompile contract config with stubbed abstract functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

package bls

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/precompile/contract"
	"github.com/ava-labs/subnet-evm/vmerrs"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"

	_ "embed"

	"github.com/ethereum/go-ethereum/common"
	
)

const (
	// Gas costs for each function. These are set to 1 by default.
	// You should set a gas cost for each function in your contract.
	// Generally, you should not set gas costs very low as this may cause your network to be vulnerable to DoS attacks.
	// There are some predefined gas costs in contract/utils.go that you can use.
	VerifyAndFormatBLSSignatureGasCost uint64 = 50000 /* SET A GAS COST HERE */
	VerifyBLSSignatureGasCost          uint64 = 40000 /* SET A GAS COST HERE */
)

// CUSTOM CODE STARTS HERE
// Reference imports to suppress errors from unused imports. This code and any unnecessary imports can be removed.
var (
	_ = abi.JSON
	_ = errors.New
	_ = big.NewInt
	_ = vmerrs.ErrOutOfGas
	_ = common.Big0
)

// Singleton StatefulPrecompiledContract and signatures.
var (

	// BlsRawABI contains the raw ABI of Bls contract.
	//go:embed contract.abi
	BlsRawABI string

	BlsABI = contract.ParseABI(BlsRawABI)

	BlsPrecompile = createBlsPrecompile()
)

type VerifyAndFormatBLSSignatureInput struct {
	Message   string
	Signature []byte
	PublicKey []byte
}

type VerifyBLSSignatureInput struct {
	Message   string
	Signature []byte
	PublicKey []byte
}

// UnpackVerifyAndFormatBLSSignatureInput attempts to unpack [input] as VerifyAndFormatBLSSignatureInput
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackVerifyAndFormatBLSSignatureInput(input []byte) (VerifyAndFormatBLSSignatureInput, error) {
	inputStruct := VerifyAndFormatBLSSignatureInput{}
	err := BlsABI.UnpackInputIntoInterface(&inputStruct, "verifyAndFormatBLSSignature", input, false)

	return inputStruct, err
}

// PackVerifyAndFormatBLSSignature packs [inputStruct] of type VerifyAndFormatBLSSignatureInput into the appropriate arguments for verifyAndFormatBLSSignature.
func PackVerifyAndFormatBLSSignature(inputStruct VerifyAndFormatBLSSignatureInput) ([]byte, error) {
	return BlsABI.Pack("verifyAndFormatBLSSignature", inputStruct.Message, inputStruct.Signature, inputStruct.PublicKey)
}

// PackVerifyAndFormatBLSSignatureOutput attempts to pack given result of type string
// to conform the ABI outputs.
func PackVerifyAndFormatBLSSignatureOutput(result string) ([]byte, error) {
	return BlsABI.PackOutput("verifyAndFormatBLSSignature", result)
}

// UnpackVerifyAndFormatBLSSignatureOutput attempts to unpack given [output] into the string type output
// assumes that [output] does not include selector (omits first 4 func signature bytes)
func UnpackVerifyAndFormatBLSSignatureOutput(output []byte) (string, error) {
	res, err := BlsABI.Unpack("verifyAndFormatBLSSignature", output)
	if err != nil {
		return "", err
	}
	unpacked := *abi.ConvertType(res[0], new(string)).(*string)
	return unpacked, nil
}

func verifyAndFormatBLSSignature(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
    if remainingGas, err = contract.DeductGas(suppliedGas, VerifyAndFormatBLSSignatureGasCost); err != nil {
        return nil, 0, err
    }
    inputStruct, err := UnpackVerifyAndFormatBLSSignatureInput(input)
    if err != nil {
        return nil, remainingGas, err
    }

    // Verify the signature
    publicKey, err := bls.PublicKeyFromBytes(inputStruct.PublicKey)
    if err != nil {
        return nil, remainingGas, err
    }
    signature, err := bls.SignatureFromBytes(inputStruct.Signature)
    if err != nil {
        return nil, remainingGas, err
    }
    messageHash := bls.Hash([]byte(inputStruct.Message))
    
    if !bls.Verify(publicKey, signature, messageHash) {
        return nil, remainingGas, errors.New("invalid signature")
    }

    // Format the signature
    formattedSignature := signature.String()

    packedOutput, err := PackVerifyAndFormatBLSSignatureOutput(formattedSignature)
    if err != nil {
        return nil, remainingGas, err
    }

    return packedOutput, remainingGas, nil
}

// UnpackVerifyBLSSignatureInput attempts to unpack [input] as VerifyBLSSignatureInput
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackVerifyBLSSignatureInput(input []byte) (VerifyBLSSignatureInput, error) {
	inputStruct := VerifyBLSSignatureInput{}
	err := BlsABI.UnpackInputIntoInterface(&inputStruct, "verifyBLSSignature", input, false)

	return inputStruct, err
}

// PackVerifyBLSSignature packs [inputStruct] of type VerifyBLSSignatureInput into the appropriate arguments for verifyBLSSignature.
func PackVerifyBLSSignature(inputStruct VerifyBLSSignatureInput) ([]byte, error) {
	return BlsABI.Pack("verifyBLSSignature", inputStruct.Message, inputStruct.Signature, inputStruct.PublicKey)
}

// PackVerifyBLSSignatureOutput attempts to pack given isValid of type bool
// to conform the ABI outputs.
func PackVerifyBLSSignatureOutput(isValid bool) ([]byte, error) {
	return BlsABI.PackOutput("verifyBLSSignature", isValid)
}

// UnpackVerifyBLSSignatureOutput attempts to unpack given [output] into the bool type output
// assumes that [output] does not include selector (omits first 4 func signature bytes)
func UnpackVerifyBLSSignatureOutput(output []byte) (bool, error) {
	res, err := BlsABI.Unpack("verifyBLSSignature", output)
	if err != nil {
		return false, err
	}
	unpacked := *abi.ConvertType(res[0], new(bool)).(*bool)
	return unpacked, nil
}

func verifyBLSSignature(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
    if remainingGas, err = contract.DeductGas(suppliedGas, VerifyBLSSignatureGasCost); err != nil {
        return nil, 0, err
    }
    inputStruct, err := UnpackVerifyBLSSignatureInput(input)
    if err != nil {
        return nil, remainingGas, err
    }

    // Verify the signature
    publicKey, err := bls.PublicKeyFromBytes(inputStruct.PublicKey)
    if err != nil {
        return nil, remainingGas, err
    }
    signature, err := bls.SignatureFromBytes(inputStruct.Signature)
    if err != nil {
        return nil, remainingGas, err
    }
    messageHash := bls.Hash([]byte(inputStruct.Message))
    
    isValid := bls.Verify(publicKey, signature, messageHash)

    packedOutput, err := PackVerifyBLSSignatureOutput(isValid)
    if err != nil {
        return nil, remainingGas, err
    }

    return packedOutput, remainingGas, nil
}

// createBlsPrecompile returns a StatefulPrecompiledContract with getters and setters for the precompile.

func createBlsPrecompile() contract.StatefulPrecompiledContract {
	var functions []*contract.StatefulPrecompileFunction

	abiFunctionMap := map[string]contract.RunStatefulPrecompileFunc{
		"verifyAndFormatBLSSignature": verifyAndFormatBLSSignature,
		"verifyBLSSignature":          verifyBLSSignature,
	}

	for name, function := range abiFunctionMap {
		method, ok := BlsABI.Methods[name]
		if !ok {
			panic(fmt.Errorf("given method (%s) does not exist in the ABI", name))
		}
		functions = append(functions, contract.NewStatefulPrecompileFunction(method.ID, function))
	}
	// Construct the contract with no fallback function.
	statefulContract, err := contract.NewStatefulPrecompileContract(nil, functions)
	if err != nil {
		panic(err)
	}
	return statefulContract
}