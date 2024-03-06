package svc

import (
	"chainmaker-did-api/cmclient"
	"chainmaker-did-api/internal/config"
	sdk "chainmaker.org/chainmaker/sdk-go/v3"
)

type ServiceContext struct {
	Config   config.Config
	CmClient *sdk.ChainClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	cmc, e := cmclient.CreateCMClient(c.SdkConfigPath)
	if e != nil {
		panic(e)
	}
	//_, e = cmc.GetChainInfo()
	//if e != nil {
	//	panic(e)
	//}
	return &ServiceContext{
		Config:   c,
		CmClient: cmc,
	}
}
