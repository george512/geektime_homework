//go:build wireinject
// +build wireinject

package main

import (
	"geektime_homework/forth/configs"
	"geektime_homework/forth/internal/biz"
	"geektime_homework/forth/internal/data"
	"geektime_homework/forth/internal/server"
	"geektime_homework/forth/internal/service"
	"geektime_homework/forth/pkg/app"
	"github.com/google/wire"
)

func initApp(conf *configs.Conf) (*app.App, error) {
	panic(wire.Build(server.ProviderSet, service.ProvideSet, biz.ProvideSet, data.ProvideSet, newApp))
}
