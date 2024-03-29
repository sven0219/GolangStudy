package main

import (
	"fmt"
	"os"
)

// 1. 文件对象的类型
// 2. 获取文件对象的详细信息

func main() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		return
	}
	fmt.Printf("%T\n", fileObj) // 文件的类型
	fileInfo, err := fileObj.Stat()
	if err != nil {
		return
	}
	//获取文件大小
	fmt.Println(fileInfo.Size())
}
