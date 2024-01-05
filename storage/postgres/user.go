package postgres

import (
	"database/sql"
	"fmt"
	"myproject/models"
	"github.com/google/uuid"
)

type userRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) userRepo {
	return userRepo{
		DB: db,
	}
}

// inserting by using method here in postgres
func (u userRepo) Insert(user models.User) (string, error) {
	fmt.Println("user:", user)
	id := uuid.New()
	if _, err := u.DB.Exec(`insert into users values ($1,$2,$3,$4,$5)`, id, user.FirstName, user.LastName, user.Email, user.Phone); err != nil {
		return "", err
	}
	return id.String(), nil
}

// deleting user by using method here in postgres
func (u userRepo) Delete(id uuid.UUID) error {
	if _, err := u.DB.Exec(`delete from users where id = $1`, id); err != nil {
		return err
	}
	return nil
}

// updating user by using method here in postgres
func (u userRepo) Update(user models.User) error {
	if _, err := u.DB.Exec(`update users set first_name=$2,last_name=$3,email=$4,phone=$5 where id=$1`, user.ID, user.FirstName, user.LastName, user.Email, user.Phone); err != nil {
		return err
	}
	return nil 
}

// get user by id by using method here in postgres
func (u userRepo) GetByID(id uuid.UUID) (models.User, error) {
	user := models.User{}

	if err := u.DB.QueryRow(`select id,first_name,last_name,email,phone from users where id=$1`, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
	); err != nil {
		return models.User{}, err
	}
	return user, nil
}

// get all users by using method here in postgres
func (u userRepo) GetListUsers() ([]models.User, error) {
	users := []models.User{}

	rows, err := u.DB.Query(`select * from users`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := models.User{}

		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
