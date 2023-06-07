package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	g "b.carriage.fun/global"
)

func SetupDatabase() {
	var err error
	g.DB, err = gorm.Open(sqlite.Open("./config/main.sqlite"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	err = g.DB.AutoMigrate(&User{}, &BangumiItem{})
	if err != nil {
		panic(err.Error())
	}
}
