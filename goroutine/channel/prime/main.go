package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i := 1; i <= 200000; i++ {
		intChan <- i
	}
	close(intChan)
}

func isPrimeChan(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primeNum(primeChan, intChan chan int, exitChan chan bool) {
	for {
		num, ok := <-intChan
		if !ok {
			fmt.Println("有一个协程退出！")
			break
		}
		if isPrimeChan(num) {
			primeChan <- num
		}
	}
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 100000)
	exitChan := make(chan bool, 4)


	start := time.Now().Unix()
	go putNum(intChan)

	for i := 0; i < 4; i++ {
		go primeNum(primeChan, intChan, exitChan)
	}

	func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		end := time.Now().Unix()
		fmt.Println("time=",end - start)
		close(primeChan)
	}()


	//for {
	//	_, ok := <-primeChan
	//	if !ok {
	//		break
	//	}
	//
	//	fmt.Printf("素数=%v\n", res)
	//}

	fmt.Println("main线程退出")

}
