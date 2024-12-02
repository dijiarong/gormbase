package model

import "gormbase/example/user"

type DB interface {
	Transaction(transaction func(tx DB) error) error
	// 注入对应实体
	UserModel() user.Model
}
