package src

import (
	"github.com/trilobit/go-chat/src/api"
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
		),

		fx.Invoke(
			api.NewApi,
			ws.New,
		),
	)

	app.Run()
}
