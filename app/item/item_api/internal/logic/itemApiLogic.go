package logic

import (
	"context"
	"micro/app/item/item_rpc/item_rpc"

	"micro/app/item/item_api/internal/svc"
	"micro/app/item/item_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Item_apiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewItem_apiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Item_apiLogic {
	return &Item_apiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Item_apiLogic) Item_api(req *types.Request) (resp *types.Response, err error) {
	pong, err := l.svcCtx.ItemRpc.Ping(l.ctx, &item_rpc.Request{Ping: "from item api:" + req.Name})
	if err != nil {
		return nil, err
	}
	return &types.Response{
		Message: pong.GetPong(),
	}, nil
}
