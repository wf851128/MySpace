package mysql

import (
	"MySpace/settings"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		settings.Conf.MysqlConfig.User,
		settings.Conf.MysqlConfig.Password,
		settings.Conf.MysqlConfig.Host,
		settings.Conf.MysqlConfig.Port,
		settings.Conf.MysqlConfig.DBName,
	)
	if db, err = sqlx.Connect("mysql", dsn); err != nil {
		zap.L().Error("connect DB failed, err %s/n", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(settings.Conf.MysqlConfig.MaxOpenConn)
	db.SetMaxIdleConns(settings.Conf.MysqlConfig.MaxIdleConn)
	return
}

func Close() {
	_ = db.Close()
}
