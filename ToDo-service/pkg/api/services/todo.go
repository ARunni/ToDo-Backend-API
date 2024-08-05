package services

import (
	"context"
	"errors"

	"github.com/ARunni/ToDo-Backend-API/ToDo-service/pkg/pb"
	"github.com/ARunni/ToDo-Backend-API/ToDo-service/pkg/repository"
	"github.com/ARunni/ToDo-Backend-API/ToDo-service/pkg/utils/helper"
	"github.com/ARunni/ToDo-Backend-API/ToDo-service/pkg/utils/models"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	Repository *repository.DBServer
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Signup(ctx context.Context, detail *pb.SignupRequest) (*pb.SignupResponse, error) {
	email, err := s.Repository.CheckUserExistsByEmail(detail.Email)
	if err != nil {
		return &pb.SignupResponse{}, errors.New("error with server")
	}
	if email != nil {
		return &pb.SignupResponse{}, errors.New("user with this email is already exists")
	}

	hashPassword, err := helper.PasswordHash(detail.Password)
	if err != nil {
		return &pb.SignupResponse{}, errors.New("error in hashing password")
	}
	details := models.SignUpRequest{
		Name:     detail.Name,
		Email:    detail.Email,
		Password: hashPassword,
	}
	userData, err := s.Repository.UserSignUp(details)
	if err != nil {
		return &pb.SignupResponse{}, errors.New("could not add the user")
	}
	accessToken, err := helper.GenerateAccessTokenUser(userData)
	if err != nil {
		return &pb.SignupResponse{}, errors.New("couldn't create access token due to error")
	}
	RefreshToken, err := helper.GenerateRefreshTokenUser(userData)
	if err != nil {
		return &pb.SignupResponse{}, errors.New("couldn't create access token due to error")
	}
	userInfo := &pb.UserInfo{
		Id:    int64(userData.Id),
		Name:  userData.Name,
		Email: userData.Email,
	}
	return &pb.SignupResponse{
		Info:         userInfo,
		Accesstoken:  accessToken,
		Refreshtoken: RefreshToken,
	}, nil
}

func (s *Server) Login(ctx context.Context, details *pb.LoginRequest) (*pb.LoginResponse, error) {
	email, err := s.Repository.CheckUserExistsByEmail(details.Email)
	if err != nil {
		return &pb.LoginResponse{}, errors.New("error with server")
	}
	if email == nil {
		return &pb.LoginResponse{}, errors.New("email doesn't exist")
	}
	detail := models.LoginRequest{
		Email:    details.Email,
		Password: details.Password,
	}
	userdeatils, err := s.Repository.FindUserByEmail(detail)
	if err != nil {
		return &pb.LoginResponse{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userdeatils.Password), []byte(details.Password))
	if err != nil {
		return &pb.LoginResponse{}, errors.New("password not matching")
	}
	var userData models.UserResponse
	err = copier.Copy(&userData, &userdeatils)
	if err != nil {
		return &pb.LoginResponse{}, err
	}
	accessToken, err := helper.GenerateAccessTokenUser(userData)
	if err != nil {
		return &pb.LoginResponse{}, errors.New("couldn't create access token due to error")
	}
	RefreshToken, err := helper.GenerateRefreshTokenUser(userData)
	if err != nil {
		return &pb.LoginResponse{}, errors.New("couldn't create access token due to error")
	}
	userInfo := &pb.UserInfo{
		Id:    int64(userData.Id),
		Name:  userData.Name,
		Email: userData.Email,
	}
	return &pb.LoginResponse{
		Info:         userInfo,
		Accesstoken:  accessToken,
		Refreshtoken: RefreshToken,
	}, nil
}

func (s *Server) CreateTodo(ctx context.Context, req *pb.AddTodoRequest) (*pb.AddTodoResponse, error) {
	userExist := s.Repository.UserIDExist(req.UserID)
	if !userExist {
		return &pb.AddTodoResponse{}, errors.New("user couldn't exist")
	}
	todo, err := s.Repository.CreateTodo(req.Title, req.Description, req.UserID)
	if err != nil {
		return &pb.AddTodoResponse{}, err
	}
	return &pb.AddTodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
	}, nil
}

func (s *Server) ListTodo(ctx context.Context, details *pb.ListTodoRequest) (*pb.ListTodoResponse, error) {
	userExist := s.Repository.UserIDExist(details.UserID)
	if !userExist {
		return &pb.ListTodoResponse{}, errors.New("user couldn't exist")
	}
	page := int(details.Page)
	count := int(details.Count)
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * count
	todos, err := s.Repository.ListTodo(count, offset, int(details.UserID))
	if err != nil {
		return &pb.ListTodoResponse{
			ListTodos: []*pb.TodoDetails{},
		}, err
	}

	var TodoDetails pb.ListTodoResponse
	for _, v := range todos {
		TodoDetails.ListTodos = append(TodoDetails.ListTodos, &pb.TodoDetails{
			ID:          int64(v.ID),
			Title:       v.Title,
			Description: v.Description,
		})
	}
	return &TodoDetails, nil
}

func (s *Server) GetTodoByID(ctx context.Context, details *pb.TodoIDRequest) (*pb.TodoItemResponse, error) {
	todoExist := s.Repository.TodoExist(details.ID, details.UserID)
	if !todoExist {
		return &pb.TodoItemResponse{}, errors.New("todolist is couldn't exist")
	}
	todo, err := s.Repository.GetTodoByID(details.ID, details.UserID)
	if err != nil {
		return &pb.TodoItemResponse{}, err
	}
	return &pb.TodoItemResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
	}, nil
}

func (s *Server) DeleteTodo(ctx context.Context, details *pb.TodoIDRequest) (*pb.DeleteTodoResponse, error) {
	todoExist := s.Repository.TodoExist(details.ID, details.UserID)
	if !todoExist {
		return &pb.DeleteTodoResponse{}, errors.New("todolist is couldn't exist")
	}
	err := s.Repository.DeleteTodoByID(details.ID, details.UserID)
	if err != nil {
		return &pb.DeleteTodoResponse{}, err
	}
	return &pb.DeleteTodoResponse{
		Status: "Deleted Successfully",
	}, nil
}

func (s *Server) UpdateTodo(ctx context.Context, details *pb.UpdateTodoRequest) (*pb.UpdateTodoResponse, error) {
	todoExist := s.Repository.TodoExist(details.ID, details.UserID)
	if !todoExist {
		return &pb.UpdateTodoResponse{}, errors.New("todolist is couldn't exist")
	}
	err := s.Repository.UpdateTodoByID(details.ID, details.UserID, details.Title, details.Description)
	if err != nil {
		return &pb.UpdateTodoResponse{}, err
	}

	todo, err := s.Repository.GetTodoByID(details.ID, details.UserID)
	if err != nil {
		return &pb.UpdateTodoResponse{}, err
	}

	return &pb.UpdateTodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
	}, nil
}
