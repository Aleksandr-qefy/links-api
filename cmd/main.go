package main

import (
	api "github.com/Aleksandr-qefy/links-api"
	"github.com/Aleksandr-qefy/links-api/internal/handler"
	"github.com/Aleksandr-qefy/links-api/internal/repository"
	"github.com/Aleksandr-qefy/links-api/internal/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srvr := &api.Server{}
	if err := srvr.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
