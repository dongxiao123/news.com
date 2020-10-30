package command

import (
	"github.com/astaxie/beego/orm"
	"news.com/service/news"
	"news.com/utils"
	"sync"
)

func getBaiduNews() error {
	utils.Logs.Info("start getBaiduNews ..")
	news := news.NewBaiduNews()
	titles := news.GetTitleData()
	o := orm.NewOrm()
	o.Using("default")

	var wg sync.WaitGroup
	wg.Add(len(titles))
	for _, t := range titles {
		go func() {
			_, _, err := o.ReadOrCreate(&t, "md_5_code_title")
			if err != nil {
				utils.Logs.Alert("ReadOrCreate error : ", t)
			}
			wg.Done()
		}()
	}
	wg.Wait() // 等待，直到计数为0
	utils.Logs.Info("end getBaiduNews ..")
	return nil
}
