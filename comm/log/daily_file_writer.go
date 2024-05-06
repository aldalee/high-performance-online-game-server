package log

import (
	"fmt"
	"io"
	"os"
	"path"
	"sync"
	"time"
)

type dailyFileWriter struct {
	fileName       string      // 日志文件名称
	lastDate       int         // 上一次写入日期
	outputFile     *os.File    // 输出文件
	fileSwitchLock *sync.Mutex // 文件交换锁
}

func (w *dailyFileWriter) Write(b []byte) (int, error) {
	if b == nil || len(b) <= 0 {
		return 0, nil
	}
	outputFile, err := w.getOutputFile()
	if err != nil {
		return 0, err
	}
	_, _ = os.Stdout.Write(b)
	_, _ = outputFile.Write(b)
	return len(b), nil
}

// 获取输出文件，每天创建一个新的日志文件
func (w *dailyFileWriter) getOutputFile() (io.Writer, error) {
	date := time.Now().YearDay()
	if w.lastDate == date && w.outputFile != nil {
		return w.outputFile, nil
	}
	// 只在跨天时加锁，保证并发安全
	w.fileSwitchLock.Lock()
	defer w.fileSwitchLock.Unlock()
	w.lastDate = date
	err := os.MkdirAll(path.Dir(w.fileName), os.ModePerm)
	if err != nil {
		return nil, err
	}
	// 定义日志文件名称
	newDailyFile := w.fileName + "-" + time.Now().Format("20060102") + ".log"
	outputFile, err := os.OpenFile(
		newDailyFile,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644, // rw-r--r--
	)
	if err != nil || outputFile == nil {
		return nil, fmt.Errorf("open file %s failed: %v", newDailyFile, err)
	}
	if w.outputFile != nil {
		_ = w.outputFile.Close() // 关闭原来的文件
	}
	w.outputFile = outputFile
	return outputFile, nil
}
