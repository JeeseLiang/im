package xlog

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	// 日志文件路径
	LogPath = "logs/service.log"
)

// ServiceColors 定义服务名称对应的颜色
var ServiceColors = map[string]string{
	"user-rpc":  ColorGreen,
	"user-api":  ColorBlue,
	"group-rpc": ColorYellow,
	"group-api": ColorPurple,
	"msg-rpc":   ColorCyan,
	"msg-api":   ColorRed,
}

// ServiceWriter 是一个自定义的 io.Writer，用于处理服务输出
type ServiceWriter struct {
	serviceName string
	color       string
	logFile     *os.File
}

// NewServiceWriter 创建一个新的 ServiceWriter
func NewServiceWriter(serviceName string, logFile *os.File) *ServiceWriter {
	color := ServiceColors[serviceName]
	if color == "" {
		color = ColorReset
	}
	return &ServiceWriter{
		serviceName: serviceName,
		color:       color,
		logFile:     logFile,
	}
}

// Write 实现 io.Writer 接口
func (w *ServiceWriter) Write(p []byte) (n int, err error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")

	// 带颜色的终端输出
	coloredOutput := fmt.Sprintf("%s[%s]%s %s %s", w.color, w.serviceName, ColorReset, timestamp, string(p))
	fmt.Print(coloredOutput)

	// 不带颜色的文件输出
	fileOutput := fmt.Sprintf("[%s] %s %s", w.serviceName, timestamp, string(p))
	w.logFile.WriteString(fileOutput)

	return len(p), nil
}

// NewLogFile 创建日志文件
func NewLogFile(filename string) (*os.File, error) {
	// 确保日志目录存在
	dir := "logs"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("创建日志目录失败: %v", err)
		}
	}
	return os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
}

// GetServiceWriter 获取服务的标准输出和错误输出
func GetServiceWriter(serviceName string, logFile *os.File) (io.Writer, io.Writer) {
	writer := NewServiceWriter(serviceName, logFile)
	errorWriter := NewServiceWriter(serviceName, logFile)
	errorWriter.color = ColorRed
	return writer, errorWriter
}
