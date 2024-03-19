package log

import (
	"context"
	"github.com/blocktock/go-pkg/tracex"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"time"
)

var logger *zap.Logger

func LoggerWithContext(ctx context.Context) *zap.Logger {
	v := ctx.Value(tracex.TraceIDCtx{})
	traceId := ""
	if v != nil {
		if tmpId, ok := v.(string); ok {
			traceId = tmpId
		}
	}
	return logger.With(zap.String("trace_id", traceId))
}

func InitLogger(outputPath string) *zap.Logger {

	encoder := getEncoder()

	//两个interface,判断日志等级
	//warnlevel以下归到info日志
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})

	//warnlevel及以上归到warn日志
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	infoWriter := getLogWriter(outputPath + "/info")
	warnWriter := getLogWriter(outputPath + "/warn")

	//创建zap.Core,for logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, infoWriter, infoLevel),
		zapcore.NewCore(encoder, warnWriter, warnLevel),
	)

	logger = zap.New(core, zap.AddCaller())

	//生成Logger
	//logger := zap.New(core, zap.AddCaller())
	//sugarLogger = logger.Sugar()

	return logger
}

// getEncoder
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 设置日志记录中时间的格式
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05+08:00")
	// 日志Encoder 还是JSONEncoder，把日志行格式化成JSON格式的
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 得到LogWriter
func getLogWriter(filePath string) zapcore.WriteSyncer {
	warnIoWriter := getWriter(filePath)
	return zapcore.AddSync(warnIoWriter)
}

// 日志文件切割
func getWriter(filename string) io.Writer {

	// 保存日志30天，每24小时分割一次日志

	hook, err := rotatelogs.New(
		filename+"_%Y%m%d.log",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*30),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	//保存日志30天，每1分钟分割一次日志
	//hook, err := rotatelogs.New(
	//	filename+"_%Y%m%d%H%M.log",
	//	rotatelogs.WithLinkName(filename),
	//	rotatelogs.WithMaxAge(time.Hour*24*30),
	//	rotatelogs.WithRotationTime(time.Minute*1),
	//)
	if err != nil {
		panic(err)
	}
	return hook
}
