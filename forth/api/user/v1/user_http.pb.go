package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func RegisterUserHttpServer(engine *gin.Engine, server UserServiceServer) {
	g := engine.Group(Prefix())
	{
		g.POST("/add", AddUserTransfer(server.AddUser))
		g.GET("/get/:id", GetUserTransfer(server.GetUser))
		g.POST("/list", ListUserTransfer(server.ListUser))
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
	return func(c *gin.Context) {
		var in AddUserRequest
		if err := c.ShouldBind(&in); err != nil {
			c.String(http.StatusBadRequest, "invalid argument")
			return
		}
		user, err := f(context.Background(), &in)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.AsciiJSON(http.StatusOK, user)
	}
}

func GetUserTransfer(f func(ctx context.Context, in *GetUserRequest) (*GetUserResponse, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		in := new(GetUserRequest)
		id := c.Param("id")
		if id == "" {
			c.String(http.StatusBadRequest, "invalid id")
			return
		}
		in.Id = id
		customer, err := f(context.Background(), in)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, customer)
	}
}

func ListUserTransfer(f func(ctx context.Context, in *ListUserRequest) (*ListUserResponse, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in ListUserRequest
		if err := c.ShouldBind(&in); err != nil {
			c.String(http.StatusBadRequest, "invalid argument")
			return
		}
		users, err := f(context.Background(), &in)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, users)
	}
}
