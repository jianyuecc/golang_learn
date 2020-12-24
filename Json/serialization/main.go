package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"name"`
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2011-11-11",
		Sal:      8000.0,
		Skill:    "鼻子大",
	}
	//序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 err=%v", err)
	}
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v", string(data))
}

//将map进行序列化
func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["addr"] = "天安门"
	//序列化
	data, err := json.Marshal(&a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v", err)
	}
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v", string(data))
}

//将切片序列化
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "上海"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["address"] = "北京"
	slice = append(slice, m2)

	//序列化
	data, err := json.Marshal(&slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v", err)
	}
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v", string(data))

}

func main() {
	testStruct()
	testMap()
	testSlice()

}
