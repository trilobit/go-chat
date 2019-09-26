package src

import (
	"github.com/trilobit/go-chat/src/api"
	"github.com/trilobit/go-chat/src/providers"
	"github.com/trilobit/go-chat/src/repositories"
	"github.com/trilobit/go-chat/src/services"
	"go.uber.org/fx"
)

func Run() {
	app := fx.New(
		fx.Provide(
			providers.NewConfig,
			providers.NewDB,
			providers.NewLogger,
			services.NewAccount,
			repositories.NewUser,
		),

		fx.Invoke(
			api.NewApi,
		),
	)

	app.Run()
}
