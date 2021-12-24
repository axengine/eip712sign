package eip712

import (
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core"
	"math/big"
)

var typesClaim = core.Types{
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
			Name: "salt",
			Type: "string",
		},
	},
	"Claim": {
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

// GenClaimTypedData 生成Claim的签名数据
// name 活动名称,业务协商
func GenClaimTypedData(name string, chainId *big.Int, salt string, value *big.Int, deadline int64) core.TypedData {
	return core.TypedData{
		Types:       typesClaim,
		PrimaryType: "Claim",
		Domain: core.TypedDataDomain{
			Name:              name,
			Version:           "1",
			ChainId:           math.NewHexOrDecimal256(chainId.Int64()),
			VerifyingContract: "",
			Salt:              salt,
		},
		Message: core.TypedDataMessage{
			"value":    (*math.HexOrDecimal256)(value),
			"deadline": math.NewHexOrDecimal256(deadline),
		},
	}
}
