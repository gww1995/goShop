package middlewares

import (
	"github.com/gin-gonic/gin"
	"goShop/web_user/models"
	"net/http"
)

func IsAdminAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		claims, _ := context.Get("claim")
		currentUser := claims.(models.CustomClaims)

		if currentUser.AuthorityId != 2 {
			context.JSON(http.StatusForbidden, gin.H{
				"msg": "无权限",
			})
			context.Abort()
			return
		}
		context.Next()
	}

}
