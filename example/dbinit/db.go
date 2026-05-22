package dbinit

import (
	"github.com/lontten/lrdb"
	"github.com/redis/go-redis/v9"
)

var DB lrdb.Engine

func init() {
	engine := lrdb.NewClient(redis.Options{
		Addr:     "127.0.0.1:6379",
		Username: "",
		Password: "",
		DB:       0,
	})
	DB = engine
}
