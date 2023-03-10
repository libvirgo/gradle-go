package logic

import (
	"context"
	"micro/app/item/item_rpc/item_rpc"

	"micro/app/item/item_api/internal/svc"
	"micro/app/item/item_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.Response, err error) {
	login, err := l.svcCtx.ItemRpc.Login(l.ctx, &item_rpc.LoginRequest{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &types.Response{Message: login.GetPong()}, nil
}
