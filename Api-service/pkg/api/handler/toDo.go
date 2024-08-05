package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ARunni/ToDo-Backend-API/Api-service/pkg/pb"
	"github.com/ARunni/ToDo-Backend-API/Api-service/pkg/utils/models"
	"github.com/gin-gonic/gin"
)

// @Summary Create a new todo item
// @Description Create a new todo item with the provided details
// @Accept     json
// @Produce    json
// @Security   BearerTokenAuth
// @Param b body models.CreateRequestBody true "Todo item details"
// @Success 201 {object} pb.AddTodoResponse "Returns the newly created todo item"
// @Failure 400 {object} models.ErrorResponse "Error response"
// @Failure 502 {object} models.ErrorResponse "Bad gateway"
// @Router /todo [POST]
func CreateTodo(ctx *gin.Context, asc pb.AuthServiceClient) {
	userID, _ := ctx.Get("user_id")
	UserID := userID.(int)
	b := models.CreateRequestBody{}
	if err := ctx.BindJSON(&b); err != nil {
		errResp := models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
		}
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}

	res, err := asc.CreateTodo(context.Background(), &pb.AddTodoRequest{
		Title:       b.Title,
		Description: b.Description,
		UserID:      int64(UserID),
	})
	if err != nil {
		errResp := models.ErrorResponse{
			Code:    http.StatusBadGateway,
			Message: "Failed to create todo item",
		}
		ctx.JSON(http.StatusBadGateway, errResp)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

// @Summary Delete a todo item
// @Description Delete a todo item with the provided ID
// @Accept     json
// @Produce    json
// @Security   BearerTokenAuth
// @Param 		id 	query 	string true 	"todo id"
// @Success 200 {object} pb.DeleteTodoResponse "Returns the deletion status"
// @Failure 400 {object} models.ErrorResponse "Error response"
// @Failure 502 {object} models.ErrorResponse "Bad gateway"
// @Router /todo [DELETE]
func DeleteTodo(ctx *gin.Context, asc pb.AuthServiceClient) {
	userID, _ := ctx.Get("user_id")
	UserID := userID.(int)
	id := ctx.Query("id")
	todoID, err := strconv.Atoi(id)
	if err != nil {
		errResp := models.ErrorResponse{
			Code:    http.StatusBadGateway,
			Message: "todo_id not in right format",
		}
		ctx.JSON(http.StatusBadGateway, errResp)
		return
	}
	res, err := asc.DeleteTodo(context.Background(), &pb.TodoIDRequest{
		ID:     int64(todoID),
		UserID: int64(UserID),
	})
	if err != nil {
		errResp := models.ErrorResponse{
			Code:    http.StatusBadGateway,
			Message: "Failed to delete todo item",
		}
		ctx.JSON(http.StatusBadGateway, errResp)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary Get a todo item by ID
// @Description Get a todo item details by the provided ID
// @Accept     json
// @Produce    json
// @Security   BearerTokenAuth
// @Param 		id 	query 	string true 	"todo id"
// @Success 200 {object} pb.TodoItemResponse "Returns the fetched todo item"
// @Failure 400 {object} models.ErrorResponse "Error response"
// @Failure 502 {object} models.ErrorResponse "Bad gateway"
// @Router /todo/get [GET]
func GetTodo(ctx *gin.Context, asc pb.AuthServiceClient) {
	userID, _ := ctx.Get("user_id")
	UserID := userID.(int)
	id := ctx.Query("id")
	todoID, err := strconv.Atoi(id)
	if err != nil {
		errResp := models.ErrorResponse{
			Code:    http.StatusBadGateway,
			Message: "todo_id not in right format",
		}
		ctx.JSON(http.StatusBadGateway, errResp)
		return
	}
	res, err := asc.GetTodoByID(context.Background(), &pb.TodoIDRequest{
		ID:     int64(todoID),
		UserID: int64(UserID),
	})
	if err != nil {
		errResp := models.ErrorResponse{
			Code:    http.StatusBadGateway,
			Message: "Failed to fetch todo item",
		}
		ctx.JSON(http.StatusBadGateway, errResp)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary List all todo items
// @Description List all todo items with pagination support
// @Accept     json
// @Produce    json
// @Security   BearerTokenAuth
// @Param 	page 	query 	string	 false	 "Page number"
// @Param 	count 	query 	string 	false	 "Page size"
// @Success 200 {object} pb.ListTodoResponse "Returns the list of todo items"
// @Failure 400 {object} models.ErrorResponse "Error response"
// @Failure 502 {object} models.ErrorResponse "Bad gateway"
// @Router /todo [GET]
func ListAllTodo(ctx *gin.Context, asc pb.AuthServiceClient) {
	userID, _ := ctx.Get("user_id")
	UserID := userID.(int)
	pageStr := ctx.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errResp := models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid page parameter",
		}
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}

	countStr := ctx.DefaultQuery("count", "100")
	pageSize, err := strconv.Atoi(countStr)
	if err != nil {
		errResp := models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid count parameter",
		}
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}

	res, err := asc.ListTodo(context.Background(), &pb.ListTodoRequest{
		Page:   int64(page),
		Count:  int64(pageSize),
		UserID: int64(UserID),
	})
	if err != nil {
		errResp := models.ErrorResponse{
			Code:    http.StatusBadGateway,
			Message: "Failed to list todo items",
		}
		ctx.JSON(http.StatusBadGateway, errResp)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary Update a todo item
// @Description Update a todo item with the provided details
// @Accept     json
// @Produce    json
// @Security   BearerTokenAuth
// @Param request body models.UpdateRequestBody true "Todo item details to update"
// @Success 200 {object} pb.UpdateTodoResponse "Returns the updated todo item"
// @Failure 400 {object} models.ErrorResponse "Error response"
// @Failure 502 {object} models.ErrorResponse "Bad gateway"
// @Router /todo [PUT]
func UpdateTodo(ctx *gin.Context, asc pb.AuthServiceClient) {
	userID, _ := ctx.Get("user_id")
	UserID := userID.(int)
	b := models.UpdateRequestBody{}
	if err := ctx.BindJSON(&b); err != nil {
		errResp := models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
		}
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}

	res, err := asc.UpdateTodo(context.Background(), &pb.UpdateTodoRequest{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		UserID:      int64(UserID),
	})
	if err != nil {
		errResp := models.ErrorResponse{
			Code:    http.StatusBadGateway,
			Message: "Failed to update todo item",
		}
		ctx.JSON(http.StatusBadGateway, errResp)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
