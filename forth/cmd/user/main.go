package main

import (
	"flag"
	"geektime_homework/forth/configs"
	"geektime_homework/forth/internal/biz"
	"geektime_homework/forth/internal/data"
	"geektime_homework/forth/internal/server"
	"geektime_homework/forth/internal/service"
	"geektime_homework/forth/pkg/app"
	"geektime_homework/forth/pkg/transport/grpc"
	"geektime_homework/forth/pkg/transport/http"
)

var (
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs/config.json", "config path, eg: -conf config.json")
}

func newApp(gs *grpc.Server, hs *http.Server) *app.App {
	return app.New(
		hs,
		gs,
	)
}

func main() {
	flag.Parse()

	c := configs.NewConf(flagconf)
	if err := c.Load(); err != nil {
		panic(err)
	}
	app, err := initApp(c)
	if err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func initApp(conf *configs.Conf) (*app.App, error) {
	userRepo := data.NewUserRepoInMemory()
	userUseCase := biz.NewUserUseCase(userRepo)
	userService := service.NewUserService(userUseCase)
	grpcServer := server.NewGrpcServer(conf, userService)
	httpServer := server.NewHTTPServer(conf, userService)
	appApp := newApp(grpcServer, httpServer)
	return appApp, nil
}
