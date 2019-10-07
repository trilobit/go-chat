package services

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/trilobit/go-chat/src/providers"
	"github.com/trilobit/go-chat/src/repositories"
)

func TestAccountService_Register_OK(t *testing.T) {
	t.Parallel()

	service := initService(t)
	user, err := service.Register(repositories.TestEmail, repositories.TestPassword)
	require.Nil(t, err)
	require.NotNil(t, user)
	require.Equal(t, user.Email, repositories.TestEmail)
}

func TestAccountService_Register_Fail_Email(t *testing.T) {
	t.Parallel()

	service := initService(t)

	_, err := service.Register("123123", repositories.TestPassword)
	require.NotNil(t, err)
	require.Equal(t, err, ErrMalformedEMail)
}

func TestAccountService_Register_Fail_Password(t *testing.T) {
	t.Parallel()

	service := initService(t)

	_, err := service.Register(repositories.TestEmail, "111")
	require.NotNil(t, err)
	require.Equal(t, err, ErrTooShorPassword)
}

func TestAccountService_Authorize_OK(t *testing.T) {
	t.Parallel()

	service := initService(t)

	user, err := service.Authorize(repositories.TestEmail, repositories.TestPassword)
	require.Nil(t, err)
	require.NotNil(t, user)
	require.Equal(t, user.Email, repositories.TestEmail)
}

func TestAccountService_Authorize_Fail_Email(t *testing.T) {
	t.Parallel()

	service := initService(t)

	var err error

	_, err = service.Authorize("123123", repositories.TestPassword)
	require.NotNil(t, err)
	require.Equal(t, err, ErrUnauthorized)
}

func TestAccountService_Authorize_Fail_Password(t *testing.T) {
	t.Parallel()

	service := initService(t)

	var err error

	_, err = service.Authorize(repositories.TestEmail, "111")
	require.NotNil(t, err)
	require.Equal(t, err, ErrUnauthorized)
}

func initService(t *testing.T) Account {
	logger, err := providers.NewLogger()
	if err != nil {
		t.Errorf("TestAccountService_Register failed")
	}

	repo := repositories.NewUserRepositoryMock()
	options := AccountOptions{
		Logger: logger,
		Repo:   repo,
	}

	return NewAccount(options)
}
