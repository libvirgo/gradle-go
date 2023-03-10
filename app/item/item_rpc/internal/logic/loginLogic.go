package logic

import (
	"context"

	"micro/app/item/item_rpc/internal/svc"
	"micro/app/item/item_rpc/item_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *item_rpc.LoginRequest) (*item_rpc.Response, error) {
	// 从数据库获取帐号密码并进行用户信息的返回
	if in.GetName() == "test" && in.GetPassword() == "test1pass" {
		return &item_rpc.Response{Pong: "test user is login success"}, nil
	}
	if in.GetName() == "test1" && in.GetPassword() == "test2pass" {
		return &item_rpc.Response{Pong: "test1 user is login success"}, nil
	}
	return &item_rpc.Response{Pong: "login failed, please check name or password"}, nil
}
