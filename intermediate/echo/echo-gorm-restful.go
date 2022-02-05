package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
	"time"
)

var (
	db            *gorm.DB
	sqlConnection = "root:123456@tcp(127.0.0.1:3306)/fw_go?charset=utf8&parseTime=true"
)

//初始化
func init() {
	//打开数据库连接
	var err error
	db, err = gorm.Open("mysql", sqlConnection)
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&SysUser{})
	db.LogMode(true)
}

func main() {
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.POST("/api/v2/user", createSysUser)       //POST方法，创建新用户
	app.GET("/api/v2/user", fetchAllSysUser)      //GET方法，获取所有用户
	app.GET("/api/v2/user/:id", fetchSysUser)     //GET方法，获取某一个用户，形如：/api/v2/user/1
	app.PUT("/api/v2/user/:id", updateSysUser)    //PUT方法，更新用户，形如：/api/v2/user/1
	app.DELETE("/api/v2/user/:id", deleteSysUser) //DELETE方法，删除用户，形如：/api/v2/user/1
	app.Start(":8081")
}

type (
	//db 操作
	SysUser struct {
		Id         uint      `gorm:"column:id"  comment:"主键" sql:"bigint(20) unsigned,PRI" `
		CreateTime time.Time `gorm:"column:create_time"  comment:"创建时间" sql:"datetime" `
		UpdateTime time.Time `gorm:"column:update_time"  comment:"更新时间" sql:"datetime" `
		CreateUser int       `gorm:"column:create_user"  comment:"创建人" sql:"bigint(20)" `
		UpdateUser int       `gorm:"column:update_user"  comment:"更新人" sql:"bigint(20)" `
		DeleteFlag int       `gorm:"column:delete_flag"  comment:"删除标志" sql:"tinyint(4)" `
		Status     int       `gorm:"column:code"  comment:"启用标志" sql:"tinyint(4)" `
		UserName   string    `gorm:"column:user_name"  comment:"用户名" sql:"varchar(50)" `
		Email      string    `gorm:"column:email"  comment:"邮箱" sql:"varchar(50)" `
		Phone      string    `gorm:"column:phone"  comment:"手机号" sql:"varchar(20),UNI" `
		Password   string    `gorm:"column:password"  comment:"密码" sql:"varchar(100)" `
	}

	//响应返回的结构体
	SysUserRes struct {
		Id         uint      `json:"id"`
		Phone      string    `json:"phone"`
		UserName   string    `json:"userName"`
		Email      string    `json:"email"`
		CreateTime time.Time `json:"createTime"`
	}
	//请求体
	SysUserReq struct {
		Phone    string `json:"phone"`
		UserName string `json:"userName"`
		Email    string `json:"email"`
	}
	//返回消息
	Result struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}
)

//md5加密
func md5Password(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 创建新用户
func createSysUser(c echo.Context) error {

	var req SysUserReq
	err := c.Bind(&req)
	if err != nil {
		fmt.Println(err)
	}
	user := SysUser{
		Phone:      req.Phone,
		UserName:   req.UserName,
		Password:   md5Password("123456"), //用户密码
		Email:      req.Email,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		DeleteFlag: 0,
		Status:     1,
	}
	db.Save(&user) //保存到数据库
	return c.JSON(http.StatusCreated, Result{Code: http.StatusCreated, Msg: "用户创建成功!", Data: user.Id})
}

// 获取所有用户
func fetchAllSysUser(c echo.Context) error {
	var user []SysUser        //定义一个数组去数据库总获取数据
	var _userRes []SysUserRes //定义一个响应数组用户返回数据到客户端

	db.Find(&user)

	if len(user) <= 0 {
		return c.JSON(http.StatusNotFound, Result{Code: http.StatusNotFound, Msg: "用户不存在!"})
	}

	//循环遍历，追加到响应数组
	for _, item := range user {
		_userRes = append(_userRes,
			SysUserRes{
				Id:         item.Id,
				Phone:      item.Phone,
				UserName:   item.UserName,
				Email:      item.Email,
				CreateTime: item.CreateTime,
			})
	}
	return c.JSON(http.StatusOK, Result{Code: http.StatusOK, Data: _userRes})
}

// 获取单个用户
func fetchSysUser(c echo.Context) error {
	var user SysUser                       //定义SysUser结构体
	data, _ := strconv.Atoi(c.Param("id")) //获取参数id

	db.First(&user, data)

	if user.Id == 0 { //如果用户不存在，则返回
		return c.JSON(http.StatusNotFound, Result{Code: http.StatusNotFound, Msg: "用户不存在!"})
	}
	//返回响应结构体
	res := SysUserRes{
		Id:         user.Id,
		Phone:      user.Phone,
		UserName:   user.UserName,
		Email:      user.Email,
		CreateTime: user.CreateTime,
	}
	return c.JSON(http.StatusOK, Result{Code: http.StatusOK, Data: res})
}

//更新用户
func updateSysUser(c echo.Context) error {
	var user SysUser                         //定义SysUser结构体
	userID, _ := strconv.Atoi(c.Param("id")) //获取参数id
	db.First(&user, userID)                  //查找数据库

	if user.Id == 0 { //如果数据库不存在，则返回
		return c.JSON(http.StatusNotFound, Result{Code: http.StatusNotFound, Msg: "用户不存在!"})
	}

	var req SysUserReq
	err := c.Bind(&req)
	if err != nil {
		fmt.Println(err)
	}
	if req.Email == "" || req.UserName == "" || req.Phone == "" {
		return c.JSON(http.StatusBadRequest, Result{Code: http.StatusBadRequest, Msg: "参数必填!"})
	}
	//更新对应的字段值
	db.Model(&user).Update(SysUser{Email: req.Email, UserName: req.UserName, Phone: req.Phone})
	return c.JSON(http.StatusOK, Result{Code: http.StatusOK, Msg: "更新用户成功!"})
}

// 删除用户
func deleteSysUser(c echo.Context) error {
	var user SysUser                         //定义SysUser结构体
	userID, _ := strconv.Atoi(c.Param("id")) //获取参数id

	db.First(&user, userID) //查找数据库

	if user.Id == 0 { //如果数据库不存在，则返回
		return c.JSON(http.StatusNotFound, Result{Code: http.StatusNotFound, Msg: "用户不存在!"})
	}

	//删除用户
	db.Delete(&user)
	return c.JSON(http.StatusOK, Result{Code: http.StatusOK, Msg: "用户删除成功!"})
}
