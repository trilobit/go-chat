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
	"time"
)

type (
	Message struct {
		From     string    `json:"from"`
		To       string    `json:"to"`
		Text     string    `json:"text"`
		DateTime time.Time `json:"date_time"`
	}

	User struct {
		Model *models.User
		Conn  *websocket.Conn
	}

	Websocket struct {
		logger   *zap.SugaredLogger
		config   *viper.Viper
		userRepo repositories.User
		history  []Message
		hub      map[string]User
		hmu      sync.RWMutex
	}

	Options struct {
		fx.In

		Logger   *zap.SugaredLogger
		Config   *viper.Viper
		Lc       fx.Lifecycle
		UserRepo repositories.User
	}
)

func New(options Options) {
	socket := &Websocket{
		logger:   options.Logger,
		config:   options.Config,
		hub:      make(map[string]User),
		userRepo: options.UserRepo,
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
