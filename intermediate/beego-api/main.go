package main

import (
	_ "beego-api/routers"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	sqlconn, err := beego.AppConfig.String("sqlconn")
	if err != nil {
		fmt.Println("异常", err)
	}
	orm.RegisterDataBase("default", "mysql", sqlconn)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
