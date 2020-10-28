package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"news.com/controllers/command"
	_ "news.com/routers"
)

func init() {
	//初始化自动化脚本
	command.Init()
}
func main() {
	beego.Run()
	defer toolbox.StopTask()
}
