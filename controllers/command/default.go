package command

import (
	"github.com/astaxie/beego/toolbox"
	"news.com/utils"
)

func Init() {
	getBaiduNews := toolbox.NewTask("getBaiduNews", "0 */2 * * * *", getBaiduNews)
	err := getBaiduNews.Run()
	if err != nil {
		utils.Logs.Warning(err.Error())
	}
	toolbox.AddTask("getBaiduNews", getBaiduNews)
	toolbox.StartTask()
}
