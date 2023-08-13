package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	g "b.carriage.fun/server/global"
	"b.carriage.fun/server/utils/logger/operationCode"
)

// func newLogger(filename string) *zap.Logger {
// 	hook := lumberjack.Logger{
// 		Filename: "logs/" + filename + ".log.json", // 日志文件路径
// 		MaxSize:  100,                              // 每个日志文件保存的最大尺寸 单位：M
// 		MaxAge:   30,                               // 文件最多保存多少天
// 		Compress: true,                             // 是否压缩
// 	}

// 	atomicLevel := zap.NewAtomicLevel()

// 	config := zap.NewProductionEncoderConfig()
// 	config.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05:00Z0700")
// 	encoder := zapcore.NewJSONEncoder(config)
// 	core := zapcore.NewCore(
// 		encoder,
// 		zapcore.AddSync(&hook),
// 		atomicLevel,
// 	)

// 	if filename == "internal" {
// 		return zap.New(core, zap.AddCaller())
// 	}
// 	return zap.New(core)
// }

func SetupLogger() {
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05:00Z0700")
	config := zap.NewProductionConfig()
	config.EncoderConfig = encoder
	g.Logger, _ = config.Build()
}

func LogNotAuthorizeOperation(id uint32, operationID operationCode.Operation) {
	g.Logger.Info("no authoriazation operation", zap.Uint32("userId", id), zap.String("operation", operationCode.OperationToStr[operationID]))
}

func LogInternalError(err error) {
	g.Logger.Error("internal error", zap.String("error", err.Error()))
}
