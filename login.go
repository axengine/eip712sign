package eip712

import (
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core"
	"math/big"
)

var typesLogin = core.Types{
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
	"Login": {
		{
			Name: "account",
			Type: "address",
		},
		{
			Name: "deadline",
			Type: "uint256",
		},
	},
}

// GenLoginTypedData 生成Login的签名数据
func GenLoginTypedData(name string, chainId *big.Int, salt string, account string, deadline int64) core.TypedData {
	return core.TypedData{
		Types:       typesLogin,
		PrimaryType: "Login",
		Domain: core.TypedDataDomain{
			Name:              name,
			Version:           "1",
			ChainId:           math.NewHexOrDecimal256(chainId.Int64()),
			VerifyingContract: "",
			Salt:              salt,
		},
		Message: core.TypedDataMessage{
			"account":  account,
			"deadline": math.NewHexOrDecimal256(deadline),
		},
	}
}
