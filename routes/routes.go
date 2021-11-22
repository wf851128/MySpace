package routes

import (
	"MySpace/controllers"
	"MySpace/logger"
	"MySpace/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup(mode string) (g *gin.Engine) {
	//判断成本版本，设定 gin 的模式
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//注册业务
	r.POST("/signup", controllers.SignUpHandler)
	//登录业务
	r.POST("/login", controllers.LoginHandler)

	r.POST("/ping", middlewares.JWTAuthorizationMiddleware(), func(context *gin.Context) {
		//判断请求头中是否有有效的 JWT token
		context.String(http.StatusOK, "pong")
	})
	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
