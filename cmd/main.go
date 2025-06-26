package main

import (
	"os"

	"github.com/TimmyTurner98/sharing"
	"github.com/TimmyTurner98/sharing/pkg/handler"
	"github.com/TimmyTurner98/sharing/pkg/repository"
	"github.com/TimmyTurner98/sharing/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load("env/app.env"); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	router := handlers.InitRoutes()

	s := &sharing.Server{}

	if err := s.Run("8080", router); err != nil {
		logrus.Fatalf("Ошибка запуска сервера: %s", err.Error())
	}

}
