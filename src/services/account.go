package services

import (
	"errors"
	"strings"

	"github.com/trilobit/go-chat/src/models"
	"github.com/trilobit/go-chat/src/providers"
	"github.com/trilobit/go-chat/src/repositories"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type (
	// Account is interface for implementing registration and authorization functions
	Account interface {
		Register(email, password string) (*models.User, error)
		Authorize(email, password string) (*models.User, error)
	}

	accountService struct {
		repo   repositories.User
		logger *zap.SugaredLogger
		crypt  providers.Crypt
	}

	// AccountOptions is config for creating AccountService
	AccountOptions struct {
		fx.In

		Logger *zap.SugaredLogger
		Repo   repositories.User
		Crypt  providers.Crypt
	}
)

var (
	// ErrMalformedEMail is incorrect email address
	ErrMalformedEMail = errors.New("malformed email")

	// ErrTooShorPassword password is too short
	ErrTooShorPassword = errors.New("password must be 6 symbols at least")

	// ErrUnauthorized user is unauthorized
	ErrUnauthorized = errors.New("unauthorized")
)

// NewAccount creates instance of accountService
func NewAccount(options AccountOptions) Account {
	return &accountService{
		repo:   options.Repo,
		logger: options.Logger.Named("account_service"),
		crypt:  options.Crypt,
	}
}

func (as *accountService) Register(email, password string) (*models.User, error) {
	if !strings.Contains(email, "@") {
		return nil, ErrMalformedEMail
	}

	if len(password) < 6 {
		return nil, ErrTooShorPassword
	}

	pswdHash, err := as.crypt.Hash(password)
	if err != nil {
		return nil, err
	}

	return as.repo.Create(email, pswdHash)
}

func (as *accountService) Authorize(email, password string) (*models.User, error) {
	user, err := as.repo.FindByEmail(email)
	if err != nil {
		as.logger.Errorf("error 1: %v", err)
		return nil, ErrUnauthorized
	}
	if !as.crypt.Compare(password, user.Pswd) {
		as.logger.Error("error 2:")
		return nil, ErrUnauthorized
	}

	token, err := as.crypt.Hash(email)
	if err != nil {
		as.logger.Errorf("error 3: %v", err)
		return nil, err
	}

	if err := as.repo.UpdateToken(user, token); err != nil {
		as.logger.Errorf("error 4: %v", err)
		return nil, err
	}

	return user, nil
}
