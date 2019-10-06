package repositories

import (
	"time"

	"github.com/trilobit/go-chat/src/models"
)

const (
	TestEmail    = "admin@example.com"
	TestPassword = "123123"
	passwordHash = "$2a$14$h9Ta1.4t/LDpYYV12xcGW.jiJ2S9bRVR0IzqGDkSzg5cFfgOenNha"
	createdTime  = "2019-09-25 00:00:00"
	updatedTime  = "2019-09-25 03:14:15"
	testToken    = "$2a$04$NQrCMTOg8GVKqmMLHPvK/OicDOlF7s7AVK2BDFb24peQBvPSK3jIG"
)

type (
	mockUserRepository struct{}
)

func NewUserRepositoryMock() *mockUserRepository {
	return &mockUserRepository{}
}

func (u *mockUserRepository) Create(email, pswdHash string) (*models.User, error) {
	created := toTime(createdTime)
	user := models.User{
		ID:        1,
		Email:     email,
		Pswd:      pswdHash,
		CreatedAt: created,
		UpdatedAt: created,
	}
	return &user, nil
}

func (u *mockUserRepository) FindByEmail(email string) (*models.User, error) {
	created := toTime(createdTime)
	updated := toTime(updatedTime)

	user := models.User{
		ID:        1,
		Email:     TestEmail,
		Pswd:      passwordHash,
		CreatedAt: created,
		UpdatedAt: updated,
	}
	return &user, nil
}

func (u *mockUserRepository) FindByToken(token string) (*models.User, error) {
	created := toTime(createdTime)
	updated := toTime(updatedTime)
	user := models.User{
		ID:        1,
		Email:     TestEmail,
		Pswd:      passwordHash,
		CreatedAt: created,
		UpdatedAt: updated,
	}
	return &user, nil
}

func (u *mockUserRepository) UpdateToken(user *models.User, token string) error {
	user.Token = token
	return nil
}

func toTime(dateStr string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", dateStr)
	return t
}
