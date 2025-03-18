package global

import (
	ut "github.com/go-playground/universal-translator"
	"goShop/web_user/proto"

	"goShop/web_user/config"
)

var (
	Trans         ut.Translator
	ServerConfig  *config.ServerConfig = &config.ServerConfig{}
	NacosConfig   *config.NacosConfig  = &config.NacosConfig{}
	UserSrvClient proto.UserClient
)
