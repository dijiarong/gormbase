package model

import "github.com/dijiarong/gormbase/example/user"

type DB interface {
	Transaction(transaction func(tx DB) error) error
	// 注入对应实体
	UserModel() user.Model
}
