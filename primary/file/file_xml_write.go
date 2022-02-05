package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type (
	DbConfig2 struct {
		XMLName  xml.Name  `xml:"config"`
		DbIp     string    `xml:"dbIp"`
		DbPost   int       `xml:"dbPost"`
		UserName string    `xml:"userName"`
		Pwd      string    `xml:"pwd"`
		Online   DbOnline2 `xml:"online"`
	}

	DbOnline2 struct {
		Flag        string   `xml:"flag,attr"`
		Description []string `xml:"description"`
	}
)

func main() {
	v := &DbConfig2{DbIp: "127.0.0.1", DbPost: 3306, UserName: "root", Pwd: "123456"}
	var description []string
	description = append(description, "127.0.0.1")
	description = append(description, "127.0.0.2")
	description = append(description, "127.0.0.3")
	description = append(description, "127.0.0.4")
	v.Online = DbOnline2{"true", description}
	output, err := xml.MarshalIndent(v, "", "	")
	if err != nil {
		fmt.Printf("error %v", err)
		return
	}
	file, _ := os.Create("db_config2.xml")
	defer file.Close()
	file.Write([]byte(xml.Header))
	file.Write(output)
}
