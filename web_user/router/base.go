package router

import (
	"github.com/gin-gonic/gin"
	"goShop/web_user/api"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("captcha", api.GetCaptcha)
		//BaseRouter.POST("send_sms", api.SendSms)
	}

}
