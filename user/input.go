package user

type RegisterUserInput struct {
	Name 	  string  `json:"name" binding:"required"`
	Email 	  string  `json:"email" binding:"required,email"`
	Username  string  `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	ProfileId int `json:"profile_id" binding:"required"`
}
