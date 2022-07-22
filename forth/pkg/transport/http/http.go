package http

import (
    "context"
    "errors"
    "net"
    "net/http"
)

type Server struct {
    *http.Server
    lis     net.Listener
    err     error
    baseCtx context.Context
    address string
}

func NewServer(address string) *Server {
    srv := &Server{
        address: address,
    }

    srv.Server = &http.Server{}
    srv.err = srv.Listen()

    return srv
}
func (s *Server) Start(ctx context.Context) error {
    if s.err != nil {
        return s.err
    }
    s.baseCtx = ctx
    err :=s.Serve(s.lis)
    if !errors.Is(err, http.ErrServerClosed){
        return err
    }
    return nil
}

func (s *Server) Stop(ctx context.Context) error {
    return s.Shutdown(ctx)
}

func (s *Server) Listen() error {
    if s.lis == nil {
        lis, err := net.Listen("tcp", s.address)
        if err != nil {
            return err
        }
        s.lis = lis
    }
    return nil
}
