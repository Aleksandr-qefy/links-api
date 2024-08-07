package main

import (
	api "github.com/Aleksandr-qefy/links-api"
	"github.com/Aleksandr-qefy/links-api/internal/handler"
	"github.com/Aleksandr-qefy/links-api/internal/repository"
	"github.com/Aleksandr-qefy/links-api/internal/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srvr := &api.Server{}
	if err := srvr.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
