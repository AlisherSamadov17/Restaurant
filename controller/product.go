package controller

import (
	"fmt"
	"myproject/models"

	"github.com/google/uuid"
)

func (c Controller) CreateProduct()  {
product := getProductInfo_1()
id,err := c.Store.ProductStorage.Insert(product)
if err != nil {
	fmt.Println("error while creating data inside controller product.go(1) :",err.Error())
}
fmt.Println("id is :",id)
}

func (c Controller) DeleteProduct()  {
	product := getProductInfo_2().ID
	if err := c.Store.ProductStorage.Delete(product);err != nil {
		panic(err)
	}
	fmt.Println("product deleted !")
}

func (c Controller) UpdateProduct() {
	product := getProductInfo_3()
	err := c.Store.ProductStorage.Update(product)
	if err != nil {
		fmt.Println("error while updating product !",err)
		return
	}
	fmt.Println("Successully updated !")
	return
}

func (c Controller) GetProductByID()  {
	idStr := ""
	fmt.Print("enter id to find :")
	fmt.Scan(&idStr)

	id,err := uuid.Parse(idStr)
	if err != nil {
	fmt.Println("id is not found :",err.Error())	
	return
}
product,err := c.Store.ProductStorage.GetByID(id)
if err != nil {
	fmt.Println("error while getting product by id :",err.Error())
    return
}
fmt.Println("searched product is",product)
}



func (c Controller) GetAllProducts()  {
	products,err := c.Store.ProductStorage.GetListProducts()
  
if err != nil {
	fmt.Println("error while getting all of the product",err.Error())
  return
}
fmt.Println(products)
}



func getProductInfo_1() models.Product {
	var (
		price int
		name string
	)
	fmt.Print("enter price :")
	fmt.Scan(&price)

	fmt.Print("enter the name :")
	fmt.Scan(&name)

	return models.Product{
		Price: price,
		Name: name,
	}
}

func getProductInfo_2() models.Product {
	var (
		id string
	  )	
	   fmt.Print("insert 'ID' to delete : ")
	   fmt.Scan(&id)
  
	   return models.Product{
		ID: uuid.MustParse(id),
	   }
}
func getProductInfo_3() models.Product {
	var (
		id string
		price int
		name string
	 )	
	 fmt.Print("enter id to update new one :") 
	 fmt.Scan(&id)
 
	 fmt.Print("enter price to update : ")
	 fmt.Scan(&price)
	
	 fmt.Print("enter your name to update:")
	 fmt.Scan(&name)
	
	
	  return models.Product{
		ID: uuid.MustParse(id),	
		Price: price,
		Name: name,
		
	  }
}