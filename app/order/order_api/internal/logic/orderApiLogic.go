package logic

import (
	"context"
	"fmt"
	"micro/app/item/item_rpc/item_rpc"
	"micro/app/order/order_rpc/order_rpc"

	"micro/app/order/order_api/internal/svc"
	"micro/app/order/order_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Order_apiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrder_apiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Order_apiLogic {
	return &Order_apiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Order_apiLogic) Order_api(req *types.Request) (resp *types.Response, err error) {
	itemPong, err := l.svcCtx.ItemRpc.Ping(l.ctx, &item_rpc.Request{Ping: "from order api:" + req.Name})
	if err != nil {
		return nil, err
	}
	orderPong, err := l.svcCtx.OrderRpc.Ping(l.ctx, &order_rpc.Request{Ping: "from order api:" + req.Name})
	if err != nil {
		return nil, err
	}
	return &types.Response{
		Message: fmt.Sprintf("item pong:%s, order pong:%s", itemPong, orderPong),
	}, nil
}
