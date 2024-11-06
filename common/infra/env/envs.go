package env

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Envs struct{}

func (e Envs) Load() (err error) {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	_ = filepath.Join(dir, ".env")

	return godotenv.Load()
}

func (e Envs) Get(key string) string {
	return os.Getenv(key)
}
