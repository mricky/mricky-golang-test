package user



type UserFormater struct {
	ID int			  `json:"id"`
	NAME string 	  `json:"name"`
	Email string 	  `json:"email"`
	Token string      `json:"token"`
}

func FormaterUser (user User,token string) UserFormater{
	formater := UserFormater{
		ID: user.Id,
		NAME: user.Email,
		Email: user.Email,
		Token: token,
		
	}

	return formater
}

