package logic

import (
	"context"
	"fmt"

	"micro/app/item/item_rpc/internal/svc"
	"micro/app/item/item_rpc/item_rpc"

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

func (l *PingLogic) Ping(in *item_rpc.Request) (*item_rpc.Response, error) {
	return &item_rpc.Response{Pong: fmt.Sprintf("pong from rpc:%s", in.GetPing())}, nil
}
