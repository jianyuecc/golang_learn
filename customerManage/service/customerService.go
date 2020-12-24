package service

import "golang_learn/customerManage/model"

//增删改查
type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	//初始化
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(
		1,
		"张三",
		"男",
		20,
		"112",
		"zs@suhu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService

	return NewCustomerService()
}

func (customerService *CustomerService) Delete(id int) bool {
	index := customerService.FindIndexById(id)
	if index == -1 {
		return false
	} else {
		customerService.customers = append(customerService.customers[:index], customerService.customers[index+1:]...)
		return true
	}
}

func (customerService *CustomerService) List() []model.Customer {
	return customerService.customers
}

func (customerService *CustomerService) Add(customer model.Customer) bool {
	customerService.customerNum++
	customer.Id = customerService.customerNum
	customerService.customers = append(customerService.customers, customer)

	return true
}

func (customerService *CustomerService) FindIndexById(id int) int {
	for i, v := range customerService.customers {
		if v.Id == id { //找到了
			return i
		}
	}
	return -1
}

func (customerService *CustomerService) Update(customer model.Customer, index int) {
	customerService.customers[index] = customer
}
