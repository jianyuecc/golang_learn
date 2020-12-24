package main

import (
	"fmt"
	"reflect"
)

func main() {

	var str string = "tom"
	fs := reflect.ValueOf(&str)
	fs.Elem().SetString("jack")
	fmt.Printf("%v\n",str)
}

func transform(b interface{})  {
	rVal := reflect.ValueOf(b)

	fmt.Printf("rVal=%v,rType=%v,rKind=%v\n",rVal,rVal.Type(),rVal.Kind())

	b = rVal.Interface()

	a := rVal.Float()
	fmt.Println(a)
}

