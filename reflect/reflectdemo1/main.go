package main

import (
	"fmt"
	"reflect"
)

//专门演示反射
func reflectTest01(b interface{}){
	//通过反射获取传入的变量的type，kind，值
	//1. 先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=",rTyp)

	//2. 获取到reflect.Value
	rVal := reflect.ValueOf(b)

	n2 := 2 + rVal.Int()
	fmt.Println("n2=",n2)

	fmt.Printf("rVal=%v rVal type=%T\n",rVal,rVal)

	//将 rVal 转成 interface{}
	iV := rVal.Interface()
	//将 interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println(num2)


}


//专门演示反射[结构体]
func reflectTest02(b interface{})  {
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=",rTyp)

	rVal := reflect.ValueOf(b)
	iV := rVal.Interface()
	fmt.Printf("iV=%v iV type=%T\n",iV,iV)

	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}


}

type Monster struct {
	Name string
	Age int
}
type Student struct {
	Name string
	Age int
}

func main() {
	//请编写一个案例
	//演示对（基本数据类型、interface{}、reflect.Value）进行反射的基本操作

	//1. 先定义一个int
	var num int = 100
	reflectTest01(num)

	//2. 定义一个Student结构体
	stu := Student{
		Name: "tom",
		Age:  20,
	}

	reflectTest02(stu)

}

