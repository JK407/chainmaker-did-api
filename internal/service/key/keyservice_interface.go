package key

type KeyService interface {
	GenerateKey(algo, path, name string) error
	ExportPublicKey(keyPath, path, name string) error
}

type keyService struct{}

func (k keyService) ExportPublicKey(keyPath, path, name string) error {
	return exportPublicKey(keyPath, path, name)
}

func (k keyService) GenerateKey(algo, path, name string) error {
	return generatePrivateKey(algo, path, name)
}

func NewKeyService() KeyService {
	return &keyService{}
}
