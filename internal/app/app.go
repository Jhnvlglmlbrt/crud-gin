package app

import (
	"fmt"
	"log"

	"github.com/Jhnvlglmlbrt/crud-gin/config"
	"github.com/Jhnvlglmlbrt/crud-gin/internal/controller"
	"github.com/Jhnvlglmlbrt/crud-gin/internal/model"
	"github.com/Jhnvlglmlbrt/crud-gin/internal/repository"
	"github.com/Jhnvlglmlbrt/crud-gin/internal/service"
	server "github.com/Jhnvlglmlbrt/crud-gin/package/http"
	"github.com/Jhnvlglmlbrt/crud-gin/package/postgres"
	"github.com/go-playground/validator/v10"
)

func Run(cfg *config.Config) {

	postgresConnection, err := postgres.Connect(&cfg.PG)
	if err != nil {
		fmt.Printf("Error at Postgre connection: %v", err)
	}

	validate := validator.New()

	postgresConnection.Table("tags").AutoMigrate(&model.Tags{})

	// Repostiory
	tagsRepository := repository.NewTagsRepositoryImpl(postgresConnection)

	// Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	// Controller
	tagsController := controller.NewTagsController(tagsService)

	httpServer := server.NewServer(tagsController)

	serverStartingError := httpServer.Start()
	if serverStartingError != nil {
		log.Fatalf("Error at server starting: %v", serverStartingError)
	}
}
