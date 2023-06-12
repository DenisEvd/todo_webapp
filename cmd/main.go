package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"todo_webapp"
	"todo_webapp/pkg/handler"
	"todo_webapp/pkg/repository"
	"todo_webapp/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("error with config file")
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("error with env variables")
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("error with db authorization")
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo_webapp.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal("error while running server")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
