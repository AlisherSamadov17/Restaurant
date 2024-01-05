package controller

import (
	"fmt"
	"myproject/models"
	"github.com/google/uuid"
)

func (c Controller) CreateOrder()  {
	order:=getOrderInfo_1()

	id,err := c.Store.OrderStorage.Insert(order)
	if err != nil {
		fmt.Println("error while creating data inside controller in order.go(1) err :",err.Error())
	}
	fmt.Println("id is :",id)
}

func (c Controller) DeleteOrder()  {
	order := getOrderInfo_2().ID

	if err := c.Store.OrderStorage.Delete(order);err != nil{
panic(err)
	}
	fmt.Println("order deleted !")
}
func (c Controller) UpdateOder()  {
	order := getOrderInfo_3()
	err := c.Store.OrderStorage.Update(order)
	if err != nil {
		fmt.Println("error while updating order !",err)
	   return
	}
	fmt.Println("Successfully updated !")
	return
}
 
func (c Controller) GetOrderByID()  {
	idStr := ""
	fmt.Print("enter any user's id to find :")
	fmt.Scan(idStr)

	id,err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("id is not uuid :",err.Error())
	 return
	}
    order,err:=c.Store.OrderStorage.GetByID(id)
    if err != nil {
		fmt.Println("error while getting order by id :",err.Error())
	  return
	}
	fmt.Println("searched order is :",order)
}

  func (c Controller) GetAllOders()  {
	orders,err := c.Store.OrderStorage.GetListOrders()
    if err != nil {
		fmt.Println("error while getting all of orders",err.Error())
       return
	}
	fmt.Println(orders)
}

func getOrderInfo_1() models.Order {
	var(
	
	amount string
	user_id string
	)

	fmt.Print("enter the amount :")
	fmt.Scan(&amount)

	fmt.Print("enter user id: ")
    fmt.Scan(&user_id)


	return models.Order {
		Amount: amount,
		User_id: uuid.MustParse(user_id),
	}
}

func getOrderInfo_2() models.User {
	var (
	  id string
	)	
	 fmt.Print("insert 'ID' to delete : ")
	 fmt.Scan(&id)

	 return models.User{
	  ID: uuid.MustParse(id),
	 }
}

func getOrderInfo_3() models.Order {
	var (
		amount string
	 )	 
	 fmt.Print("amount: ")
	 fmt.Scan(&amount)
	
	  return models.Order{
		Amount: amount,
	  }
}