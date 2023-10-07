package envconfig

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	TOKEN            string
	ServerPort       string
}

func NewEnv() (*Env, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &Env{
		TOKEN:            os.Getenv("TOKEN"),
		ServerPort:       os.Getenv("PORT"),
	}, nil
}
