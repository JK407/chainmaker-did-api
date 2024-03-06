package key

import (
	"github.com/test-go/testify/assert"
	"testing"
)

func TestKeyService_GenerateKey(t *testing.T) {
	algo := "SM2"
	path := "../../../user"
	name := "user"
	err := generatePrivateKey(algo, path, name)
	assert.NoError(t, err)
}
