package initialize

import (
	"github.com/MephistoSolsist/mysql-practice/global"
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", global.CONFIG.MySqlConfig.Dsn())
	if err != nil {
		return nil
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return db
}
