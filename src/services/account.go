package services

import (
	"errors"
	"strings"
	"time"

	"github.com/trilobit/go-chat/src/models"
	"github.com/trilobit/go-chat/src/repositories"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type (
	Account interface {
		Register(email, password string) (*models.User, error)
		Authorize(email, password string) (*models.User, error)
	}

	accountService struct {
		repo   repositories.User
		logger *zap.SugaredLogger
	}

	AccountOptions struct {
		fx.In

		Logger *zap.SugaredLogger
		Repo   repositories.User
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

func NewAccount(options AccountOptions) Account {
	return &accountService{
		repo:   options.Repo,
		logger: options.Logger.Named("account_service"),
	}
}

func (as *accountService) Register(email, password string) (*models.User, error) {
	if !strings.Contains(email, "@") {
		return nil, ErrMalformedEMail
	}

	if len(password) < 6 {
		return nil, ErrTooShorPassword
	}

	pswdHash, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	return as.repo.Create(email, pswdHash)
}

func (as *accountService) Authorize(email, password string) (*models.User, error) {
	user, err := as.repo.FindByEmail(email)
	if err != nil {
		// as.logger.Errorf("Error in search user: %v", err)
		return nil, ErrUnauthorized
	}
	if !checkPassword(password, user.Pswd) {
		return nil, ErrUnauthorized
	}

	token, err := generateToken()
	if err != nil {
		return nil, err
	}

	if err := as.repo.UpdateToken(user, token); err != nil {
		return nil, err
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(time.Now().String()), 4)
	return string(bytes), err
}
