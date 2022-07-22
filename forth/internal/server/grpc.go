package server

import (
    v1 "geektime_homework/forth/api/user/v1"
    "geektime_homework/forth/configs"
    "geektime_homework/forth/internal/service"
    "geektime_homework/forth/pkg/transport/grpc"
)

func NewGrpcServer(c *configs.Conf, service *service.UserService) *grpc.Server {
    if c.GrpcAddress == "" {
        c.GrpcAddress = "0.0.0.0:8090"
    }

    srv := grpc.NewServer(c.GrpcAddress)
    v1.RegisterUserServiceServer(srv, service)
    return srv
}
