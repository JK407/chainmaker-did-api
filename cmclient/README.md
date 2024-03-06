# 使用Go-SDK开发ChainMakerClient


## 如何使用

### 1.配置文件
将客户端的配置文件放在testdata中

### 2.在go-zero的etc中的config.yaml中加入以下配置
```yaml
SdkConfigPath: ../testdata/sdk_config.yml
```

### 3.在config.go中加入以下配置
```go
type Config struct {
    SdkConfigPath string
}
```

### 4.在service_context.go中加入以下代码
```go
type ServiceContext struct {
	Config config.Config
	CMClient *sdk.ChainClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	cmClient, err := cmclient.CreateCMClient(c.SdkConfigPath)
	if err != nil {
		panic("create client Failed,error:" + err.Error())
	}
	return &ServiceContext{
		Config: c,
		CMClient: cmClient,
	}
}
```

### 这样就可以在logic中通过logic将client传入到service中