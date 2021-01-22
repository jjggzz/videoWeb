package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	CustomerRepo CustomerRepository
}

func NewRepo(db *sqlx.DB) Repository {
	return Repository{
		CustomerRepo: NewCustomerRepo(db),
	}
}

func NewCustomerRepo(db *sqlx.DB) CustomerRepository {
	return &customerRepo{
		db: db,
	}
}

type CustomerRepository interface {
	Count() (int64, error)
	Insert(*Customer) (int64, error)
	UpdateByPrimaryKey(int64, *Customer) (int64, error)
	DeleteByPrimaryKey(int64) (int64, error)
	SelectByPrimaryKey(int64) (*Customer, error)
	UpdateByExample(*CustomerExample, *Customer) (int64, error)
	DeleteByExample(*CustomerExample) (int64, error)
	SelectByExample(*CustomerExample) ([]*Customer, error)
}
