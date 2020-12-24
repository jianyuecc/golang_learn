package main

import (
	"fmt"
	"reflect"
)

// 通过反射修改
// num int 的值
// 修改student的值

func reflect01(b interface{})  {
	rVal := reflect.ValueOf(b)
	fmt.Println(rVal.Kind())
	rVal.Elem().SetInt(20)
}

func main() {
	var num int = 10
	reflect01(&num)
	fmt.Println(num)
}


