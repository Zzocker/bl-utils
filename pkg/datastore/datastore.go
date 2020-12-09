package datastore

import (
	"github.com/Zzocker/bl-utils/config"
	"github.com/Zzocker/bl-utils/pkg/errors"
)

var (
	dbMap = map[string]datastoreBuilder{
		config.MYSQL: &sqlDSBuilder{},
	}
)

type datastoreBuilder interface {
	Build(cfg *config.DatastoreConfig) (interface{}, *errors.Er)
}

// FromFactory returns datastore builder
func FromFactory(code string) datastoreBuilder {
	return dbMap[code]
}
