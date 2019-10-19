package providers

import "os"

// Config contains settings for application
type Config struct {
	DbUser     string
	DbPassword string
	DbName     string
	ListenAddr string
	Complexity int
}

const (
	defaultDbUser     = "postgres"
	defaultDbPassword = "password"
	defaultDbName     = "go_chat"
	defaultAddr       = ":9090"
	defaultComplexity = 10
)

// NewConfig creates instance of config
func NewConfig() (*Config, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASWORD")
	dbName := os.Getenv("DB_NAME")
	addr := os.Getenv("LISTEN_ADDR")

	if len(dbUser) == 0 {
		dbUser = defaultDbUser
	}
	if len(dbPassword) == 0 {
		dbPassword = defaultDbPassword
	}
	if len(dbName) == 0 {
		dbName = defaultDbName
	}
	if len(addr) == 0 {
		addr = defaultAddr
	}

	return &Config{
		DbUser:     dbUser,
		DbPassword: dbPassword,
		DbName:     dbName,
		ListenAddr: addr,
		Complexity: defaultComplexity,
	}, nil
}
