package initialization

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"goShop/web_user/global"
)

func InitConfig() {
	pro := GetEnvInfo("GOSHOP_DEBUG")

	v := viper.New()
	configFileName := fmt.Sprintf("web_user/config-debug.yaml")
	if pro {
		configFileName = fmt.Sprintf("web_user/config-pro.yaml")
	}
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	v.Unmarshal(global.NacosConfig)
	zap.S().Infof("配置信息：&v", global.NacosConfig)

	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   global.NacosConfig.Port,
		},
	}

	cc := constant.ClientConfig{
		TimeoutMs:           5000,
		NamespaceId:         global.NacosConfig.Namespace,
		CacheDir:            "tmp/nacos/cache",
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		LogLevel:            "debug",
	}

	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	config, err := client.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group,
	})

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(config), &global.ServerConfig)
	zap.S().Infof("nacos：%s", config)
	if err != nil {
		zap.S().Fatalf("读取 nacos 配置失败：%s", err.Error())
		fmt.Println(&global.ServerConfig)
	}
}

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}
