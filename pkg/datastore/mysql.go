package datastore

import (
	"database/sql"
	"fmt"

	"github.com/Zzocker/bl-utils/pkg/errors"
	_ "github.com/go-sql-driver/mysql"
)

type sqlDSBuilder struct{}

func (s *sqlDSBuilder) Build(cfg DatastoreConfig) (interface{}, *errors.Er) {
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.Username, cfg.Password, cfg.URL, cfg.DBName)
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db, nil
}
