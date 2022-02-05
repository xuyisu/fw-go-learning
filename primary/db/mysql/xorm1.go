package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"xorm.io/builder"
)

// User 用户信息
type UserXorm struct {
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

	//设置缓存  采用了LRU算法的一个缓存，缓存方式是存放到内存中，缓存struct的记录数为1000条，缓存针对的范围是所有具有主键的表，没有主键的表中的数据将不会被缓存
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	// 启用一个全局的内存缓存
	db.SetDefaultCacher(cacher)
	// 针对某个对象
	//db.MapCacher(&UserXorm{}, cacher)
	//由于存在ORM和RAW两种方式操作数据库，故在使用Exec方法执行了方法之后（RAW方式），可能会导致缓存与数据库不一致的地方。因此如果启用缓存，尽量避免使用类Exec的方法。如果必须使用，则需要在使用了Exec之后调用ClearCache手动做缓存清除的工作
	//db.ClearCache(new(UserXorm))

	// 显示sql
	db.ShowSQL(true)
	//插入
	insert(db)
	insertWithEvent(db)

	//主键查询
	selectById(db)
	//列表查询
	//List(err, db)
	//列表查询Builder
	//ListWithBuilder(err, db)
	//更新
	//updateById(db)

	// 删除
	//deleteById2(db)
	//事务
	//SaveWithTransaction(db, err)

}

func SaveWithTransaction(db *xorm.Engine, err error) {
	session := db.NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		log.Fatal(err)
	}
	user := &UserXorm{4, "aaaaaa", "18888888888"}
	before := func(bean interface{}) {
		fmt.Println("写入前的事件:", bean)
	}
	insert, err := db.Before(before).Insert(user)
	if err != nil {
		session.Rollback()
		return
	} else {
		session.Commit()
		fmt.Println("写入数据:", insert)
		return
	}
}

func insertWithEvent(db *xorm.Engine) {
	user := &UserXorm{3, "aaaaaa", "18888888888"}
	before := func(bean interface{}) {
		fmt.Println("写入前的事件:", bean)
	}
	db.Before(before).Insert(user)
}

func ListWithBuilder(err error, db *xorm.Engine) {
	list := make([]UserXorm, 0)
	err = db.Where(builder.NotIn("name", "a", "b").And(builder.In("id", 1, 2, 3))).Find(&list)
	if err == nil {
		fmt.Printf("%v\n", list)
	} else {
		log.Fatal("ormFindRows error", err)
	}
}

func deleteById2(db *xorm.Engine) {
	user := new(UserXorm)
	user.Id = 1
	delete, _ := db.Delete(user)
	fmt.Println("删除行数:", delete)
}

func updateById(db *xorm.Engine) {
	user := new(UserXorm)
	user.Id = 1
	user.Name = "myname2"
	affected, _ := db.Where("id=?", 1).Update(user)
	fmt.Println("影响行数:", affected)
}

//批量查询
func List(err error, db *xorm.Engine) {
	list := make([]UserXorm, 0)
	err = db.Cols("id", "name", "phone").Where("id>?", 0).Limit(2).Asc("id").Find(&list)

	if err == nil {
		fmt.Printf("%v\n", list)
	} else {
		log.Fatal("ormFindRows error", err)
	}
}

func selectById(db *xorm.Engine) {
	user := &UserXorm{Id: 1}
	flag, _ := db.Get(user)

	if flag {
		fmt.Printf("%v\n", *user)
	}
}

func insert(db *xorm.Engine) {
	u1 := &UserXorm{1, "刘备", "18888888888"}
	// 新增
	db.Insert(u1)
}
