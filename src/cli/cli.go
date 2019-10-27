package cli

import (
	"database/sql"
	"fmt"
	"github.com/pressly/goose"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

type (
	Options struct {
		fx.In

		Logger *zap.SugaredLogger
		Config *viper.Viper
	}
)

func Migrate(cmd string, opts Options) error {
	conn := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		opts.Config.GetString("db.user"),
		opts.Config.GetString("db.password"),
		opts.Config.GetString("db.addr"),
		opts.Config.GetString("db.db"),
	)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return err
	}

	if err := goose.Run(cmd, db, opts.Config.GetString("migrations.dir")); err != nil {
		return err
	}

	return nil
}
