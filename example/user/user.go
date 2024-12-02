package user

import (
	"github.com/gormbase"

	"gorm.io/gorm"
)

// 对应表实体，内嵌GormBase
type Model interface {
	gormbase.GormBase[int64, *DBObject]
}

type model struct {
	gormbase.GormBase[int64, *DBObject]
}

func New(db *gorm.DB) Model {
	return &model{
		gormbase.NewModelBase[int64, *DBObject](db),
	}
}

type DBObject struct {
	ID   int64  `gorm:"column:id;type:int;primaryKey;autoIncrement"` // 主键id
	Name string `gorm:"column:name;type:varchar"`                    // 名称
}

func (d *DBObject) TableName() string {
	return "user"
}

func (d *DBObject) GetID() int64 {
	return d.ID
}
