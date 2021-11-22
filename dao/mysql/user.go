package mysql

import (
	"MySpace/models"
	"MySpace/settings"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

/**
把每一步数据库操作封装成函数待logic层根据业务需求来进行调用
*/
var (
	ErrorUserExist       = errors.New("用户已存在") //用户已存在
	ErrorUserNoExist     = errors.New("用户不存在") //用户不存在
	ErrorInvalidPassword = errors.New("密码错误")  //密码错误
)

// CheckUserExits 检查用户是否存在
func CheckUserExits(userName string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	err = db.Get(&count, sqlStr, userName)
	if err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 向数据库插入用户记录
func InsertUser(user *models.User) (err error) {
	user.Password = md5Password(user.Password)
	//对用户密码进行加密
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.UserName, user.Password)
	if err != nil {
		return err
	}
	return
}
func md5Password(oPassword string) string {
	h := md5.New()
	h.Write([]byte(settings.Conf.AppConfig.PasswordSecretKey))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(p *models.User) (err error) {
	//用户输入的密码
	oPassword := p.Password
	sqlStr := `select username,password from user where username = ?`
	err = db.Get(p, sqlStr, p.UserName)
	if err == sql.ErrNoRows {
		return ErrorUserNoExist
	}
	if err != nil {
		return err
	}
	//将用户输入密码转换成 md5
	password := md5Password(oPassword)
	//	判断密码是否正确
	if password != p.Password {
		return ErrorInvalidPassword
	}
	return
}
