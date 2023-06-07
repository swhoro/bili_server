package main

import (
	"b.carriage.fun/global"
	"b.carriage.fun/model"
	"b.carriage.fun/router"
	"b.carriage.fun/utils/logger"
)

func init() {
	model.SetupDatabase()
	logger.SetupLogger()
	router.RegisterRouter()
}

var listenAddr string

func main() {
	err := global.App.Listen(listenAddr)
	if err != nil {
		panic(err)
	}
}
