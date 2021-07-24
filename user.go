package api_go

type User struct {
	Id       int    `json:"_" db:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	IsActive bool   `json:"is_active" db:"is_active"`
}

type PatchUserInput struct {
	IsActive   bool `json:"is_active"`
}

type UpdateUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
