package datastore

import (
	"github.com/Zzocker/bl-utils/config"
	"github.com/Zzocker/bl-utils/pkg/errors"
)

// DatastoreConfig represents datastore config struct
// config for connecting datastore
type DatastoreConfig struct {
	Code     string
	URL      string
	Username string
	Password string
	DBName   string
}

var (
	dbMap = map[string]datastoreBuilder{
		config.MYSQL: &sqlDSBuilder{},
	}
)

type datastoreBuilder interface {
	Build(cfg DatastoreConfig) (interface{}, *errors.Er)
}

// FromFactory returns datastore builder
func FromFactory(code string) datastoreBuilder {
	return dbMap[code]
}
