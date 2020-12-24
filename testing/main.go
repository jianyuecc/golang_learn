package main

import "fmt"

//被测试的函数

func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func main() {
	//传统的测试方法
	res := addUpper(10)
	if res != 55 {
		fmt.Printf("addUpper错误 返回值=%v 期望值=%v", res, 55)
	} else {
		fmt.Printf("addUpper正确 返回值=%v 期望值=%v", res, 55)
	}
}
