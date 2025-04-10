package global

import (
	ut "github.com/go-playground/universal-translator"
	"goShop/web_order/config"
	"goShop/web_order/proto"
)

var (
	Trans ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	NacosConfig *config.NacosConfig = &config.NacosConfig{}

	GoodsSrvClient web_order.GoodsClient

	OrderSrvClient web_order.OrderClient

	InventorySrvClient web_order.InventoryClient
)
