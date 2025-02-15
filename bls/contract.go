// Code generated
// This file is a generated precompile contract config with stubbed abstract functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

package bls

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/precompile/contract"
	"github.com/ava-labs/subnet-evm/vmerrs"

	_ "embed"

	"github.com/ethereum/go-ethereum/common"
)

const (
	// Gas costs for each function. These are set to 1 by default.
	// You should set a gas cost for each function in your contract.
	// Generally, you should not set gas costs very low as this may cause your network to be vulnerable to DoS attacks.
	// There are some predefined gas costs in contract/utils.go that you can use.
	
	// 40.000 is a moderate to high computational cost for a cryptographic operation like BLS signature verification.
	VerifyBLSSignatureGasCost uint64 = 40000 /* SET A GAS COST HERE */
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

type VerifyBLSSignatureInput struct {
	Message   string
	Signature []byte
	PublicKey []byte
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
	publicKey, err := bls.PublicKeyFromCompressedBytes(inputStruct.PublicKey)
	if err != nil {
		return nil, remainingGas, err
	}
	signature, err := bls.SignatureFromBytes(inputStruct.Signature)
	if err != nil {
		return nil, remainingGas, err
	}
	messageHash := []byte(inputStruct.Message)

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
		"verifyBLSSignature": verifyBLSSignature,
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
