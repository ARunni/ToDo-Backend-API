package handler

import (
	"context"
	"net/http"

	"github.com/ARunni/ToDo-Backend-API/Api-service/pkg/pb"
	"github.com/ARunni/ToDo-Backend-API/Api-service/pkg/utils/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// @Summary		User Signup
// @Description	user can signup by giving their details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			UserSignupDetail   body  models.SignUpRequest  true  "User Signup"
// @Success 201 {object} pb.SignupResponse "Returns the newly created user"
// @Failure 400 {object} models.ErrorResponse "Error response"
// @Failure 502 {object} models.ErrorResponse "Bad gateway"
// @Router			/signup    [POST]
func Signup(ctx *gin.Context, c pb.AuthServiceClient) {
	var UserSignupDetail models.SignUpRequest
	if err := ctx.ShouldBindJSON(&UserSignupDetail); err != nil {
		errResp := models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
		}
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}
	if err := validator.New().Struct(UserSignupDetail); err != nil {
		errResp := models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Validation error: " + err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}
	user, err := c.Signup(context.Background(), &pb.SignupRequest{
		Name:     UserSignupDetail.Name,
		Email:    UserSignupDetail.Email,
		Password: UserSignupDetail.Password,
	})
	if err != nil {
		errResp := models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to register user: " + err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

// @Summary		User Login
// @Description	user can log in by giving their details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			UserLoginDetail  body  models.LoginRequest  true	"User Login"
// @Success		200	{object} pb.LoginResponse "Login successfully"
// @Failure		500	{object} models.ErrorResponse  "Bad request"
// @Router			/login     [POST]
func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	var UserLoginDetail models.LoginRequest
	if err := ctx.ShouldBindJSON(&UserLoginDetail); err != nil {
		errResp := models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
		}
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}
	if err := validator.New().Struct(UserLoginDetail); err != nil {
		errResp := models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Validation error: " + err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}
	user, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    UserLoginDetail.Email,
		Password: UserLoginDetail.Password,
	})
	if err != nil {
		errResp := models.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Authentication failed: " + err.Error(),
		}
		ctx.JSON(http.StatusUnauthorized, errResp)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
