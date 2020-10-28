package main

import (
	"github.com/astaxie/beego"
	_ "news.com/controllers/command"
	_ "news.com/routers"
)

func main() {
	beego.Run()
}
