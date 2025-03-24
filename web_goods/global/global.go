package global

import (
	ut "github.com/go-playground/universal-translator"
	"goShop/web_goods/config"
	"goShop/web_goods/proto"
)

var (
	Trans ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	NacosConfig *config.NacosConfig = &config.NacosConfig{}

	GoodsSrvClient web_goods.GoodsClient
)
