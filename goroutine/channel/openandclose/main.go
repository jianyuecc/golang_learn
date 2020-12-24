package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200

	//这时不能够再写入数据到channel
	//intChan<- 300

	//当管道关闭后,是可以读取数据的
	n1 := <-intChan
	fmt.Println(n1)

	//遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2 //放入100个数据
	}

	//遍历管道不能使用普通的 for 循环
	//for i := 0;i < len(intChan2); i++ {
	//
	//}

	//在遍历时,如果channel没有关闭,则会出现deadlock的错误
	//在遍历时,如果channel已经关闭,则会正常遍历,遍历完后,就会退出遍历
	close(intChan2)
	for v := range intChan2 {
		fmt.Printf("v=%v\n", v)
	}

}
