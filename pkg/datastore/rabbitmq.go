package datastore

import (
	"fmt"

	"github.com/Zzocker/bl-utils/pkg/errors"
	"github.com/streadway/amqp"
)

type rabbitMQBuilder struct {
}

func (r *rabbitMQBuilder) Build(cfg DatastoreConfig) (interface{}, *errors.Er) {
	url := fmt.Sprintf("amqp://%s:%s@%s", cfg.Username, cfg.Password, cfg.URL)
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, errors.New(errors.INTERNAL, err)
	}
	if conn.IsClosed() {
		return nil, errors.NewMsgln(errors.INTERNAL, "rabbitmq connection is closed")
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, errors.New(errors.INTERNAL, err)
	}
	return ch, nil
}
