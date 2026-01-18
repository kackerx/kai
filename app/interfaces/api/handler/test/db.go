package test

import (
	"github.com/kackerx/kai/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, func(), error) {
	var dialector gorm.Dialector
	dialector = mysql.Open("root:Wasd4044@tcp(localhost:3306)/kai?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	db = db.Debug()

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}
	cleanFunc := func() {
		err := sqlDB.Close()
		if err != nil {
			logger.Errorf("Gorm db close error: %s", err.Error())
		}
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, cleanFunc, err
	}

	return db, cleanFunc, nil
}
