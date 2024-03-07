package constants

import (
	"chainmaker-did-api/internal/types"
	"fmt"
)

const (
	// SuccessCode 通用成功码
	SuccessCode = 200
	// NewAccountSuccessCode NewAccountResponseSuccessCode
	NewAccountSuccessCode = SuccessCode + iota

	// ErrorCode 通用错误码
	ErrorCode = 500
	// MkdirAllErrorCode NewAccountResponseErrorCode
	MkdirAllErrorCode = ErrorCode + iota
	GenerateKeyErrorCode
	ExportPublicKeyErrorCode
)

var (
	// MkdirAllError NewAccountResponseError
	MkdirAllError        = types.NewAccountResponse{}.BuildErrorResponse(MkdirAllErrorCode, fmt.Errorf("创建文件夹失败"))
	GenerateKeyError     = types.NewAccountResponse{}.BuildErrorResponse(GenerateKeyErrorCode, fmt.Errorf("生成私钥失败"))
	ExportPublicKeyError = types.NewAccountResponse{}.BuildErrorResponse(ExportPublicKeyErrorCode, fmt.Errorf("生成公钥失败"))
	// NewAccountSuccess NewAccountResponseSuccess
	NewAccountSuccess = types.NewAccountResponse{}.BuildSuccessResponse(NewAccountSuccessCode, "账户创建成功", types.NewAccountData{})
)
