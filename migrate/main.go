package main

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"b.carriage.fun/datamodel"
)

func main() {
	c, err := os.Stat(datamodel.DatabaseDir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(datamodel.DatabaseDir, os.ModePerm)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
		if !c.IsDir() {
			os.Remove(datamodel.DatabaseDir)
			err = os.MkdirAll(datamodel.DatabaseDir, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	}

	_, err = os.Stat(datamodel.DatabasePath)
	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create(datamodel.DatabasePath)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	db, err := gorm.Open(sqlite.Open(datamodel.DatabasePath), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&datamodel.User{}, &datamodel.BangumiItem{})
	if err != nil {
		panic(err.Error())
	}
}
