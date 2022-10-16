/**
* @Author: yibo_LastDay
* @Date: 2022/10/16 12:43
 */

package ylog

import (
	"github.com/natefinch/lumberjack"
	"quick_frame/const_data"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"quick_frame/util"
)

var Log *zap.SugaredLogger

func InitLogger(logLevel zapcore.Level) {

	core := zapcore.NewCore(getEncoder(), logSync(const_data.LogDir), logLevel)
	logger := zap.New(core, zap.AddCaller())
	Log = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderCfg)
}

func logSync(path string) zapcore.WriteSyncer {
	if !util.IsDirExists(path) {
		util.MakeDir(path)
	}
	//file, err := os.Create("./log/qLog.log")
	//if err != nil {
	//	log.Fatal(fmt.Sprintf("create error file error, error:%v \n", err))
	//}
	jackLog := &lumberjack.Logger{
		Filename:   "./log/qLog.log",
		MaxSize:    15, //日志大小15MB
		MaxBackups: 30,
		MaxAge:     60,   //保存60天
		Compress:   true, //日志开启压缩
	}

	return zapcore.AddSync(jackLog)
}
