package command

import (
	"github.com/astaxie/beego/orm"
	"news.com/service/news"
	"news.com/utils"
	"time"
)

func getBaiduNews() error {
	news := news.NewBaiduNews()
	titles := news.GetTitleData()
	o := orm.NewOrm()
	o.Using("default")
	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	for _, t := range titles {
		go func() {
			_, err := o.InsertOrUpdate(&t, "updated_at='"+updatedAt+"'")
			if err != nil {
				utils.Logs.Alert("InsertOrUpdate", t)
			}
		}()
	}
	return nil
}
