package utils

import (
	"github.com/astaxie/beego"
	beelog "github.com/astaxie/beego/logs"
	"time"
)

var (
	Logs = beego.BeeLogger
)

type BLogs struct {
	beelog.BeeLogger
}

func init() {
	filePath := beego.AppConfig.String("LOG_PATH") + "/" + time.Now().Format("2006-01-02_15")
	filename := filePath + ".log"
	beelog.SetLogger(beelog.AdapterFile, `{"filename":"`+filename+`"}`)
}
func (bl *BLogs) Write(p []byte) (n int, err error) {
	Logs.Warning(string(p))

	return len(p), nil
}
