package main

import (
	"fmt"
	"golang_learn/customerManage/model"
	"golang_learn/customerManage/service"
)

type customerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

//主菜单

func (cv *customerView) mainMean() {
	for cv.loop {
		fmt.Println("---------------客户信息管理系统---------------")
		fmt.Println("               1 添 加 客 户")
		fmt.Println("               2 修 改 客 户")
		fmt.Println("               3 删 除 客 户")
		fmt.Println("               4 客 户 列 表")
		fmt.Println("               5 退      出")
		fmt.Print("请选择(1-5)：")
		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			cv.update()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.exit()
		default:
			fmt.Println("您的输入有误，请重新输入...")

		}
	}
}

func main() {
	customerView := customerView{
		key:  "",
		loop: true,
	}
	customerView.customerService = service.NewCustomerService()
	customerView.mainMean()

}

func (cv customerView) delete() {
	var id int
	fmt.Print("请输入要删除用户编号：")
	fmt.Scanln(&id)
	if cv.customerService.Delete(id) {
		fmt.Println("-------------删 除 成 功-------------")
	} else {
		fmt.Println("删除失败，用户不存在！")
	}
}

func (cv customerView) add() {
	var name string
	var gender string
	var age int
	var phone string
	var email string

	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入性别：")
	fmt.Scanln(&gender)
	fmt.Print("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Print("请输入手机号：")
	fmt.Scanln(&phone)
	fmt.Print("请输入邮箱：")
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name, gender, age, phone, email)

	if cv.customerService.Add(customer) {
		fmt.Println("-------------添 加 完 成-------------")
	} else {
		fmt.Println("-------------添 加 失 败-------------")
	}

}

func (cv customerView) update() {
	var id int
	var name string
	var gender string
	var age int
	var phone string
	var email string
	fmt.Print("请输入需要修改的用户编号：")
	fmt.Scanln(&id)
	index := cv.customerService.FindIndexById(id)
	if index == -1 {
		fmt.Println("用户不存在！")
	} else {
		fmt.Print("请输入姓名：")
		fmt.Scanln(&name)
		fmt.Print("请输入性别：")
		fmt.Scanln(&gender)
		fmt.Print("请输入年龄：")
		fmt.Scanln(&age)
		fmt.Print("请输入手机号：")
		fmt.Scanln(&phone)
		fmt.Print("请输入邮箱：")
		fmt.Scanln(&email)
		customer := model.NewCustomer(id, name, gender, age, phone, email)
		cv.customerService.Update(customer, index)
	}

}

func (cv customerView) list() {
	customers := cv.customerService.List()
	//显示
	fmt.Println("---------------客 户 列 表---------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].ShowInfo())
	}
	fmt.Println("---------------客户列表完成---------------")

}

func (cv customerView) exit() {
	fmt.Print("确认是否退出(Y/N)：")
	for {
		fmt.Scanln(&cv.key)
		if cv.key == "Y" || cv.key == "y" || cv.key == "N" || cv.key == "n" {
			break
		}
		fmt.Println("你的输入有误，确认是否退出(Y/N)：")
	}
	cv.loop = false
}
