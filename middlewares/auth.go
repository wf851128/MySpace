package middlewares

import (
	"MySpace/controllers"
	"MySpace/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

const UseridKey = "userID"

func JWTAuthorizationMiddleware() func(context *gin.Context) {
	return func(context *gin.Context) {
		//	通常 token 存放在1.请求头 2.请求体 3.URL
		//	如果在请求头中，格式为 Bearer eyJhbG.....1NiI
		authHeader := context.Request.Header.Get("Authorization")
		if authHeader == "" {
			controllers.ResponseError(context, controllers.CodeNeedLogin)
			context.Abort()
			return
		}
		//对authorization 进行分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controllers.ResponseError(context, controllers.CodeInvalidToken)
			context.Abort()
			return
		}
		//parts[1]是对获取到的 tokenString分割后的记过，使用 jwt 进行解析
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controllers.ResponseError(context, controllers.CodeInvalidToken)
			context.Abort()
			return
		}
		//将当前请求的 userID 保存到请求上下文context上
		context.Set(UseridKey, mc.UserID)
		//后续处理可以通过 context.get(UseridKey)来获取当前请求用户信息
		context.Next()
	}
}
