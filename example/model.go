package model

import (
	"github.com/gormbase/example/user"

	"gorm.io/gorm"
)

type gormDB struct {
	DB *gorm.DB
}

func New(db *gorm.DB) DB {
	return &gormDB{
		DB: db,
	}
}

func (g *gormDB) Transaction(transaction func(tx DB) error) error {
	return g.DB.Transaction(
		func(tx *gorm.DB) error {
			return transaction(New(tx))
		},
	)
}

func (g *gormDB) UserModel() user.Model {
	return user.New(g.DB)
}
