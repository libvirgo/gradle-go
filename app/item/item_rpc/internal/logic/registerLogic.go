package logic

import (
	"context"

	"micro/app/item/item_rpc/internal/svc"
	"micro/app/item/item_rpc/item_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *item_rpc.RegisterRequest) (*item_rpc.Response, error) {
	// todo: add your logic here and delete this line

	return &item_rpc.Response{}, nil
}
