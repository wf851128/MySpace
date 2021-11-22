package routes

import (
	"MySpace/controllers"
	"MySpace/logger"
	"MySpace/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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

	r.POST("/ping", JWTAuthorizationMiddleware(), func(context *gin.Context) {
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

func JWTAuthorizationMiddleware() func(context *gin.Context) {
	return func(context *gin.Context) {
		//	通常 token 存放在1.请求头 2.请求体 3.URL
		//	如果在请求头中，格式为 Bearer eyJhbG.....1NiI
		authHeader := context.Request.Header.Get("Authorization")
		if authHeader == "" {
			context.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中 Authorization 为空",
			})
			context.Abort()
			return
		}
		//对authorization 进行分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			context.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中 Authorization 格式有误",
			})
			context.Abort()
			return
		}
		//parts[1]是对获取到的 tokenString分割后的记过，使用 jwt 进行解析
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "请求头中 Authorization 无效",
			})
			context.Abort()
			return
		}
		//将当前请求的 userID 保存到请求上下文context上
		context.Set("userID", mc.UserID)
		//后续处理可以通过 context.get("userID")来获取当前请求用户信息
		context.Next()
	}
}
