package repository

import (
	"errors"

	"github.com/ARunni/ToDo-Backend-API/ToDo-service/pkg/db"
	"github.com/ARunni/ToDo-Backend-API/ToDo-service/pkg/domain"
	"github.com/ARunni/ToDo-Backend-API/ToDo-service/pkg/utils/models"
	"gorm.io/gorm"
)

type DBServer struct {
	H db.Handler
}

func NewDBServer(handler db.Handler) *DBServer {
	return &DBServer{H: handler}
}

func (s *DBServer) CheckUserExistsByEmail(email string) (*domain.User, error) {
	var user domain.User
	res := s.H.DB.Where(&domain.User{Email: email}).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return &user, nil
}

func (s *DBServer) UserSignUp(user models.SignUpRequest) (models.UserResponse, error) {
	var signupDetails models.UserResponse
	err := s.H.DB.Raw(`INSERT INTO users (name,email,password,created_at) VALUES(?,?,?,NOW()) RETURNING id,name,email`, user.Name, user.Email, user.Password).Scan(&signupDetails).Error
	if err != nil {
		return models.UserResponse{}, err
	}
	return signupDetails, nil
}

func (s *DBServer) FindUserByEmail(user models.LoginRequest) (models.UserResponsewithPassword, error) {
	var userDetails models.UserResponsewithPassword
	err := s.H.DB.Raw("SELECT * FROM users WHERE email=?", user.Email).Scan(&userDetails).Error
	if err != nil {
		return models.UserResponsewithPassword{}, errors.New("error checking user details")
	}
	return userDetails, nil
}

func (s *DBServer) CreateTodo(title, description string, userID int64) (domain.Todo, error) {
	var todo domain.Todo
	err := s.H.DB.Raw("INSERT INTO todos (title, description,created_at,user_id) VALUES(?, ?,Now(),?) RETURNING id, title, description", title, description, userID).Scan(&todo).Error
	return todo, err
}

func (s *DBServer) ListTodo(limit, offset, userID int) ([]models.Todo, error) {
	var todos []models.Todo
	err := s.H.DB.Raw(`SELECT id, title, description FROM todos WHERE user_id = ? LIMIT ? OFFSET ?`, userID, limit, offset).Scan(&todos).Error
	return todos, err
}

func (s *DBServer) GetTodoByID(id, userID int64) (models.Todo, error) {
	var todo models.Todo
	err := s.H.DB.Raw("SELECT id, title, description FROM todos WHERE id = ? AND user_id = ?", id, userID).Scan(&todo).Error
	return todo, err
}

func (s *DBServer) DeleteTodoByID(id, userID int64) error {
	return s.H.DB.Exec("DELETE FROM todos WHERE id = ? AND user_id = ?", id, userID).Error
}

func (s *DBServer) UpdateTodoByID(id, userID int64, title, description string) error {
	err := s.H.DB.Exec("UPDATE todos SET title = $1, description = $2 WHERE id = $3 AND user_id = $4", title, description, id, userID).Error
	return err
}

func (s *DBServer) UserIDExist(userID int64) bool {
	var count int
	err := s.H.DB.Raw(`SELECT COUNT(*) FROM users WHERE id = ?`, userID).Scan(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}

func (s *DBServer) TodoExist(id, userID int64) bool {
	var count int
	err := s.H.DB.Raw(`SELECT COUNT(*) FROM todos WHERE id = ? AND user_id = ?`, id, userID).Scan(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}
