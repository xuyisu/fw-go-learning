package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL) //数据库类型设计
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/fw_go?charset=utf8")
	// 需要在init中注册定义的model
	orm.RegisterModel(new(BeegoUser))
	//debug 模式
	orm.Debug = true
}

type BeegoUser struct {
	Id    int //默认主健为id
	Name  string
	Phone string
}

func main() {

	o := orm.NewOrm()
	//写入
	//insertBeeGoUser(o)

	//查询
	//selectByUserId(o)
	//selectByUserName(o)

	//更新
	//updateByUserId(o)
	//updateByUserName(o)
	//删除
	//deleteByUserId(o)
	//deleteByUserName(o)
	//事务
	beeGoTransaction(o)
}

func updateByUserId(o orm.Ormer) {
	user := BeegoUser{}
	// 先对主键id赋值, 查询数据的条件就是where id=4
	user.Id = 4
	user.Name = "关羽"
	user.Phone = "15966668888"

	num, err := o.Update(&user)
	if err != nil {
		fmt.Println("更新失败")
	} else {
		fmt.Println("更新数据影响的行数:", num)
	}
}

func updateByUserName(o orm.Ormer) {
	user := BeegoUser{}
	// 先对主键id赋值, 查询数据的条件就是where id=4
	user.Id = 4
	user.Name = "关羽"
	user.Phone = "15966668888"

	num, err := o.Update(&user, "name")
	if err != nil {
		fmt.Println("更新失败")
	} else {
		fmt.Println("更新数据影响的行数:", num)
	}
}

func deleteByUserId(o orm.Ormer) {
	user := BeegoUser{}
	// 先对主键id赋值, 查询数据的条件就是where id=2
	user.Id = 2

	if num, err := o.Delete(&user); err != nil {
		fmt.Println("删除失败")
	} else {
		fmt.Println("删除数据影响的行数:", num)
	}
}

func deleteByUserName(o orm.Ormer) {
	user := BeegoUser{}
	// 先对主键id赋值, 查询数据的条件就是where id=2
	user.Name = "关羽"

	if num, err := o.Delete(&user, "name"); err != nil {
		fmt.Println("删除失败")
	} else {
		fmt.Println("删除数据影响的行数:", num)
	}
}

func selectByUserId(o orm.Ormer) {
	user := BeegoUser{}
	// 先对主键id赋值, 查询数据的条件就是where id=6
	user.Id = 2

	// 通过Read函数查询数据
	// 等价sql: select * from beego_user where id = 6
	err := o.Read(&user)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.Id, user.Name)
	}
}

func selectByUserName(o orm.Ormer) {
	user := BeegoUser{}
	// 先对主键id赋值, 查询数据的条件就是where id=6
	user.Name = "刘备"

	// 通过Read函数查询数据
	// 等价sql: select * from beego_user where id = 6
	err := o.Read(&user, "name")

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else {
		fmt.Println(user.Id, user.Name)
	}
}

func beeGoTransaction(o orm.Ormer) {
	//事务
	o.Begin()
	user1 := BeegoUser{}
	// 赋值
	user1.Id = 5
	user1.Name = "关羽"

	user2 := BeegoUser{}
	// 赋值
	user2.Id = 6
	user2.Name = "貂蝉"

	_, err1 := o.Update(&user1)
	_, err2 := o.Insert(&user2)
	// 检测事务执行状态
	if err1 != nil || err2 != nil {
		// 如果执行失败，回滚事务
		o.Rollback()
	} else {
		// 任务执行成功，提交事务
		o.Commit()
	}
	//事务
}

func insertBeeGoUser(o orm.Ormer) {
	user := new(BeegoUser)
	user.Name = "刘备"
	user.Phone = "18888888888"
	insert, _ := o.Insert(user)
	//输出最后新增的id
	fmt.Println(insert)
}
