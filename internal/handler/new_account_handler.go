package handler

import (
	"net/http"

	"chainmaker-did-api/internal/logic"
	"chainmaker-did-api/internal/svc"
	"chainmaker-did-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func NewAccountHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NewAccountRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewNewAccountLogic(r.Context(), svcCtx)
		resp, err := l.NewAccount(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
