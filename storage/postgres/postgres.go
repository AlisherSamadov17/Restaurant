package postgres

import (
	"database/sql"
	"fmt"
	"myproject/config"
)

type Store struct {
	DB          *sql.DB
	UserStorage userRepo
	OrderStorage orderRepo
	ProductStorage productRepo
}

func New(cfg config.Config) (Store, error) {
	url := fmt.Sprintf(`host=%s port=%s user=%s password=%s database=%s sslmode=disable`, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return Store{}, err
	}

	userRepo := NewUserRepo(db)
	orderRepo := NewOrderRepo(db)
	productRepo :=NewProductRepo(db)

	return Store{
		DB:          db,
		UserStorage: userRepo,
		OrderStorage: orderRepo,
		ProductStorage: productRepo,
	}, nil
}
