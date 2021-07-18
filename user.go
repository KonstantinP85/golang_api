package api_go

type User struct {
	Id 		   int `json:"-"`
	Email 	   string `json:"email"`
	Password   string `json:"password"`
	IsActive   bool `json:"is_active"`
}
