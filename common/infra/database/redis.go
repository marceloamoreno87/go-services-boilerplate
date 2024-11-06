package database

import "github.com/redis/go-redis/v9"


type Redis struct{}


func (d Redis) InitDB(dsn string) (conn *redis.Client) {
	url := dsn
	opts, err := redis.ParseURL(url)
	if err != nil {
		// Retorna o erro se a conexão falhar
		return
	}

	return redis.NewClient(opts)
}
