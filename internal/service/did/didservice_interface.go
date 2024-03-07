package did

import "chainmaker.org/chainmaker/common/v3/crypto"

type DidService interface {
	GetNowString() string
	GetPubKeyPem(path string) []byte
	GetPubKey(pubKeyPath string) crypto.PublicKey
	GetPrivateKey(path string) crypto.PrivateKey
	GetAddress(pk crypto.PublicKey) string
}

type didService struct{}

func (d didService) GetNowString() string {
	return getNowString()
}

func (d didService) GetPubKeyPem(path string) []byte {
	return getPubKeyPem(path)
}

func (d didService) GetPubKey(pubKeyPath string) crypto.PublicKey {
	return getPubKey(pubKeyPath)
}

func (d didService) GetPrivateKey(path string) crypto.PrivateKey {
	return getPrivateKey(path)
}

func (d didService) GetAddress(pk crypto.PublicKey) string {
	return getAddress(pk)
}

func NewDidService() DidService {
	return &didService{}
}
