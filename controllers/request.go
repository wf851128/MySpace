package controllers

import (
	"MySpace/middlewares"
	"errors"
	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户未登录")

//getCurrentUserID 获取当前用户登录 ID
func getCurrentUserID(context *gin.Context) (userID int64, err error) {
	uid, ok := context.Get(middlewares.UseridKey)
	if !ok {
		err = ErrorUserNotLogin
		return 0, err
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return 0, err
	}
	return userID, nil
}
