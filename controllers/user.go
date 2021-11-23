package controllers

import (
	"MySpace/dao/mysql"
	"MySpace/logic"
	"MySpace/models"
	"MySpace/pkg/jwt"
	"MySpace/pkg/validatorTrans"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 用户注册
func SignUpHandler(context *gin.Context) {
	var p = new(models.ParamSignUp)
	//	参数校验
	if err := context.ShouldBindJSON(&p); err != nil {
		//log 中记录错误信息
		zap.L().Error("controllers.SignUpHandler.ShouldBindQuery error,", zap.Error(err))
		//检查err是不是 validatorTrans 的类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(context, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(context,
				CodeInvalidParam,
				validatorTrans.RemoveTopStruct(errs.Translate(validatorTrans.Trans)),
			)
		}
		return
	}
	//	请求参数有误,进行异常处理
	//	业务处理
	err := logic.SignUp(p)
	if err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(context, CodeUserExist)
			return
		}
		ResponseError(context, CodeServerBusy)
		return
	}
	//	返回响应
	ResponseSuccess(context, nil)
}

// LoginHandler 用户登录
func LoginHandler(context *gin.Context) {
	var p = new(models.User)
	//	获取请求参数及参数校验
	if err := context.ShouldBindJSON(&p); err != nil {
		//log 中记录错误信息
		zap.L().Error("controllers.LoginHandler.ShouldBindJSON error,", zap.Error(err))
		//检查err是不是 validatorTrans 的类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(context, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(context,
				CodeInvalidParam,
				validatorTrans.RemoveTopStruct(errs.Translate(validatorTrans.Trans)),
			)
		}
	}
	//	业务逻辑处理
	aToken, rToken, err := logic.Login(p)
	if err != nil {
		//登录失败
		zap.L().Error("controllers.LoginHandler.Login error,",
			zap.String("username", p.UserName),
			zap.Error(err),
		)
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(context, CodeUserExist)
			return
		}
		ResponseError(context, CodeInvalidUserNameOrPassword)
	} else {
		//登录成功
		ResponseSuccess(context, aToken, rToken)
	}
}

func RefreshTokenHandler(context *gin.Context) {
	var p = new(models.RefreshToken)

	//	获取参数
	if err := context.ShouldBindJSON(&p); err != nil {
		//log 中记录错误信息
		zap.L().Error("controllers.RefreshTokenHandler.ShouldBindJSON error,", zap.Error(err))
		//检查err是不是 validatorTrans 的类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(context, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(context,
				CodeInvalidParam,
				validatorTrans.RemoveTopStruct(errs.Translate(validatorTrans.Trans)),
			)
		}
	}
	//	验证参数 通过 rToken获取新的 aToken
	aToken, err := jwt.RefreshToken(p.AccessToken, p.RefreshToken)
	if err != nil {
		//刷新失败
		zap.L().Error("controllers.LoginHandler.RefreshToken error,",
			zap.String("AccessToken", p.AccessToken),
			zap.Error(err),
		)
		ResponseError(context, CodeInvalidUserNameOrPassword)
	} else {
		//刷新成功
		ResponseSuccess(context, aToken)
	}
}
