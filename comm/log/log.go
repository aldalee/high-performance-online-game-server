package log

import (
	"fmt"
	"log"
)

var (
	writer      *dailyFileWriter
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
)

func Config(fileName string) {
	writer = &dailyFileWriter{
		fileName: fileName,
		lastDate: -1,
	}
	debugLogger = log.New(writer, "[DEBUG] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	infoLogger = log.New(writer, "[INFO] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	warnLogger = log.New(writer, "[WARN] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	errorLogger = log.New(writer, "[ERROR] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}

func Debug(format string, args ...any) {
	_ = debugLogger.Output(2, fmt.Sprintf(format, args...))
}

func Info(format string, args ...any) {
	_ = infoLogger.Output(2, fmt.Sprintf(format, args...))
}

func Warn(format string, args ...any) {
	_ = warnLogger.Output(2, fmt.Sprintf(format, args...))
}

func Error(format string, args ...any) {
	_ = errorLogger.Output(2, fmt.Sprintf(format, args...))
}
