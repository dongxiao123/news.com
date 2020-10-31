package command

import (
	"github.com/astaxie/beego/orm"
	"news.com/service/news"
	"news.com/utils"
)

func getBaiduNews() error {
	utils.Logs.Info("start getBaiduNews ..")
	news := news.NewBaiduNews()
	titles := news.GetTitleData()
	o := orm.NewOrm()
	o.Using("default")

	for _, t := range titles {
		_, _, err := o.ReadOrCreate(&t, "Md5CodeTitle")
		if err != nil {
			utils.Logs.Alert("ReadOrCreate error : ", err.Error())
		}
	}
	utils.Logs.Info("end getBaiduNews ..")
	return nil
}
