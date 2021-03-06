package command

import (
	"github.com/astaxie/beego/orm"
	"news.com/service/news"
	"news.com/utils"
	"time"
)

func spiderNews() error {
	utils.Logs.Info("start spiderNews ..")

	titles, err := news.Spider()
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	o.Using("default")
	updatedAt := time.Now().Format("2006-01-02 15:04:05")

	for _, t := range titles {
		_, err := o.InsertOrUpdate(&t, "updated_at='"+updatedAt+"'")
		if err != nil {
			utils.Logs.Alert("InsertOrUpdate", t)
		}
	}

	utils.Logs.Info("end spiderNews ..")
	return nil
}
