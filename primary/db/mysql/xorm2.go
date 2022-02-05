package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// User 用户信息
type UserXorm2 struct {
	Id    int
	Name  string
	Phone string
}

func main() {
	db, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/fw_go?"+
		"charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// 显示sql
	db.ShowSQL(true)

	//原生新增
	insertXorm(db)

	//原生sql 查询
	selectByIdXorm(db)

	//原生更新
	updateXorm(db)

	//原生删除
	deleteXorm(db)

}

func insertXorm(db *xorm.Engine) {
	result, err := db.Exec("insert into user_xorm(name,phone) values (?,?)", "刘备", 13966666666)
	if err != nil {
		fmt.Println(err)
	}
	affected, _ := result.RowsAffected()
	fmt.Println("影响行数:", affected)

}

func updateXorm(db *xorm.Engine) {
	result, err := db.Exec("UPDATE user_xorm SET name = ?, phone = ? where id=?", "貂蝉", "13666669999", 5)
	if err != nil {
		fmt.Println(err)
	}
	affected, _ := result.RowsAffected()
	fmt.Println("影响行数:", affected)

}

func deleteXorm(db *xorm.Engine) {
	result, err := db.Exec("delete from  user_xorm  where id=?", 5)
	if err != nil {
		fmt.Println(err)
	}
	affected, _ := result.RowsAffected()
	fmt.Println("影响行数:", affected)

}

func selectByIdXorm(db *xorm.Engine) {
	//原生sql查询
	result, err := db.QueryString("select id,name,phone from `user_xorm` where id=?", 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	//需要定义成切片
	var user []UserXorm2
	err = db.SQL("select id,name,phone from `user_xorm` where id=?", 1).Find(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}
