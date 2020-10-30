package command

import (
	"github.com/astaxie/beego/orm"
	"news.com/service/news"
	"news.com/utils"
	"sync"
	"time"
)

func spiderNews() error {
	utils.Logs.Info("start spiderNews ..")

	titles, err := news.Spider()
	if err!=nil{
		return err
	}
	o := orm.NewOrm()
	o.Using("default")
	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	var wg sync.WaitGroup

	wg.Add(len(titles))
	for _, t := range titles {
		go func() {
			_, err := o.InsertOrUpdate(&t, "updated_at='"+updatedAt+"'")
			if err != nil {
				utils.Logs.Alert("InsertOrUpdate", t)
			}
			wg.Done()
		}()
	}
	wg.Wait() // 等待，直到计数为0
	utils.Logs.Info("end spiderNews ..")
	return nil
}
