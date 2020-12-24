package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().Unix()
	for i := 1; i <= 200000; i++ {
		if isPrimeChan(i) {

		}
	}

	end := time.Now().Unix()

	fmt.Println("time ", end - start)
}
func isPrimeChan(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}