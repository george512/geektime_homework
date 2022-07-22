package v1

import (
    "context"
    "github.com/gin-gonic/gin"
    "os"
    "runtime"
    "strings"
)

func RegisterUserHttpServer(engine *gin.Engine, server UserServiceServer) {
    g := engine.Group(Prefix())
    {
        g.POST("/add", AddUserTransfer(server.AddUser))
        g.GET("/get/:id", GetUserTransfer(server.GetUser))
        g.POST("/list",ListUserTransfer(server.ListUser))
    }
}

// get api prefix according file location
func Prefix() string {
    dir, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    _, file, _, ok := runtime.Caller(0)
    if !ok {
        panic("failed to assign router prefix")
    }
    dirFrag := strings.Split(strings.Replace(file, dir, "", -1), "/")
    dirFrag = dirFrag[:len(dirFrag)-1]
    return strings.Join(dirFrag, "/")
}

func AddUserTransfer(f func(ctx context.Context, in *AddUserRequest) (*AddUserResponse, error)) gin.HandlerFunc {

}

func GetUserTransfer(f func(ctx context.Context, in *GetUserRequest) (*GetUserResponse, error)) gin.HandlerFunc {

}

func ListUserTransfer(f func(ctx context.Context, in *ListUserRequest) (*ListUserResponse, error)) gin.HandlerFunc {

}
