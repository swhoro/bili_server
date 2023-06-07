package global

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

type DefaultLogger struct {
	Level  zapcore.Level
	Logger zap.Logger
}

func (l *DefaultLogger) DefaultLog(message string) {
	switch l.Level {
	case zap.WarnLevel:
		l.Logger.Warn(message)
	case zap.InfoLevel:
		l.Logger.Info(message)
	}
}

var (
	// DB 全局数据库
	DB *gorm.DB
	// Secret 全局jwt令牌
	Secret string
	// App 全局fiber应用
	App    *fiber.App
	Logger *zap.Logger
)
