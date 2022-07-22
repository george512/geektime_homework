package app

import (
    "context"
    "geektime_homework/forth/pkg/transport"
    "golang.org/x/sync/errgroup"
    "sync"
)

type App struct {
    ctx    context.Context
    cancel func()
    srvs   []transport.Server
}

func New(httpserver transport.Server, grpcserver transport.Server) *App {
    ctx, cancel := context.WithCancel(context.Background())
    return &App{
        ctx:    ctx,
        cancel: cancel,
        srvs:   []transport.Server{httpserver, grpcserver},
    }
}

func (a *App) Run() error {
    eg, ctx := errgroup.WithContext(a.ctx)
    wg := sync.WaitGroup{}

}
