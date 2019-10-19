package services

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/trilobit/go-chat/src/models"
	"github.com/trilobit/go-chat/src/providers"
	mock_providers "github.com/trilobit/go-chat/src/providers/mocks"
	mock_repositories "github.com/trilobit/go-chat/src/repositories/mocks"
)

const (
	TestEmail    = "admin@example.com"
	TestPassword = "123123"
	passwordHash = "$2a$14$h9Ta1.4t/LDpYYV12xcGW.jiJ2S9bRVR0IzqGDkSzg5cFfgOenNha"
	createdTime  = "2019-09-25 00:00:00"
	updatedTime  = "2019-09-25 03:14:15"
	testToken    = "test_token"
)

func TestAccountService(t *testing.T) {
	t.Parallel()

	service := initService(t)

	t.Run("AccountService", func(t *testing.T) {
		t.Run("Register", func(t *testing.T) {
			t.Run("Success", func(t *testing.T) {
				user, err := service.Register(TestEmail, TestPassword)
				require.Nil(t, err)
				require.NotNil(t, user)
				require.Equal(t, user.Email, TestEmail)
			})

			t.Run("ErrMalformedEMail", func(t *testing.T) {
				_, err := service.Register("123123", TestPassword)
				require.Error(t, err)
				require.Equal(t, err, ErrMalformedEMail)
			})

			t.Run("ErrMalformedEMail", func(t *testing.T) {
				_, err := service.Register(TestEmail, "111")
				require.Error(t, err)
				require.Equal(t, err, ErrTooShorPassword)
			})
		})

		t.Run("Authorize", func(t *testing.T) {
			t.Run("Success", func(t *testing.T) {
				user, err := service.Authorize(TestEmail, TestPassword)
				require.Nil(t, err)
				require.NotNil(t, user)
				require.Equal(t, user.Email, TestEmail)
			})
			t.Run("Incorrect password", func(t *testing.T) {
				_, err := service.Authorize(TestEmail, "wrongpass")
				require.NotNil(t, err)
			})
			t.Run("Incorrect email", func(t *testing.T) {
				_, err := service.Authorize("wrongmail", TestPassword)
				require.NotNil(t, err)
			})
		})
	})

}

func initService(t *testing.T) Account {
	logger, err := providers.NewLogger()
	if err != nil {
		t.Errorf("TestAccountService_Register failed")
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUser := mock_repositories.NewMockUser(ctrl)
	mockCrypt := mock_providers.NewMockCrypt(ctrl)

	created := toTime(createdTime)

	testUser := &models.User{
		ID:        1,
		Email:     TestEmail,
		Pswd:      passwordHash,
		CreatedAt: created,
		UpdatedAt: created,
	}

	mockUser.EXPECT().Create(TestEmail, passwordHash).Return(testUser, nil).AnyTimes()
	mockUser.EXPECT().FindByEmail(TestEmail).Return(testUser, nil).AnyTimes()
	mockUser.EXPECT().FindByEmail("wrongmail").Return(nil, ErrUnauthorized).AnyTimes()
	mockUser.EXPECT().UpdateToken(testUser, testToken).Return(nil).AnyTimes()

	mockCrypt.EXPECT().Hash(TestPassword).Return(passwordHash, nil).AnyTimes()
	mockCrypt.EXPECT().Compare(TestPassword, passwordHash).Return(true).AnyTimes()

	mockCrypt.EXPECT().Compare("wrongpass", passwordHash).Return(false).AnyTimes()

	mockCrypt.EXPECT().Hash(TestEmail).Return(testToken, nil).AnyTimes()

	// repo := repositories.NewUserRepositoryMock()
	options := AccountOptions{
		Logger: logger,
		Repo:   mockUser,
		Crypt:  mockCrypt,
	}

	return NewAccount(options)
}

func toTime(dateStr string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", dateStr)
	return t
}
