package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"server/conf"
	"time"
)
var MysqlCenterDb *sql.DB
var MysqlServerDb *sql.DB
var Err error
func init() {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", conf.User, conf.Pwd, conf.Host, conf.Port, conf.DataBase, conf.Charset)
	// 打开连接失败
	//dbDSN :=conf.User+":"+conf.Pwd+"@"+conf.Host+"/"+conf.DataBase
	MysqlCenterDb, Err = sql.Open("mysql", dbDSN)
	if Err != nil {
		panic("数据源配置错误: " + Err.Error())
	}
	MysqlCenterDb.SetMaxOpenConns(1000)
	// 闲置连接数
	MysqlCenterDb.SetMaxIdleConns(200)
	// 最大连接周期
	MysqlCenterDb.SetConnMaxLifetime(100 * time.Second)
	if Err = MysqlCenterDb.Ping(); nil != Err {
		panic("数据库链接失败: " + Err.Error())
	}
	dbServerDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", conf.User, conf.Pwd, conf.Host, conf.Port, conf.DataServerBase, conf.Charset)
	// 打开连接失败
	MysqlServerDb, Err = sql.Open("mysql", dbServerDSN)
	if Err != nil {
		panic("数据源配置错误: " + Err.Error())
	}
	MysqlServerDb.SetMaxOpenConns(1000)
	// 闲置连接数
	MysqlServerDb.SetMaxIdleConns(200)
	// 最大连接周期
	MysqlServerDb.SetConnMaxLifetime(100 * time.Second)
	if Err = MysqlServerDb.Ping(); nil != Err {
		panic("数据库链接失败: " + Err.Error())
	}

}