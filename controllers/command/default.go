package command

import (
	"github.com/astaxie/beego/toolbox"
	"news.com/utils"
)

func Init() {
	getBaiduNews := toolbox.NewTask("getBaiduNews", "0 */5 * * * *", getBaiduNews)
	spiderNews := toolbox.NewTask("getBaiduNews", "0 */5 * * * *", spiderNews)
	err := getBaiduNews.Run()
	if err != nil {
		utils.Logs.Warning(err.Error())
	}
	err = spiderNews.Run()
	if err != nil {
		utils.Logs.Warning(err.Error())
	}
	toolbox.AddTask("getBaiduNews", getBaiduNews)
	toolbox.AddTask("spiderNews", spiderNews)
	toolbox.StartTask()
}
