package user

type RegisterUserInput struct {
	Name 	  string  `json:"name" binding:"required"`
	Email 	  string  `json:"email" binding:"required,email"`
	Username  string  `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	ProfileId int `json:"profile_id" binding:"required"`
}

type LoginInput struct {
	Email 	  string  `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}
