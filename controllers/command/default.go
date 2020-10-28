package command

import (
	"github.com/astaxie/beego/toolbox"
	"news.com/utils"
	"time"
)

func init() {
	tk := toolbox.NewTask("tk", "* * * * * *", func() error {
		utils.Logs.Warning("cron start:", time.Now().Format("2006-01-02 15:04:05"))
		return nil
	})
	err := tk.Run()
	if err != nil {
		utils.Logs.Warning(err.Error())
	}
	toolbox.AddTask("tk", tk)
	toolbox.StartTask()
	defer toolbox.StopTask()
}
