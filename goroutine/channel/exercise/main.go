package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	var personChan chan interface{}
	personChan = make(chan interface{}, 10)
	for i := 0; i < 10; i++ {
		name := rand.Intn(100)
		age := rand.Intn(30)
		address := rand.Intn(10000)

		person := Person{
			Name:    strconv.Itoa(name),
			Age:     age,
			Address: strconv.Itoa(address),
		}
		personChan <- person
	}

	fmt.Println(len(personChan))
	for len(personChan) > 0 {
		person := <-personChan
		a := person.(Person)
		fmt.Printf("Name=%v,Age=%v,Address=%v\n", a.Name, a.Age, a.Address)
	}

	intChan := make(chan int,2)

	intChan <- 10
	intChan <- 10
	close(intChan)
	for i:=1;i<10;i++{
		v, ok := <-intChan
		fmt.Printf("内容=%v，ok=%v\n",v,ok)
	}
}
