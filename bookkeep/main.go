package main

import (
	"fmt"
	"golang_learn/bookkeep/model"
)

func main() {
	test(model.OUTLAY)
	key := ""
	loop := true
	account := model.NewAccount()
	for loop {
		mainMean()
		fmt.Scanln(&key)
		switch key {
		case "1":
			detail(account)
		case "2":
			checkIn(account, model.INCOME)
		case "3":
			checkIn(account, model.OUTLAY)
		case "4":
			fmt.Println("你退出了家庭收支记账软件")
			loop = false
		default:
			fmt.Println("输入错误,请输入正确选项！")
		}
	}
}

func detail(account *model.Account) {
	fmt.Println("---------------当前收支明细记录---------------")
	fmt.Println("收支\t账户余额\t收支金额\t说明")
	var s string
	s = fmt.Sprintf("%v\t%v\t%v\t%v", account.GetTurnover(), account.Deposit(), account.Money(), account.Notes())

	fmt.Println(s)
}

func checkIn(account *model.Account, turnover model.Turnover) {
	var money float64
	var notes string
	fmt.Print("请输入金额：")
	fmt.Scanln(&money)
	fmt.Print("请输入说明：")
	fmt.Scanln(&notes)
	account.SetMoney(money)
	account.SetNotes(notes)
	account.SetTurnover(turnover)
	if turnover == model.INCOME {
		account.SetDeposit(account.Deposit() + money)
	} else if turnover == model.OUTLAY {
		account.SetDeposit(account.Deposit() - money)
	}
}

func mainMean() {
	fmt.Println("---------------家庭收支记账软件---------------")
	fmt.Println("               1. 收支明细")
	fmt.Println("               2. 登记收入")
	fmt.Println("               3. 登记支出")
	fmt.Println("               4. 退出软件")
	fmt.Print("请选择：")
}

func test(turnover model.Turnover) {
	var s = fmt.Sprintf("%v", turnover)
	fmt.Println(s)
}
