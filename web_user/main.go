package main

import (
	"fmt"
	uuid2 "github.com/google/uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"goShop/web_user/global"
	"goShop/web_user/initialization"
	"goShop/web_user/utils"
	"goShop/web_user/utils/register/consul"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//1.初始化
	initialization.InitLogger()

	initialization.InitConfig()

	Router := initialization.Routers()

	initialization.InitSrvConn()

	viper.AutomaticEnv()

	debug := viper.GetBool("GOSHOP_DEBUG")
	if !debug {
		port, err := utils.GetFreePort()
		if err != nil {
			global.ServerConfig.Port = port
		}
	}

	client := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)

	serviceId := fmt.Sprintf("%s", uuid2.New())
	err := client.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败", err.Error())
	}

	zap.S().Debugf("启动服务器，端口：%d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败：", err.Error())
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

}
