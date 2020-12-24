package main

import "fmt"

func main() {
	numChan := make(chan int,100)
	resChan := make(chan int,2000)
	exitChan := make(chan bool,8)

	go writeData(numChan)

	for i := 0; i < 8; i++ {
		go readData(numChan,resChan,exitChan)
	}


	func() {
		for i:=0;i<8;i++{
			<- exitChan
		}
		close(resChan)
		fmt.Println("main主线程退出")
	}()

	for  {
		res,ok := <- resChan
		if !ok {
			break
		}
		fmt.Println(res)
	}
}

func readData(numChan, resChan chan int,exitChan chan bool)  {
	for {
		num,ok := <- numChan
		if !ok {

			fmt.Println("协程退出！")
			break
		}
		resChan <- res(num)
	}
	exitChan <- true
}

func writeData(numChan chan int) {
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
	close(numChan)
}

func res(num int) int {
	var res int
	for i := 1; i < num; i++ {
		res += i
	}
	return res
}