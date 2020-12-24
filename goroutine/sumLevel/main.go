package main

import (
	"fmt"
	"sync"
	"time"
)

//思路
//1. 编写一个函数，来计算各个数的阶乘法，并放入到map中
//2. 我们启动的协程多个，统一将结果放入到map中
//3. map应该做出一个全局的

var (
	sumLevel, accumulator = make(map[int]int, 10), Accumulator(1)

	//全局互斥锁
	//sync 是包: synchronized 同步
	lock sync.Mutex
)

func test(n int) {

	res := 1
	for i := 1; i <= n; i++ {
		res *= i
		lock.Lock()
		accumulator()
		lock.Unlock()
	}
	//将res放入map中
	lock.Lock()
	sumLevel[n] = res
	lock.Unlock()
}

func main() {

	// 开启多个协程
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 5)

	lock.Lock()
	//输出结果
	for i := 1; i <= 20; i++ {
		fmt.Printf("map[%d]=%d\n", i, sumLevel[i])
	}
	lock.Unlock()
	fmt.Println(accumulator())

}

func Accumulator(iValue int) func() int {
	return func() int {
		iValue++
		return iValue
	}

}
