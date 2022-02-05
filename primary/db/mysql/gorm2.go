package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 为了演示，实际是可以直接复用 UserGorms
type UserGorms2 struct {
	Id    int
	Name  string
	Phone string
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/fw_go?"+
		"charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	//输出详细日志
	db.LogMode(true)
	//新增
	insert2(db)
	//查询
	query2(db)
	//更新
	update2(db)
	//删除
	delete2(db)

}

func delete2(db *gorm.DB) *gorm.DB {
	return db.Exec("delete from user_gorms  where id=?", 5)
}

func update2(db *gorm.DB) *gorm.DB {
	return db.Exec("UPDATE user_gorms SET name = ?, phone = ? where id=?", "貂蝉", "13666669999", 5)
}

func query2(db *gorm.DB) {
	var result UserGorms2
	db.Raw("SELECT id, name, phone FROM user_gorms WHERE id = ?", 3).Scan(&result)
	fmt.Printf("查询查找%v\n", result)

	// Use GORM API build SQL
	row := db.Table("user_gorms").Where("id = ?", 3).Row()
	row.Scan(&result)
	fmt.Printf("查询查找%v\n", result)
}

func insert2(db *gorm.DB) {
	//新增
	db.Exec("insert into user_gorms(name,phone) values (?,?)", "刘备", 13966666666)
}
