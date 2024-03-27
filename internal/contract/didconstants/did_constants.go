package didconstants

const SdkConfigPath = "./testdata/sdk_config_pk_test.yml"

const (
	ContractName      = "DID1"
	DidMethod         = "DidMethod"         //获取DID Method
	IsValidDid        = "IsValidDid"        //判断DID URL是否合法
	AddDidDocument    = "AddDidDocument"    //添加DID文档
	GetDidDocument    = "GetDidDocument"    //根据DID URL获取DID文档
	GetDidByPubkey    = "GetDidByPubkey"    //根据公钥获取DID URL
	GetDidByAddress   = "GetDidByAddress"   //根据地址获取DID URL
	UpdateDidDocument = "UpdateDidDocument" //更新DID文档

	GasLimit              = 999999999
	DefaultDelegateAction = "sign"
	True                  = "true"
)
