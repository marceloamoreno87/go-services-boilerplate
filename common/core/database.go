package core

import (
	"database/sql"
	"log"

	"sendzap-checkout/common/infra/database"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var POSTGRESCONN *sql.DB

var MONGOCONN *mongo.Client

var REDISCONN *redis.Client

func NewPostgres(dsn string) (instance *database.Postgres) {
	instance = &database.Postgres{}
	conn, err := instance.InitDB(dsn)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	POSTGRESCONN = conn
	return instance
}

func NewMongo(dsn string) (instance *database.Mongo) {
	instance = &database.Mongo{}
	conn, err := instance.InitDB(dsn)
	if err != nil {

		log.Fatalf("Failed to initialize database: %v", err)
	}

	MONGOCONN = conn
	return instance
}

func NewRedis(dsn string) (instance *database.Redis) {
	instance = &database.Redis{}
	conn := instance.InitDB(dsn)

	REDISCONN = conn
	return instance
}
