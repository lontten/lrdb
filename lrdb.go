package lrdb

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type Engine struct {
	db1 *redis.Client
	db2 *redis.ClusterClient
}

func NewClient(c redis.Options) Engine {
	db1 := redis.NewClient(&c)
	err := db1.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
	return Engine{db1: db1}
}

func New() Engine {
	//  单节点客户端
	db1 := redis.NewClient(&redis.Options{
		// 基础连接
		Addr:     "localhost:6379",
		Password: "", // 无密码
		DB:       0,  // 默认DB
	})

	//  Redis 集群客户端
	db2 := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"192.168.1.10:6379", "192.168.1.11:6379", "192.168.1.12:6379"},
		// 其他配置...
	})

	return Engine{db1: db1, db2: db2}
}

func Set(db Engine, key string, value any, expire time.Duration) error {
	ctx := context.Background()
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	v := string(bytes)

	if db.db1 != nil {
		return db.db1.Set(ctx, key, v, expire).Err()
	} else if db.db2 != nil {
		return db.db2.Set(ctx, key, v, expire).Err()
	}
	return errors.New("redis db is nil")
}
func Get[T any](db Engine, key string) (*T, error) {
	ctx := context.Background()
	var v string
	var err error
	if db.db1 != nil {
		v, err = db.db1.Get(ctx, key).Result()
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}
	} else if db.db2 != nil {
		v, err = db.db2.Get(ctx, key).Result()
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("redis db is nil")
	}

	var obj T
	err = json.Unmarshal([]byte(v), &obj)
	if err != nil {

		return nil, err
	}
	return &obj, nil
}
