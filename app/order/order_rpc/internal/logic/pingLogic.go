package logic

import (
	"context"
	"fmt"

	"micro/app/order/order_rpc/internal/svc"
	"micro/app/order/order_rpc/order_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *order_rpc.Request) (*order_rpc.Response, error) {
	return &order_rpc.Response{Pong: fmt.Sprintf("pong from order rpc:%s", in.GetPing())}, nil
}
