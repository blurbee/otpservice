package store

import (
	"context"
	"time"

	"caluxor.com/api"
	"caluxor.com/util"
	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client  *redis.Client
	options *redis.Options
}

type redis_creds struct {
	Url string `envconfig:"REDIS_URL"`
}

func (r *RedisStore) Init(cfg *util.Config) (err api.StatusCode) {
	url := cfg.GetSecret("REDIS_URL")
	if url != "" {
		util.Error("Unable to load redis config")
		return api.CONFIG_ERROR
	}

	var er error
	r.options, er = redis.ParseURL(url)
	if er != nil {
		util.Error("Unable to connect to redis:", er)
		return api.CONN_FAILED
	}

	r.client = redis.NewClient(r.options)
	if r.client != nil {
		util.Info("Connected to redis")
	} else {
		util.Error("Unable to connect to redis")
		err = api.CONN_FAILED
		return
	}

	// Perform basic diagnostic to check if the connection is working
	// Expected result > ping: PONG
	// If Redis is not running, error case is taken instead
	status, er := r.client.Ping(context.Background()).Result()
	if er != nil {
		util.Error("Redis connection was refused:", status, er)
		err = api.CONN_FAILED
		return
	}
	return api.OK
}

func (r *RedisStore) StoreKey(sessionid string, otpcode string,
	expires time.Duration) (err api.StatusCode) {
	_, error := r.client.Set(context.Background(), sessionid, otpcode, expires).Result()
	if error != nil {
		util.Error("Setting key failed:", error)
		err = api.STORE_ERROR
		return
	}
	return api.OK
}

func (r *RedisStore) RetrieveKey(sessionid string) (otpcode string, err api.StatusCode) {
	otpcode, error := r.client.Get(context.Background(), sessionid).Result()
	if error != nil {
		util.Error("Getting key failed:", error)
		err = api.STORE_ERROR
		return
	}

	err = api.OK
	return
}

func (r *RedisStore) Close() (err api.StatusCode) {
	er := r.client.Close()
	if er != nil {
		util.Error("Error while closing redis connection:", er)
		return api.CONN_FAILED
	}
	return api.OK
}
