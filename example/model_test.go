package model

import (
	"context"
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestModel(t *testing.T) {
	mysqlInGorm, err := gorm.Open(
		mysql.New(
			mysql.Config{DSN: "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"}),
		&gorm.Config{SkipDefaultTransaction: true,
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		panic(err)
	}
	gormDB := New(mysqlInGorm)
	res, err := gormDB.UserModel().List(context.Background())
	if err != nil {
		panic(err)
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
