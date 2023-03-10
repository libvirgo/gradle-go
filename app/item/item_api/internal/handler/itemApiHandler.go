package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"micro/app/item/item_api/internal/logic"
	"micro/app/item/item_api/internal/svc"
	"micro/app/item/item_api/internal/types"
)

func Item_apiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewItem_apiLogic(r.Context(), svcCtx)
		resp, err := l.Item_api(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
