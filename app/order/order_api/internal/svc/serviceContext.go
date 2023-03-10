package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"micro/app/item/item_rpc/itemrpcclient"
	"micro/app/order/order_api/internal/config"
	"micro/app/order/order_rpc/orderrpcclient"
)

type ServiceContext struct {
	Config   config.Config
	ItemRpc  itemrpcclient.ItemRpc
	OrderRpc orderrpcclient.OrderRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		ItemRpc:  itemrpcclient.NewItemRpc(zrpc.MustNewClient(c.ItemRpcConf)),
		OrderRpc: orderrpcclient.NewOrderRpc(zrpc.MustNewClient(c.OrderRpcConf)),
	}
}
