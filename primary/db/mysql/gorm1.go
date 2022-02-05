package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserGorms 用户信息
type UserGorms struct {
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

	// 自动迁移 最好手动创建表
	db.AutoMigrate(&UserGorms{})

	//输出详细日志
	db.LogMode(true)

	// 新增
	save(db)

	var userGorm UserGorms
	// 查询
	query(db, userGorm)
	//更新
	update(db, userGorm)
	//事务操作
	saveWithTransaction(db)
	//删除
	deleteById(db, userGorm)
}

func deleteById(db *gorm.DB, userGorm UserGorms) {
	//删除
	db.Delete(&userGorm, 1)
}

func saveWithTransaction(db *gorm.DB) {
	//开启事务
	tx := db.Begin()
	user := UserGorms{
		Phone: "16666666666",
		Name:  "张飞",
	}
	if err := tx.Create(&user).Error; err != nil {
		//事务回滚
		tx.Rollback()
		fmt.Println(err)
	}
	db.Model(&user).Updates(UserGorms{Name: "关羽", Phone: "19999999999"})
	//事务提交
	tx.Commit()
}

func update(db *gorm.DB, userGorm UserGorms) {
	//更新
	db.Model(&userGorm).Where("id=?", 1).Update("name", "关羽")
	//更新多个字段
	db.Model(&userGorm).Where("id=?", 1).Updates(UserGorms{Name: "关羽", Phone: "19999999999"})
}

func query(db *gorm.DB, userGorm UserGorms) {
	//根据主键查找
	db.First(&userGorm, 1)
	fmt.Printf("主键查找%v\n", userGorm)
	db.First(&userGorm, "name = ?", "刘备")
	fmt.Printf("普通条件查找%v\n", userGorm)
}

func save(db *gorm.DB) *gorm.DB {
	return db.Create(&UserGorms{1, "刘备", "18888888888"})
}
