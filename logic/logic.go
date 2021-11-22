package logic

import (
	"MySpace/dao/mysql"
	"MySpace/models"
	"MySpace/pkg/jwt"
	"MySpace/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	//	判断用户是否存在
	err = mysql.CheckUserExits(p.UserName)
	if err != nil {
		return err
	}
	//	生成 uid
	userID := snowflake.GenID()
	//构造一个 user 实例
	user := &models.User{
		UserID:   userID,
		UserName: p.UserName,
		Password: p.Password,
	}
	//	保存至数据库
	return mysql.InsertUser(user)
}

func Login(p *models.User) (token string, err error) {
	user := &models.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	//	判断用户是否存在
	if err = mysql.Login(user); err != nil {
		return "", err
	}
	//用户存在的情况下，生成 token
	if err != nil {
		return "", err
	}
	return jwt.GenToken(p.UserID, p.UserName)
}
