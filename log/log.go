package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SugarLogger *zap.SugaredLogger

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	SugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./metrics.log",
		MaxSize:    300,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
		//Filename: 日志文件的位置
		//MaxSize：在进行切割之前，日志文件的最大大小（以MB为单位）
		//MaxBackups：保留旧文件的最大个数
		//MaxAges：保留旧文件的最大天数
		//Compress：是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}
