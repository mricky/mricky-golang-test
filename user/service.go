package user

import "golang.org/x/crypto/bcrypt"

type UserService interface {
	RegisterUser(input RegisterUserInput)(User, error)
}

type service struct {
	reporsitory UserRepository
}

func ImplUserService(reporsitory UserRepository) *service {
	return &service{reporsitory}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error){
	user := User{}
	
	user.Name = input.Name
	user.Email = input.Email
	user.Username = input.Username
	user.Password = input.Password
	passwordHash,err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	user.ProfileId = 1
	
	newUser, err := s.reporsitory.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}