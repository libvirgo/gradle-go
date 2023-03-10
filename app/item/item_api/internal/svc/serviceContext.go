package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"micro/app/item/item_api/internal/config"
	"micro/app/item/item_rpc/itemrpcclient"
)

type ServiceContext struct {
	Config  config.Config
	ItemRpc itemrpcclient.ItemRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		ItemRpc: itemrpcclient.NewItemRpc(zrpc.MustNewClient(c.ItemRpcConf)),
	}
}
