package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (c Cal) GetSub(name string) {
	res := c.Num1 - c.Num2
	fmt.Printf("%v完成了减法运算，%v - %v = %v\n", name,c.Num1,c.Num2,res)
}

func main()  {
	var cal Cal

	val := reflect.ValueOf(&cal)
	val.Elem().Field(0).SetInt(8)
	val.Elem().Field(1).SetInt(2)

	inValue := make([]reflect.Value,1)
	inValue[0] = reflect.ValueOf("tom")
	val.Elem().Method(0).Call(inValue)
}




