package providers

import (
	"github.com/go-pg/pg"
	"github.com/spf13/viper"
)

func NewDB(config *viper.Viper) (*pg.DB, error) {
	conn := pg.Connect(&pg.Options{
		Addr:     config.GetString("db.addr"),
		User:     config.GetString("db.user"),
		Password: config.GetString("db.password"),
		Database: config.GetString("db.db"),
	})

	if _, err := conn.ExecOne("SELECT 1"); err != nil {
		return nil, err
	}

	return conn, nil
}
