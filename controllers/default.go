package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"news.com/service/news"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	news := news.NewBaiduNews()
	titles := news.GetTitleData()
	o := orm.NewOrm()
	o.Using("default")
	successNums, err := o.InsertMulti(100, titles)
	fmt.Println(titles)
	fmt.Println(successNums)
	fmt.Println(err)
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
