package contract

import (
	"chainmaker-did-api/internal/contract/didconstants"
	"chainmaker.org/chainmaker/pb-go/v3/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v3"
	"fmt"
)

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

type DIDContractService struct {
	cmClient     sdk.SDKInterface
	contractName string
}

func (D DIDContractService) IsValidDid(did string) (bool, error) {
	kvs := []*common.KeyValuePair{
		{
			Key:   "did",
			Value: []byte(did),
		},
	}
	r, err := D.cmClient.QueryContract(D.contractName, didconstants.IsValidDid, kvs, -1)
	if err != nil {
		return false, err
	}
	if r.ContractResult.Code != 0 {
		err = fmt.Errorf("query [contractMethod:%s] failed, [code:%d]/[msg:%s]/[result:%s]",
			didconstants.IsValidDid, r.ContractResult.Code, r.ContractResult.Message, r.ContractResult.Result)
		return false, err
	}
	return string(r.ContractResult.Result) == didconstants.True, nil
}

func (D DIDContractService) AddDidDocument(didDocument string) error {
	kvs := []*common.KeyValuePair{
		{
			Key:   "didDocument",
			Value: []byte(didDocument),
		},
	}
	gasLimit := &common.Limit{
		GasLimit: didconstants.GasLimit,
	}
	r, err := D.cmClient.InvokeContractWithLimit(D.contractName, didconstants.AddDidDocument,
		"", kvs, -1, true, gasLimit)
	if err != nil {
		return err
	}
	if r.ContractResult.Code != 0 {
		err = fmt.Errorf("invoke [contractMethod:%s] failed, [code:%d]/[msg:%s]/[result:%s]",
			didconstants.AddDidDocument, r.ContractResult.Code, r.ContractResult.Message, r.ContractResult.Result)
		return err
	}
	return nil
}

func (D DIDContractService) GetDidDocument(did string) (string, error) {
	kvs := []*common.KeyValuePair{
		{
			Key:   "did",
			Value: []byte(did),
		},
	}
	r, err := D.cmClient.QueryContract(D.contractName, didconstants.GetDidDocument, kvs, -1)
	if err != nil {
		return "", err
	}
	if r.ContractResult.Code != 0 {
		err = fmt.Errorf("query [contractMethod:%s] failed, [code:%d]/[msg:%s]/[result:%s]",
			didconstants.GetDidDocument, r.ContractResult.Code, r.ContractResult.Message, r.ContractResult.Result)
		return "", err
	}
	return string(r.ContractResult.Result), nil
}

func (D DIDContractService) GetDidByPubkey(pk string) (string, error) {
	kvs := []*common.KeyValuePair{
		{
			Key:   "pubKey",
			Value: []byte(pk),
		},
	}
	r, err := D.cmClient.QueryContract(D.contractName, didconstants.GetDidByPubkey, kvs, -1)
	if err != nil {
		return "", err
	}
	if r.ContractResult.Code != 0 {
		err = fmt.Errorf("query [contractMethod:%s] failed, [code:%d]/[msg:%s]/[result:%s]",
			didconstants.GetDidByPubkey, r.ContractResult.Code, r.ContractResult.Message, r.ContractResult.Result)
		return "", err
	}
	return string(r.ContractResult.Result), nil
}

func (D DIDContractService) GetDidByAddress(address string) (string, error) {
	kvs := []*common.KeyValuePair{
		{
			Key:   "address",
			Value: []byte(address),
		},
	}
	r, err := D.cmClient.QueryContract(D.contractName, didconstants.GetDidByAddress, kvs, -1)
	if err != nil {
		return "", err
	}
	if r.ContractResult.Code != 0 {
		err = fmt.Errorf("query [contractMethod:%s] failed, [code:%d]/[msg:%s]/[result:%s]",
			didconstants.GetDidByAddress, r.ContractResult.Code, r.ContractResult.Message, r.ContractResult.Result)
		return "", err
	}
	return string(r.ContractResult.Result), nil
}

func (D DIDContractService) UpdateDidDocument(didDocument string) error {
	kvs := []*common.KeyValuePair{
		{
			Key:   "didDocument",
			Value: []byte(didDocument),
		},
	}
	gasLimit := &common.Limit{GasLimit: didconstants.GasLimit}
	r, err := D.cmClient.InvokeContractWithLimit(D.contractName, didconstants.UpdateDidDocument,
		"", kvs, -1, true, gasLimit)
	if err != nil {
		return err
	}
	if r.ContractResult.Code != 0 {
		err = fmt.Errorf("invoke [contractMethod:%s] failed, [code:%d]/[msg:%s]/[result:%s]",
			didconstants.UpdateDidDocument, r.ContractResult.Code, r.ContractResult.Message, r.ContractResult.Result)
		return err
	}
	return nil
}

func NewDIDContractService(cmClient sdk.SDKInterface, contractName string) DIDContractFunction {
	return &DIDContractService{
		cmClient:     cmClient,
		contractName: contractName,
	}
}
