package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"im_message/common/xlog"
)

func main() {
	// 创建日志文件
	logFile, err := xlog.NewLogFile(xlog.LogPath)
	if err != nil {
		log.Fatal("无法创建日志文件:", err)
	}
	defer logFile.Close()

	// 创建上下文，用于优雅退出
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 启动顺序：user -> group -> msg
	// 每个服务按照 rpc -> api 的顺序启动
	services := []struct {
		name string
		path string
	}{
		{"user-rpc", "app/user/rpc/user.go"},
		{"user-api", "app/user/api/user.go"},
		{"group-rpc", "app/group/rpc/group.go"},
		{"group-api", "app/group/api/group.go"},
		{"msg-rpc", "app/msg/rpc/msg.go"},
		{"msg-api", "app/msg/api/msg.go"},
	}

	// 用于存储所有启动的进程
	var processes []*exec.Cmd

	// 启动所有服务
	for _, service := range services {
		cmd := exec.CommandContext(ctx, "go", "run", service.path)

		// 获取服务的标准输出和错误输出
		stdout, stderr := xlog.GetServiceWriter(service.name, logFile)
		cmd.Stdout = stdout
		cmd.Stderr = stderr

		log.Printf("%s[%s]%s 正在启动服务\n", xlog.ServiceColors[service.name], service.name, xlog.ColorReset)
		if err := cmd.Start(); err != nil {
			log.Printf("%s[%s]%s 启动失败: %v\n", xlog.ColorRed, service.name, xlog.ColorReset, err)
			// 如果启动失败，停止所有已启动的服务
			for _, p := range processes {
				p.Process.Kill()
			}
			os.Exit(1)
		}

		processes = append(processes, cmd)
		// 等待一小段时间确保服务正常启动
		time.Sleep(2 * time.Second)
	}

	log.Println("所有服务已启动")

	// 等待中断信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("正在关闭所有服务...")

	// 优雅关闭所有服务
	for _, cmd := range processes {
		if cmd.Process != nil {
			cmd.Process.Signal(syscall.SIGTERM)
		}
	}

	// 等待所有进程结束
	for _, cmd := range processes {
		cmd.Wait()
	}

	log.Println("所有服务已关闭")
}

// ServiceWriter 是一个自定义的 io.Writer，用于处理服务输出
type ServiceWriter struct {
	serviceName string
	color       string
	logFile     *os.File
}

func (w *ServiceWriter) Write(p []byte) (n int, err error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	output := fmt.Sprintf("%s[%s]%s %s %s", w.color, w.serviceName, xlog.ColorReset, timestamp, string(p))

	// 同时输出到控制台和文件
	fmt.Print(output)
	w.logFile.WriteString(output)

	return len(p), nil
}
