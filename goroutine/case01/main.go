package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
在主线程中，开启一个goroutine，该协程每1秒输出“Hello World”
在主线程中也每隔1秒输出“Hello golang”，输出10次后，退出程序
要求主线程和goroutine同时执行
*/

func test() {
	for i := 1; i < 10; i++ {
		fmt.Println("test() hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test() //开启一个协程

	for i := 1; i < 10; i++ {
		fmt.Println("main() hello golang" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}
