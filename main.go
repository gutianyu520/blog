package main

import (
	"github.com/astaxie/beego/client/orm"
	beego "github.com/astaxie/beego/server/web"
	"hello/models"
	_ "hello/routers"
)

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Run()
}

func init() {
	models.RegisterDB()
}
