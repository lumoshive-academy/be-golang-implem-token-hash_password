package main

import (
	"fmt"
	"log"
	"net/http"
	"session-22/database"
	"session-22/handler"
	"session-22/repository"
	"session-22/router"
	"session-22/service"
	"session-22/utils"
)

func main() {
	config, err := utils.ReadConfiguration()
	if err != nil {
		log.Fatal("error file configration")
	}
	fmt.Println(config)
	db, err := database.InitDB(config.DB)
	if err != nil {
		panic(err)
	}

	logger, err := utils.InitLogger(config.PathLogging, config.Debug)

	repo := repository.NewRepository(db, logger)
	service := service.NewService(repo)
	handler := handler.NewHandler(service, config)

	r := router.NewRouter(handler, service, logger)

	fmt.Println("server running on port " + config.Port)
	if err := http.ListenAndServe(":"+config.Port, r); err != nil {
		log.Fatal("error server")
	}
}
