package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

func main() {
	// 创建
	u1 := uuid.NewV4()
	fmt.Printf("UUID is : %s\n", u1)
	// 解析
	u2, err := uuid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("UUID is : %s\n", u2)
}
