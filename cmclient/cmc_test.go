package cmclient

import (
	"github.com/test-go/testify/assert"
	"testing"
)

const sdkConfigOrg1Client1Path = "../chainconfig/sdk_config.yml"

func TestCMC(t *testing.T) {
	cmc, err := CreateCMClient(sdkConfigOrg1Client1Path)
	assert.NoError(t, err)
	t.Logf("ChainMakerClient:%v", cmc)
}
