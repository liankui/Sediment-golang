package main

import "fmt"

type myStruct struct {
	myField int
}

func main() {
	var value myStruct
	value.myField = 3
	var pointer *myStruct = &value
	//fmt.Println(*pointer.myField) // 报错，go认为myField必须是一个指针
	fmt.Println((*pointer).myField)	// 获取指针指向的struct的值，然后访问struct的字段
	fmt.Println(pointer.myField)	// 点运算符允许通过struct的指针来访问字段

	pointer.myField = 9 // 也可以通过指针来赋值给struct字段
}

