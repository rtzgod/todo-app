package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	todo "github.com/rtzgod/todo-app"
	"github.com/rtzgod/todo-app/internal/handler"
	"github.com/rtzgod/todo-app/internal/repository"
	"github.com/rtzgod/todo-app/internal/service"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing config file, %s", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env, %s", err)
	}

	db, err := repository.NewPostgres(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("Error initializing DB, %s", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
