package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//响应对象进行抽象

type ResponseData struct {
	Code ResCode     `json:"code"` //错误代码
	Msg  interface{} `json:"msg"`  //提示信息
	Data interface{} `json:"data"` //数据
}

// ResponseError 在 code 列表中的常规错误
func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

// ResponseErrorWithMsg 自定义错误及错误信息
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// ResponseSuccess 请求成功应答
func ResponseSuccess(c *gin.Context, data ...interface{}) {

	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}
