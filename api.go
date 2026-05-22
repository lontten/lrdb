package lrdb

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Stringer interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd

	//
	//Append(ctx context.Context, key, value string) *IntCmd
	//Decr(ctx context.Context, key string) *IntCmd
	//DecrBy(ctx context.Context, key string, decrement int64) *IntCmd
	//DelExArgs(ctx context.Context, key string, a DelExArgs) *IntCmd
	//Digest(ctx context.Context, key string) *DigestCmd
	//GetRange(ctx context.Context, key string, start, end int64) *StringCmd
	//GetSet(ctx context.Context, key string, value interface{}) *StringCmd
	//GetEx(ctx context.Context, key string, expiration time.Duration) *StringCmd
	//GetDel(ctx context.Context, key string) *StringCmd
	//Incr(ctx context.Context, key string) *IntCmd
	//IncrBy(ctx context.Context, key string, value int64) *IntCmd
	//IncrByFloat(ctx context.Context, key string, value float64) *FloatCmd
	//LCS(ctx context.Context, q *LCSQuery) *LCSCmd
	//MGet(ctx context.Context, keys ...string) *SliceCmd
	//MSet(ctx context.Context, values ...interface{}) *StatusCmd
	//MSetNX(ctx context.Context, values ...interface{}) *BoolCmd
	//MSetEX(ctx context.Context, args MSetEXArgs, values ...interface{}) *IntCmd
	//SetArgs(ctx context.Context, key string, value interface{}, a SetArgs) *StatusCmd
	//SetEx(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd
	//SetIFEQ(ctx context.Context, key string, value interface{}, matchValue interface{}, expiration time.Duration) *StatusCmd
	//SetIFEQGet(ctx context.Context, key string, value interface{}, matchValue interface{}, expiration time.Duration) *StringCmd
	//SetIFNE(ctx context.Context, key string, value interface{}, matchValue interface{}, expiration time.Duration) *StatusCmd
	//SetIFNEGet(ctx context.Context, key string, value interface{}, matchValue interface{}, expiration time.Duration) *StringCmd
	//SetIFDEQ(ctx context.Context, key string, value interface{}, matchDigest uint64, expiration time.Duration) *StatusCmd
	//SetIFDEQGet(ctx context.Context, key string, value interface{}, matchDigest uint64, expiration time.Duration) *StringCmd
	//SetIFDNE(ctx context.Context, key string, value interface{}, matchDigest uint64, expiration time.Duration) *StatusCmd
	//SetIFDNEGet(ctx context.Context, key string, value interface{}, matchDigest uint64, expiration time.Duration) *StringCmd
	//SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd
	//SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd
	//SetRange(ctx context.Context, key string, offset int64, value string) *IntCmd
	//StrLen(ctx context.Context, key string) *IntCmd
}
