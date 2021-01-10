package controllers

import (
	beego "github.com/astaxie/beego/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["flag"] = true
	c.TplName = "index.html"
	//s, _ := beego.AppConfig.String("appname")
	////c.Ctx.WriteString("appname:"+s)
	//c.Ctx.WriteString(beego.AppPath)
	//logger := logs.GetBeeLogger()
	//logger.SetLogger("console")
	type u struct {
		Name string
		Age  int
		Sex  string
	}
	user := &u{"lisi", 14, "man"}

	c.Data["user"] = user

	nums := []int{1, 2, 3, 4, 5}
	c.Data["nums"] = nums

	c.Data["tplVal"] = "val"

	c.Data["Html"] = "<div>hello go</div>"

	c.Data["Html2"] = "<div>hello go</div>"
}
