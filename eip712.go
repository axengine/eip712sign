package eip712

import (
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core"
	"strings"
)

// TypedDataHash 计算typedData的签名hash
func TypedDataHash(typedData core.TypedData) ([]byte, error) {
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return nil, err
	}

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return nil, err
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	sighash := crypto.Keccak256(rawData)
	return sighash, nil
}

// Sign 对typedData签名
func Sign(typedData core.TypedData, privateKey string) ([]byte, error) {
	sigHash, err := TypedDataHash(typedData)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(privateKey, "0x") {
		privateKey = string([]byte(privateKey)[2:])
	}

	private, err := crypto.ToECDSA(ethcmn.Hex2Bytes(privateKey))
	if err != nil {
		return nil, err
	}

	signature, err := btcec.SignCompact(btcec.S256(), (*btcec.PrivateKey)(private), sigHash, false)
	if err != nil {
		return nil, err
	}

	// Convert to Ethereum signature format with 'recovery id' v at the end.
	v := signature[0]
	copy(signature, signature[1:])
	signature[64] = v
	return signature, nil
}

// Verify 验签
func Verify(address, signature string, hash []byte) (bool, error) {
	signer, err := recover(signature, hash)
	if err != nil {
		return false, err
	}

	if ethcmn.HexToAddress(address) != signer {
		return false, fmt.Errorf("addresses do not match")
	}

	return true, nil
}

// Recover 恢复出signer
func Recover(signature string, hash []byte) (ethcmn.Address, error) {
	return recover(signature, hash)
}

func recover(signature string, hash []byte) (ethcmn.Address, error) {
	signatureByte, err := hexutil.Decode(signature)
	if err != nil {
		return ethcmn.Address{}, err
	}
	if len(signatureByte) != 65 {
		return ethcmn.Address{}, fmt.Errorf("invalid signature length: %d", len(signatureByte))
	}

	if signatureByte[64] != 27 && signatureByte[64] != 28 {
		return ethcmn.Address{}, fmt.Errorf("invalid recovery id: %d", signatureByte[64])
	}
	signatureByte[64] -= 27

	pubKey, err := crypto.SigToPub(hash, signatureByte)
	if err != nil {
		return ethcmn.Address{}, fmt.Errorf("invalid signature: %s", err.Error())
	}

	return crypto.PubkeyToAddress(*pubKey), nil
}
