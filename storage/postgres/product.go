package postgres

import (
	"database/sql"
	"fmt"
	"myproject/models"
	"github.com/google/uuid"
)


type productRepo struct {
	DB *sql.DB
}

func NewProductRepo(db *sql.DB) productRepo {
	return productRepo{
		DB: db,
	}
}

// insert
func (p productRepo) Insert(product models.Product) (string,error) {
fmt.Println("product:",product)
id := uuid.New()
  if _,err := p.DB.Exec(`insert into products values ($1,$2,$3)`,id,product.Price,product.Name);err != nil {
	return "",err
  }
  return id.String(),nil
}

// delete
func (p productRepo) Delete(id uuid.UUID) error {
if _,err :=p.DB.Exec(`delete from products where id=$1`,id);err != nil {
	return err
}	
return nil
}
// update
func (p productRepo) Update(product models.Product) error {
if _, err := p.DB.Exec(`update products set price=$2,name=$3 where id=$1`,product.ID,product.Price,product.Name);err != nil {
     return err
  }
  return nil
}


// get users by id
func (p productRepo) GetByID(id uuid.UUID) (models.Product,error)  {
product :=models.Product{}
if err := p.DB.QueryRow(`select id,price,name from products where id=$1`,id).Scan(
	&product.ID,
	&product.Price,
	&product.Name,
 );err != nil {
return models.Product{},err
 }
 return product,nil
}



// get all products
func (p productRepo) GetListProducts()([]models.Product,error)  {
products := []models.Product{}

rows, err := p.DB.Query(`select * from products`)
if err != nil {
return nil,err
}
for rows.Next(){
	product:=models.Product{}

	if err := rows.Scan(product.ID,product.Price,product.Name); err != nil {
		return nil,err
	}
	products = append(products, product)
}
return products,nil
}













