package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	// 错误写法
	for _, stu := range stus {
		m[stu.Name] = &stu // 指向同一个指针，最终该指针的值为遍历的最后一个struct的值拷贝
						   // 循环中使用的这个变量 &stu 会在每一次迭代被重新赋值，在赋值时也发生了拷贝
	}
	for _, stu := range stus {
		stu.Age = stu.Age + 10
	}

	for k, v := range m {
		fmt.Println(k, "->", v)
	}
	fmt.Println("-----")

	// 正确写法
	for k, stu := range stus {
		m[stu.Name] = &stus[k]
	}
	for k, v := range m {
		fmt.Println(k, "->", v)
	}

	ints := make([]int, 0)
	fmt.Println(reflect.TypeOf(ints))
	i := new([2]int)
	fmt.Println(reflect.TypeOf(i))
}
