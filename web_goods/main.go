package main

import (
	"fmt"
	uuid2 "github.com/google/uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"goShop/web_goods/global"
	"goShop/web_goods/initialization"

	"goShop/web_user/utils"
	"goShop/web_user/utils/register/consul"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//1.初始化
	initializetion.InitLogger()

	initializetion.InitConfig()

	Router := initializetion.Routers()

	if err := initializetion.InitTrans("zh"); err != nil {
		panic(err)
	}

	initializetion.InitSrvConn()
	viper.AutomaticEnv()

	pro := viper.GetBool("GOSHOP_DEBUG")
	if pro {
		port, err := utils.GetFreePort()
		if err != nil {
			global.ServerConfig.Port = port
		}
	}

	//Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:85.0) Gecko/20100101 Firefox/85.0
	//scrapy requests
	//Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.190 Safari/537.36
	/*
		1. S()可以获取一个全局的sugar，可以让我们自己设置一个全局的logger
		2. 日志是分级别的，debug， info ， warn， error， fetal
		3. S函数和L函数很有用， 提供了一个全局的安全访问logger的途径
	*/
	registerClient := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid2.New())
	err := registerClient.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败:", err.Error())
	}
	zap.S().Debugf("启动服务器, 端口： %d", global.ServerConfig.Port)
	go func() {
		if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
			zap.S().Panic("启动失败:", err.Error())
		}
	}()

	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = registerClient.DeRegister(serviceId); err != nil {
		zap.S().Info("注销失败:", err.Error())
	} else {
		zap.S().Info("注销成功:")
	}

}
