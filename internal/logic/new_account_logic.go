package logic

import (
	"chainmaker-did-api/internal/service/key"
	"chainmaker-did-api/internal/svc"
	"chainmaker-did-api/internal/types"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/zeromicro/go-zero/core/logx"
)

type NewAccountLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	keyService key.KeyService
}

func NewNewAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NewAccountLogic {
	return &NewAccountLogic{
		Logger:     logx.WithContext(ctx),
		ctx:        ctx,
		svcCtx:     svcCtx,
		keyService: key.NewKeyService(),
	}
}

func (l *NewAccountLogic) NewAccount(req *types.NewAccountRequest) (resp *types.NewAccountResponse, err error) {
	// 根据名称创建文件夹
	dir := fmt.Sprintf("./%s", req.Name)
	if err = os.MkdirAll(dir, 0755); err != nil {
		resp = &types.NewAccountResponse{
			Code: -1,
			Msg:  err.Error(),
		}
		return resp, nil
	}

	// 生成私钥
	if err = l.keyService.GenerateKey(req.Algo, dir, req.Name); err != nil {
		resp = &types.NewAccountResponse{
			Code: -1,
			Msg:  err.Error(),
		}
		return resp, nil
	}
	prvKeyPath := filepath.Join(dir, fmt.Sprintf("%s.key", req.Name))

	// 生成公钥
	if err = l.keyService.ExportPublicKey(prvKeyPath, dir, req.Name); err != nil {
		resp = &types.NewAccountResponse{
			Code: -1,
			Msg:  err.Error(),
		}
		return resp, nil
	}
	pubKeyPath := filepath.Join(dir, fmt.Sprintf("%s.pem", req.Name))

	// 返回成功响应
	resp = &types.NewAccountResponse{
		Code: 0,
		Msg:  "Account created successfully",
		Data: types.NewAccountData{
			PrivKeyPath: prvKeyPath,
			PubKeyPath:  pubKeyPath,
		},
	}
	return resp, nil
}
