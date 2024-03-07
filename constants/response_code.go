package constants

import (
	"chainmaker-did-api/internal/types"
	"fmt"
)

const (
	// 通用成功码
	SuccessCode = 200
	//NewAccountResponseSuccessCode
	NewAccountSuccessCode = SuccessCode + iota

	// 通用错误码
	ErrorCode = 500
	// NewAccountResponseErrorCode
	MkdirAllErrorCode = ErrorCode + iota
	GenerateKeyErrorCode
	ExportPublicKeyErrorCode
)

var (
	// NewAccountResponseError
	MkdirAllError        = types.NewAccountResponse{}.BuildErrorResponse(MkdirAllErrorCode, fmt.Errorf("创建文件夹失败"))
	GenerateKeyError     = types.NewAccountResponse{}.BuildErrorResponse(GenerateKeyErrorCode, fmt.Errorf("生成私钥失败"))
	ExportPublicKeyError = types.NewAccountResponse{}.BuildErrorResponse(ExportPublicKeyErrorCode, fmt.Errorf("生成公钥失败"))
	// NewAccountResponseSuccess
	NewAccountSuccess = types.NewAccountResponse{}.BuildSuccessResponse(NewAccountSuccessCode, "账户创建成功", types.NewAccountData{})
)
