package main

import "fmt"

/*
var 变量名 chan 数据类型

举例:
var intChan chan int(intChan用于存放int数据)
var mapChan chan map[int]string(mapChan用于存放map[int]string类型)
var perChan chan Person
var perChan2 chan * Person

说明
1)channel是引用类型
2)channel必须初始化才能写入数据,即make后才能使用
3)管道是有类型的,intChan只能写入整数int

*/

//channel 是线程安全的,多个协程操作同一个管道时,不会发生资源竞争问题
func main() {

	//演示管道的使用
	//1. 创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	//2. 看看intChan是什么
	fmt.Printf("intChan的值=%v.intChan本身的地址=%p\n", intChan, &intChan)

	//3. 向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num

	//注意,当给管道写入数据时,不能超过cap(容量)
	intChan <- 50

	//4. 看看管道的长度和cap(容量)
	fmt.Printf("channel len=%v\nchannel cap=%v\n", len(intChan), cap(intChan))

	//5. 读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len=%v\nchannel cap=%v\n", len(intChan), cap(intChan))

	//6. 在没有使用协程的情况下,如果管道数据全部取出,再取就会报告deadlock
	num3 := <-intChan
	num4 := <-intChan
	num5 := <-intChan
	fmt.Println("num3=", num3, "num4=", num4, "num5=", num5)
}
