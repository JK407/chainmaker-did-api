package did

import "testing"

func TestDidService_GetNowString(t *testing.T) {
	t.Log(NewDidService().GetNowString())
}

func TestDidService_GetPubKeyPem(t *testing.T) {
	t.Log(NewDidService().GetPubKeyPem("../../../accounts/user/user.pem"))
}

func TestDidService_GetPubKey(t *testing.T) {
	t.Log(NewDidService().GetPubKey("../../../accounts/user/user.pem"))
}

func TestDidService_GetPrivateKey(t *testing.T) {
	t.Log(NewDidService().GetPrivateKey("../../../accounts/user/user.key"))
}

// f0234677d99856824af0b2214611ac634acda707
func TestDidService_GetAddress(t *testing.T) {
	t.Log(NewDidService().GetAddress(NewDidService().GetPubKey("../../../accounts/user1/user1.pem")))
}
