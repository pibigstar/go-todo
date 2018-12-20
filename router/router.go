package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pibigstar/go-todo/controller"
	"github.com/pibigstar/go-todo/middleware"
)

// Route 路由配置
func Route(router *gin.Engine) {

	api := router.Group("/api", middleware.RefreshTokenCookie)
	{
		api.GET("/wxLogin/:code", controller.WxLogin)
	}
}
