package server

import (
	"github.com/ARunni/ToDo-Backend-API/Api-service/pkg/api/handler"
	"github.com/ARunni/ToDo-Backend-API/Api-service/pkg/api/middleware"
	"github.com/ARunni/ToDo-Backend-API/Api-service/pkg/client"
	"github.com/ARunni/ToDo-Backend-API/Api-service/pkg/config"
	"github.com/ARunni/ToDo-Backend-API/Api-service/pkg/pb"
	"github.com/gin-gonic/gin"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func Routes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: client.InitServiceClient(c),
	}
	r.POST("/signup", svc.SignUp)
	r.POST("/login", svc.Login)
	r.Use(middleware.UserAuthMiddleware())
	{
		routes := r.Group("/todo")
		{
			routes.POST("", svc.CreateTodo)
			routes.DELETE("", svc.DeleteTodo)
			routes.GET("/get", svc.GetTodo)
			routes.GET("", svc.ListAll)
			routes.PUT("", svc.updateTodo)
		}
	}
	return svc
}

func (svc *ServiceClient) SignUp(ctx *gin.Context) {
	handler.Signup(ctx, svc.Client)
}
func (svc *ServiceClient) Login(ctx *gin.Context) {
	handler.Login(ctx, svc.Client)
}
func (svc *ServiceClient) CreateTodo(ctx *gin.Context) {
	handler.CreateTodo(ctx, svc.Client)
}
func (svc *ServiceClient) DeleteTodo(ctx *gin.Context) {
	handler.DeleteTodo(ctx, svc.Client)
}
func (svc *ServiceClient) GetTodo(ctx *gin.Context) {
	handler.GetTodo(ctx, svc.Client)
}
func (svc *ServiceClient) ListAll(ctx *gin.Context) {
	handler.ListAllTodo(ctx, svc.Client)
}
func (svc *ServiceClient) updateTodo(ctx *gin.Context) {
	handler.UpdateTodo(ctx, svc.Client)
}
