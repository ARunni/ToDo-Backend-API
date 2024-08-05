package models

type SignUpRequest struct {
	Name     string `json:"name" validate:"gte=3"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=6,max=20"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Todo struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UserResponsewithPassword struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
