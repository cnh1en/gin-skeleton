package initialize

import (
	"context"
	"fmt"

	"github.com/cnh1en/lets_go/global"
	"github.com/cnh1en/lets_go/pkg/utils"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func InitRedis() {
	rs := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", rs.Host, rs.Port),
		Password: rs.Password,
		DB:       rs.Db,
		PoolSize: rs.PoolSize,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		utils.CheckErrorPanic(err, "init redis")
	}

	global.Logger.Info("init redis success")

	global.Rdb = rdb

}
