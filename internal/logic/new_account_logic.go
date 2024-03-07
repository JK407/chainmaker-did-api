package logic

import (
	"chainmaker-did-api/constants"
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
	dir := fmt.Sprintf("./accounts/%s", req.Name)
	if err = os.MkdirAll(dir, 0755); err != nil {
		l.Logger.Errorf("MkdirAll error: ", err)
		return &constants.MkdirAllError, nil
	}

	// 生成私钥
	if err = l.keyService.GenerateKey(req.Algo, dir, req.Name); err != nil {
		l.Logger.Error("GenerateKey error: ", err)
		return &constants.GenerateKeyError, nil
	}
	//prvKeyPath := filepath.Join(dir, fmt.Sprintf("%s.key", req.Name))
	prvKeyPath := filepath.ToSlash(filepath.Join(dir, fmt.Sprintf("%s.key", req.Name)))

	// 生成公钥
	if err = l.keyService.ExportPublicKey(prvKeyPath, dir, req.Name); err != nil {
		l.Logger.Errorf("ExportPublicKey error: ", err)
		return &constants.ExportPublicKeyError, nil
	}
	//pubKeyPath := filepath.Join(dir, fmt.Sprintf("%s.pem", req.Name))
	pubKeyPath := filepath.ToSlash(filepath.Join(dir, fmt.Sprintf("%s.pem", req.Name)))
	l.Logger.Infof("success to create new account.")
	r := constants.NewAccountSuccess.SetData(types.NewAccountData{
		PrivKeyPath: prvKeyPath,
		PubKeyPath:  pubKeyPath,
	})
	return &r, nil
}
