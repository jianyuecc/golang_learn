package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()

	fmt.Println("cpuNum=", cpuNum)

	//自己设置使用多少个cpu
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")

}
