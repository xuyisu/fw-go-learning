package main

import (
	"fmt"
	"strconv"
)

func main() {
	//int转换为字符串
	int2string := strconv.Itoa(100)
	fmt.Println("int转换为字符串:", int2string)
	//string转换为int,失败会返回err 信息
	string2int, _ := strconv.Atoi(int2string)
	fmt.Println("string转换为int:", string2int)
	//字符串解析为bool型,失败会返回err 信息
	// 它接受真值：1, t, T, TRUE, true, True
	// 它接受假值：0, f, F, FALSE, false, False
	parseBool, _ := strconv.ParseBool("true")
	fmt.Println("字符串解析为bool型:", parseBool)
	//字符串解析为float型,失败会返回err 信息 bitSize参数表示转换为什么位的int/uint，有效值为0、8、16、32、64。当bitSize=0的时候，表示转换为int或uint类型。例如bitSize=8表示转换后的值的类型为int8或uint8
	parseFloat, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Println("字符串解析为float型:", parseFloat)
	//字符串转整形 base 表述参数表示以什么进制的方式去解析给定的字符串 bitSize 跟上面含义一样
	parseInt, _ := strconv.ParseInt("-159666666666", 10, 64)
	fmt.Println("字符串解析为int型:", parseInt)
	//字符串转无符号整形,负数返回0
	parseUint, _ := strconv.ParseUint("159666666666", 10, 64)
	fmt.Println("字符串解析为uint型:", parseUint)
	//bool 转字符串
	formatBool := strconv.FormatBool(true)
	fmt.Println("bool 转字符串:", formatBool)
	//float 转字符串
	//fmt表示格式：'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）、'g'（指数很大时用'e'格式，否则'f'格式）、'G'（指数很大时用'E'格式，否则'f'格式）
	//prec控制精度（排除指数部分）：对'f'、'e'、'E'，它表示小数点后的数字个数；对'g'、'G'，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f
	formatFloat := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Println("float 转字符串:", formatFloat)
	//int 转字符串 base 代表进制
	formatInt := strconv.FormatInt(-159666666666, 10)
	fmt.Println("int 转字符串:", formatInt)
	//uint 转字符串 base 代表进制
	formatUint := strconv.FormatUint(159666666666, 10)
	fmt.Println("uint 转字符串:", formatUint)
	// CanBackquote 判断字符串 s 是否可以表示为一个单行的“反引号”字符串
	// 字符串中不能含有控制字符（除了 \t）和“反引号”字符，否则返回 false
	canBackquote := strconv.CanBackquote("C:\\go\t")
	fmt.Println("canBackquote:", canBackquote)
	// AppendInt 将 int 型整数 i 转换为字符串形式，并追加到 dst 的尾部
	// i：要转换的字符串
	// base：进位制
	// 返回追加后的 []byte
	b := make([]byte, 0)
	b = strconv.AppendInt(b, 100, 10)
	b = strconv.AppendInt(b, 500, 10)
	fmt.Printf("%s", b)
}
