package datastore

import (
	"github.com/Zzocker/bl-utils/pkg/errors"
	"github.com/gomodule/redigo/redis"
)

type redisBuilder struct{}

func (r *redisBuilder) Build(cfg DatastoreConfig) (interface{}, *errors.Er) {
	// creating pool
	pool := redis.Pool{
		MaxIdle:   80,
		MaxActive: 0,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp", cfg.URL,
				redis.DialUsername(cfg.Username),
				redis.DialPassword(cfg.Password),
			)
		},
	}
	// initial check
	conn := pool.Get()
	defer conn.Close()
	if _, err := conn.Do("PING"); err != nil {
		return nil, errors.New(errors.INTERNAL, err)
	}
	return &pool, nil
}
