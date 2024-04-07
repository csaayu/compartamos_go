package main

import (
	"compartamos_go/config"
	"compartamos_go/controller"
	"compartamos_go/helper"
	"compartamos_go/model"
	"compartamos_go/repository"
	"compartamos_go/router"
	"compartamos_go/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Servidor corriendo...:)")

	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("users").AutoMigrate(&model.Users{})

	//Repository
	userRepository := repository.NewUsersRePositoryImpl(db)

	// services
	userService := service.NewUsersServiceImpl(userRepository, validate)

	//controller
	userController := controller.NewUsersController(userService)

	//router

	routes := router.NewRouter(userController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()

	helper.ErrorSync(err)

}
