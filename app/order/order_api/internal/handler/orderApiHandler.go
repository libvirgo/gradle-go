package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"micro/app/order/order_api/internal/logic"
	"micro/app/order/order_api/internal/svc"
	"micro/app/order/order_api/internal/types"
)

func Order_apiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOrder_apiLogic(r.Context(), svcCtx)
		resp, err := l.Order_api(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
