package infra

import (
	"database/sql"
)

var DB *sql.DB

func InitDatabase(config *Config) {
	var err error
	DB, err = sql.Open(config.DbDriver, config.DbUrl)

	if err != nil {
		panic(err)
	}
}
