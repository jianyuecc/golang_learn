package main

import (
	"flag"
	"fmt"
)

func main() {
	//fmt.Println(len(os.Args))
	//
	//for i,v := range os.Args {
	//	fmt.Printf("第%v个参数为%v\n",i,v)
	//}

	var user string
	var pwd string
	var host string
	var port int
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口，默认3306")

	flag.Parse()
	fmt.Printf("user=%v pwd=%v host=%v port=%v", user, pwd, host, port)

}
