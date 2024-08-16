package core_cache

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"strings"
	"time"
)

type RedisCacheConfig struct {
	Address  string `json:"address" yaml:"address"`
	Database int    `json:"database" yaml:"database"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

type RedisCacheHandler struct {
	config *RedisCacheConfig
}

func NewRedisCacheHandler(config *RedisCacheConfig) (*RedisCacheHandler, error) {
	if len(config.Address) < 1 {
		return nil, errors.New("redis配置地址不能为空")
	}

	cacheHandler := &RedisCacheHandler{
		config,
	}

	return cacheHandler, nil
}

func (handler *RedisCacheHandler) newClient() redis.UniversalClient {
	//判断单机、集群，如果地址中有逗号(,)分割，则为集群
	address := strings.Split(handler.config.Address, ",")

	return redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    address,
		DB:       handler.config.Database,
		Username: handler.config.Username,
		Password: handler.config.Password,
	})
}

func (handler *RedisCacheHandler) Keys(prefix string) ([]string, error) {
	client := handler.newClient()

	ctx := context.Background()

	res, _, err := client.Scan(ctx, 0, prefix+"*", 0).Result()

	return res, err
}

func (handler *RedisCacheHandler) Has(key string) (bool, error) {
	client := handler.newClient()

	ctx := context.Background()

	res, err := client.Exists(ctx, key).Result()

	if err != nil {
		return false, err
	}

	return res > 0, err
}

func (handler *RedisCacheHandler) Get(key string, out any) (bool, error) {
	client := handler.newClient()

	ctx := context.Background()

	res, err := client.Get(ctx, key).Result()

	if err != nil {
		return false, err
	}

	if len(res) < 1 {
		return false, nil
	}

	return true, json.Unmarshal([]byte(res), &out)
}

//
//func (handler *RedisCacheHandler) Add(key string, value any, expire int64) (bool, error) {
//	client := handler.newClient()
//
//	ctx := context.Background()
//
//	res, err := client.SetNX(ctx, key, value, time.Duration(expire)*time.Second).Result()
//
//	return res, err
//}

func (handler *RedisCacheHandler) Set(key string, value any, expire int64) error {
	client := handler.newClient()

	ctx := context.Background()

	_, err := client.Set(ctx, key, value, time.Duration(expire)*time.Second).Result()

	return err
}

func (handler *RedisCacheHandler) Del(key string) (bool, error) {
	client := handler.newClient()

	ctx := context.Background()

	res, err := client.Del(ctx, key).Result()

	return res > 0, err
}
