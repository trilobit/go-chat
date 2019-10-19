package api

import (
	"context"

	"github.com/labstack/echo"
	"github.com/trilobit/go-chat/src/providers"
	"github.com/trilobit/go-chat/src/repositories"
	"github.com/trilobit/go-chat/src/services"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type (
	Api struct {
		logger         *zap.SugaredLogger
		config         *providers.Config
		accountService services.Account
		userRepo       repositories.User
	}

	ApiOptions struct {
		fx.In

		Logger         *zap.SugaredLogger
		Config         *providers.Config
		AccountService services.Account
		UserRepo       repositories.User
		Lc             fx.Lifecycle
	}

	UserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func NewApi(options ApiOptions) {
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
			options.Logger.Infof("starting server at: %s", options.Config.ListenAddr)

			go e.Start(options.Config.ListenAddr) // todo why?

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}
