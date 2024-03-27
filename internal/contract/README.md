
```go
type DIDContractFunction interface {

// IsValidDid 判断DID是否有效
IsValidDid(did string) (bool, error)

// AddDidDocument 添加DID文档
AddDidDocument(didDocument string) error

// GetDidDocument 获取DID文档
GetDidDocument(did string) (string, error)

// GetDidByPubkey 根据公钥获取DID
GetDidByPubkey(pk string) (string, error)

// GetDidByAddress 根据地址获取DID
GetDidByAddress(address string) (string, error)

// UpdateDidDocument 更新DID文档
UpdateDidDocument(didDocument string) error
}
```

