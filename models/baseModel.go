package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"news.com/utils"
)

type BaseModel struct {
}
type BaseModelInterface interface {
	TableName() string
}

func init() {
	dbHost, dbUser, dbPass, dbName, debug := initDBConfig()
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbUser+":"+dbPass+"@tcp("+dbHost+")/"+dbName+"?charset=utf8")
	orm.DebugLog = orm.NewLog(utils.Logs)
	orm.Debug = debug
}

//初始化数据库相关配置
func initDBConfig() (string, string, string, string, bool) {
	runmode := beego.AppConfig.DefaultString("runmode", "dev")
	dbHost := beego.AppConfig.String(runmode + "::DB_HOST")
	if dbHost == "" {
		err := "config dbHost is empty"
		utils.Logs.Error(err)
		panic(err)
	}
	dbUser := beego.AppConfig.String(runmode + "::DB_USER")
	if dbUser == "" {
		err := "config DB_USER is empty"
		utils.Logs.Error(err)
		panic(err)
	}
	dbPass := beego.AppConfig.String(runmode + "::DB_PASS")
	if dbPass == "" {
		err := "config DB_PASS is empty"
		utils.Logs.Error(err)
		panic(err)
	}
	dbName := beego.AppConfig.String(runmode + "::DB_DATABASE")
	if dbName == "" {
		err := "config DB_DATABASE is empty"
		utils.Logs.Error(err)
		panic(err)
	}
	debug := beego.AppConfig.DefaultBool(runmode+"::Debug", false)

	return dbHost, dbUser, dbPass, dbName, debug
}
