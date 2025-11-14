package core

import (
	"FishingSpotMarker/core/conf"
	"FishingSpotMarker/core/log"
	"FishingSpotMarker/core/router"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Init() error {
	if err := conf.Init(); err != nil {
		return err
	}

	if err := log.Init(); err != nil {
		return err
	}

	if err := router.Init(); err != nil {
		return err
	}

	return nil
}

func Run() error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.GetConf().Server.Port),
		Handler: router.GetRouter(),
	}

	// 在 goroutine 中启动服务器
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.StdErrorf("Listening failed: %s", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 设置优雅关闭超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 关闭服务器
	if err := server.Shutdown(ctx); err != nil {
		log.StdErrorf("Server Shutdown failed: %s", err)
	}

	return nil
}
