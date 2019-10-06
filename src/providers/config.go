package providers

import "os"

type Config struct {
	DbUser     string
	DbPassword string
	DbName     string
	ListenAddr string
}

func NewConfig() (*Config, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASWORD")
	dbName := os.Getenv("DB_NAME")
	addr := os.Getenv("LISTEN_ADDR")

	if len(dbUser) == 0 {
		dbUser = "postgres"
	}
	if len(dbPassword) == 0 {
		dbPassword = "password"
	}
	if len(dbName) == 0 {
		dbName = "go_chat"
	}
	if len(addr) == 0 {
		addr = ":9090"
	}

	return &Config{
		DbUser:     dbUser,
		DbPassword: dbPassword,
		DbName:     dbName,
		ListenAddr: addr,
	}, nil
}
