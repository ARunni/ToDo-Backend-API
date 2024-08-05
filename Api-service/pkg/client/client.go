package client

import (
	"fmt"

	"github.com/ARunni/ToDo-Backend-API/Api-service/pkg/config"
	"github.com/ARunni/ToDo-Backend-API/Api-service/pkg/pb"
	"google.golang.org/grpc"
)

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.TodoSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect:", err)
	}
	return pb.NewAuthServiceClient(cc)
}
