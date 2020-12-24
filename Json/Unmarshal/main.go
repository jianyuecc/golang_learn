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

func unmarshalMap() {
	var str = "{\"addr\":\"天安门\",\"age\":30,\"name\":\"红孩儿\"}"
	var a map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("ummarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 monster=%v", a)

}

func unmarshalStruct() {
	var str = "{\"name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"鼻子大\"}"

	//定义一个Monster实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("ummarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 monster=%v", monster)

}

func unmarshalSlice() {
	var str = "[{\"address\":\"上海\",\"age\":\"7\",\"name\":\"jack\"}," +
		"{\"address\":\"北京\",\"age\":\"20\",\"name\":\"tom\"}]"
	var slice []map[string]interface{}

	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("ummarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 monster=%v", slice)
}

func main() {

	unmarshalMap()
	unmarshalStruct()
	unmarshalSlice()
}
