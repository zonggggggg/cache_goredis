package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// var logger *logrus.Entry
var logger = logrus.New()

const (
	_path            = "../LOG/record.log"
	_infoPath        = "info.log"
	_errorPath       = "error.log"
	_rotationHour    = 24
	_maxFile         = 30
	_timeStampFormat = "2006-01-02-15:04:05"
)

type customFormatter struct{}

func (s *customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format(_timeStampFormat)
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel:
		levelColor = 36 // blue
	case logrus.InfoLevel:
		levelColor = 32 // green
	case logrus.WarnLevel:
		levelColor = 93 // yellow
	case logrus.ErrorLevel:
		levelColor = 31 // red
	case logrus.FatalLevel:
		levelColor = 41 // background red
	case logrus.PanicLevel:
		levelColor = 45 // background purple
	default:
		levelColor = 37
	}
	// var file string
	// var line int
	// if entry.Caller != nil {
	// 	file = filepath.Base(entry.Caller.File)
	// 	line = entry.Caller.Line
	// }
	// msg := fmt.Sprintf("%s [%s]%-1s[%s:%d] %1s\n", timestamp, strings.ToUpper(entry.Level.String()), "\t", file, line, entry.Message)
	// msg := fmt.Sprintf("%s [%s]%-1s%s\n", timestamp, strings.ToUpper(entry.Level.String()), "\t", entry.Message)
	msg := fmt.Sprintf("\x1b[%dm%s [%s] %s\n\x1b[0m", levelColor, timestamp, strings.ToUpper(entry.Level.String()), entry.Message)

	return []byte(msg), nil
}

func init() {
	// 日誌分割相關
	// `WithLinkName` 建立日誌
	// `WithRotationTime` 設定日誌分割的時間
	// WithMaxAge 和 WithRotationCount只能設定一個
	//  `WithMaxAge` 設定檔案清理前最長儲存的時間
	//  `WithRotationCount` 設定檔案清理最多儲存的數量
	// 下面配置每隔1天建立新檔案，保留最近30天的日誌檔案，其他的則清理掉

	writerFile, _ := rotatelogs.New(
		_path+".%Y%m%d",
		rotatelogs.WithLinkName(_path),
		// rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
		rotatelogs.WithRotationCount(_maxFile),
		rotatelogs.WithRotationTime(time.Duration(_rotationHour)*time.Hour),
	)
	writeStdout := os.Stdout

	logger.SetOutput(io.MultiWriter(writerFile, writeStdout))
	// logger.SetReportCaller(true)
	logger.SetFormatter(new(customFormatter))

	// 設定log level
	// logger.SetLevel(logrus.DebugLevel)

	// 分級分檔案
	// infoWriter, err := rotatelogs.New(
	// 	_infoPath+".%Y-%m-%d",
	// 	rotatelogs.WithLinkName(_infoPath),
	// 	rotatelogs.WithRotationTime(time.Duration(_rotationHour)*time.Hour),
	// 	rotatelogs.WithRotationCount(_countFile),
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// errorWriter, err := rotatelogs.New(
	// 	_errorPath+".%Y-%m-%d",
	// 	rotatelogs.WithLinkName(_errorPath),
	// 	rotatelogs.WithRotationTime(time.Duration(_rotationHour)*time.Hour),
	// 	rotatelogs.WithRotationCount(_countFile),
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// logger.AddHook(lfshook.NewHook(
	// 	lfshook.WriterMap{
	// 		logrus.InfoLevel:  infoWriter,
	// 		logrus.ErrorLevel: errorWriter,
	// 	},
	// 	&logrus.TextFormatter{
	// 		DisableColors:   true,
	// 		TimestampFormat: _timeStampFormat,
	// 	},
	// ))

}

// Set level
func SetLevel(level int) {
	// DEBUG: 5, INFO: 4, WARNING: 3, ERROR: 2, FATAL: 1
	logger.SetLevel(logrus.Level(level))
}

// Debug : log.Debug
func Debug(input ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	file = filepath.Base(file)
	msg := fmt.Sprintf("[%s:%d] %+v", file, line, input)
	logger.Debug(msg)
}

// Info : log.info
func Info(input ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	file = filepath.Base(file)
	msg := fmt.Sprintf("[%s:%d] %+v", file, line, input)
	logger.Info(msg)
}

// Warn : log.Warn
func Warn(input ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	file = filepath.Base(file)
	msg := fmt.Sprintf("[%s:%d] %+v", file, line, input)
	logger.Warn(msg)
}

// Error : log.Error
func Error(input ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	file = filepath.Base(file)
	msg := fmt.Sprintf("[%s:%d] %+v", file, line, input)
	logger.Error(msg)
}

// Fatal : log.Fatal
func Fatal(input ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	file = filepath.Base(file)
	// msg := fmt.Sprintf("[%s:%d] %1s", file, line, input)
	msg := fmt.Sprintf("[%s:%d] %+v", file, line, input)
	logger.Fatal(msg)
}

// Panic : log.Panic
func Panic(input ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	file = filepath.Base(file)
	msg := fmt.Sprintf("[%s:%d] %+v", file, line, input)
	logger.Panic(msg)
}

// func main() {
// 	// for i := 0; i < 15; i++ {
// 	for i := 0; ; i++ {
// 		if i%10 == 0 {
// 			log.Error("ERRORRRRRR!")
// 		} else {
// 			log.Info("hello, world!")
// 		}
// 		log.Debug("DDDDDDDDDDD")
// 		log.Warn("WWWWWWWWWWWW")

// 		time.Sleep(time.Duration(1) * time.Second)
// 	}
// }
