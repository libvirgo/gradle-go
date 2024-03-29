// Code generated by goctl. DO NOT EDIT.
// Source: item_rpc.proto

package itemrpcclient

import (
	"context"

	"micro/app/item/item_rpc/item_rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginRequest    = item_rpc.LoginRequest
	RegisterRequest = item_rpc.RegisterRequest
	Request         = item_rpc.Request
	Response        = item_rpc.Response

	ItemRpc interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error)
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Response, error)
	}

	defaultItemRpc struct {
		cli zrpc.Client
	}
)

func NewItemRpc(cli zrpc.Client) ItemRpc {
	return &defaultItemRpc{
		cli: cli,
	}
}

func (m *defaultItemRpc) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := item_rpc.NewItemRpcClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultItemRpc) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error) {
	client := item_rpc.NewItemRpcClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultItemRpc) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Response, error) {
	client := item_rpc.NewItemRpcClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}
