package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"b.carriage.fun/datamodel"
	g "b.carriage.fun/server/global"
	"b.carriage.fun/server/router"
	"b.carriage.fun/server/utils/logger"
)

func setupDatabase() {
	var err error
	g.DB, err = gorm.Open(sqlite.Open("./config/main.sqlite"), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})
	if err != nil {
		panic(err)
	}

	err = g.DB.AutoMigrate(&datamodel.User{}, &datamodel.BangumiItem{})
	if err != nil {
		panic(err.Error())
	}
}

func init() {
	setupDatabase()
	logger.SetupLogger()
	router.RegisterRouter()
}

var listenAddr string

func main() {
	err := g.App.Listen(listenAddr)
	if err != nil {
		panic(err)
	}
}
