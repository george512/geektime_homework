package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"strconv"
	"sync"
	"time"
)

var (
	size = []int{
		10000,
		50000,
		100000,
		200000,
		300000,
		400000,
		500000,
		1000000,
	}
	beforePrefix = "=============Before=============\n"
	afterPrefix  = "==============After=============\n"
)

func main() {
	var (
		err    error
		conn   *redis.Client
		ctx    context.Context
		before string
		after  string
		wg     sync.WaitGroup
	)
	ctx = context.Background()

	if conn, err = initRedis(); err != nil {
		goto END
	}
	defer conn.Close()

	for _, s := range size {
		conn.FlushDB(ctx)

		before, err = conn.Info(ctx, "memory").Result()
		if err != nil {
			goto END
		}

		for i := 0; i < s; i++ {
			wg.Add(1)
			go func() {
				id, _ := uuid.NewUUID()
				conn.Set(ctx, id.String(), 0, 0)
			}()
		}

		wg.Wait()

		after, err = conn.Info(ctx, "memory").Result()
		if err != nil {
			goto END
		}
		err = ioutil.WriteFile(
			"./result/result_"+strconv.Itoa(s),
			[]byte(beforePrefix+"\n"+before+"\n\n"+afterPrefix+"\n"+after),
			0644,
		)
		if err != nil {
			goto END
		}
		conn.FlushDB(ctx)
		time.Sleep(3 * time.Second)
	}
END:
	log.Print(err.Error())
}

func initRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	return client, client.Ping(context.TODO()).Err()
}
