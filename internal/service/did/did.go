package did

import (
	"chainmaker.org/chainmaker/common/v3/crypto"
	"chainmaker.org/chainmaker/common/v3/crypto/asym"
	"chainmaker.org/chainmaker/common/v3/evmutils"
	"encoding/hex"
	"os"
	"time"
)

func getNowString() string {
	currentTime := time.Now().UTC()
	formattedTime := currentTime.Format(time.RFC3339)
	return formattedTime
}

func getPubKeyPem(path string) []byte {
	pem, _ := os.ReadFile(path)
	return pem
}
func getPubKey(pubKeyPath string) crypto.PublicKey {
	pubKey, err := asym.PublicKeyFromPEM(getPubKeyPem(pubKeyPath))
	if err != nil {
		panic(err)
	}
	return pubKey
}
func getPrivateKey(path string) crypto.PrivateKey {
	pem, _ := os.ReadFile(path)
	privKey, err := asym.PrivateKeyFromPEM(pem, nil)
	if err != nil {
		panic(err)
	}
	return privKey
}
func getAddress(pk crypto.PublicKey) string {
	pkBytes, err := evmutils.MarshalPublicKey(pk)
	if err != nil {
		panic(err)
	}
	data := pkBytes[1:]
	bytesAddr := evmutils.Keccak256(data)
	addr := hex.EncodeToString(bytesAddr)[24:]
	return addr
}
