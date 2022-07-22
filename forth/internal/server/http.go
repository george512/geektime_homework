package server

import (
    v1 "geektime_homework/forth/api/user/v1"
    "geektime_homework/forth/configs"
    "geektime_homework/forth/internal/service"
    "geektime_homework/forth/pkg/transport/http"
    "github.com/gin-gonic/gin"
)

func NewHTTPServer(c *configs.Conf, service *service.UserService) *http.Server {
    engine := gin.Default()
    v1.RegisterUserHttpServer(engine, service)
    if c.HttpAddress == "" {
        c.HttpAddress = "0.0.0.0:8080"
    }

    srv := http.NewServer(c.HttpAddress)
    srv.Server.Handler = engine

    return srv
}
