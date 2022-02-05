package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type (

	// Db 连接配置
	DbConfig struct {
		XMLName  xml.Name `xml:"config"`
		DbIp     string   `xml:"dbIp"`
		DbPost   int      `xml:"dbPost"`
		UserName string   `xml:"userName"`
		Pwd      string   `xml:"pwd"`
		Online   DbOnline `xml:"online"`
	}

	//在线Ip
	DbOnline struct {
		Flag        string   `xml:"flag,attr"`
		Description []string `xml:"description"`
	}
)

func main() {
	file, err := os.Open("db_config.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := DbConfig{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	marshal, err := json.Marshal(v)
	fmt.Println("JSON Data  is: ", string(marshal))
	fmt.Println("DbIp is : ", v.DbIp)
	fmt.Println("DbPost is : ", v.DbPost)
	fmt.Println("UserName is : ", v.UserName)
	fmt.Println("Pwd is : ", v.Pwd)
	fmt.Println("Online.Flag is : ", v.Online.Flag)
	fmt.Println("online ip is : ")
	for i, element := range v.Online.Description {
		fmt.Println(i, element)
	}
}
