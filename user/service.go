package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(input RegisterUserInput)(User, error)
	Login(input LoginInput)(User, error)
	GetUserByID(userId int) (User, error)

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

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.reporsitory.FindByEmail(email)

	if err != nil {
		return user, err
	}
	if user.Id == 0 {
		return user, errors.New("No found email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))

	if err != nil {
		return user,err
	}

	return user, nil
}

func (s *service) GetUserByID(userId int) (User, error){
	
	user, err := s.reporsitory.FindByID(userId)

	if err != nil {
		return user, err
	}
	return user, nil
}