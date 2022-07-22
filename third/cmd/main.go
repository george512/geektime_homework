package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	// pprof服务
	serverDebug := http.Server{Addr: ":8001"}

	// app服务
	serverAppMux := http.NewServeMux()
	serverAppMux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})
	serverApp := http.Server{Handler: serverAppMux, Addr: ":8080"}

	// 构建server切片
	servers := []*http.Server{&serverDebug, &serverApp}

	ctx, cancel := context.WithCancel(context.Background())
	eg, errCtx := errgroup.WithContext(ctx)

	var wg sync.WaitGroup
	for _, srv := range servers {
		// 使srv只在当前循环下可见, 防止因为闭包所有goroutine获取到同一个srv
		srv := srv
		// 关闭服务goroutine
		eg.Go(func() error {
			<-errCtx.Done()                                             // 等待cancel()函数进行统一的关闭
			stopCtx, sCancel := context.WithTimeout(ctx, time.Second*2) // 设置超时关闭
			defer sCancel()
			return srv.Shutdown(stopCtx)
		})
		wg.Add(1)
		// 启动服务goroutine
		eg.Go(func() error {
			wg.Done()
			return srv.ListenAndServe()
		})
	}
	// 让主goroutine等待其他子gouritne启动
	wg.Wait()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	eg.Go(func() error {
		for {
			select {
			// 捕获到系统退出信号, 调用cancel停止所有子goroutine
			case <-quit:
				cancel()
			case <-errCtx.Done(): //子goroutine退出
				return nil
			}
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		fmt.Errorf("errgroup.Wait:%v", err)
	}
	fmt.Println("all servers shutdown.")
}
