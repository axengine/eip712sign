package eip712

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/signer/core"
	"math/big"
	"testing"
	"time"
)

func TestHashStruct(t *testing.T) {
	var typedData = core.TypedData{
		Types:       typesStandard,
		PrimaryType: primaryTypeStandard,
		Domain:      domainStandard,
		Message:     messageStandard,
	}
	hash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		t.Fatal(err)
	}
	mainHash := fmt.Sprintf("0x%s", common.Bytes2Hex(hash))
	fmt.Println(mainHash)
	if mainHash != "0xc52c0ee5d84264471806290a3f2c4cecfc5490626bf912d01f240d7a274b371e" {
		t.Errorf("Expected different hashStruct result (got %s)", mainHash)
	}

	hash, err = typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		t.Error(err)
	}
	domainHash := fmt.Sprintf("0x%s", common.Bytes2Hex(hash))
	fmt.Println(domainHash)
	if domainHash != "0xf2cee375fa42b42143804025fc449deafd50cc031ca257e0b194a650a912090f" {
		t.Errorf("Expected different domain hashStruct result (got %s)", domainHash)
	}
}

func TestEncodeType(t *testing.T) {
	var typedData = GenClaimTypedData("activty", big.NewInt(97), "", big.NewInt(1), time.Now().Unix())
	domainTypeEncoding := string(typedData.EncodeType("EIP712Domain"))
	if domainTypeEncoding != "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)" {
		t.Errorf("Expected different encodeType result (got %s)", domainTypeEncoding)
	}

	mailTypeEncoding := string(typedData.EncodeType(typedData.PrimaryType))
	if mailTypeEncoding != "Claim(uint256 value,uint256 deadline)" {
		t.Errorf("Expected different encodeType result (got %s)", mailTypeEncoding)
	}
}

func TestSignAndVerifyClaim(t *testing.T) {
	//deadline := time.Now().Unix()
	//fmt.Println(deadline)
	var typedData = GenClaimTypedData("activity", big.NewInt(97), "BD15}{|SD", big.NewInt(1), 1639531996)
	signHash, err := TypedDataHash(typedData)
	if err != nil {
		t.Fatal(err)
	}

	signature, err := Sign(typedData, "8f08198852c63b7894251d940a7b6884bbfc02a8468701ffd246e3c5b4092382")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature))

	b, err := Verify("0x3fcb0d10f7F6589F47c527FDF48582502A0F90EC", hexutil.Encode(signature), signHash)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(b)
}

func TestVerifyClaim(t *testing.T) {
	var typedData = GenClaimTypedData("Claim", big.NewInt(8724), "1", big.NewInt(1), 1640069436)
	signHash, err := TypedDataHash(typedData)
	if err != nil {
		t.Fatal(err)
	}

	b, err := Verify("0xeBf70e73198D6B18ea8026A1864c0f298D4e23A7", "0xeb635cd63d85698db98dd0d199d8841e6e58efd982d5d9fa22e866940a1055fb0701e7b6cd47c3c2e375a843e242862bdc41d4e5ab6a61a08068588a2e21bef81c", signHash)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(b)
}

func TestSignAndVerifyLogin(t *testing.T) {
	var typedData = GenLoginTypedData("activity", big.NewInt(97), "BD15}{|SD", "0xe5DaF2824B43d8b0C961225Ab9992baf39F5F835", time.Now().Unix())
	signHash, err := TypedDataHash(typedData)
	if err != nil {
		t.Fatal(err)
	}

	signature, err := Sign(typedData, "8f08198852c63b7894251d940a7b6884bbfc02a8468701ffd246e3c5b4092382")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature))

	b, err := Verify("0x3fcb0d10f7F6589F47c527FDF48582502A0F90EC", hexutil.Encode(signature), signHash)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(b)
}

func TestSignAndVerifyPermit(t *testing.T) {
	var typedData = GenPermitTypedData("activity", big.NewInt(97), "0xe5DaF2824B43d8b0C961225Ab9992baf39F5F830", "BD15}{|SD",
		"0xe5DaF2824B43d8b0C961225Ab9992baf39F5F831", "0xe5DaF2824B43d8b0C961225Ab9992baf39F5F832", big.NewInt(1), time.Now().Unix())
	signHash, err := TypedDataHash(typedData)
	if err != nil {
		t.Fatal(err)
	}

	signature, err := Sign(typedData, "8f08198852c63b7894251d940a7b6884bbfc02a8468701ffd246e3c5b4092382")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature))

	b, err := Verify("0x3fcb0d10f7F6589F47c527FDF48582502A0F90EC", hexutil.Encode(signature), signHash)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(b)
}
