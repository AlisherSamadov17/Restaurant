package main

import (
	// "fmt"
	"log"
	"myproject/config"
	"myproject/controller"
	"myproject/storage/postgres"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()
	store, err := postgres.New(cfg)
	if err != nil {
		log.Fatalln("error while connecting to db", err.Error())
		return
	}
	defer store.DB.Close()

	con := controller.New(store)

	// __ __user__ __ \\
	// fmt.Println("\n    Create user-1     ")
	// fmt.Println("\n    Delete user-2     ")
	// fmt.Println("\n    Update user-3     ")
	// fmt.Println("\n    Get user by ID-4  ")
	// fmt.Println("\n   Get all users-5 \n ")

	// var cmd int
	// fmt.Print("enter the command to select :")
	// fmt.Scan(&cmd)
	// switch cmd {
	// case 1:
	// 	con.CreateUser()

	// case 2:
	// 	con.DeleteUser()

	// case 3:
	// 	con.UpdateUser()

	// case 4:
	// 	con.GetUserByID()

	// case 5:
	// 	con.GetAllUsers()
	// }


//  __Order__
	// con.CreateOrder()
	// con.DeleteOrder()
	// con.GetAllOders()
	// con.GetOrderByID()

// _Product_
con.CreateProduct()
con.DeleteProduct()
con.GetAllProducts()
con.GetProductByID()
}