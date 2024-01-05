package postgres

import (
	"database/sql"
	"fmt"
	"myproject/models"
	"github.com/google/uuid"
   _"golang.org/x/tools/go/analysis/passes/timeformat"
)


type orderRepo struct {
	DB *sql.DB
}

func NewOrderRepo(db *sql.DB) orderRepo {
	return orderRepo{
		DB: db,
	}
}

// inserting by using method here in postgres
func (o orderRepo) Insert(order models.Order) (string,error) {
	fmt.Println("order :",order)
	id := uuid.New()
	timeformat :=order.Created_at
   if _,err := o.DB.Exec(`insert into orders values ($1,$2,$3,$4)`,id,order.Amount,order.User_id,timeformat);err != nil {
	return "",err
   }
   return id.String(),nil
}

// deleting 
func (o orderRepo) Delete(id uuid.UUID) error {
	if _,err := o.DB.Exec(`delete from orders where id=$1`,id);err != nil{
        return err
	}
	return nil
}

// updating
func (o orderRepo) Update(order models.Order) error {
 if _,err := o.DB.Exec(`update orders set id=$1,user_id=$3,created_at=$4 where amount=$2`,order.ID,order.Amount,order.User_id,order.Created_at);err != nil{
	return err
 }
 return nil
}

// get order by id
func (o orderRepo) GetByID(id uuid.UUID) (models.Order,error) {
	order := models.Order{}

  if err := o.DB.QueryRow(`select id,amount,user_id,created_at where id=$1`,id).Scan(
    &order.ID,
	&order.Amount,
	&order.User_id,
	&order.Created_at,
	);err != nil{
		return models.Order{},err
	}
	return order, nil
}
// get all orders
func (o orderRepo) GetListOrders() ([]models.Order,error) {
orders := []models.Order{}

rows,err := o.DB.Query(`select * from orders`)
if err != nil {
	return nil,err
}
for rows.Next(){
	order := models.Order{}

	if err := rows.Scan(&order.ID,&order.Amount,&order.User_id,&order.Created_at);err != nil{
		return nil,err
	}
	orders = append(orders, order)
}
return orders,nil
}