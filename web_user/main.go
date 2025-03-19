package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	uuid2 "github.com/google/uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"goShop/web_user/global"
	"goShop/web_user/initialization"
	"goShop/web_user/utils"
	"goShop/web_user/utils/register/consul"
	vaild "goShop/web_user/validator"
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

	pro := viper.GetBool("GOSHOP_DEBUG")
	if pro {
		port, err := utils.GetFreePort()
		if err != nil {
			global.ServerConfig.Port = port
		}
	}

	//初始化验证器的翻译
	if err := initialization.InitTrans("zh"); err != nil {
		panic(err)
	}
	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", vaild.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
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
