package main

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
)

type User struct {
	Uid      int
	Name     string
	Phone    string
	Email    string
	Password string
}

//定义一个全局变量
var u User

func main() {
	//定义导出的文件名
	filename := "./exportUsers.csv"

	//从数据库中获取数据
	users := queryMultiRow()
	//定义一个二维数组
	column := [][]string{{"手机号", "用户UID", "Email", "用户名"}}
	for _, u := range users {
		str := []string{}
		str = append(str, u.Phone)
		str = append(str, strconv.Itoa(u.Uid))
		str = append(str, u.Email)
		str = append(str, u.Name)
		column = append(column, str)
	}
	//导出
	ExportCsv(filename, column)
}

//导出csv文件
func ExportCsv(filePath string, data [][]string) {
	fp, err := os.Create(filePath) //创建文件句柄
	if err != nil {
		log.Fatalf("创建文件["+filePath+"]句柄失败,%v", err)
		return
	}
	defer fp.Close()
	fp.WriteString("\xEF\xBB\xBF") //写入UTF-8 BOM
	w := csv.NewWriter(fp)         //创建一个新的写入文件流
	w.WriteAll(data)
	w.Flush()
}

// 查询多条数据
func queryMultiRow() []User {
	user := User{
		Uid:      rand.Int(),
		Name:     "测试" + strconv.Itoa(rand.Int()),
		Phone:    "123" + strconv.Itoa(rand.Int()) + "456",
		Email:    "123@qq.com",
		Password: "123456",
	}
	// 循环读取结果集中的数据
	users := []User{}
	var i int
	for i = 1; i < 10; i++ {
		users = append(users, user)
	}
	return users
}
