package models

type CreateRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateRequestBody struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
