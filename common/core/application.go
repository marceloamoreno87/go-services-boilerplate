package core

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type Application struct{}

func (a Application) Run(router IRouter) (err error) {
	fmt.Println("Starting application...")

	dsnPostgres := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&search_path=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_NAME"),
		os.Getenv("POSTGRES_SSL_MODE"),
		os.Getenv("POSTGRES_SCHEMA"),
	)

	pgInstance := NewPostgres(dsnPostgres)

	dsnRedis := fmt.Sprintf(
		"redis://:%s@%s:%s",
		os.Getenv("REDIS_PASSWORD"),
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
	)

	_ = NewRedis(dsnRedis)

	err = pgInstance.RunMigrate(dsnPostgres)
	if err != nil {
		fmt.Printf("Error migrating database: %v\n", err)
		return
	}
	fmt.Println("Migrated database...")

	NewObservability()

	NewJWT(os.Getenv("JWT_SECRET"))

	mux := NewMux()
	muxWithRoutes := router.GetRoutes(mux)

	srv := &http.Server{
		Addr:              ":" + os.Getenv("PORT"),
		Handler:           muxWithRoutes,
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	go NewMetrics()

	fmt.Println("Starting web application on port", os.Getenv("PORT"), "...")
	fmt.Printf("http://localhost:%s\n", os.Getenv("PORT"))
	fmt.Println("Press CTRL+C to stop the application...")

	if err = srv.ListenAndServe(); err != nil {
		fmt.Printf("Error starting web application: %v\n", err)
		return
	}
	return
}

//teste
