package main

import "fmt"

func main() {
	username := ""
	sex := ""
	fmt.Scanln(&username, &sex)
	fmt.Println("username:", username)
	fmt.Println("sex:", sex)
}
