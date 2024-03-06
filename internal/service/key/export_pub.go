package key

import (
	"chainmaker.org/chainmaker/common/v3/crypto/asym"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func exportPublicKey(keyPath, path, name string) error {
	key, err := os.ReadFile(keyPath)
	if err != nil {
		return err
	}
	pri, err := asym.PrivateKeyFromPEM(key, nil)
	if err != nil {
		return err
	}
	pubPem, err := pri.PublicKey().String()
	if err != nil {
		return err
	}
	fileName := fmt.Sprintf("%s.pem", name)
	if err = ioutil.WriteFile(filepath.Join(path, fileName),
		[]byte(pubPem), 0600); err != nil {
		return fmt.Errorf("failed to save public key, path = %s,  err = %s", name, err.Error())
	}
	return nil
}
