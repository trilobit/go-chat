package api

import (
	"context"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"github.com/trilobit/go-chat/src/repositories"
	"github.com/trilobit/go-chat/src/services"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type (
	Api struct {
		logger         *zap.SugaredLogger
		config         *viper.Viper
		accountService services.Account
		userRepo       repositories.User
	}

	Options struct {
		fx.In

		Logger         *zap.SugaredLogger
		Config         *viper.Viper
		AccountService services.Account
		UserRepo       repositories.User
		Lc             fx.Lifecycle
	}

	UserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func NewApi(options Options) {
	a := &Api{
		logger:         options.Logger.Named("api"),
		config:         options.Config,
		accountService: options.AccountService,
		userRepo:       options.UserRepo,
	}

	e := echo.New()
	e.HidePort = true
	e.HideBanner = true

	e.GET("/", a.home)
	e.POST("/sign-up", a.register)
	e.POST("/sign-in", a.login)
	e.GET("/profile", a.profile, a.authMiddleware)

	// Start & Stop server
	options.Lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			options.Logger.Infof("starting server at: %s", options.Config.GetString("api.addr"))

			go e.Start(options.Config.GetString("api.addr"))

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}
