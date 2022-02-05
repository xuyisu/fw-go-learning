package main

import (
	"fmt"
	"time"
)

func main() {
	//当前时间戳
	fmt.Println("当前时间戳:", time.Now().Unix())
	//str格式化时间
	fmt.Println("str格式化时间:", time.Now().Format("2006-01-02 15:04:05")) // 必须是这个时间点, 据说是go诞生之日
	//时间戳转str格式化时间
	strTime := time.Unix(1640937278, 0).Format("2006-01-02 15:04:05")
	fmt.Println("时间戳转str格式化时间:", strTime)
	//str格式化时间转时间戳
	localTime, err := time.Parse("2006-01-02 15:04:05", "2014-01-08 09:04:41")
	if err == nil {
		unixTime := localTime.Unix()
		fmt.Println("str格式化时间转时间戳:", unixTime)
	}

}
