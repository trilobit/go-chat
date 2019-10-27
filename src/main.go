package src

import (
	"context"
	"github.com/trilobit/go-chat/src/api"
	"github.com/trilobit/go-chat/src/cli"
	"github.com/trilobit/go-chat/src/providers"
	"github.com/trilobit/go-chat/src/repositories"
	"github.com/trilobit/go-chat/src/services"
	"github.com/trilobit/go-chat/src/ws"
	"go.uber.org/fx"
)

func Run() {
	app := fx.New(
		fx.Provide(
			providers.NewConfig,
			providers.NewDB,
			providers.NewLogger,
			providers.NewCryptByBCrypt,
			services.NewAccount,
			repositories.NewUser,
			repositories.NewHistory,
		),

		fx.Invoke(
			api.NewApi,
			ws.New,
		),
	)

	app.Run()
}

func Migrate(cmd string) {
	app := fx.New(
		fx.Provide(
			providers.NewConfig,
			providers.NewLogger,
		),
		fx.Invoke(func(opts cli.Options) {

			if err := cli.Migrate(cmd, opts); err != nil {
				opts.Logger.Errorf("error migrating: %v", err)
			}

			opts.Logger.Infof("migration successful")
		}),
	)

	if err := app.Start(context.Background()); err != nil {
		panic(err)
	}
}
