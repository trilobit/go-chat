package providers

import (
	"golang.org/x/crypto/bcrypt"
)

type (
	// Crypt is interface for implementing hash and validate functions
	Crypt interface {
		Hash(str string) (string, error)
		Compare(str, hash string) bool
	}

	// CryptByBCrypt uses bcrypt for hashing
	CryptByBCrypt struct {
		complexity int
	}

	// CryptByBCryptOptions struct {
	// 	fx.In

	// 	Complexity int
	// }
)

// NewCryptByBCrypt creates a new hasher with bcrypt inside
func NewCryptByBCrypt(config *Config) Crypt {
	return &CryptByBCrypt{
		complexity: config.Complexity,
	}
}

// Hash method creates hash given string using bcrypt
func (b *CryptByBCrypt) Hash(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), b.complexity)
	return string(bytes), err
}

// Compare checks string and hash using bcrypt
func (b *CryptByBCrypt) Compare(str, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}
