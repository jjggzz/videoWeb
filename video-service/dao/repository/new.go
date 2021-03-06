// Code generated by sqlxgen. DO NOT EDIT.
package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	VideoRepo VideoRepository
}

func NewRepo(db *sqlx.DB) Repository {
	return Repository{
		VideoRepo: NewVideoRepo(db),
	}
}

func NewVideoRepo(db *sqlx.DB) VideoRepository {
	return &videoRepo{
		db: db,
	}
}

type VideoRepository interface {
	Count() (int64, error)
	Insert(*Video) (int64, error)
	UpdateByPrimaryKey(int64, *Video) (int64, error)
	DeleteByPrimaryKey(int64) (int64, error)
	SelectByPrimaryKey(int64) (*Video, error)
	UpdateByExample(*VideoExample, *Video) (int64, error)
	DeleteByExample(*VideoExample) (int64, error)
	SelectByExample(*VideoExample) ([]*Video, error)
}
