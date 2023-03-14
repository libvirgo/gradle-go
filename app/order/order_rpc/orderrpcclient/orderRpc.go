// Code generated by goctl. DO NOT EDIT!
// Source: order_rpc.proto

package orderrpcclient

import (
	"context"

	"micro/app/order/order_rpc/order_rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = order_rpc.Request
	Response = order_rpc.Response

	OrderRpc interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultOrderRpc struct {
		cli zrpc.Client
	}
)

func NewOrderRpc(cli zrpc.Client) OrderRpc {
	return &defaultOrderRpc{
		cli: cli,
	}
}

func (m *defaultOrderRpc) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := order_rpc.NewOrderRpcClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
