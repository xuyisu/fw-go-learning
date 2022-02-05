package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL) //数据库类型设计
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/fw_go?charset=utf8")
	// 需要在init中注册定义的model
	orm.RegisterModel(new(BeegoUser2))
	//debug 模式
	orm.Debug = true
}

type BeegoUser2 struct {
	Id    int //默认主健为id
	Name  string
	Phone string
}

func main() {
	o := orm.NewOrm()
	//原生新增
	insertBeeGo(o)

	//原生根据名字更新
	updateByNameBeeGo(o)

	//原生根据主键查询
	selectByIdBeeGo(o)

	//原生根据主键删除
	deleteBeeGo(o)
}

func selectByIdBeeGo(o orm.Ormer) {
	var user []BeegoUser2
	var _ orm.RawSeter
	rows, _ := o.Raw("select * from  beego_user WHERE id = ?", 4).QueryRows(&user)
	fmt.Println("用户nums:", rows)
	fmt.Println(user)
}

func updateByNameBeeGo(o orm.Ormer) {
	var _ orm.RawSeter
	res, err := o.Raw("UPDATE beego_user SET name = ? WHERE id = ?", "关羽", 4).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
}

func insertBeeGo(o orm.Ormer) {
	var _ orm.RawSeter
	res, err := o.Raw("insert into beego_user(name,phone) values (?,?)", "刘备", "15866669999").Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
}

func deleteBeeGo(o orm.Ormer) {
	var _ orm.RawSeter
	res, err := o.Raw("delete from beego_user WHERE id = ?", 4).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
}
