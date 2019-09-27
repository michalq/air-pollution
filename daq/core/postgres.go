package core

import "github.com/go-pg/pg/v9"

type PostgresConfig struct {
	Address  string
	User     string
	Password string
	Database string
}

func CreatePostgresConnection(conf PostgresConfig) *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     conf.Address,
		User:     conf.User,
		Password: conf.Password,
		Database: conf.Database,
	})
}
