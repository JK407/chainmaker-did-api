package key

import (
	"chainmaker.org/chainmaker/common/v3/cert"
	"chainmaker.org/chainmaker/common/v3/crypto"
	"fmt"
	"strings"
)

func generatePrivateKey(algo, path, name string) error {
	a := strings.ToUpper(algo)
	switch a {
	case "SM2", "ECC_P256":
		if keyType, ok := crypto.AsymAlgoMap[a]; ok {
			fileName := fmt.Sprintf("%s.key", name)
			_, err := cert.CreatePrivKey(keyType, path, fileName, true)
			return err
		}
		return fmt.Errorf("unsupported algorithm %s", algo)
	default:
		return fmt.Errorf("unsupported algorithm %s", algo)
	}
}
