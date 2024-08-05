package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ARunni/ToDo-Backend-API/ToDo-service/pkg/api/services"
	"github.com/ARunni/ToDo-Backend-API/ToDo-service/pkg/config"
	"github.com/ARunni/ToDo-Backend-API/ToDo-service/pkg/db"
	"github.com/ARunni/ToDo-Backend-API/ToDo-service/pkg/pb"
	"github.com/ARunni/ToDo-Backend-API/ToDo-service/pkg/repository"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", c.DBHost, c.DBUser, c.DBName, c.DBPort, c.DBPassword)
	h := db.Init(psqlInfo)
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	fmt.Println("Todo Svc on", c.Port)

	repo := repository.NewDBServer(h)
	s := &services.Server{
		Repository: repo,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
