package eip712

import (
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core"
	"math/big"
)

var typesPermit = core.Types{
	"EIP712Domain": {
		{
			Name: "name",
			Type: "string",
		},
		{
			Name: "version",
			Type: "string",
		},
		{
			Name: "chainId",
			Type: "uint256",
		},
		{
			Name: "verifyingContract",
			Type: "address",
		},
		{
			Name: "salt",
			Type: "string",
		},
	},
	"Permit": {
		{
			Name: "owner",
			Type: "address",
		},
		{
			Name: "spender",
			Type: "address",
		},
		{
			Name: "value",
			Type: "uint256",
		},
		{
			Name: "deadline",
			Type: "uint256",
		},
	},
}

// GenPermitTypedData 生成Permit的签名数据
func GenPermitTypedData(name string, chainId *big.Int, verifyingContract string, salt string,
	owner, spender string, value *big.Int, deadline int64) core.TypedData {
	valueBig := (math.HexOrDecimal256)(*value)
	return core.TypedData{
		Types:       typesPermit,
		PrimaryType: "Permit",
		Domain: core.TypedDataDomain{
			Name:              name,
			Version:           "1",
			ChainId:           math.NewHexOrDecimal256(chainId.Int64()),
			VerifyingContract: verifyingContract,
			Salt:              salt,
		},
		Message: core.TypedDataMessage{
			"owner":    owner,
			"spender":  spender,
			"value":    &valueBig,
			"deadline": math.NewHexOrDecimal256(deadline),
		},
	}
}
