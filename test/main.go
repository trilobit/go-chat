package test

import (
	"github.com/trilobit/go-chat/src/providers"
	"github.com/trilobit/go-chat/src/services"
	"github.com/trilobit/go-chat/test/repositories"
)

func InitService() (services.Account, error) {
	logger, err := providers.NewLogger()
	if err != nil {
		return nil, err
	}

	repo := repositories.NewUserRepositoryMock()
	options := services.AccountOptions{
		Logger: logger,
		Repo:   repo,
	}

	return services.NewAccount(options), nil
}
