package ws

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"github.com/trilobit/go-chat/src/models"
	"github.com/trilobit/go-chat/src/repositories"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
	"sync"
)

type (
	User struct {
		Model *models.User
		Conn  *websocket.Conn
	}

	Websocket struct {
		logger      *zap.SugaredLogger
		config      *viper.Viper
		userRepo    repositories.User
		historyRepo repositories.History
		history     []models.Message
		hub         map[string]User
		hmu         sync.RWMutex
	}

	Options struct {
		fx.In

		Logger      *zap.SugaredLogger
		Config      *viper.Viper
		Lc          fx.Lifecycle
		UserRepo    repositories.User
		HistoryRepo repositories.History
	}
)

func New(options Options) {
	// Load history of public messages
	history, err := options.HistoryRepo.Load()
	if err != nil {
		options.Logger.Errorf("error on load history from db: %v", err)
	}

	socket := &Websocket{
		logger:      options.Logger,
		config:      options.Config,
		hub:         make(map[string]User),
		userRepo:    options.UserRepo,
		historyRepo: options.HistoryRepo,
		history:     history,
	}

	options.Lc.Append(fx.Hook{
		OnStart: func(context context.Context) error {
			addr := options.Config.GetString("ws.addr")
			options.Logger.Infof("starting websocket server at: %s", addr)
			go func() {
				if err := http.ListenAndServe(addr, socket); err != nil {
					options.Logger.Errorf("error starting websocket server: %v", err)
				}
			}()
			return nil
		},
		OnStop: nil,
	})
}
