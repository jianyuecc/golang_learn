package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age int
}


func main() {

	student := Student{"jack", 20}
	student.Name = "张三"
	test(student)





}

func test(exam interface{}) {
	exam = reflect.ValueOf(exam)
	student := exam.(Student)
	fmt.Println(student.Name)
}