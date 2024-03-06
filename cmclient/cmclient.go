package cmclient

import (
	sdk "chainmaker.org/chainmaker/sdk-go/v3"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

// CreateCMClient 创建长安链客户端 chainMaker
func CreateCMClient(sdkConfigPath string, opts ...sdk.ChainClientOption) (*sdk.ChainClient, error) {
	newOptions := make([]sdk.ChainClientOption, 0, len(opts)+2)
	newOptions = append(newOptions, sdk.WithConfPath(sdkConfigPath))
	newOptions = append(newOptions, sdk.WithChainClientLogger(NewLogAdapter(logx.WithContext(context.Background()))))
	newOptions = append(newOptions, opts...)
	client, err := sdk.NewChainClient(
		newOptions...,
	)
	if err != nil {
		return nil, err
	}
	_, err = client.GetChainInfo()
	if err != nil {
		return nil, err
	}
	return client, nil
}
