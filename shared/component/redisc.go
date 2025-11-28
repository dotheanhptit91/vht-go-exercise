package sharedcomponent

import (
	"context"
	"encoding/json"
	"flag"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	sctx "github.com/viettranx/service-context"
)

type IRedisComp interface {
	Get(ctx context.Context, key string, pointer interface{}) error
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Del(ctx context.Context, key string) error
}

type redisc struct {
	id       string
	uri      string
	password string
	logger   sctx.Logger
	rdb      *redis.Client
}

func NewRedisComp(id string) *redisc {
	return &redisc{id: id}
}

func (r *redisc) ID() string {
	return r.id
}

func (r *redisc) InitFlags() {
	flag.StringVar(&r.uri, "redis-uri", "localhost:6379", "Redis URI")
	flag.StringVar(&r.password, "redis-password", "", "Redis password")
}

func (r *redisc) Activate(sctx sctx.ServiceContext) error {
	r.logger = sctx.Logger("redis")

	rdb := redis.NewClient(&redis.Options{
		Addr:     r.uri,
		Password: "", // no password
		DB:       0,  // use default DB
		Protocol: 2,
	})

	r.rdb = rdb

	return nil
}

func (r *redisc) Stop() error {
	return nil
}

func (r *redisc) Get(ctx context.Context, key string, pointer interface{}) error {
	cmd := r.rdb.Get(ctx, key)
	if err := cmd.Err(); err != nil {
		return errors.Wrap(err, "failed to get value from Redis")
	}

	dataBytes, err := cmd.Result()
	if err != nil {
		return errors.Wrap(err, "failed to get value from Redis")
	}

	err = json.Unmarshal([]byte(dataBytes), pointer)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal value from JSON")
	}

	return nil
}

func (r *redisc) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	dataBytes, err := json.Marshal(value)
	if err != nil {
		return errors.Wrap(err, "failed to marshal value to JSON")
	}

	return r.rdb.Set(ctx, key, dataBytes, expiration).Err()
}

func (r *redisc) Del(ctx context.Context, key string) error {
	return r.rdb.Del(ctx, key).Err()
}