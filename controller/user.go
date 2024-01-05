package controller

import (
	"fmt"
	"myproject/models"
	"github.com/google/uuid"
)
func (c Controller) CreateUser() {
user := getUserInfo_1()

id,err := c.Store.UserStorage.Insert(user)
if err != nil {
	fmt.Println("error while creating data inside controller in user.go(1):",err.Error())
}
fmt.Println("id is :",id)
}

func (c Controller) DeleteUser() {
  user := getUserInfo_2().ID
  if err := c.Store.UserStorage.Delete(user);err != nil {
	 panic(err)
  }
	fmt.Println("user deleted !")
}

func (c Controller) UpdateUser()  {
	user := getUserInfo_3()
	err := c.Store.UserStorage.Update(user)
	if err != nil {
		fmt.Println("eror while updating user !",err)
		return
	}
	  fmt.Println("Successfully updated !")
	  return
}

func (c Controller) GetUserByID()  {
	idStr := ""
	fmt.Print("enter any user's id to find he or she :")
	fmt.Scan(&idStr)

	id,err := uuid.Parse(idStr)
	if err != nil {
	fmt.Println("id is not not uuid :",err.Error())
     return
	}
	user,err := c.Store.UserStorage.GetByID(id)
	if err != nil {
		fmt.Println("error while getting car by id err :",err.Error())
	    return
	}
	fmt.Println("searched user is :",user)
}

func (c Controller) GetAllUsers()  {
	users,err := c.Store.UserStorage.GetListUsers()
    if err != nil {
		fmt.Println("error while getting all of the users",err.Error())
	return
	}
	fmt.Println(users)
}




// adding new user by using getUserInfo_1
func getUserInfo_1() models.User {
 var (
	firstname string
	lastname string
	email string
	phone string
 )	
  
 fmt.Print("enter your first and last name: ")
 fmt.Scan(&firstname,&lastname)

 fmt.Print("enter your email:")
 fmt.Scan(&email)

 fmt.Print("enter your phone number:")
 fmt.Scan(&phone)

  return models.User{
    FirstName: firstname,
	LastName: lastname,
	Email: email,
	Phone: phone,
  }
}

// for deleting user by using id
func getUserInfo_2() models.User {
	var (
	  id string
	)	
	 fmt.Print("insert 'ID' to delete : ")
	 fmt.Scan(&id)

	 return models.User{
	  ID: uuid.MustParse(id),
	 }
}

// for updating user by 
   func getUserInfo_3() models.User {
	var (
       id string
	   firstname string
	   lastname string
	   email string
	   phone string 
	)	
	fmt.Print("enter id to update new one :") 
	fmt.Scan(&id)

	fmt.Print("enter your first and last name to update : ")
	fmt.Scan(&firstname,&lastname)
   
	fmt.Print("enter your email to update:")
	fmt.Scan(&email)
   
	fmt.Print("enter your phone number to update:")
	fmt.Scan(&phone)
   
	 return models.User{
	   ID: uuid.MustParse(id),	
	   FirstName: firstname,
	   LastName: lastname,
	   Email: email,
	   Phone: phone,
	 }
}   