package datastore

import (
	"testing"

	"github.com/Zzocker/bl-utils/config"
	"github.com/gomodule/redigo/redis"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

func TestRedis(t *testing.T) {
	cfg := DatastoreConfig{
		Code: config.REDIS,
		URL:  "localhost:6379",
	}
	rawPool, err := FromFactory(cfg.Code).Build(cfg)
	assert.Nil(t, err)
	assert.NotNil(t, rawPool)
	_, ok := rawPool.(*redis.Pool)
	assert.True(t, ok)

}

func TestRabbitMQ(t *testing.T) {
	cfg := DatastoreConfig{
		Code:     config.RABBITMQ,
		URL:      "localhost:5672",
		Username: "user",
		Password: "passwor",
	}
	rawCh, err := FromFactory(cfg.Code).Build(cfg)
	assert.Nil(t, err)
	_, ok := rawCh.(*amqp.Channel)
	assert.True(t, ok)
}
