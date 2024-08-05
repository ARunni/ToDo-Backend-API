package main

import (
	"log"

	_ "github.com/ARunni/ToDo-Backend-API/Api-service/cmd/docs"
	server "github.com/ARunni/ToDo-Backend-API/Api-service/pkg/api/server"
	"github.com/ARunni/ToDo-Backend-API/Api-service/pkg/config"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go + Gin TODO APP_Microservice_CleanArchitecture
// @version 1.0.0
// @description ToDo is a list of activities that is to be done by a particular individual
// @contact.name API Support
// @securityDefinitions.apikey BearerTokenAuth
// @in header
// @name Authorization
// @host localhost:8000
// @BasePath /
// @query.collection.format multi
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the env file")
	}
	config, Err := config.LoadConfig()
	if Err != nil {
		log.Fatal("cannot load config : ", Err)
	}
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.Routes(r, &config)
	r.Run(config.Port)
}
