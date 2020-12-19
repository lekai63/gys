package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

// Init 初始化数据库连接
func Init(c db.Connection) {

	// orm, err = gorm.Open("postgres", c.GetDB("default"))
	sqlDB := c.GetDB("default")
	DB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic("initialize orm failed")
	}

}
