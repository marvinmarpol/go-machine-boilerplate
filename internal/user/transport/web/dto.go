package web

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}

type CreateUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
