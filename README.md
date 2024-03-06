
goctl api -o chainmaker.api

goctl api go --api .\api\chainmaker.api --style=go_zero -dir ./




### 模块

- [x] NewAccount // 新建账户 
- > 新建账户接口 <font face="隶书" color=#be905c size=3>**POST**</font>
```shell
// POST
curl --location '127.0.0.1:8888/account/new' \
--form 'algo="SM2"' \
--form 'name="user"'